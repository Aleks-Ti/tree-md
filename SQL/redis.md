---
tags:
  - redis
  - key-value
  - cheatsheet
  - draft
up: "[[MOC Базы данных]]"
---

> [!info] Раздел: [[MOC Базы данных]]

# Redis

In-memory key-value хранилище, см. тип «Key-Value» в [[Виды БД]]. Типично: кэш, сессии, счётчики, очереди, rate limit.

## Запуск и консоль

```bash
docker run -d --name redis -p 6379:6379 redis:7
docker exec -it redis redis-cli
```

## Структуры данных

| Тип | Команды | Пример использования |
| --- | --- | --- |
| String | `SET k v`, `GET k`, `INCR k`, `SET k v EX 60` | кэш, счётчик, TTL |
| Hash | `HSET user:1 name Bob`, `HGETALL user:1` | объект |
| List | `LPUSH`, `RPOP`, `LRANGE q 0 -1` | простая очередь |
| Set | `SADD`, `SMEMBERS`, `SISMEMBER` | уникальные значения |
| Sorted set | `ZADD board 10 bob`, `ZRANGE board 0 -1 WITHSCORES` | рейтинги |

## Полезное

```text
TTL k          # сколько секунд жить
EXPIRE k 30    # задать TTL
DEL k
KEYS pattern   # только для отладки! на проде — SCAN
FLUSHALL       # стереть всё (осторожно)
```

## Persistence

- **RDB** — снимки по расписанию;
- **AOF** — журнал операций (надёжнее, больше места).

> [!warning] Не публикуй порт 6379 наружу без пароля — смотри [[Порты, и действия с ними]].

---

## Связанное

- [[Виды БД]]
- [[psql]]
- [[base_commands]]
