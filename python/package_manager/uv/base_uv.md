# UV

## Навигация

- [Использование UV](#использование-uv)
- - [Основные команды](#основные-команды)
- - [Python versions](#python-versions)
- - [Scripts](#scripts)
- - [Projects](#projects)
- - [Tools](#tools)
- - [The pip interface](#the-pip-interface)
- - [Utility](#utility)
- [Настройки в проекте](#настройки-для-pyprojecttoml)

## Использование UV

### Основные команды

- `uv venv`: Установить окружение
- `uv pip sync pyproject.toml`: установить зависимости из pyproject.toml
- - `uv pip sync requirements.txt`
- `uv pip install package_name`: добавить зависимость
- - `uv add package_name`
- `uv pip remove package_name`: удалить зависимости
- - `uv remove package_name`
- `uv tree`: посмотреть дерево зависимостей
- `uv pip list`: глянуть установленные пакеты в окружении

***

### Python versions

Installing and managing Python itself.

- `uv python install`: Install Python versions.
- `uv python list`: View available Python versions.
- `uv python find`: Find an installed Python version.
- `uv python pin`: Pin the current project to use a specific Python version.
- `uv python uninstall`: Uninstall a Python version.

See the [guide on installing Python](https://docs.astral.sh/uv/guides/install-python/) to get started.

***

### Scripts

Executing standalone Python scripts, e.g., example.py. (Выполнение автономных сценариев Python, например, example.py.)

- `uv run`: Run a script.
- `uv add` --script: Add a dependency to a script
- `uv remove` --script: Remove a dependency from a script

See the [guide on running scripts](https://docs.astral.sh/uv/guides/scripts/) to get started.

***

### Projects

Creating and working on Python projects, i.e., with a pyproject.toml.

- `uv init`: Create a new Python project.
- `uv add`: Add a dependency to the project.
- `uv remove`: Remove a dependency from the project.
- `uv sync`: Sync the project's dependencies with the environment.
- `uv lock`: Create a lockfile for the project's dependencies.
- `uv run`: Run a command in the project environment.
- `uv tree`: View the dependency tree for the project.

See the [guide on projects](https://docs.astral.sh/uv/guides/projects/) to get started.

***

### Tools

Running and installing tools published to Python package indexes, e.g., ruff or black.

- `uvx / uv tool run`: Run a tool in a temporary environment.
- `uv tool install`: Install a tool user-wide.
- `uv tool uninstall`: Uninstall a tool.
- `uv tool list`: List installed tools.
- `uv tool update-shell`: Update the shell to include tool executables. (Обновите оболочку, чтобы включить в нее исполняемые файлы инструментов)

See the [guide on tools](https://docs.astral.sh/uv/guides/tools/) to get started.

***

### The pip interface

Ручное управление средами и пакетами - предназначено для использования в унаследованных рабочих процессах или в случаях, когда команды высокого уровня не обеспечивают достаточного контроля.

Creating virtual environments (replacing `venv` and `virtualenv`):

- `uv venv`: Create a new virtual environment.
See the documentation on [using environments](https://docs.astral.sh/uv/pip/environments/) for details.

Управление пакетами в окружении (замена `pip` и `pipdeptree`):

- `uv pip install`: Установка пакетов в текущее окружение.
- `uv pip show`: Показать подробную информацию об установленном пакете.
- `uv pip freeze`: Список установленных пакетов и их версий.
- `uv pip check`: Проверьте, есть ли в текущем окружении совместимые пакеты.
- `uv pip list`: List installed packages.
- `uv pip uninstall`: Uninstall packages.
- `uv pip tree`: Просмотрите дерево зависимостей для среды.

See the documentation on [managing packages](https://docs.astral.sh/uv/pip/packages/) for details.

Locking packages in an environment (replacing [pip-tools](https://github.com/jazzband/pip-tools)):

- `uv pip compile`: Скомпилируйте требования в файл блокировки.
- `uv pip sync`: Синхронизируйте окружение с файлом блокировки.
См. документацию по [locking environments](https://docs.astral.sh/uv/pip/compile/) for details.

### Utility

Управление и проверка состояния uv, например, кэша, каталогов хранения или выполнение самообновления:

- `uv cache clean`: Удаление записей кэша.
- `uv cache prune`: Удалите устаревшие записи кэша.
- `uv cache dir`: Показать путь к каталогу uv-кэша.
- `uv tool dir`: Показать путь к каталогу uv-инструмента.
- `uv python dir`: Показать путь к установленным в uv версиям Python.
- `uv self update`: Обновите uv до последней версии.

## Настройки для pyproject.toml

```toml
[project]
name = "webserver"
version = "0.1.0"
description = "webserver on python check"
requires-python = ">=3.12"
readme = "README.md"
dependencies = ["aiofiles>=24.1.0"]


[project.optional-dependencies]
# cli = ["ruff>=0.6.4"]

[tool.uv]
dev-dependencies = ["ruff>=0.6.5"]

[tool.uv.sources]
src = { workspace = true }

[tool.uv.workspace]
members = ["webserver/*"]
exclude = ["webserver/seeds"]

[project.scripts]


[tool.ruff]
line-length = 135
dummy-variable-rgx = "^(_+|(_+[a-zA-Z0-9_]*[a-zA-Z0-9]+?))$"
target-version = "py312"
select = ["E", "W", "F", "B", "I", "Q", "COM"]
fixable = ["I", "W", "COM", "Q"]
flake8-quotes.inline-quotes = "double"
flake8-quotes.docstring-quotes = "double"
ignore = ["D100"]
exclude = [
    "settings.py",
    "manage.py",
    ...
    ...
    "migrations",
]

[tool.ruff.per-file-ignores]
"src/settings.py" = ["E501"]

[tool.ruff.pydocstyle]
convention = "google"
```
