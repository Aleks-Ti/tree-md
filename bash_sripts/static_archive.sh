#!/bin/bash
echo "Архивация статики."
echo "Начало работы скрипта: $(date +%Y-%m-%d_%H-%M-%S)"

# Путь к директории со статическими файлами
STATIC_DIR="path/to_dir/static"  # Указывать путь нeобходимо абсолютный! Узнать - pwd - из которой забираем статику.
# Путь к директории, где будут храниться архивы
BACKUP_DIR="path/to_dir/backups_static"  # Указывать путь нeобходимо абсолютный! Узнать - pwd - куда будем складировать архивы статики.
# Имя архива
ARCHIVE_NAME="static-$(date +%Y-%m-%d_%H-%M-%S).tar.gz"
# Максимальное количество архивов
MAX_ARCHIVES=3

# Создаем архив
tar -czf "$BACKUP_DIR/$ARCHIVE_NAME" -C "$STATIC_DIR" .

# Удаляем старые архивы, если их количество превышает MAX_ARCHIVES
find "$BACKUP_DIR" -name 'static-*.tar.gz' -type f | sort | head -n -"$MAX_ARCHIVES" | xargs rm -f

echo "Архивация статики завершена: $BACKUP_DIR/$ARCHIVE_NAME"
echo "Время окончания архивации: $(date +%Y-%m-%d_%H-%M-%S)"


# дать разрешение на выполнение скрипта
# chmod +x path/to_dir/static_archive.sh - путь обязательно относительный, через pwd.
# crontav -e чтобы добавить команду запуска скрипта
# 2 ночи каждый день
# 0 2 * * * path/to_dir/static_archive.sh >> /tmp/cron.log 2>&1
                                # куда пишутся логи >> /tmp/cron.log 2>&1
