# OPENVPN на серваке

- гайд которому следовал +- `https://www.youtube.com/watch?v=XpY1Jbe7UO4`

## VPS

aeza

## OPenVPN

- Перед всем, неплохобы обновить пустую машину:

```bash
apt update
...
apt updrage
```

- Завести Акк, пройти сюда `https://as-portal.openvpn.com/quick-start` и выбрать `Get Access Sever`

Будет предложен выбор серверов - скипаем, vps уже существует.

Еще ниже будет сегенрированна ссылка на установки openvpn на сервак, копируем, идем ставить на серв

```bash
bash <(curl -fsS https://as-repository.openvpn.net/as/install.sh)
```

не помешало бы ребутнуть машину

```bash
sudo reboot
```

## Работы на VPS-ке

Креды находятся по пути >> `/usr/local/openvpn_as/init.log`

- Чекнуть пользователя `openvpn_as`

Глянуть список пользователей на машине можно тут:

```bash
cat /etc/passwd
```

- зайти на пользователя, но это не нужно по сути, так для справки, тем более у этого пользователя с вероятность 120 нет пароля и не разрешенны подкулючения и тп

```bash
su - openvpn_as
```

## Работ с UI OpenVpn как ADMIN

Заглянуть сюда, через браузер.

`https://your_ip:port/admin/`

Креды находятся по пути >> `/usr/local/openvpn_as/init.log`

Зайти в STATUS в меню слева, и проверить, работает ли VPN, если он работает, будет активна кнопк4а `Stop VPN services`

- Нужно изменить настройки маршрутизации VPN.
- - Зайти >> `CONFIGURATION` >> `VPN Settings`
- - - Прокрутить до подменю `Routing`
- - - - Поставить `NO` для `Should VPN clients have access to private subnets (non-public networks on the server side)?`
- - - Сохранить новые настройки
- - - Прожать вверху в всплывающем предупреждении после сохранения >> `Update Running Server`
- - - Вернуться на вкладку `STATUS`
- - - - Оставновить VPN >> нажать `Stop VPN services`
- - - - Затем запустить, с ту же кнопку но с дргим NAME

## Работ с UI OpenVpn как USER

зайти по адресу `https://193.233.133.151:943/login`

залогинется как пользователь уже, а не админ.

После авторизации, предложит сразу же скачать программу `OpenVPN`
Качаем, если не было до этого установлено.
