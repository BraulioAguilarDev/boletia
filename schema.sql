CREATE TABLE currency (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "code" VARCHAR(28) NOT NULL,
  "value" VARCHAR(100) NOT NULL,
  "updated_at" TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);