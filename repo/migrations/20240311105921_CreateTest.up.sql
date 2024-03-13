CREATE TABLE test (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    bio text
); 

CREATE INDEX IF NOT EXISTS "idx_test_id" ON "test" ("id");