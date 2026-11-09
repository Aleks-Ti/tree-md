---
tags:
  - nginx
  - cheatsheet
up: "[[MOC Инфраструктура]]"
---

> [!info] Раздел: [[MOC Инфраструктура]]

# Nginx: базовые команды и расположение настроек

## Где что лежит

| Путь | Назначение |
| --- | --- |
| `/etc/nginx/nginx.conf` | главный конфиг (подключает остальные) |
| `/etc/nginx/sites-available/` | все описанные сайты |
| `/etc/nginx/sites-enabled/` | активные сайты (симлинки на `sites-available`) |
| `/etc/nginx/conf.d/*.conf` | дополнительные конфиги |
| `/var/log/nginx/access.log`, `error.log` | логи, см. [[Nginx logs]] |

Включить сайт: `sudo ln -s /etc/nginx/sites-available/mysite /etc/nginx/sites-enabled/`

## Команды

| Команда | Что делает |
| --- | --- |
| `sudo nginx -t` | проверить синтаксис конфига (**всегда перед reload**) |
| `sudo nginx -T` | проверить и вывести итоговый конфиг целиком |
| `sudo systemctl status nginx` | статус службы |
| `sudo systemctl reload nginx` | применить конфиг без обрыва соединений |
| `sudo systemctl restart nginx` | полный перезапуск |
| `sudo systemctl stop nginx` / `start nginx` | остановить / запустить |
| `sudo systemctl enable nginx` | автозапуск при загрузке |

> [!tip] Порядок правки конфига
> Поправил → `nginx -t` → `reload`. Если `-t` ругается, служба продолжит работать на старом конфиге.

Если порт 80/443 занят — см. [[Порты, и действия с ними]]; управление службами — [[systemd]].

---

## Связанное

- [[Nginx logs]]
- [[Nginx certbot]]
- [[Nginx http2]]
- [[балансировщик]]
- [[systemd]]
- [[Порты, и действия с ними]]
