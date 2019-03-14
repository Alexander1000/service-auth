create table auth_credentials_type
(
  type_id integer not null,
  name character varying,
  title character varying,
  constraint auth_credentials_type_pkey primary key (type_id)
);
