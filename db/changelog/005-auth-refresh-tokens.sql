create table auth_refresh_tokens
(
  refresh_token_id bigint not null,
  token_id bigint not null,
  status_id integer not null,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  token character varying,
  expire_at timestamp with time zone,
  constraint auth_refresh_tokens_pkey primary key (refresh_token_id)
);

CREATE SEQUENCE auth_refresh_tokens_refresh_token_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY auth_refresh_tokens ALTER COLUMN refresh_token_id SET DEFAULT nextval('auth_refresh_tokens_refresh_token_id_seq'::regclass);

create unique index auth_refresh_tokens_token_ux on auth_refresh_tokens (token);

create unique index auth_refresh_tokens_token_id_ux on auth_refresh_tokens(token_id);
