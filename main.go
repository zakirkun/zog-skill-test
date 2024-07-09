package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/zakirkun/zot-skill-test/bootstrap"
	"github.com/zakirkun/zot-skill-test/pkg/config"
	"github.com/zakirkun/zot-skill-test/pkg/database"
	"github.com/zakirkun/zot-skill-test/pkg/server"
	"github.com/zakirkun/zot-skill-test/router"
	"golang.org/x/exp/rand"
)

var configFile *string

var (
	c = promauto.NewCounter(prometheus.CounterOpts{
		Name: "zog_news_app_sample_metric",
		Help: "Sample metric for News Services",
	})

	h = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "zog_news_app_sample_histogram",
		Help: "Sample histogram for News Services",
	})

	d = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "zog_news_app_sample_devices",
		Help: "Sample counter opts devices for News Services"}, []string{"device"})

	e = promauto.NewCounter(prometheus.CounterOpts{
		Name: "zog_news_app_push_metric",
		Help: "Sample metric for News Services course (push)",
	})
)

func init() {
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()
}

func main() {
	setConfig()
	setMaxprocs()

	go func() {
		for {
			rand.Seed(uint64(time.Now().UnixNano()))
			h.Observe(float64(rand.Intn(100-0+1) + 0))
			d.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
			c.Inc()
			fmt.Print(".")
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			// Example of metric push
			err := push.New("http://pushgateway:9091", "zog_news_job").Collector(e).Add()
			if err != nil {
				_ = fmt.Errorf("%v", err)
			}
			e.Inc()
			fmt.Print("_")
			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	initApp := bootstrap.NewInfrastructure(SetDatabase(), SetWebServer())
	initApp.Database()
	initApp.WebServer()
}

func setMaxprocs() {
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
}

func setConfig() {
	cfg := config.NewConfig(*configFile)
	if err := cfg.Initialize(); err != nil {
		log.Fatalf("Error reading config : %v", err)
		os.Exit(1)
	}
}

func SetDatabase() database.DBModel {
	return database.DBModel{
		ServerMode:   config.GetString("server.mode"),
		Driver:       config.GetString("database.db_driver"),
		Host:         config.GetString("database.db_host"),
		Port:         config.GetString("database.db_port"),
		Name:         config.GetString("database.db_name"),
		Username:     config.GetString("database.db_username"),
		Password:     config.GetString("database.db_password"),
		MaxIdleConn:  config.GetInt("pool.conn_idle"),
		MaxOpenConn:  config.GetInt("pool.conn_max"),
		ConnLifeTime: config.GetInt("pool.conn_lifetime"),
	}
}

func SetWebServer() server.ServerContext {
	return server.ServerContext{
		Host:         ":" + config.GetString("server.port"),
		Handler:      router.NewRouter(),
		ReadTimeout:  time.Duration(config.GetInt("server.http_timeout")),
		WriteTimeout: time.Duration(config.GetInt("server.http_timeout")),
	}
}
