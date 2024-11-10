"""Фабрика зависимостей для FastApi приложения."""

from abc import abstractmethod
from typing import Any


class ContestRepository:
    """from src.contests.repository import ContestRepository"""
    ...


class ContestsService:
    """from src.contests.service import ContestsService"""
    ...


def main_metrics_service():
    """from src.metrics.dependencies import main_metrics_service"""
    ...


class UserRepository:
    """from src.user.repository import UserRepository"""
    ...


class DependsFactory:
    """
    Обязательный параметр authorised_mode, предупреждает ошибки программиста, опечатки и прочее, \
    чтобы передаваемые флаги были действительными.
    Нужно заранее объявлять, какие флаги существуют для вашей фабрики.
    Пример:
    ```
    def __init__(self, modes: tuple):
        self.authorised_mode = {"metrics"}
        super().__init__(self.authorised_mode, modes)
    ...
    ```
    """

    def __init__(self, authorised_mode: set, modes: tuple[str] | None = None,):
        self.authorised_mode = authorised_mode
        self.additional_services: list = []
        """Контейнер для складывания зависимостей, которые потом разом прокидываются в параметры класса."""
        self.modes: tuple | None = modes
        self.validate_parameter()
        self.validate_modes()
        self.validate_authorised_mode()

    def validate_parameter(self) -> None:
        if self.modes:
            if not isinstance(self.modes, tuple):
                raise ValueError(
                    f"Было передано {type(self.modes)}, когда ожидает кортеж строк >> tuple(str) | None\n"
                    "`mode` передаваемые поля, должны быть строкового типа.\n"
                    'Пример: `Depends(lambda: *_service("metrics", "user", ...))`'
                )

    def validate_modes(self) -> None:
        if self.modes:
            if not all(isinstance(mode, str) for mode in self.modes):
                raise ValueError(
                    "`mode` передаваемые поля, должны быть строкового типа.\n"
                    'Пример: `Depends(lambda: *_service("metrics", "user", ...))`'
                )
            for mode in self.modes:
                if mode not in self.authorised_mode:
                    raise ValueError(
                        f"mode >> {mode}, не объявлен в разрешенных значениях `authorised_mode`, если mode `{mode}` "
                        f'является валидным флагом для фабрики зависимостей, просто объявите его в authorised_mode = {{"{mode}"}},'
                        f"либо устраните опечатку флага или не валидный флаг {mode}"
                    )

    def validate_authorised_mode(self) -> None:
        if not self.authorised_mode:
            raise ValueError(
                f"Error {self.__class__.__name__}\n"
                "`authorised_mode` должен быть множество с флагами значений.\n"
                "Определите authorised_mode в классе потомке.\n"
                "Пример:\n"
                "class ContestFactory(DependsFactory):\n"
                '   authorised_mode = {"metrics", "user"}'
            )
        if not isinstance(self.authorised_mode, set):
            raise ValueError(
                f"Error {self.__class__.__name__}\n"
                "`authorised_mode` должен быть кортеж с флагами значений.\n"
                "Пример:\n"
                "class ContestFactory(DependsFactory):\n"
                '   authorised_mode = {"metrics", "user"}'
            )

    @abstractmethod
    def get_service(self) -> Any:
        """
        Для опционального добавления N количество сервисов или репозиториев,
        используйте паттерн декоратор через self.additional_services: list.<br>
        `self.additional_services` объявлен в наследуемом классе DependsFactory:
        ```
        def __init__(self, authorised_mode: set, modes: tuple[str] | None = None):
            ...
            self.additional_services: list = []
            ...
        ```

        Пример реализации:

        ```
        class ContestFactory(DependsFactory):
            def __init__(self, modes: tuple):
                self.authorised_mode = {"metrics", "user"}
                super().__init__(self.authorised_mode, modes)

            def get_service(self) -> ContestsService:
                contest_service = ContestsService(ContestRepository)
                for mode in self.modes:
                    if mode == "metrics":
                        self.additional_services.append(main_metrics_service)
                    if mode == "user":
                        self.additional_services.append(UserRepository)

                if self.additional_services:
                    contest_service = ContestsService(ContestRepository, *self.additional_services)

                return contest_service
        ```
        """
        raise NotImplementedError(f"Вы обязаны реализовать метод get_service в классе наследнике {self.__class__.__name__}")


class ContestFactory(DependsFactory):
    def __init__(self, modes: tuple):
        self.authorised_mode = {"metrics", "user"}
        super().__init__(self.authorised_mode, modes)

    def get_service(self) -> ContestsService:
        contest_service = ContestsService(ContestRepository)
        if not self.modes:
            return contest_service

        for mode in self.modes:
            if mode == "metrics":
                self.additional_services.append(main_metrics_service)
            if mode == "user":
                self.additional_services.append(UserRepository)

        if self.additional_services:
            contest_service = ContestsService(ContestRepository, *self.additional_services)

        return contest_service


def contest_service(*modes: str) -> ContestsService:
    """
    Всегда объявлять как callable, даже без параметров.

    Пример без параметров:
    ```
    async def create_contest(
        ...,
        contest_service: Annotated[ContestsService, Depends(lambda: _contest_service())],
        ...,
    ):
    ```
    Пример с параметрами:
    ```
    async def create_contest(
        ...,
        contest_service: Annotated[ContestsService, Depends(lambda: _contest_service("metrics", "users", ..., ...,))],
        ...,
    ):
    ```
    Params:
    - "metrics" >> set service MainMetricsService
    - "user" >> set service UserRepository
    """

    factory = ContestFactory(modes)
    return factory.get_service()
