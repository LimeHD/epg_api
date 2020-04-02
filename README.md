# API для EPG

### Build

- Build: `go build`, опционально флаг `-o` указывает куда билдить: `go build -o ../gobuilds`
- После вы можете запустить его: `./main`

### Without build

- Run: `go run .`

#### Flags & Env easy to use

Можно комбинировать одни и те же параметы через аргументы командной строи и через переменные окружения

- in command line `go run . --host localhost --port 1337`
- in environment vars `todo`

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

