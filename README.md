# API для EPG

### Build

- Build: `go build`, опционально флаг `-o` указывает куда билдить: `go build -o ../gobuilds`
- После вы можете запустить его: `./main`

### Without build

- Run: `go run .`

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

### Tests

- Запуск тестов: `go test`

