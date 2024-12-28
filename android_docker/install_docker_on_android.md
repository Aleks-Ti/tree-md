# Установка

кейс - по фану поставить докер и заставить крутить приложения

## Подготовка телефона.

Подопытный Самсунг а12

- включение Режима Разработчика

Зайти - Настройки>>Сведения о телефоне>>Сведения о ПО>>`прокликать 7 раз в строке Номер сборки`
- - Пройдет оповещение, что включен режим разраба
- затем нужно выйти назад в Настройки.
- - В самом низу появится подменю `Параметры разработчика` - заходим
- - - Включаем галочку на `Заводская разблакировка` и `Отладка по USB`

Выходим/Подключаемся по USB к компу. Выбираем USB для передачи файлов.

- Идем в ПК, видим свое устройство в МОЙ компьютер.

Нужно поставить pdb:

Скачайте и установите последнюю версию Android SDK Platform Tools (https://developer.android.com/studio/releases/platform-tools) на ваш компьютер.

- Добавьте папку с инструментами `adb` в переменную среды `PATH` (у меня это получилась папка `platform-tools`).
- - скачал, распаковал на рабочий стол, и добавил эту папку в PATH винды

затем через виндовый терминал юзал `adb`

```bash
# Введите следующую команду для проверки подключенных устройств:
adb devices
# зайти в устройство в терминальном режиме
adb shell
```

## Установка Termux

Крч так не поставишь особо докер и тп. Потому пришлось ставить эту утилиту.

Можно поставить как и через гугл плэй так и ручками.

Я ставил поскольку уже был установлен adb поставил через него.

- Сначала скачал с сайта `aok` https://f-droid.org/fr/packages/com.termux/

```shell
cd path/to/termux.apk  # чет через абсолютный путь не запускало, пришлось идти в папку с файлом
adb install termux-app_v0.119.0-beta.1+apt-android-7-github-debug_universal.apk
# При установке может подвиснуть, но это не подвис, а телефон будет стрематься ставить че то сторонне, в телефоне подтверждаем установку
```

### Поскольку с телефона отстой что либо делать, чутка мучаемся, но профит не заставит себя ждать

Заходим в Termux на телефоне, ставим тулзы чтобы подрубится под SSH

```bash
pkg install openssh
```

ставим пароль

```bash
passwd
```

запускаем ssh server

```bash
sshd
```

По умолчанию SSH-сервер будет слушать на порту `8022`. Вы можете изменить порт, если это необходимо, отредактировав файл конфигурации.

Подключитесь с компьютера:
Теперь вы можете подключиться к вашему Android-устройству с компьютера через SSH. Используйте терминал на вашем ПК и выполните команду:

- Поддержка постоянного доступа
Если вы хотите, чтобы SSH-сервер автоматически запускался при каждом запуске Termux, вы можете использовать termux-services. Для этого создайте директорию для сервиса и файл конфигурации:

```bash
mkdir -p ~/.termux/boot
echo "sshd" > ~/.termux/boot/start-sshd.sh
chmod +x ~/.termux/boot/start-sshd.sh
```

- Залетаем по SSH

```bash
ssh username@<ваш_IP_адрес> -p 8022
```

Эта команда даст айпишку:

```bash
ifconfig
```

Эта команда выведет имя текущего пользователя.

```bash
whoami
```

## Увы на андроид не поставишь так просто ДОКЕР :(

### Ставим Ubuntu :)

```bash
pkg update
pkg upgrade
pkg install proot-distro
proot-distro install ubuntu
proot-distro login ubuntu
```

залетаем в итоге на ubuntu в андроиде под рутом после эти команд

уже в терминальной среде ubuntu ставим докер

```bash
apt update
apt upgrade
apt install docker.io
```

запускам

```bash
service docker start
docker --version
```

Начались приключения. Докер не стартует, но установился.
Какие то пермишины отстреливает

```bash
INFO[2024-12-28T21:47:17.668128662Z] loading plugin "io.containerd.tracing.processor.v1.otlp"...  type=io.containerd.tracing.processor.v1
INFO[2024-12-28T21:47:17.668459739Z] skip loading plugin "io.containerd.tracing.processor.v1.otlp"...  error="no OpenTelemetry endpoint: skip plugin" type=io.containerd.tracing.processor.v1
INFO[2024-12-28T21:47:17.668619585Z] loading plugin "io.containerd.internal.v1.tracing"...  type=io.containerd.internal.v1
INFO[2024-12-28T21:47:17.668704893Z] skipping tracing processor initialization (no tracing plugin)  error="no OpenTelemetry endpoint: skip plugin"
INFO[2024-12-28T21:47:17.670155816Z] loading plugin "io.containerd.grpc.v1.healthcheck"...  type=io.containerd.grpc.v1
INFO[2024-12-28T21:47:17.670417046Z] loading plugin "io.containerd.nri.v1.nri"...  type=io.containerd.nri.v1
INFO[2024-12-28T21:47:17.670537200Z] NRI interface is disabled by configuration.
INFO[2024-12-28T21:47:17.675319816Z] serving...                                    address=/run/docker/containerd/containerd-debug.sock
INFO[2024-12-28T21:47:17.678573585Z] serving...                                    address=/run/docker/containerd/containerd.sock.ttrpc
INFO[2024-12-28T21:47:17.681870739Z] serving...                                    address=/run/docker/containerd/containerd.sock
INFO[2024-12-28T21:47:17.682263046Z] containerd successfully booted in 0.282551s
WARN[2024-12-28T21:47:18.330480739Z] unable to modify root key limit, number of containers could be limited by this quota: open /proc/sys/kernel/keys/root_maxkeys: permission denied
WARN[2024-12-28T21:47:18.438677200Z] Failed to configure golang's threads limit: open /proc/sys/kernel/threads-max: permission denied
INFO[2024-12-28T21:47:18.545633893Z] stopping healthcheck following graceful shutdown  module=libcontainerd
INFO[2024-12-28T21:47:18.545746739Z] stopping event stream following graceful shutdown  error="context canceled" module=libcontainerd namespace=plugins.moby
failed to start daemon: Devices cgroup isn't mounted
```

- это провал.
- - чтобы попытаться поставить все это барахло, слишком много вложенностей получилось, да и для \
докера нужны cgroups, а тулза proot-distro на которой запущена ubuntu в которой я ставил докер, \
не вывозит без нужных прав и доступов. Потому проще даже не ставить ubuntu а просто в termux \
запускать код, если необходимо
