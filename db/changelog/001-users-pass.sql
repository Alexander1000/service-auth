create table users_pass
(
  user_id bigint not null,
  pass_hash character varying,
  pass_salt character varying,
  created_at timestamp with timezone,
  updated_at timestamp with timezone
);
