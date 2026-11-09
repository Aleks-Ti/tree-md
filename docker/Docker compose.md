---
tags:
  - docker
up: "[[MOC Инфраструктура]]"
---

> [!info] Раздел: [[MOC Инфраструктура]]

# Docker compose: команды

Команды `docker compose` нужно выполнять из директории, в которой лежит файл конфигурации.

## Основные команды

| Команда | Что делает |
|---|---|
| `docker compose up` | поднять compose-файл |
| `docker compose stop` | остановить контейнеры, сети и volumes остаются (удобно для перезапуска/обновления) |
| `docker compose down` | остановить и удалить контейнеры, сети и анонимные volumes; именованные **volumes** сохраняются |
| `docker compose down -v` | то же, но вместе с volumes |
| `docker compose logs` | логи запущенных контейнеров |
| `docker compose ps` | список контейнеров compose-проекта |

Если после `docker compose down` снова создать контейнеры, данные на **volumes** будут доступны.

## Перезапуск

```bash
docker compose stop && docker compose up --build
# или
docker compose down
# уронить образы и можно вместе с volume и т.п.
docker compose down -v
```

## Команды для backend (dev)

```bash
docker compose exec backend python manage.py collectstatic
```

```bash
docker compose exec backend cp -r /app/collected_static/. /backend_static/static/
```

```bash
docker compose exec backend python manage.py migrate
```

```bash
docker compose exec backend python manage.py createsuperuser
```

## Прод: отдельный compose-файл

Поднять YML для прода (образы берутся с Docker Hub):

```bash
docker compose -f docker-compose.production.yml up
```

Собрать статику:

```bash
docker compose -f docker-compose.production.yml exec backend python manage.py collectstatic
```

```bash
docker compose -f docker-compose.production.yml exec backend cp -r /app/collected_static/. /backend_static/static/
```

Миграции:

```bash
docker compose -f docker-compose.production.yml exec backend python manage.py migrate
```

---

## Связанное

- [[Docker dump]]
- [[Docker установка на сервер]]
- [[base_commands]]
- [[docker_example_description]]
- [[YAML base]]
