-- name: CreateUser :exec
    INSERT INTO users (user_id, username) VALUES ($1, $2);