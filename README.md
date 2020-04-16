[![Build Status](https://ci.iptv2022.com/app/rest/builds/buildType(id:backend_root_EpgApi_EpgApiMaster)/statusIcon)](https://ci.iptv2022.com/viewType.html?buildTypeId=backend_root_EpgApi_EpgApiMaster)

# API для EPG

HTTP-сервис по выдаче телепрограммы по дням на все телеканалы Лайм и Премиум Лайм. Разработан в рамках разделения общего бекенда на независимые единицы.

## TODO

* [ ] Логи в ./log/application.log
* [ ] add bugsnag
* [ ] избавиться от билда на разворачиваемом сервере

### Endpoints

- `/channels` - return list of all channels
- `/channel/{id}/programm` - return tv programm for target channel, option get params: `curdate`, `tz`, `msk` all `integer`
- `/docs/swagger` - View swagger docs

### Usage

`./epg_api --dbuser {username_here} --dbpass {passowd_here} --dbname {db_name_here} --dbhost {@} --bugsnag_key {key_here}`

### Swagger docs

- `make swagger-init`
- `make swagger-generate`
- `make swagger-${OS}`, where `$OS` one of [`windows`, `linux`]

### Tests

- Запуск тестов: `go test`

## Разворачивание не сервере

Перед разворачиванием соберите приложение локально:

```
go get && go build -a
```

Скрипт разворачивания копирует приложение из текущего каталога.

### Разворачивание впервые на свежем сервере

```
bundle exec cap STAGE systemd:go:setup
```

### Типовой деплой

```
bundle exec cap STAGE deploy
```

Где STAGE = production|reproduction

Например разворачиваем ветку master на боевом сервере:

```
BRANCH=master bundle exec cap production deploy
```

### Зайти на боевой сервер и посмотреть что там да как

```
cap production shell
tail -f log/epg_api.log
```

### Список всех команд выполняемых на сервере:


```
cap production -T
```
