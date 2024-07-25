# Make

Коротко.
Чтобы поставить make на windows, нужно сначала установить Chocolatey(choco - виндовый пакетный менеджер)

## Ubuntu(если не стоит с коробки)

```bash
sudo apt install make
```

## Windows

## Ставим на винду `Chocolatey`

ссылка на доку `https://chocolatey.org/install`

выполнить команду в терминале `Shell`:

```bash
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

Проверить командой `choco` если все ок, будет отзыв.

- Теперь можно поставить `make`

```bash
choco install make
```
