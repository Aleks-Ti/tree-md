#!/bin/bash
echo "Создание дампа. Начала работы."

# Загрузка переменных из файла .env
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Путь к папке с дампами относительно скрипта
DUMPS_DIR="./dumps"
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

echo "Создание дампа завершено. Работы окончены."
