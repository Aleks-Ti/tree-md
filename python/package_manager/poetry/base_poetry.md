# Poetry

Методичка по использованию и установки poetry для операционной системы windows 10

## Навигация

- [Установка в ОС](#poetry-install-in-windows)
- [Добавление в переменные операционной системы](#install-path)
- [Установка auto окружения venv в папке проекта](#auto-создание-окружение-в-папке-проекта)

***

- [Работа с poetry](#использование-poetry)
- - [Создание нового проекта](#создание-нового-проекта)
- - [Установка/инициализация poetry для проекта](#инициализация-существующего-проекта)
- - [Update зависимостей](#update-зависимостей)

***

- [Базовые команды](#base-commands)
- - [Популярные команды](#most-popular-commands)
- - [Остальные команды](#other-commands)

***

- [Ссылка на офф доку по установке](#poetry-install-docs)

## poetry install in windows

в windows 11 PowerShell прописать команду для установки:

```rust
(Invoke-WebRequest -Uri https://install.python-poetry.org -UseBasicParsing).Content | py -
```

Теперь можно проверить poetry командой:

```yaml
poetry --version
# Poetry (version 1.7.1)
```

## install path

затем нужно => прокликать **"ПКМ"** на **<Мой компьютер>**, и выбрать свойства, в открывшемся меню выбрать => **<Дополнительные параметры системы>**.

В открывшемся меню, на вкладке <Дополнительно> выбрать подменю <Переменные среды> \

- далее в открывшемся окне в окошке <Переменные среды пользователя для **User**> \
найти **Path** и выбрать кликом **"ЛКМ"**, затем прожать меню *Изменить*
- в новом открывшемся меню <Изменить переменную среды> выбрать действие **Создать** \
туда помещаем значение пути до **poetry**

```yaml
C:\Users\User\AppData\Roaming\Python\Scripts
```

Теперь можно проверить poetry командой в терминал gitbush для проекта, например:

```yaml
poetry --version
# Poetry (version 1.7.1)
```

Если все Ок, можно идти дальше.

## Auto создание окружение в папке проекта

Для того что poetry создавал окружение в проекте, необходимо прописать в терминале папки проекта команду:

```yaml
poetry config virtualenvs.in-project true
```

## Использование poetry

### Создание нового проекта

- Вниманией! Чтобы установился нужной версии python, нужно предварительно прописать ```poetry env use python3.11```

для того чтобы создать проект с нуля, необходимо прописать команду:

```javascript
poetry new my-app
```

будет создана подобная структура проекта:

```javascript

my-app/
├── README.md
├── my_app
│   └── __init__.py
├── pyproject.toml
└── tests
    └── __init__.py
 
2 directories, 4 files
```

***
***

### Инициализация существующего проекта

Если pyproject.toml уже создан и в нем прописан конфиг на зависимости и прочее:

- Вниманией! Чтобы установился нужной версии python, нужно предварительно прописать ```poetry env use python3.11```

```yaml
poetry install
```

poetry установит все прописанные в конфиге pyproject.toml зависимости и создаст окружение venv в папке проекта.

В противном случае, будет создан/инициализирован дефолтный проект

***
***

### Update зависимостей

Для того чтобы добавить/удалить зависимости проекта в poetry, необходимо обратится к файлу pyproject.toml и в разделе **[tool.poetry.dependencies]** убрать или добавить пакет/модуль/фреймфорк.

Пример:

```toml
[tool.poetry.dependencies]
python = "3.12"
passlib = { extras = ["bcrypt"], version = "^1.7.4" }
python-jose = { extras = ["cryptography"], version = "^3.3.0" }
fastapi = "^0.109.0"
uvicorn = "^0.25.0"
...
```

Затем нужно выполнить команду в терминале:

```yaml
poetry update
```

Будет выполнена установка/удаление зависимостей.

***
***

## Base commands

### most popular commands

```go
poetry new // Создает новый проект Poetry.

poetry init // Инициализирует проект в текущем каталоге, добавляя pyproject.toml.

poetry add // Добавляет зависимость к проекту.

poetry install // Устанавливает зависимости проекта.

poetry remove // Удаляет зависимость из проекта.

poetry update // Обновляет зависимости проекта.

poetry run // Запускает скрипт из секции [tool.poetry.scripts] в pyproject.toml.

poetry shell // Активирует виртуальное окружение проекта.

poetry build // Собирает пакет проекта.

poetry publish // Публикует пакет на PyPI.
```

### other commands

Сборка проекта, зачем бы он не нужна была :/

```yaml
poetry build  # собираем как sdist, так и wheel
poetry build --format sdist  # собираем только sdist
poetry build --format wheel  # собираем только wheel
```

Для публикации на PyPI предварительно получаем API token и устанавливаем его

```yaml
poetry config pypi-token.pypi <my-token>
```

публикуем:

```yaml
poetry publish
```

## Poetry install docs

```bash
https://python-poetry.org/docs/#installing-with-the-official-installer
```
