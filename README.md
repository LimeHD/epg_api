[![Build Status](https://ci.iptv2022.com/app/rest/builds/buildType(id:backend_root_EpgApi_EpgApiMaster)/statusIcon)](https://ci.iptv2022.com/viewType.html?buildTypeId=backend_root_EpgApi_EpgApiMaster)

# API для EPG

HTTP-сервис по выдаче телепрограммы по дням на все телеканалы Лайм и Премиум Лайм. Разработан в рамках разделения общего бекенда на независимые единицы.

### Endpoints

- `/channels` - return list of all channels
- `/channel/{id}/programm` - return tv programm for target channel, option get params: `curdate`, `tz`, `msk` all `integer`

### Usage

`./epg_api --dbuser {username_here} --dbpass {passowd_here} --dbname {db_name_here} --dbhost @`

### Swagger docs

- `make swagger-init`
- `make swagger-generate`
- `make swagger-${OS}`, where `$OS` one of [`windows`, `linux`]

### Tests

- Запуск тестов: `go test`

## Разворачивание не сервере

Текущую или любую ветку (запросит при деплое)

> bundle exec cap production deploy

Ветку master

> BRANCH=master bundle exec cap production deploy

Зайти на боевой сервер и посмотреть что там да как

> cap production shell

