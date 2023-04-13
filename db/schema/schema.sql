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

create table public.grupo_produtos
(
    id        uuid
        constraint grupo_produtos_pk
            primary key,
    nome      text,
    id_tenant uuid
        constraint grupo_produtos_tenant_id_fk
            references public.tenant
);

alter table public.grupo_produtos
    alter column id_tenant set not null;


alter table public.usuarios
    alter column id_tenant set not null;

alter table public.grupo_produtos
    alter column nome set not null;

alter table public.grupo_produtos
    alter column id set default uuid_generate_v4();
