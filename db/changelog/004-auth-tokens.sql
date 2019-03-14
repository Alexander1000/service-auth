create table auth_tokens
(
  token_id bigint not null,
  auth_id bigint not null,
  token character varying,
  created_at timestamp with timezone,
  updated_at timestamp with timezone,
  status_id integer not null,
  expire_at timestamp with timezone
);
