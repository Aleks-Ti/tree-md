---
tags:
  - linux
  - systemd
  - cheatsheet
  - draft
up: "[[MOC Linux]]"
---

> [!info] Раздел: [[MOC Linux]]

# systemd и journalctl

## Управление службами

| Команда | Действие |
| --- | --- |
| `systemctl status <name>` | состояние и последние логи |
| `systemctl start / stop / restart <name>` | запуск / остановка / перезапуск |
| `systemctl reload <name>` | перечитать конфиг без остановки |
| `systemctl enable --now <name>` | автозапуск + сразу стартовать |
| `systemctl disable <name>` | убрать из автозапуска |
| `systemctl list-units --type=service --state=running` | что запущено |
| `systemctl --failed` | упавшие службы |
| `systemctl daemon-reload` | после правки unit-файлов |

## Логи

```bash
journalctl -u nginx                 # логи службы
journalctl -u nginx -f              # следить в реальном времени
journalctl -u nginx --since "1 hour ago"
journalctl -p err -b                # ошибки с последней загрузки
journalctl --disk-usage             # сколько занимают
sudo journalctl --vacuum-time=7d    # почистить старше 7 дней
```

См. также [[журналы и чистка]].

## Свой unit-файл

`/etc/systemd/system/myapp.service`:

```ini
[Unit]
Description=My app
After=network.target

[Service]
User=app
WorkingDirectory=/opt/myapp
ExecStart=/opt/myapp/.venv/bin/uvicorn main:app --port 8000
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now myapp
```

---

## Связанное

- [[Nginx base]]
- [[журналы и чистка]]
- [[Порты, и действия с ними]]
- [[OC]]
