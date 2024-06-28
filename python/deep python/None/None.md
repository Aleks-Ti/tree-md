# None

```bash
dir(None)
[
    '__bool__',
    '__class__',
    '__delattr__',
    '__dir__',
    '__doc__',
    '__eq__',
    '__format__',
    '__ge__',
    '__getattribute__',
    '__gt__',
    '__hash__',
    '__init__',
    '__init_subclass__',
    '__le__',
    '__lt__',
    '__ne__',
    '__new__',
    '__reduce__',
    '__reduce_ex__',
    '__repr__',
    '__setattr__',
    '__sizeof__',
    '__str__',
    '__subclasshook__'
]
 ```

```None``` - умеет хэшироваться. Не изменяемый.

```python
None.__hash__()
# -9223363241436679553
```

```python
None.__bool__()
# False
```

```python
None.__str__()
# "None"
```

```python
None.__sizeof__()
# 16
```

```python
None.__gt__(1)
# None.__gt__(1)
```

## Плохая практика использования None

- Пример когда плохо:

```python
def main(arrgs: list[int] | None = None):
    ...
```

- - Типо тут мы позволяем не передавать конкретно значений. Иногда это конечно полезно, все понятно, но если мы пишем код не для внутрянки, а для библиотеки или тп, что то значимое, то это оч плохой варик. Мы позволяем прокинуть None, а если мы позволили, то это обязательно кто то да и сделает!)

Решение:

есть такой паттерн, назвается - ```santinel value```, можно воспользоваться им

```python
from typing import Any

_sentinel: Any = object()

def accepts_list(x: list[int] = _sentinel):
    if x is _sentinel:
        x = [1, 2, 3]
    print(x)

accepts_list()
accepts_list([4, 5, 7])
# [1, 2, 3]
# [4, 5, 7]
```
