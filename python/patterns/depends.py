from typing import Optional, Callable, Any
from typing_extensions import Annotated, get_origin, get_args
import inspect


class InnerDepends:
    def __init__(self, dependency: Optional[Callable[..., Any]] = None) -> None:
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


def fabric_inject(func) -> Callable[..., Any]:
    def wrapper(*args, **kwargs) -> Any:
        signature = inspect.signature(func)
        for param_name, param in signature.parameters.items():
            # Без Annotated не продолжаем, нужно соблюдать соглашение по тому, как мы передаем значения для инъекции
            if get_origin(param.annotation) == Annotated:
                # Извлекаем метаданные из Annotated
                dependency = get_args(param.annotation)[1]
                if isinstance(dependency, InnerDepends):
                    # Создаем экземпляр зависимости и просто передаем его в kwargs,
                    # поулчается что в параметр функции service передано value >> экземпляр класса Service
                    if param_name not in kwargs:
                        kwargs[param_name] = dependency.dependency()
        return func(*args, **kwargs)
    return wrapper


class Service:
    def hello(self) -> None:
        print("hello, depends!")


def service() -> Service:
    return Service()


@fabric_inject
def main(service: Annotated[Service, FabricDepends(service)]) -> None:
    """_summary_

    Args:
        service (Annotated[Service, FabricDepends): _description_

    Returns:
        _type_: _description_
    """
    return service.hello()


if __name__ == "__main__":
    main()
