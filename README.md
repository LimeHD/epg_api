# API для EPG

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

### Tests

- Запуск тестов: `go test`
