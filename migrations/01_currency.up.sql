SET timezone = 'America/Mexico_City';

-- Enable uuid_generate_v4
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Currency table
CREATE TABLE currency (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "code" VARCHAR(5) NOT NULL,
  "value" FLOAT NOT NULL,
  "updated_at" TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

-- Log table
CREATE TABLE log (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "duration" INTEGER NOT NULL,
  "code" INTEGER NOT NULL,
  "request" TEXT NOT NULL,
  "error" TEXT DEFAULT '' NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);