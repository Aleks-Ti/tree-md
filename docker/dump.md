# Свалка команд и заметок

## Навигация

- [File_1](#compose-финал)
- [File_2](#docker_commandstxt)
- [File_3](#composetxt)

## compose финал

```bash
docker compose stop && docker compose up --build
```

```bash
docker compose exec backend python manage.py collectstatic
```

```bash
docker compose exec backend cp -r /app/collected_static/. /backend_static/static/ 
```

```docker compose exec backend python manage.py migrate```

```docker compose exec backend python manage.py createsuperuser```


```bash
# перезапуск
docker compose stop && docker compose up --build
# или
docker compose down
# уронить ваще эти образы и можно вместе с валуме и тп
docker compose down -v
```

## docker_commands.txt

- ```docker run --name hello-container hello-world``` - хелло ворлд / создаст мини образ из докер хаба и запустит

- ```docker run python``` - Докер-клиент отправит команду докер-демону, докер-демон скачает докер-образ python и запустит докер-контейнер на его основе.

- ```docker run --name nginx-container nginx``` - запустите контейнер из этого образа.
Остановите контейнер с Nginx командой Ctrl+C и удалите его командой docker container rm nginx-container

- Запуск нового контейнер с портами:

```bash
docker run -p 80:80 --name nginx-container nginx
```

- Остановите контейнер с Nginx, удалите его и запустите заново — но с другим портом:

```bash
docker container rm nginx-container
docker run -p 8080:80 --name nginx-container nginx
```

- Чекнуть список запущенных контейнеров:
```docker container ls```

- Чекнуть список всех запущенный контейнеров, в том числе и остановленных:
```docker container ls -a```

- Запустить/остановить контейнер:

```bash
docker container stop(start) nginx-container
# or
docker start nginx-container (старая команда до релиза 1.13)
```

- Запустить терминал внутри контейнера:

```bash
docker exec -it nginx-container bash
Покинуть контейнер: exit or ctrl + D
```

- Создать какой нить файл внутри контейнера:
```docker exec nginx-container touch /my-file```

- Подсказка для комманд КОНТЕЙНЕРОВ:
```docker container --help```

- Подсказка для комманд образов IMAGE(ОБРАЗЫ):

```bash
docker image --help
# Удалить образ:
docker image rm hello-world
```

- Запуск контейнера БЭКА:
    ```docker run --name taski_backend_container --rm -p 8000:8000 taski_backend```

- Запуск контейнера ФРОНТА:
    ```docker run --rm -it -p 8000:8000 --name taski_frontend_test taski_frontend```

- Выполнить миграции в контейнере БЭКА:
    ```docker exec taski_backend_container python manage.py migrate```

- Создание БД контейнера:
    ```docker volume create sqlite_data```

- Найти БЕЗЫМЯННЫЕ образы:
    ```docker image ls -f "dangling=true" -q```

- Снести БЕЗЫМЯННЫЕ образы:
    ```docker image rm $(docker image ls -f "dangling=true" -q)```

## compose.txt

```docker compose up``` - поднять компос файл

Остановить и запустить компост - ```docker compose stop && docker compose up```

В директории ```taski-docker/``` выполните команду ```**docker compose down**```. По этой команде будут удалены все контейнеры и
    другие связанные сущности докера — например, сети.
При этом **volumes** сохранятся, и если после выполнения этой команды вновь создать
    контейнеры — все данные на **volumes** будут доступны.

- Поднять YML для ПРОДА и из IMAGE с GitDOCKER     ->    ```docker compose -f docker-compose.production.yml up```

- собрать статику в прод -> ```docker compose -f docker-compose.production.yml exec backend python manage.py collectstatic```

```bash
docker compose -f docker-compose.production.yml exec backend cp -r /app/collected_static/. /backend_static/static/
```

МИГРАЦИИ:

```bash
docker compose -f docker-compose.production.yml exec backend python manage.py migrate
```

### УСТАНОВКА НА СЕРВ COMPOSE

```sudo apt update```

```sudo apt install curl```

```curl -fSL https://get.docker.com -o get-docker.sh```

```sudo sh ./get-docker.sh```

```sudo apt-get install docker-compose-plugin```

- Запускаем Docker Compose на сервере
Есть и другой вариант: создайте на сервере пустой файл ```docker-compose.production.yml``` и с помощью редактора ```nano``` добавьте
    в него содержимое из локального ```docker-compose.production.yml```.
Скопируйте файл ```.env``` на сервер, в директорию ```taski/```.
Когда на сервере будут запущены контейнеры, они должны будут продолжить работу и после того,
    как вы отключите терминал. У Docker Compose для этого есть специальный режим запуска, в котором он работает в фоне, в режиме «демона».
Для запуска Docker Compose в режиме демона команду ```docker compose up``` нужно запустить с флагом ```-d```. Выполните эту команду на сервере в папке ```taski/```:
    ```sudo docker compose -f docker-compose.production.yml up -d```

- Выполнять команды docker compose нужно из той директории, в которой размещён файл конфигурации.
Основные команды, которые вам понадобятся для управления:

```bash
docker compose stop # — остановит все контейнеры, но оставит сети и volume. Эта команда пригодится, чтобы перезагрузить или обновить приложения.

docker compose down # — остановит все контейнеры, удалит их, сети и анонимные volumes. Можно будет начать всё заново.
docker compose logs # — просмотр логов запущенных контейнеров.
```

- Проверьте, что все нужные контейнеры запущены:
    ```sudo docker compose -f docker-compose.production.yml ps```

```yaml
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

## step by step


- собрать image(образ) бэка:
    ```docker build -t kitty_backend .```

- запустить контэйнер бэка с удалением после остановки:

```docker run --name kitty_backend_container --rm -p 8080:8080 kitty_backend```

- а это - ```docker run -p 8080:8080 --rm kitty_backend```

- запустить миграции для бэка:
```docker exec kitty_backend_container python manage.py migrate```

- зайти в контейнер в режиме командной строки если надо:
```docker exec -it trusting_hypatia bash```

- собрать image(образ) фронта:
```docker build -t kitty_frontend .```

- запустить контейнер для фронта с удаление после остановки:
```docker run --rm -it -p 8080:8080 --name kitty_frontend_test kitty_frontend```


- Постгресс
- - По этой команде будет создан контейнер kittygram , а в нём запустится сервер PostgreSQL.

```bash
docker run --name kittygram \
            --env-file .env \
            -v kitty_pg_data:/var/lib/postgresql/data \
            postgres:13.10
```

- подключится к БД выше

```bash
docker exec -it kittygram  psql -U kittygram_user -d kittygram
# POSTGRES_DB   POSTGRES_USER   DB_NAME=kittygram
```

Выйдите из psql, нажав ```Ctrl+D```

```docker container stop kittygram```

```docker container rm kittygram```  

- Запустить контейнер с джанго приложением:

```bash
docker run --env-file .env \
           --net django-network \
           --name kitty_backend_container \
           -p 8080:8080 kitty_backend
```
