-- name: CreateGrupoProdutos :one
insert into public.grupo_produtos (nome, id_tenant, ativo) values ($1, $2, $3) returning *;

-- name: SelectGrupoProdutos :many
SELECT id, nome, id_tenant, ativo FROM public.grupo_produtos where id_tenant = $1 and ativo = TRUE;

-- name: SelectGrupoProdutosById :one
SELECT id, nome, id_tenant, ativo FROM public.grupo_produtos where id = $1 and id_tenant = $2 AND ativo = TRUE;

-- name: UpdateGrupoProdutos :one
UPDATE public.grupo_produtos SET nome = $1, ativo = $2 WHERE id = $3 and id_tenant = $4 returning *;

-- name: SelectGrupoProdutosByNome :one
SELECT id, nome, ativo, id_tenant FROM public.grupo_produtos where nome = $1 and id_tenant = $2 AND ativo = TRUE;

-- name: DisableGrupoProdutos :one
UPDATE public.grupo_produtos SET ativo = FALSE WHERE id = $1 and id_tenant = $2 returning *;