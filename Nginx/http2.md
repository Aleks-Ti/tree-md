# HTTP2

HTTP/2 — это улучшенная версия HTTP, обеспечивающая многопоточность, сжатие заголовков и другие преимущества. Для включения нужно использовать SSL/TLS, так как большинство браузеров поддерживают HTTP/2 только через HTTPS.

## Подключение

- Активация HTTP/2

Убедитесь, что в конфигурации сервера есть директива `http2`:

```nginx
server {
    listen 443 ssl http2;
    server_name example.com www.example.com;

    # Пути к SSL-сертификатам
    ssl_certificate /etc/letsencrypt/live/example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/example.com/privkey.pem;

    # HTTP/2 настройки
    http2_push_preload on;

    location / {
        proxy_pass http://backend_servers;
    }
}
```

- HTTP/2 Push (опционально)

Механизм `HTTP/2 Server Push` позволяет серверу `"предугадывать"`, какие ресурсы (CSS, JS, изображения) понадобятся клиенту, и отправлять их заранее:

```nginx
location / {
    http2_push /static/styles.css;
    http2_push /static/script.js;
    proxy_pass http://backend_servers;
}
```

- Проверка HTTP/2

После перезапуска Nginx можно проверить, что HTTP/2 работает:

Использовать онлайн-тесты, например, https://tools.keycdn.com/http2-test.

Проверить в браузере: в `DevTools > Network` ищите колонку `Protocol` (будет указано `h2`).

## Full example

```nginx
server {
    listen 80;
    server_name example.com www.example.com;

    # Перенаправление HTTP на HTTPS
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl http2;
    server_name example.com www.example.com;

    # SSL настройки
    ssl_certificate /etc/letsencrypt/live/example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/example.com/privkey.pem;

    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers on;

    # HTTP/2 настройки
    http2_push_preload on;

    # Заголовки безопасности
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;

    location / {
        proxy_pass http://backend_servers;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```
