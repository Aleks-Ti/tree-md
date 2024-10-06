"""
Скрипт для создания дампа базы данных postgres в контейнере через docker.

В скрипт нужно обязательно добавить для константы PATH_TO_ENV ваш актуальный value с path до файла .env:

    if __name__ == "__main__":
        ...
        PATH_TO_ENV="path/to/env/file/.env"
        ...


Так же нужно будет наполнить .env соответствующими данными и указать их название для констант:

    CONTAINER_NAME = get_variable_env("YOUR_CONTAINER_NAME")  >> в файле .env должно быть так >> YOUR_CONTAINER_NAME=my_name_container_db_postgres
    PG_USER = get_variable_env("YOUR_PG_USERNAME")  >> в файле .env должно быть так >> YOUR_PG_USERNAME=my_value_for_pg_uername
    PG_DB = get_variable_env("YOUR_PG_DB_NAME")  >> в файле .env должно быть так >> YOUR_PG_DB_NAME=my_value_pg_db_name
    DIR_DUMPS = get_variable_env("YOUR_DIR_DUMPS")  >> в файле .env должно быть так >> YOUR_DIR_DUMPS=/my/path/to/dir/dumps/ >> куда будут складываться дампы


Пример использования скрипта в системе Ubuntu через crontab:

bash$ crontab -e
    ...
    0 2 * * * python3 /user/backups/script_dumps.py >> /user/backups/cron.log 2>&1
    ...
    каждый день в 2 ночи + запись в файл cron.log вывода скрипта.
"""  # noqa

import logging
import os
import subprocess

from pathlib import Path
from datetime import datetime


class NoEnvData(Exception):
    pass


def config_logging(level):
    logging.basicConfig(
        level=level,
        datefmt="%Y-%m-%d %H:%M:%S",
        format="[%(asctime)s.%(msecs)03d] %(module)s:%(lineno)-4s %(levelname)-7s - %(message)s",
    )


config_logging(logging.INFO)
logger = logging.getLogger(__name__)


def runner_command(command):
    result = subprocess.run(command, shell=True, text=True, capture_output=True)
    if result.returncode != 0:
        logger.error(f"Ошибка при выполнении: {command}")
        raise Exception(f"Команда завершилась с ошибкой: {result.stderr}")


def load_env_file(env_path):
    """Загрузка переменных из .env файла и добавление их в окружение скрипта."""

    if not os.path.exists(env_path):
        raise FileNotFoundError(
            "Файл .env найден. Убедитесь что вы указали корректный путь до файла. "
            'Ожидается абсолютный путь. path_to_env="correct/path/to/.env"'
        )

    keys = []
    with open(env_path) as env_file:
        for line in env_file:
            line = line.strip()
            if not line or line.startswith("#"):
                continue

            key, value = line.split("=", 1)
            os.environ[key] = value
            keys.append(key)

    keys = " ".join(keys)
    logger.info(f"Переменные окружения из .env успешно загружены:\n{keys}")


def dir_exists_or_create(path):
    """Создает конечную папку под дампы, если её ещё не существует."""

    Path(path).mkdir(parents=True, exist_ok=True)


def create_dumps(path_dir_dumps: str) -> None:
    """Создание дампа средствами docker и pg_dump."""

    dump_name = f'dump_{datetime.now().strftime("%d-%m-%Y_%H-%M-%S")}.sql'
    dump_path = Path(path_dir_dumps) / dump_name
    runner_command(
        ["docker", "exec", CONTAINER_NAME, "pg_dump", "-U", PG_USER, "-d", PG_DB, "f", dump_path]
    )
    logger.info(f"Успешно создан новый дамп: {dump_name}")


def control_count_dumps(path_dir_dumps: str, num_to_delete=5) -> None:
    files = sorted(
        Path(path_dir_dumps).glob("*.sql"),
        key=lambda x: x.stat().st_ctime
    )
    if len(files) <= num_to_delete:
        return None

    for file in files[:-num_to_delete]:
        file.unlink()
        logger.info(f"Удален старый дамп >> файл: {file}")


def main():
    dir_exists_or_create(DIR_DUMPS)
    create_dumps(DIR_DUMPS)
    control_count_dumps(DIR_DUMPS)


def get_variable_env(var: str):
    value_var = os.getenv(var)
    if value_var is None:
        raise NoEnvData(f"No .env data - needs variable {var}")
    return value_var


if __name__ == "__main__":
    logger.info("Старт сценария >> созданиe дампа postgres.")

    PATH_TO_ENV = "/path/to/env/file/.env"
    load_env_file(PATH_TO_ENV)

    CONTAINER_NAME = get_variable_env("YOUR_CONTAINER_NAME")  # Имя контейнера с Postgres
    PG_USER = get_variable_env("YOUR_PG_USERNAME")            # Имя пользователя Postgres
    PG_DB = get_variable_env("YOUR_PG_DB_NAME")               # Имя базы данных Postgres
    DIR_DUMPS = get_variable_env("YOUR_DIR_DUMPS")            # Путь для хранения дампов

    try:
        main()
        logger.info("Процесс завершен.")
    except Exception:
        logger.exception("Ошибка сценария.")
