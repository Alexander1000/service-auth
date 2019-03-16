create table auth_tokens
(
  token_id bigint not null,
  auth_id bigint not null,
  token character varying,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  status_id integer not null,
  expire_at timestamp with time zone,
  constraint auth_tokens_pkey primary key (token_id)
);

CREATE SEQUENCE auth_tokens_token_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY auth_tokens ALTER COLUMN token_id SET DEFAULT nextval('auth_tokens_token_id_seq'::regclass);

create unique index auth_tokens_token_ux on auth_tokens (token) where status_id in (0, 1);

-- queue on move from active to expired
create index auth_tokens_active_idx on auth_tokens using btree (expire_at) where status_id = 0;

-- queue on move from expired to disabled
create index auth_tokens_expired_idx on auth_tokens using btree (expire_at) where status_id = 1;
