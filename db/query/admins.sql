-- name: GetAdminByID :one
SELECT user_id
FROM admins
WHERE user_id = $1;
