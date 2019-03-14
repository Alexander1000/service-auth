create table auth_refresh_tokens
(
  refresh_token_id bigint not null,
  token_id bigint not null,
  status_id integer not null,
  created_at timestamp with timezone,
  updated_at timestamp with timezone,
  token character varying,
  expire_at timestamp with timezone,
  constraint auth_refresh_tokens_pkey primary key (refresh_token_id)
);

CREATE SEQUENCE auth_refresh_tokens_refresh_token_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY auth_refresh_tokens ALTER COLUMN refresh_token_id SET DEFAULT nextval('auth_refresh_tokens_refresh_token_id_seq'::regclass);

create unique index auth_refresh_tokens_token_ux on auth_refresh_tokens (token) where status_id = 0;

create index auth_refresh_tokens_token on auth_refresh_tokens using btree (token_id) where status_id = 0;

-- queue for move from active to expired
create index auth_refresh_tokens_active on auth_refresh_tokens using btree (expire_at) where status_id = 0;
