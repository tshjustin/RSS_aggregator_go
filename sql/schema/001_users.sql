-- Migrations are ordered - and they allow for managing changes to DB over time 
-- think of it as version control for DB schemas 
-- goose is a tool to handle migrations, similar to albemic python package 
-- Below up => up migration , below down => down migration (downgrade)


-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- goose postgres postgresql:... up | schemas > Tables > ... 

-- +goose Down 

DROP TABLE users;