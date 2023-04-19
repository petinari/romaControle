-- name: CreateTenant :one
insert into tenant default values returning *;

-- name: SelectTenantJoinUsuario :many
SELECT T.id , u.email
	FROM public.tenant as T inner join public.usuarios as u on id_tenant = T.id;