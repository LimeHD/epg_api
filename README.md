# API для EPG

### Build

- Build: `go build`, опционально флаг `-o` указывает куда билдить: `go build -o ../gobuilds`
- После вы можете запустить его: `./main`

### Without build

- Run: `go run .`

### Swagger docs

- Генерируем документацию: `swag init` - если ее нет, конечно же
- Запускаем: `swagger serve docs/swagger.json` - у вас автоматически откроется окно дефолтного бровзера

#### У меня нет swagger, что мне делать?!

Нужно сделать отдуельную доку по swagger, пока опишу тут.

1. Для начала нужно установить сам swagger, это можно сделать так: [Жмяк installation](https://goswagger.io/install.html)
2. Для windows ы можете проделать следующее:
    - Клонируем репо: git clone https://github.com/go-swagger/go-swagger
    - Переходи до директории `/cmd/swagger`
    - Делаем `go build`
    - Радуемся
    
В приципе второй пункт можно проделать для любой архитектуры

### Systemd

```

[Unit]
Description=epg_api

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/path/to/build/epg_api

[Install]
WantedBy=multi-user.target

```
