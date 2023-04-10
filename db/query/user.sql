-- name: CreateUsuario :one
insert into usuarios (email, senha, id_tenant)
values ($1, $2, $3) returning *;

-- name: SelectUsuarioPorEmail :one
select * from usuarios where email = ($1);