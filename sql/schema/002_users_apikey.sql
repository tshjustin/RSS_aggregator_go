-- +goose up 
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT(
    encode(sha256(random()::text::bytea), 'hex')  -- We need default since we set NOT NULL and we already have records inside the DB 
); -- Generate random bytes and case cast binary, then fix to 256 for fixed output and encode to hex 

-- +goose down
ALTER TABLE users DROP COLUMN api_key;