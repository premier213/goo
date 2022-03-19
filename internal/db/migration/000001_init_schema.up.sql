CREATE TABLE public.users (
    id bigserial PRIMARY KEY,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL
);
