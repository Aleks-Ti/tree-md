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
