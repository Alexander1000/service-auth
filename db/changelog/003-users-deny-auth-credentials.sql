create table users_deny_auth_credentials
(
  auth_id bigint not null,
  created_at timestamp with time zone,
  constraint users_deny_auth_credentials_pkey primary key (auth_id)
);
