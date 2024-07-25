# Конфиг для VS code

На `ctrl + s` выполняет инструкцию ниже.
Создать файл `launch.json` в папке `.vscode` (создать папку если нет) и прописать инструкцию ниже.

```json
{
    "files.autoSave": "afterDelay",
    "files.autoSaveDelay": 1000,
    "editor.formatOnSave": true,
    "[python]": {
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.fixAll": "explicit",
            "source.organizeImports.ruff": "explicit"
        },
        "editor.defaultFormatter": "charliermarsh.ruff",
    },
    "ruff.lint.args": [
        "--config=pyproject.toml" // указывает на файл где лежит конфиг ruff
    ],
    "ruff.interpreter": [
        "./.venv/bin/python3.11"
    ],
    "ruff.path": [
        "./.venv/bin/ruff" // путь до библы ruff в окружении проекта
    ],
    "python.testing.pytestArgs": [
        "./tests"
    ],
}
```
