---
tags:
  - python
  - fastapi
  - draft
up: "[[MOC Python]]"
---

> [!info] Раздел: [[MOC Python]]

# FastAPI: быстрый старт

## Установка и запуск

```bash
uv add "fastapi[standard]"
uv run fastapi dev main.py        # dev, автоперезагрузка
uv run uvicorn main:app --host 0.0.0.0 --port 8000   # прод-подобно
```

`main.py`:

```python
from fastapi import FastAPI

app = FastAPI()


@app.get("/items/{item_id}")
async def read_item(item_id: int, q: str | None = None):
    return {"item_id": item_id, "q": q}
```

Автодокументация: `/docs` (Swagger UI) и `/redoc`.

## Что где

| Задача | Инструмент |
| --- | --- |
| Валидация тела запроса | модели `pydantic.BaseModel` |
| Зависимости, auth, сессия БД | `Depends(...)` |
| Миграции | [[Alembic start]] |
| Пакеты | [[base_uv]] |
| Деплой за прокси | [[Nginx base]], [[step by step request in API back & front]] |

> [!note] Заготовка
> Расширь своими примерами (роутеры, `Depends`, тесты через `TestClient`).

---

## Связанное

- [[example_best_practice]]
- [[Alembic start]]
- [[http protocol]]
