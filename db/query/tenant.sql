-- name: CreateTenant :one
insert into tenant default values returning *;