# Работа с backend проекта

Примечание.

Первичные шаги по использованию проекта, описаны для понимающий людей в разработке бэкенда. В случае если необходимых знаний не хватает, попробуйте обратиться к старшему коллеге из команды разработки backend.

## Содержание

- [Первичные настройки: окружение/зависимости](#окружениезависимости)
- [Инициализация данных перед запуском контейнеров](#первоначальная-настройка-docker-compose)
- [Cборка контейнеров](#запуск-и-сборка-сети-контейнеров-docker)
- [Проброс порта для работы с Postgres в контейнере](#проброска-порта-для-работы-с-бд)
- [Доступ к проекту](#доступ-к-проекту)

***

### Окружение/зависимости

В проекте используется Python версии 3.11

В корне проекта **cd backend/** выполнить команды:

```bash
python -3.11 -m venv venv
source venv/Scripts/activate  #*unix - source venv/bin/activate
# (venv)...
# or
. venv/Scripts/activate
# (venv)...
```

Установка зависимостей:

```bash
pip install -r morec/requirements.txt
# or
cd morec/
pip install -r requirements.txt
```

***

### Первоначальная настройка docker compose

Запустить приложение docker desctop, или поднять машину через терминал.

```bash
cd devops
touch devops.env # поместить в файл devops.env данные из .env.example в корне проекта
```

***

### Запуск и сборка сети контейнеров docker

```bash
cd devops
docker compose -f docker-compose.local.yml down && docker compose -f docker-compose.local.yml up --build # команда подразумевает непрерывную работу с изменениями и постоянным сбросом контейнеров и пересборку. Если запуск осуществляется впервые, можно(не обязательно) отбросить первую часть команды до амперсандов(включительно &&(объеденяющая команда))
docker compose -f debug_in_container.yml up --build  # режим дэбага в контейнере(только если настроен собственный файл.yml под дэбаг)
docker compose -f debug_in_container.yml down
docker compose -f debug_in_container.yml down && docker compose -f debug_in_container.yml up --build
```

***

### Проброска порта для работы с бд

В файле docker-compose.local.yml в разделе postgres_db сделать изменения.

```bash
# ВНИАНИЕ! Эти настройки пишет только для себя, и перед пулл реквестом убираем, либо используем только в свое ветке/мастерской.
# Не несем в прод!
  postgres_db:
    image: "postgres:13.4-alpine"
    env_file: devops.env
    volumes:
      - postgres_data:/var/lib/postgresql/data/
    ports:
      - 5440:5432 # внутренний порт оставляем стандарт, а внешний, на свое усмотрение 5440 или другой.
```

Далее в pg admin или dbeaver создаем подключение к БД в котейнере:

```bash
# данные из .env для подключения postgres под управлением приложения.
Host = localhost
port = 5440  # или другой порта, в зависимости от шага выше
name_db = main_db
user = postgres
password = postgres
```

***

### Доступ к проекту

При запущенной сети контейнеров, нужно будет зайти в контейне, применить миграции, и создать себе суперпользователя, также подгрузить статику.

Вариант 1

```bash
# узнать имя контэйнера бэкенда
docker container ls
# пример:
    # devops-backend-1

# зайти в контейнер:
docker exec -it devops-backend-1 bash
# появится режим командной строки внутри контейнера:
# пример:
    # root@0450d15a0dd3:/app#

# Сначала применить миграции:
python manage.py migrate
# Затем создать суперюзера:
python manage.py createsuperuser
# Напоследок подгрузить статику:
python manage.py collectstatic
# 163 static files copied to '/app/collected_static'

# Выйти из режима контейнера
ctrl + D
```

Вариант 2

В этом случае нужно выполнить все те же команды выше начиная с миграции, но зайти в Docker Desktop, выбрать контейнер бэкенда из списка запущенный контейнеров, и перейти в раздел Terminal, далее накатить миграции, создать суперюзера, и выгурзить статику.

Вариант 3 (Pro)

```bash
cd devops
docker compose -f docker-compose.local.yml exec backend python manage.py migrate
docker compose -f docker-compose.local.yml exec backend python manage.py createsuperuser
docker compose -f docker-compose.local.yml exec backend python manage.py collectstatic
```

Сайт в полной мере доступен для разработки: http://localhost:80/admin/

***
...
