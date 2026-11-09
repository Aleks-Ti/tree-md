---
tags:
  - docker
  - compose
  - postgres
  - template
  - draft
up: "[[MOC Инфраструктура]]"
---

> [!info] Раздел: [[MOC Инфраструктура]]

# Шаблон compose: приложение + Postgres

Заготовка для нового проекта. Команды compose — в [[Docker compose]], синтаксис YAML — в [[YAML base]].

## `compose.yaml`

```yaml
services:
  db:
    image: postgres:16
    restart: unless-stopped
    env_file: .env
    environment:
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_DB: ${POSTGRES_DB}
    volumes:
      - pg_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U $${POSTGRES_USER} -d $${POSTGRES_DB}"]
      interval: 5s
      timeout: 3s
      retries: 10
    # ports: ["127.0.0.1:5432:5432"]   # только для локальной отладки, см. ниже

  app:
    build: .
    restart: unless-stopped
    env_file: .env
    environment:
      DATABASE_URL: postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}
    depends_on:
      db:
        condition: service_healthy
    ports:
      - "8000:8000"

volumes:
  pg_data:
```

## `.env` (не коммитить!)

```bash
POSTGRES_USER=app
POSTGRES_PASSWORD=change-me
POSTGRES_DB=app
```

Рядом положи `.env.example` с теми же ключами и пустыми значениями. Подробнее — [[Переменные окружения и секреты]].

## Ключевые моменты

| Что | Зачем |
| --- | --- |
| хост `db` в `DATABASE_URL` | внутри сети compose сервисы видят друг друга по имени, не `localhost` |
| `healthcheck` + `depends_on: service_healthy` | приложение стартует, когда БД реально готова |
| `volumes: pg_data` | данные переживают `docker compose down` (но не `down -v`) |
| `$$` в healthcheck | экранирование, чтобы переменную раскрыл контейнер, а не compose |
| порт `5432` наружу | только на `127.0.0.1` и только для отладки — иначе см. [[Порты, и действия с ними]] |

## Запуск и работа с БД

```bash
docker compose up -d --build
docker compose logs -f app
docker compose exec db psql -U app -d app
docker compose down                            # данные остаются
docker compose down -v                         # ⚠️ стирает volume
```

Миграции: `docker compose exec app alembic upgrade head` — см. [[Alembic start]]. Бэкап: [[practice_bash_sript_create_dump]].

---

## Связанное

- [[Docker compose]]
- [[docker_example_description]]
- [[psql]]
- [[Alembic start]]
- [[Переменные окружения и секреты]]
- [[YAML base]]
