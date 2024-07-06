# Линтер Ruff в найстройках проекта pyproject.toml

```toml
...

[tool.ruff]
line-length = 135
dummy-variable-rgx = "^(_+|(_+[a-zA-Z0-9_]*[a-zA-Z0-9]+?))$"
target-version = "py311"
select = [
    "E",  # pycodestyle
    "F",  # pyflakes
    "UP", # pyupgrade
    "Q",  # quotes
    "I",  # isort
]
fixable = ["Q", "F", "I", "E"]
flake8-quotes.inline-quotes = "double"
ignore = ["F401", "E712"]
exclude = [
    ".bzr",
    ".direnv",
    ".eggs",
    ".git",
    ".git-rewrite",
    ".hg",
    ".mypy_cache",
    ".nox",
    ".pants.d",
    ".pytype",
    ".ruff_cache",
    ".svn",
    ".tox",
    ".venv",
    "__pypackages__",
    "_build",
    "buck-out",
    "build",
    "dist",
    "node_modules",
    "venv",
    "seed_data",
    "migrations",
    "seed",
    "manage.py",
]

[tool.ruff.pydocstyle]
convention = "google"

...
```
