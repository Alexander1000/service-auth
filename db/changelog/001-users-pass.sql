create table users_pass
(
  user_id bigint not null,
  pass_hash character varying,
  pass_salt character varying,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  constraint users_pass_pkey primary key (user_id)
);
