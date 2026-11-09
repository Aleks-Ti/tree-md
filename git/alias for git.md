---
tags:
  - git
up: "[[MOC Инструменты]]"
---

> [!info] Раздел: [[MOC Инструменты]]

# Git alias

- В домашней директории как правило лежит конфиг, в `~/.gitconfig` там `nano .gitconfig`, добавляем конфиг:

```bash
...
[alias]
che = checkout
ci = commit
st = status
br = branch
sw = switch
hist = log --pretty=format:\"%h %ad | %s%d [%an]\" --graph --date=short --stat
type = cat-file -t
dump = cat-file -p
lg = log --graph --pretty --oneline -10
```

---

## Связанное

- [[Git base]]
