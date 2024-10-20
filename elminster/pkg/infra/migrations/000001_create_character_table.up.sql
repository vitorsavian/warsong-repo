BEGIN;

CREATE TABLE IF NOT EXISTS stats(
   id UUID PRIMARY KEY,
   hp INT NOT NULL,
   sp INT NOT NULL,
   str INT NOT NULL,
   dex INT NOT NULL,
   con INT NOT NULL,
   inte INT NOT NULL,
   wil INT NOT NULL,
   cha INT NOT NULL,
   sanity INT NOT NULL,
   courage INT NOT NULL
);

CREATE TABLE IF NOT EXISTS characters(
   id UUID PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   level INT NOT NULL,
   id_stats UUID NOT NULL,
   FOREIGN KEY (id_stats) REFERENCES stats(id)
);

COMMIT;