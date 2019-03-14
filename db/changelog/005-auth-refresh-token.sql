create table auth_refresh_tokens
(
  refresh_token_id bigint not null,
  token_id bigint not null,
  status_id integer not null,
  created_at timestamp with timezone,
  updated_at timestamp with timezone,
  token character varying,
  expire_at timestamp with timezone
);
