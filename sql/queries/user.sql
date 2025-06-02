-- name: GetUserById :one
select *
from user
where user.id = ?
;

-- name: GetUser :one
select *
from user
where user.email = ? or user.username = ?
;

-- name: CreateUser :one
INSERT INTO user (
    id,
    firstname,
    lastname,
    username,
    alias,
    email,
    password,
    phone_number,
    birth_date,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateUser :one
UPDATE user SET
    firstname = ?,
    lastname = ?,
    alias = ?,
    birth_date = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: UpdateEmail :one
UPDATE user SET
    email = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: UpdateUsername :one
UPDATE user SET
    username = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: UpdatePassword :one
UPDATE user SET
    password = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
UPDATE user SET 
    deleted_at = ?
WHERE id = ?;

