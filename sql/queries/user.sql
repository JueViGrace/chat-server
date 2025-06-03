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

-- name: GetEmail :one
select email
from user
where user.email = ?
;

-- name: GetUsername :one
select username
from user
where user.username = ?
;

-- name: GetPhoneNumber :one
select phone_number
from user
where user.phone_number = ?
;

-- name: CreateUser :exec
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
    role
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE user SET
    firstname = ?,
    lastname = ?,
    alias = ?,
    birth_date = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateEmail :exec
UPDATE user SET
    email = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUsername :exec
UPDATE user SET
    username = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdatePassword :exec
UPDATE user SET
    password = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUser :exec
UPDATE user SET 
    deleted_at = CURRENT_TIMESTAMP
WHERE id = ?;

