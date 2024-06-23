-- SELECT 'CREATE DATABASE store_db WITH ENCODING = ''UTF8'';' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'store_db')\gexec
------------------------------------------------------------------
--- Level 1 -> Product
------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "authors" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "nationality" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "books" (
    "id" UUID PRIMARY KEY NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "gender" VARCHAR(100) NOT NULL,
    "author_id" UUID NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);

------------------------------------------------------------------
-- INDEX
------------------------------------------------------------------

CREATE INDEX IF NOT EXISTS "idx_authors_name" ON "authors" (name);
CREATE INDEX IF NOT EXISTS "idx_books_title" ON "books" (title);

ALTER TABLE "books" ADD CONSTRAINT "fk_book_author" FOREIGN KEY ("author_id") REFERENCES "authors" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
--
--
-- ------------------------------------------------------------------
-- -- LIMPEZA DA BASE
-- ------------------------------------------------------------------
--
TRUNCATE TABLE authors RESTART IDENTITY CASCADE;
TRUNCATE TABLE books RESTART IDENTITY CASCADE;
