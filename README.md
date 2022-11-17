# go
- install go using https://go.dev/doc/install
- install vs code
- install go extension
- using a command line install packages
  - go get github.com/go-kit/kit/endpoint
  - go get github.com/go-kit/kit/transport/http
  
##Database
 - github.com/lib/pq
 - docker pull postgres
 - docker run --name my-postgres -p 5432:5432 -e POSTGRES_PASSWORD=go -d postgres
 - CREATE DATABASE users
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
 - CREATE SEQUENCE IF NOT EXISTS public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;
 - ALTER SEQUENCE public.users_id_seq
    OWNER TO postgres;
 - CREATE TABLE IF NOT EXISTS public.users
  (
    id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    name text COLLATE pg_catalog."default" NOT NULL,
    pass text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
  )
- TABLESPACE pg_default;
- ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

 - added "host  all  all 0.0.0.0/0 md5" to "C:\Program Files\PostgreSQL\15\data\pg_hba.conf" for testing purposes as appplication connects to external IP address server

 ##Testing
 - "go test" (in the current directory)

##Docker (all commands in the root dir)
 - "docker build . -t 6204931/users"
 - "docker run --name my-users -p 12345:12345 6204931/users"