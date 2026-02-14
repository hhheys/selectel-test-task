# Golang Log Linter for golangci-lint

## Тестовое задание: Backend-разработка. Golang

Линтер для проверки лог-записей в Go-коде, совместимый с **golangci-lint**, который анализирует вызовы логгеров и проверяет их соответствие установленным правилам.

**Поддерживаемые логгеры:**  
- `log/slog`  
- `go.uber.org/zap`  

**Что было сделано**
- Базовые анализаторы
- Поддержка конфиг файла
- Билд плагина для golangci-lint
- CI в github actions

---

## Установка

Соберите плагин:

```bash
go build -buildmode=plugin -o build/myplugin.so ./plugin
```


## Использование с golangci-lint

Добавьте плагин в .golangci.yml:
```
linters:
  enable:
    - myplugin

plugins:
  - path: ./build/myplugin.so
```
