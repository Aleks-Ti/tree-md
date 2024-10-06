# PathLib

## Чтение и запись файлов

```python
from pathlib import Path

# Создание объекта пути
file_path = Path('example.txt')

# Запись в файл
file_path.write_text('Hello, World!')

# Чтение из файла
content = file_path.read_text()
print(content)  # Вывод: Hello, World!
```

## Работа с директорией

```python
from pathlib import Path

# Создание объекта пути к директории
dir_path = Path('example_dir')

# Создание директории
dir_path.mkdir(parents=True, exist_ok=True)

# Проверка существования директории
if dir_path.exists():
    print(f"Directory {dir_path} exists")

# Получение списка файлов в директории
for file in dir_path.iterdir():
    print(file)
```

## Поиск файлов по шаблону

```python
from pathlib import Path

# Поиск всех файлов с расширением .txt в текущей директории
for txt_file in Path('.').glob('*.txt'):
    print(txt_file)
```

## Получение информации о файле

```python
from pathlib import Path

# Создание объекта пути
file_path = Path('example.txt')

# Получение размера файла
file_size = file_path.stat().st_size
print(f"File size: {file_size} bytes")

# Получение времени последнего изменения
modification_time = file_path.stat().st_mtime
print(f"Last modification time: {modification_time}")
```

## Создание вложенных директорий

```python
from pathlib import Path

# Создание вложенных директорий
nested_dir_path = Path('parent_dir/child_dir')
nested_dir_path.mkdir(parents=True, exist_ok=True)
```

## Удаление файла или директории

```python
from pathlib import Path

# Удаление файла
file_path = Path('example.txt')
if file_path.exists():
    file_path.unlink()

# Удаление пустой директории
dir_path = Path('example_dir')
if dir_path.exists() and dir_path.is_dir():
    dir_path.rmdir()
```

## Копирование файла

```python
from pathlib import Path
import shutil

# Копирование файла
source = Path('example.txt')
destination = Path('example_copy.txt')
shutil.copy(source, destination)
```

## Перемещение файла

- `Переименование и перемещение`: Метод `rename` выполняет `обе функции` — он может `переименовать` файл или директорию, а также `переместить` их в новое место.

```python
from pathlib import Path

# Перемещение файла
source = Path('example.txt')
destination = Path('new_location/example.txt')
source.rename(destination)
```

## Получение абсолютного пути

```python
from pathlib import Path

# Получение абсолютного пути
relative_path = Path('example.txt')
absolute_path = relative_path.resolve()
print(absolute_path)
# or 
relative_path = Path('example.txt').resolve()
```
