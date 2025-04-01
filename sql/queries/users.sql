-- Function :(returns what)

-- name: CreateUser :one 
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4) -- This would create a function CreateUser, with 4 parameters
RETURNING *;

-- sqlc generate -> Would convert the sql to .go type safe code under /internal (that is defined in our yaml)
 
