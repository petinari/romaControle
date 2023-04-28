-- name: InsertSubGrupoProduto :one
insert into public.subgrupo_produtos (nome, id_grupo, ativo, id_tenant) values ($1, $2, $3, $4) returning *;

-- name: UpdateSubGrupoProduto :one
update public.subgrupo_produtos set nome = $1, ativo = $2 where id = $3 and id_tenant = $4 returning *;

-- name: DesativarSubGrupoProduto :one
update public.subgrupo_produtos set ativo = false where id = $1 and id_tenant = $2 returning *;

-- name: AtivarSubGrupoProduto :one
update public.subgrupo_produtos set ativo = true where id = $1 and id_tenant = $2 returning *;

-- name: GetSubGrupoProdutoById :one
select * from public.subgrupo_produtos where id = $1 and id_tenant = $2 and ativo = true;

-- name: GetSubGrupoProdutoByNome :one
select * from public.subgrupo_produtos where nome = $1 and id_tenant = $2 and ativo = true;

-- name: GetSubGrupoProdutoByGrupo :many
select * from public.subgrupo_produtos where id_grupo = $1 and id_tenant = $2 and ativo = true;

