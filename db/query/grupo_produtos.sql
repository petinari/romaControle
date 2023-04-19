-- name: CreateGrupoProdutos :one
insert into grupo_produtos (nome, id_tenant)
values ($1, $2) returning *;

-- name: SelectGrupoProdutos :many
SELECT id, nome, id_tenant
FROM public.grupo_produtos where id_tenant = $1;

