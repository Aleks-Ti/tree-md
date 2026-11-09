---
tags:
  - linux
  - ssh
  - security
  - cheatsheet
  - draft
up: "[[MOC Linux]]"
---

> [!info] Раздел: [[MOC Linux]]

# SSH

## Ключи

```bash
ssh-keygen -t ed25519 -C "comment"      # создать пару ключей (~/.ssh/id_ed25519{,.pub})
ssh-copy-id user@host                   # положить публичный ключ на сервер
ssh-add ~/.ssh/id_ed25519               # добавить в ssh-agent
```

Права обязательны: `chmod 700 ~/.ssh`, `chmod 600 ~/.ssh/id_ed25519` — см. [[chmod]].

## Подключение

```bash
ssh user@host -p 22
ssh -i ~/.ssh/key user@host
scp file user@host:/path/        # копирование
rsync -avz ./dir user@host:/path/
```

## Туннели

| Тип | Команда | Смысл |
| --- | --- | --- |
| Локальный | `ssh -L 5432:localhost:5432 user@host` | порт сервера (БД) на моём localhost |
| Обратный | `ssh -R 8080:localhost:3000 user@host` | мой локальный сервис наружу через сервер |
| SOCKS | `ssh -D 1080 user@host` | прокси через сервер |

## ~/.ssh/config

```text
Host myvps
    HostName 203.0.113.10
    User deploy
    Port 22
    IdentityFile ~/.ssh/id_ed25519
```

Дальше достаточно `ssh myvps`.

## Укрепление сервера (`/etc/ssh/sshd_config`)

- `PasswordAuthentication no`
- `PermitRootLogin no`
- после правки: `sudo sshd -t && sudo systemctl reload ssh` ([[systemd]])

---

## Связанное

- [[Криптография]] — почему ключи работают (асимметричное шифрование)
- [[SShRemote]] — VS Code Remote SSH
- [[OpenVPN start]]
- [[Git base]]
- [[chmod]]
