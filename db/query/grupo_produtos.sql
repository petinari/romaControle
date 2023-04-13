-- name: CreateGrupoProdutos :one
insert into grupo_produtos (nome, id_tenant)
values ($1, $2) returning *;