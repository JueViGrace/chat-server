-- name: GetSessionById :one
select *
from session
where id = ?
;

-- name: CreateSession :exec
insert into session(
    id,
    refresh_token,
    access_token,
    user_id
)
values (?, ?, ?, ?);

-- name: UpdateSession :exec
update session set
    refresh_token = ?,
    access_token = ?
where id = ?;

-- name: DeleteSessionById :exec
delete from session
where id = ?
;

-- name: DeleteSessionByToken :exec
delete from session
where refresh_token = ? or access_token = ?
;

