create table users_auth_credentials
(
  auth_id bigint not null,
  user_id bigint not null,
  credential_id bigint not null,
  type_id integer not null,
  created_at timestamp with time zone,
  constraint users_auth_credentials_pkey primary key (auth_id)
);

CREATE SEQUENCE users_auth_credentials_auth_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY users_auth_credentials ALTER COLUMN auth_id SET DEFAULT nextval('users_auth_credentials_auth_id_seq'::regclass);
