CREATE TABLE "Tags" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Task" (
  "id" bigserial PRIMARY KEY,
  "text" text NOT NULL,
  "tag_id" bigint NOT NULL,
  "due" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "Task" ADD FOREIGN KEY ("tag_id") REFERENCES "Tags" ("id");

CREATE INDEX ON "Task" ("tag_id");

CREATE INDEX ON "Task" ("due");
