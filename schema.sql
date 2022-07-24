CREATE TABLE currency (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "code" VARCHAR(5) NOT NULL,
  "value" FLOAT NOT NULL,
  "updated_at" TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);