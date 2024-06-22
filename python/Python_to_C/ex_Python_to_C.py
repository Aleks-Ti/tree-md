# example.py

import ctypes

# Загружаем библиотеку на C
lib = ctypes.CDLL('./ci_math.exe')  # Путь к скомпилированной библиотеке .so или .dll

# Объявляем типы аргументов и возвращаемого значения функции
lib.add.argtypes = [ctypes.c_int, ctypes.c_int]
lib.add.restype = ctypes.c_int


# Обертка для функции add
def add(a, b):
    return lib.add(a, b)


# Пример использования функции
print("Результат сложения:", add(10, 20))
