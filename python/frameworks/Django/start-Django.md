# Первичные настройки Django

## Навигация

- [Начало](#стартовые-команды)
- - [Установка](#установка)
- - [Инициализация проекта](#развернуть-приложение)
- [Установка](#установка)
- [Настройка проекта](#изменение-дефолтных-настроек)
- - [Подмена ДБ](#замена-db)
- [Добавить новое приложение в проект](#новое-приложение-в-проекте)
- [Возможная проблема с импортами приложений](#не-удается-разрешить-импорты-приложений-в-ide)

## Стартовые команды

В этом блоке заметки собраны команды для старта проекта и созданий приложений.

- Коротко, базовые команды `python/app manage.py` (покажет все команды):

```bash
# создать структуру проекта
django-admin startproject app

# запуск проекта
python app/manage.py runserver
# or
python app/manage.py runserver 127.0.0.1:8001

# создать файлы миграции
python app/manage.py makemigrations

# применить миграции к БД
python app/manage.py migrate

# создать новое приложение в проекте
cd app
python ../app/manage.py startapp "your_name_app"

# создать супер юзера
python app/manage.py createsuperuser
```

### Установка

- `Ставим джангу`, либо в простой venv ```python -m venv venv``` либо poetry ```poetry install``` - в этом случае по дефолту папка создасться папка `.venv`

- Теперь нужно активировать окружение

```bash
. venv/Scripts/activate  # for unix -> . venv/bin/activate
# or
source venv/Scripts/activate  # # for unix -> source venv/bin/activate
# для poetry
poetry shell
```

- После активации, установим Django

```bash
pip install Django
```

для `poetry`:

```bash
poetry add Django
```

- В дальнешем по заметке, все команду будут приводится исключительно для `poetry`

### Развернуть приложение

- Чтобы развернуть дефолтное приложение, введем команду:

```bash
djngo-admin startproject 'your name app'
```

## Изменение дефолтных настроек

- В проект необходимо как правило поставить `python-dotenv` для подгрузки в проект переменных из `.env`

```bash
poetry add python-dotenv
```

Переходим в `app/app/settings.py` и добавляем в начало файлы к импортам

```python
...
from dotenv import load_dotenv
...

load_dotenv()  # чтобы переменные из env подгрузились для файла настроек
```

- Теперь можно добавить настроек

```python
###
#####
#######
#########
################################################################
#                {block developer settings}:

SECRET_KEY = os.getenv("SECRET_KEY")

DEBUG = True

ALLOWED_HOSTS = [
    "localhost",
    "127.0.0.1",
    "[::1]",
    "testserver",
]

AUTH_USER_MODEL = "users.User"  # нужно указывать, если уже создана своя модель User, иначе нужно удалить или закомментировать до появления модели.

MEDIA_URL = "/media/"

DEBUG = True if os.getenv("DEBUG", "False") == "True" else False

STATICFILES_DIRS = [os.path.join(BASE_DIR, "static")]

MEDIA_ROOT = os.path.join(BASE_DIR, "media")

DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"


#              {endblock developer settings}
################################################################
##########
########
######
###
```

- Так же можно поменять локализацию проекта

```bash
...
LANGUAGE_CODE = "ru"
...
```

### Замена DB

- Поменять на Postgres

```python
###
#####
#######
#########
################################################################
#                {DATABASES CONNECTION SETTINGS}

DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": os.getenv("PG_NAME_DB", "django"),
        "USER": os.getenv("PG_USERNAME", "django"),
        "PASSWORD": os.getenv("PG_PASSWORD", ""),
        "HOST": os.getenv("PG_HOST", "localhost"),
        "PORT": os.getenv("PG_PORT", 5432),
    }
}

#                {DATABASES CONNECTION SETTINGS}
################################################################
##########
########
######
###
```

## Новое приложение в проекте

- Чтобы добавить новое приложение, например `user` в проект, нужно выполнить команду:

```bash
python path/to/manage.py startapp "your app name"
# real example
py app/manage.py startapp user
```

- После того приложение будет добавлено, необходимо его добавить в `settings`

```python
...
INSTALLED_APPS = [
    "django.contrib.admin",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "django.contrib.staticfiles",
    "user",
]
...
```

## Не удается разрешить импорты приложений в IDE

- Поскольку это не касается на работоспособность приложения, а только на линтеры и отображение непосредственное, можно это по фиксить через настройки `IDE Vs Code`

Создаем папку если ее еще нет `.vscode`, а в папке файл `settings.json`, внутри прописываем настройки

```json
{
    "python.analysis.extraPaths": [
        "./app"
    ]
}
```
