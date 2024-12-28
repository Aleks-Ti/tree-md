# SertBot

## Навигация

- [SertBot](#sertbot)
- - [установка](#установка)
- - [генерация](#поставить-для-конфигагенерация-сертификата)
- - [продление](#автоматическое-продление)
- - [оптимизация](#оптимизация-безопасности)

### Установка

```bash
sudo apt update
```

```bash
sudo apt install certbot python3-certbot-nginx
```

### Поставить для конфига/Генерация сертификата

```bash
sudo certbot --nginx -d example.com -d www.example.com
```

### Автоматическое продление

Certbot автоматически настраивает задачу в cron для продления сертификатов. Но вы можете проверить это вручную:

```bash
sudo certbot renew --dry-run
```

### Оптимизация безопасности

`HSTS (HTTP Strict Transport Security)`:

Добавить заголовок, чтобы принудительно использовать HTTPS:

```bash
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
```

`OCSP Stapling`:

Ускорьте проверку сертификатов клиентами:

```bash
ssl_stapling on;
ssl_stapling_verify on;
resolver 1.1.1.1 8.8.8.8 valid=300s;
resolver_timeout 5s;
```