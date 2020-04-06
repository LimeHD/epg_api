[![Build Status](https://ci.iptv2022.com/app/rest/builds/buildType(id:backend_root_EpgApi_EpgApiMaster)/statusIcon)](https://ci.iptv2022.com/viewType.html?buildTypeId=backend_root_EpgApi_EpgApiMaster)

# API для EPG

Проект является сервисом по выдаче телепрограммы по дням на все телеканалы Лайм и Премиум Лайм. Разработан в рамках разделения общего бекенда на независимые единицы.

### Usage

`./epg_api --dbuser {username_here} --dbpass {passowd_here} --dbname {db_name_here} --dbhost @`

### Swagger docs

- `make swagger-init`
- `make swagger-generate`
- `make swagger-${OS}`, where `$OS` one of [`windows`, `linux`]

### Tests

- Запуск тестов: `go test`
