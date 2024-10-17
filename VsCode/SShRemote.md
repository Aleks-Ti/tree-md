# SSH REMOTE server

Отвалился сервер vscode на виртуалке

```bash
rm -rf ~/.vscode-server
```

попробовать заново подключится.

Бывает SSH в целом отваливается, можно вырубить демона со службой ssh:

```bash
sudo systemctl restart sshd
```

или

```bash
sudo systemctl stop sshd
# затем
sudo systemctl start sshd
```

, и заново запустить, но в таком случаем будет сделан новый IP серва.
