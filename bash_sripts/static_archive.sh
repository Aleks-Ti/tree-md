#!/bin/bash

# Путь к директории со статическими файлами
STATIC_DIR="./static"
# Путь к директории, где будут храниться архивы
BACKUP_DIR="./backups_static"
# Имя архива
ARCHIVE_NAME="static-$(date +%Y-%m-%d_%H-%M-%S).tar.gz"
# Максимальное количество архивов
MAX_ARCHIVES=3

# Создаем архив
tar -czf "$BACKUP_DIR/$ARCHIVE_NAME" -C "$STATIC_DIR" .

# Удаляем старые архивы, если их количество превышает MAX_ARCHIVES
find "$BACKUP_DIR" -name 'static-*.tar.gz' -type f | sort | head -n -"$MAX_ARCHIVES" | xargs rm -f

echo "Архивация статики завершена: $BACKUP_DIR/$ARCHIVE_NAME"
