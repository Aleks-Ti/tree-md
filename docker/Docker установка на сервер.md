---
tags:
  - docker
up: "[[MOC Инфраструктура]]"
---

> [!info] Раздел: [[MOC Инфраструктура]]

# Docker и Compose: установка на сервер

## Установка

```bash
sudo apt update
sudo apt install curl
curl -fSL https://get.docker.com -o get-docker.sh
sudo sh ./get-docker.sh
sudo apt-get install docker-compose-plugin
```

## Запуск Compose на сервере

Создайте на сервере пустой файл `docker-compose.production.yml` и с помощью `nano` добавьте в него содержимое из локального `docker-compose.production.yml` (либо скопируйте файл). Файл `.env` скопируйте на сервер в директорию проекта (`taski/`).

Чтобы контейнеры продолжили работу после отключения терминала, Docker Compose запускают в фоне («режим демона») — с флагом `-d`. Выполнять в папке проекта:

```bash
sudo docker compose -f docker-compose.production.yml up -d
```

Проверить, что все нужные контейнеры запущены:

```bash
sudo docker compose -f docker-compose.production.yml ps
```

Управление (из директории с файлом конфигурации):

```bash
docker compose stop # — остановит все контейнеры, но оставит сети и volume. Пригодится, чтобы перезагрузить или обновить приложения.

docker compose down # — остановит все контейнеры, удалит их, сети и анонимные volumes. Можно начать всё заново.
docker compose logs # — просмотр логов запущенных контейнеров.
```

## Nginx на сервере

Пример конфигурации для проксирования на контейнеры (проект kittygram):

```nginx
location /api/ {
    proxy_pass http://127.0.0.1:8080;
    client_max_body_size 20M;
}
location /admin/ {
    proxy_pass http://127.0.0.1:8080;
    client_max_body_size 20M;
}

location / {
    root   /var/www/kittygram;
    index  index.html index.htm;
    try_files $uri /index.html;
}
```

---

## Связанное

- [[Docker dump]]
- [[Docker compose]]
- [[base_commands]]
- [[Nginx base]]
