from typing import Optional, Callable, Any
from typing_extensions import Annotated, get_origin, get_args


class InnerDepends:
    def __init__(self, dependency: Optional[Callable[..., Any]] = None):
        self.dependency = dependency

    def __repr__(self) -> str:
        attr = getattr(self.dependency, "__name__", type(self.dependency).__name__)
        return f"{self.__class__.__name__}({attr})"


def FabricDepends(dependency: Optional[Callable[..., Any]] = None) -> Any:
    """
    Логика с типами взята из FastApi Depends.

    Реализация:


    """
    return InnerDepends(dependency=dependency)


def fabric_inject(func):
    def wrapper(*args, **kwargs):
        # Тащим аннотации напрямую из __annotations__, чтобы сохранить информацию об Annotated
        annotations: dict = func.__annotations__
        for param, hint in annotations.items():
            # Без Annotated не продолжаем, нужно соблюдать соглашение по тому, как мы передаем значения для инъекции
            if get_origin(hint) == Annotated:
                # Извлекаем метаданные из Annotated
                dependency = get_args(hint)[1]
                if isinstance(dependency, InnerDepends):
                    # Создаем экземпляр зависимости и просто передаем его в kwargs,
                    # поулчается что в параметр функции service передано value >> экземпляр класса Service
                    kwargs[param] = dependency.dependency()
        return func(*args, **kwargs)
    return wrapper


class Service:
    def hello(self):
        print("hello, depends!")


def service():
    return Service()


@fabric_inject
def main(service: Annotated[Service, FabricDepends(service)]):
    return service.hello()


if __name__ == "__main__":
    main()
