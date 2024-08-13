import multiprocessing


def calculate(num):
    # Имитация вычислений
    result = num ** 2
    print(f"Процесс {multiprocessing.current_process().name} вычислил {num}^2 = {result}")
    return result


if __name__ == '__main__':
    # Создаем пул процессов
    pool = multiprocessing.Pool(processes=4)

    # Передаем список чисел для вычисления
    numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    results = pool.map(calculate, numbers)

    print("Результаты вычислений:")
    for num, res in zip(numbers, results):
        print(f"{num}^2 = {res}")

    # Закрываем пул процессов
    pool.close()
    pool.join()