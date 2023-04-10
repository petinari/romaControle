create table if not exists public.tenant
(
    id uuid default uuid_generate_v4() not null
    constraint tenant_pk
    primary key
    );

alter table public.tenant
    owner to postgres;

create table if not exists public.usuarios
(
    id        uuid default uuid_generate_v4() not null
    constraint usuarios_pk
    primary key,
    email     text                            not null,
    senha     text                            not null,
    id_tenant uuid
    constraint usuarios_tenant_id_fk
    references public.tenant
    );

alter table public.usuarios
    owner to postgres;

alter table public.usuarios
    add constraint usuarios_pk2
        unique (email);
