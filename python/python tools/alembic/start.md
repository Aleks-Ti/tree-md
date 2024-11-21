# Первичная настройка алембик

## Пулл команд

- инциализировать асинхронный `alembic` в проекте:

```bash
alembic init --template async migrations
```

- инициализировать синхронный `alembic` в проекте:

```bash
alembic init migrations
```

## Настройки env.py

- Предварительные настройки для моделей в настройках - нужны будут далее для настройки `alembic`:

```python
# src >> common >> base.py

from sqlalchemy import MetaData
from sqlalchemy.ext.declarative import as_declarative
from sqlalchemy.orm import declared_attr

metadata = MetaData()


@as_declarative(metadata=metadata)
class Base:
    @classmethod
    @declared_attr
    def __tablename__(cls):
        return cls.__name__.lower()

    __allow_unmapped__ = False
```

- прокинуть настройки для `alembic`

```python
# src >> migrations >> env.py
...
from src.core.base import metadata
from src.core.connector_for_alembic_and_alchemy import DataBaseConfig
from src.user.models import (   # noqa
    User,
)
...

...
### MY SETTINGS # noqa
DATABASE_URL = DataBaseConfig().build_connection_str()
target_metadata = metadata
config.set_main_option("sqlalchemy.url", DATABASE_URL)
### END MY SETTINGS # noqa
...
```
