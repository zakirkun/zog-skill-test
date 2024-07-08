package bootstrap

import (
	"os"

	"github.com/zakirkun/zot-skill-test/app/domain/models"
	"github.com/zakirkun/zot-skill-test/pkg/database"
	"github.com/zakirkun/zot-skill-test/pkg/server"
)

type Infrastructure interface {
	Database()
	WebServer()
}

type infrastructureContext struct {
	database database.DBModel
	server   server.ServerContext
}

func NewInfrastructure(database database.DBModel,
	server server.ServerContext,
) Infrastructure {
	return infrastructureContext{
		database: database,
		server:   server,
	}
}

func (i infrastructureContext) Database() {
	db, err := i.database.OpenDB()
	if err != nil {
		os.Exit(1)
	}

	// Running Migration //
	// Create the custom ENUM type
	db.Exec("CREATE TYPE status_news AS ENUM ('draft', 'deleted', 'published');")

	// Migrate the schema
	db.AutoMigrate(&models.Topic{}, &models.News{}, &models.NewsTopic{})

	// Adding foreign key constraints manually
	db.Exec("ALTER TABLE news_topics ADD CONSTRAINT fk_news_topic FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE;")
	db.Exec("ALTER TABLE news_topics ADD CONSTRAINT fk_topic_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE;")

	database.DB = &i.database
}

func (i infrastructureContext) WebServer() {
	i.server.Run()
}
