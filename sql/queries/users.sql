-- Function :(returns what)

-- name: CreateUser :one 
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, 
    encode(sha256(random()::text::bytea), 'hex')
) -- This would create a function CreateUser, with 4 parameters
RETURNING *;

-- NOTE must run sqlc generate -> Would convert the sql to .go type safe code under /internal (that is defined in our yaml)
 
-- Now we need to update our query -- SQL would handle the creation of the API Key once we add this function inside ! 

-- name: GetUserByAPIKey :one 
SELECT * FROM users WHERE api_key = $1;