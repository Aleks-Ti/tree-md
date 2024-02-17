# Конфликт главных веток

Появилась ветка как и main так master

В ветке лежала автогенерируемая лицензия все остальное было пусто.

Когда с клонил себе проект чтобы начать разрабатывать, локально у меня появилась ветка master.
А поскольку дефолтной веткой по умолчанию осталась main, на странице проекта всегда отображалсь она,
но она было пустой и только с лицензией. Вот так и получилось, что ветку никак не удавалось с мержить норм c мастером.

Решение по ветки, спулить себе ее на локалку,

Чекнуть че как на удаленке какая ветка под каким коммитом живет:

Для начала нужно затащить себе ветку с удаленки на локалку, но не факт что это получится сделать.

Если не получится, можно создать ветку main самому, и привязать локальную ветку к удаленной:

Простая команда ```git pull``` не помогает, пишет ошибку,
так же если попытаться так ```git pull origin main``` так же ошибка, и предлагает в трэйсе сделать команду ```git push --set-upstream origin main```, но это так же вызывает ошибку:

```bash
To github.com:Aleks-Ti/python-gRPC-example.git
 ! [rejected]        main -> main (non-fast-forward)
error: failed to push some refs to 'github.com:Aleks-Ti/python-gRPC-example.git'
hint: Updates were rejected because the tip of your current branch is behind
hint: its remote counterpart. Integrate the remote changes (e.g.
hint: 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.
```

Можно почекать че кто и куда какая ветка к какому коммиту привязана, это может проснить картину:

```bash
git branch -vv
```

Решение проблемы:
Поскольку простым пулом не получалось решить проблему, помогла вот эта команда.

```bash
git pull origin main --allow-unrelated-histories
```

После нее, одинокий файл лицензии из ветки main, наконец с пулился ко мне на локальную ветку main, где уже давно было куча кода смерженного из master.
