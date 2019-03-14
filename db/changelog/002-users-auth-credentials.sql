create table users_auth_credentials
(
  auth_id bigint not null,
  user_id bigint not null,
  credential_id bigint not null,
  type_id integer not null,
  created_at timestamp with time zone
);
