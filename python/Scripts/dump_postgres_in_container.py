import logging
import os
import subprocess
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
    """Загрузка переменных из .env файла и добавление их в окружение."""

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

    os.makedirs(path, exist_ok=True)


def create_dumps(path_dir_dumps: str) -> None:
    """Создание дампа."""

    dump_name = f'dump_{datetime.now().strftime("%d-%m-%Y_%H:%M:%S")}.sql'

    runner_command(f"docker exec {CONTAINER_NAME} pg_dump -U {PG_USER} -d {PG_DB} > {str(os.path.join(path_dir_dumps, dump_name))}")
    logger.info(f"Успешно создан новый дамп: {dump_name}")


def control_count_dumps(path_dir_dumps: str, num_to_delete=5) -> None:
    files = [f for f in os.listdir(path_dir_dumps) if os.path.isfile(os.path.join(path_dir_dumps, f))]
    if len(files) <= num_to_delete:
        return None

    files.sort(key=lambda x: os.path.getctime(os.path.join(path_dir_dumps, x)))
    for file in files[:-num_to_delete]:
        file_path = os.path.join(path_dir_dumps, file)
        os.remove(file_path)
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

    CONTAINER_NAME = get_variable_env("CONTAINER_NAME")
    PG_USER = get_variable_env("PG_USERNAME")
    PG_DB = get_variable_env("PG_DB_NAME")
    DIR_DUMPS = get_variable_env("DIR_DUMPS")

    try:
        main()
        logger.info("Процесс завершен.")
    except Exception:
        logger.exception("Ошибка сценария.")
