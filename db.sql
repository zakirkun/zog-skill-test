CREATE TYPE "status_news" AS ENUM (
  'draft',
  'deleted',
  'published'
);

CREATE TABLE "topics" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "slug" varchar(255) NOT NULL,
  "created_at" datetime,
  "updated_at" datetime
);

CREATE TABLE "news" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "title" varchar(255) NOT NULL,
  "slug" varchar(255) NOT NULL,
  "thumbnail" text NOT NULL,
  "status" status_news DEFAULT 'draft',
  "content" longtext,
  "created_at" datetime,
  "updated_at" datetime
);

CREATE TABLE "news_topics" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "id_topic" int(6),
  "id_news" int(6)
);

ALTER TABLE "topics" ADD FOREIGN KEY ("id") REFERENCES "news_topics" ("id_topic");

ALTER TABLE "news_topics" ADD FOREIGN KEY ("id_news") REFERENCES "news" ("id");
