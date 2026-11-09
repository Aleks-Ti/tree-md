#!/bin/bash
echo "Создание дампа. Начала работы."

# Загрузка переменных из файла .env
if [ -f /path/to_env/.env ]; then  # Указывать путь наобходимо абсолютный! Узнать - pwd - в папке с файлом
    export $(grep -v '^#' /path/to_env/.env | xargs)  # Указывать путь наобходимо абсолютный! Узнать - pwd - в папке с файлом
fi  # Достает из env переменные PG_USER, DATABASE_NAME и CONTAINER_NAME с помощью grep и xargs - убирает комментарии с начала строки или конца.

# Путь к папке с дампами относительно скрипта
DUMPS_DIR="/path/to_dir/dumps"  # Указывать путь наобходимо абсолютный! Узнать - pwd - в папке в которую предполагается складировать дампы.
# Максимальное количество дампов
MAX_DUMPS=3

# Создать папку для дампов, если она не существует
mkdir -p $DUMPS_DIR

# Сформировать имя файла дампа
DUMP_FILE="$DUMPS_DIR/dump_$(date +%Y-%m-%d_%H-%M-%S).sql"

# Создать дамп базы данных
docker exec $CONTAINER_NAME pg_dump -U $PG_USER $DATABASE_NAME > $DUMP_FILE

# Удалить старые дампы, если их количество превышает MAX_DUMPS
DUMPS_COUNT=$(ls -1 $DUMPS_DIR/*.sql 2>/dev/null | wc -l)
if [ $DUMPS_COUNT -gt $MAX_DUMPS ]; then
  # Находим самые старые дампы и удаляем их
  ls -1tr $DUMPS_DIR/*.sql | head -n -$MAX_DUMPS | xargs rm -f
fi

echo "Создание дампа завершено. Работы окончены: $(date +%Y-%m-%d_%H-%M-%S)"

# дать разрешение на выполнение скрипта
# chmod +x path/to_dir/dump.sh - путь обязательно относительный, через pwd.
# crontav -e чтобы добавить команду запуска скрипта
# 2 ночи каждый день
# 0 2 * * * path/to_dir/dump.sh >> /tmp/cron.log 2>&1
                                # куда пишутся логи >> /tmp/cron.log 2>&1
