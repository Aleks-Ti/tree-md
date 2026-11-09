---
tags:
  - docker
up: "[[MOC Инфраструктура]]"
---

> [!info] Раздел: [[MOC Инфраструктура]]

# Docker: сборка и запуск контейнеров step by step

Пошаговый пример на проекте kittygram (бэк на Django, фронт, PostgreSQL).

## Бэкенд

Собрать образ бэка:

```bash
docker build -t kitty_backend .
```

Запустить контейнер бэка с удалением после остановки:

```bash
docker run --name kitty_backend_container --rm -p 8080:8080 kitty_backend
```

Короткий вариант:

```bash
docker run -p 8080:8080 --rm kitty_backend
```

Миграции для бэка:

```bash
docker exec kitty_backend_container python manage.py migrate
```

Зайти в контейнер в режиме командной строки, если надо:

```bash
docker exec -it trusting_hypatia bash
```

## Фронтенд

```bash
docker build -t kitty_frontend .
```

Запуск с удалением после остановки:

```bash
docker run --rm -it -p 8080:8080 --name kitty_frontend_test kitty_frontend
```

## PostgreSQL

Будет создан контейнер `kittygram`, в нём запустится сервер PostgreSQL:

```bash
docker run --name kittygram \
            --env-file .env \
            -v kitty_pg_data:/var/lib/postgresql/data \
            postgres:13.10
```

Подключиться к БД выше:

```bash
docker exec -it kittygram  psql -U kittygram_user -d kittygram
# POSTGRES_DB   POSTGRES_USER   DB_NAME=kittygram
```

Выйти из psql: `Ctrl+D`. Остановить и удалить контейнер БД:

```bash
docker container stop kittygram
docker container rm kittygram
```

## Django-приложение с env и сетью

```bash
docker run --env-file .env \
           --net django-network \
           --name kitty_backend_container \
           -p 8080:8080 kitty_backend
```

## Пример на проекте taski

```bash
# бэк
docker run --name taski_backend_container --rm -p 8000:8000 taski_backend
# фронт
docker run --rm -it -p 8000:8000 --name taski_frontend_test taski_frontend
# миграции в контейнере бэка
docker exec taski_backend_container python manage.py migrate
# volume для БД (sqlite)
docker volume create sqlite_data
```

---

## Связанное

- [[Docker dump]]
- [[Docker compose]]
- [[base_commands]]
- [[docker_example_description]]
- [[psql]]
