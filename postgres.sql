-- CREATED_AT = 21-02-2021 01:00:00
-- create table `blogs` with id using uuid
DROP TABLE IF EXISTS public.blogs;
CREATE TABLE public.blogs
(
    id char(36) not null constraint blogs_pkey primary key,
    title varchar(100) not null,
    body varchar(1000) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp
);
CREATE INDEX idx_deleted_at ON public.blogs(deleted_at);
