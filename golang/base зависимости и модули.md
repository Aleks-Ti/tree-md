# Модули и зависимости

## Что такое модули

В Go модулем принято называть любое приложение, которое можно опубликовать, версионировать, импортировать или скачать. С помощью модулей мы можем управлять зависимостями.

Добавим в пример с Hello, Hexlet пакет для логгирования. В коде это будет выглядеть так:

```golang
package main

import "github.com/sirupsen/logrus" // Указываем путь до нужного пакета внутри репозитория

func main() {
    logrus.Println("Hello, Hexlet!")
}
```

Но чтобы это код запустился, надо создать модуль.

## Как создать модуль

Чтобы превратить папку с кодом в Go-модуль, можно использовать команду go mod init:

```bash
go mod init github.com/hexlet/hello-hexlet  # Сразу указываем имя модуля

go mod init github.com/hexlet/hello-hexlet
go: creating new go.mod: module github.com/hexlet/hello-hexlet
go: to add module requirements and sums:
        go mod tidy
```

Команда сгенерировала go.mod файл со следующим содержимым:

```bash
module github.com/hexlet/hello-hexlet

go 1.23.0
```

Команда go mod tidy проверяет импорты в коде, загружает недостающие зависимости и удаляет лишние. Файл go.mod обновился и теперь включает в себя раздел с зависимостями:

```bash
module github.com/hexlet/hello-hexlet

go 1.23.0

require github.com/sirupsen/logrus v1.9.3

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
```

Также появился файл go.sum: