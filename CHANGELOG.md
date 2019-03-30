# Changelog service-auth

## 2019-03-30
### Added
  - Добавлен handler для refresh токенов
  - Обновлен TODO
  - storage logout
  - add handler for logout

## 2019-03-24
### Added
  - добавлен хендлер проверяющий авторизацию по токену
  - добавлен метод в сторадж выполняющий рефреш токена

## 2019-03-23
### Changed
  - Сохранение refresh_token
  - время жизни токенов: access: 48h, для refresh: 1 month
  - добавлен хендлер authenticate
  - из хендлера отдается пара токенов (access + refresh)
  - исправлены индексы
### Added
  - добавлена функция для авторизации по токену

## 2019-03-18
### Changed
  - Сохранение токена

## 2019-03-17
### Added
  - наброски аутентификации в storage
  - проверка пароля

## 2019-03-16
### Added
  - add handler registration

## 2019-03-15
### Added
  - add entrypoint
  - add config
  - handle os signals
  - add makefile
  - add storage package
  - go pkg include

## 2019-03-14
### Added
  - initial commit
  - add db migrations
