create table auth_tokens
(
  token_id bigint not null,
  auth_id bigint not null,
  token character varying,
  created_at timestamp with timezone,
  updated_at timestamp with timezone,
  status_id integer not null,
  expire_at timestamp with timezone,
  constraint auth_tokens_pkey primary key (token_id)
);

CREATE SEQUENCE auth_tokens_token_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY auth_tokens ALTER COLUMN token_id SET DEFAULT nextval('auth_tokens_token_id_seq'::regclass);
