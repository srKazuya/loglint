# LogLint

`loglint` — standalone-линтер для анализа Go-файлов.
Проект также поддерживает интеграцию в виде плагина для golangci-lint.

## Требования

* Go 1.25
* Task ([https://taskfile.dev](https://taskfile.dev/docs/installation)) 
    apt
    If you Set up the repository by running:
    ```bash
        curl -1sLf 'https://dl.cloudsmith.io/public/task/task/setup.deb.sh' | sudo -E bash
    ```
    Then you can install Task with:
    
    ```bash
        apt install task
    ```
    npm
    ```bash
        npm install -g @go-task/cli
    ```
    Chocolatey 
     ```bash
        choco install go-task
    ```

---
## Быстрый старт
Билд standalone + линт всех файлов testdata 
```bash
task lint-with-standalone 
```
Линт с автоисправлением 
```bash
task lint-with-standalone-fix
```
Установка  golangci-lint + линт файлов через плагин | 
```bash
task lint-with-plugin
```
---
## Основные команды: 
Taskfile.yml

| Команда                       | Описание                    |
| ----------------------------- | -----------------------     |
| task build                    | Сборка бинарника            |
| task test                     | Запуск тестов               |
| task run                      | Сборка + тесты + запуск     |
| task lint-with-standalone     | Линт всех файлов standalone |
| task lint-with-standalone-fix | Линт с автоисправлением     |
| task install-golangci-lint    | Установка golangci-lint     |
| task make-plugin              | Сборка плагина              |
| task lint-with-plugin         | Линт через плагин           | 
---

## Сборка
Сборка через Task:
```bash
task build
```
Бинарный файл будет создан в каталоге `bin/`:
* Windows: `bin/loglint.exe`
* Linux/macOS: `bin/loglint`

Сборка вручную:
```bash
go build -o bin/loglint cmd/loglint/main.go
```
---
## Тесты
Запуск тестов:

```bash
task test
```
Или:
```bash
go test -v ./...
```
---
## Запуск standalone-линтера
```bash
task lint-with-standalone
```

Команда:

* собирает проект
* запускает тесты
* запускает линтер на все файлы

---

## Примеры использования

### Проверка одного файла

```bash
./bin/loglint ./testdata/sample.go
```

### Проверка всех файлов в testdata

```bash
task lint-with-standalone
```

Эквивалент:
```bash
for f in ./testdata/**/*.go; do ./bin/loglint "$f" || true; done
```
### Проверка с автоисправлением

```bash
task lint-with-standalone-fix
```
Или:
```bash
./bin/loglint --fix ./testdata/sample.go
```

---
## Использование как плагина для golangci-lint
### Установка golangci-lint

```bash
task install-golangci-lint
```
Используемая версия:
```
v2.1.5
```
Бинарник будет установлен в каталог `bin/`.
---
### Сборка плагина
```bash
task make-plugin
```
Будет создан кастомный бинарник `custom-gcl.exe`.
---

### Запуск линтинга через плагин

```bash
task lint-with-plugin
```
Или:
```bash
./custom-gcl run ./...
```



