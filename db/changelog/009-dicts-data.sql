insert into auth_credentials_type(type_id, name, title)
values
  (0, 'login', 'Login'),
  (1, 'email', 'E-mail'),
  (2, 'phone', 'Phone');

insert into auth_token_status(status_id, title)
values
  (0, 'Active'),
  (1, 'Refreshed'),
  (2, 'Disabled');

insert into auth_refresh_token_status(status_id, title)
values
  (0, 'Active'),
  (1, 'Refreshed'),
  (2, 'Disabled');
