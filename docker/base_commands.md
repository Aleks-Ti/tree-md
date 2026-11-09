---
tags:
  - docker
up: "[[MOC Инфраструктура]]"
---

> [!info] Раздел: [[MOC Инфраструктура]]

# Базовые команды докер

## Навигация

- [(None) образы](#чистка-образов-none)

### Чистка образов (None)

Увидеть список таких `безымянных образов` (они называются `dangling`, англ. «болтающиеся») можно так:

```bash
docker image ls -f "dangling=true" -q
```

Почистить компьютер от таких образов можно такой командой:

```bash
docker image rm $(docker image ls -f "dangling=true" -q)
```

Эта команда удаляет все безымянные образы, из которых не запущены контейнеры: сначала она найдёт все такие контейнеры с помощью docker image, а полученный результат передаст docker image rm.

### Загрузить образы на Docker Hub

```bash
docker push username/project_frontend
docker push username/project_backend
docker push username/project_gateway
```

## Запуск и управление контейнерами

- Hello world (создаст мини-образ из Docker Hub и запустит):

```bash
docker run --name hello-container hello-world
```

- `docker run python` — клиент отправит команду демону, демон скачает образ `python` и запустит контейнер на его основе.

- Запуск Nginx; остановить по `Ctrl+C`, удалить командой `docker container rm`:

```bash
docker run --name nginx-container nginx
docker container rm nginx-container
```

- Запуск с пробросом портов (хост:контейнер), затем пересоздание с другим портом:

```bash
docker run -p 80:80 --name nginx-container nginx
docker container rm nginx-container
docker run -p 8080:80 --name nginx-container nginx
```

- Список запущенных контейнеров / всех, включая остановленные:

```bash
docker container ls
docker container ls -a
```

- Запустить / остановить контейнер:

```bash
docker container stop nginx-container
docker container start nginx-container
# старая команда (до релиза 1.13)
docker start nginx-container
```

- Терминал внутри контейнера (выйти: `exit` или `Ctrl+D`):

```bash
docker exec -it nginx-container bash
```

- Создать файл внутри контейнера:

```bash
docker exec nginx-container touch /my-file
```

- Подсказки по командам контейнеров и образов; удаление образа:

```bash
docker container --help
docker image --help
docker image rm hello-world
```

- Создать volume:

```bash
docker volume create sqlite_data
```

---

## Связанное

- [[Docker compose]]
- [[Docker установка на сервер]]
- [[Docker сборка и запуск контейнеров]]
- [[Docker dump]]
- [[docker_example_description]]
- [[install_docker_on_android]]
