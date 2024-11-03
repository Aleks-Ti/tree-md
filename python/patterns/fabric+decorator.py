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


class DependsFactory:
    authorised_mode: set | None = None
    additional_services: list = []

    def __init__(self, modes: list[str] | None = None):
        self.modes: list | None = modes
        self.validate_parameter()
        self.validate_modes()
        self.validate_authorised_mode()

    def validate_parameter(self) -> None:
        if self.modes:
            if not isinstance(self.modes, list):
                raise ValueError(
                    f"Было передано {type(self.modes)}, когда ожидает список строк >> list[str] | None\n"
                    "`mode` передаваемые поля, должны быть строкового типа.\n"
                    'Пример: `Depends(lambda: *_service(["metrics", "user", ...]))`'
                )

    def validate_modes(self) -> None:
        if self.modes:
            if not all(isinstance(mode, str) for mode in self.modes):
                raise ValueError(
                    "`mode` передаваемые поля, должны быть строкового типа.\n"
                    'Пример: `Depends(lambda: *_service(["metrics", "user", ...]))`'
                )

    def validate_authorised_mode(self) -> None:
        if not self.authorised_mode:
            raise ValueError(
                f"Error {self.__class__.__name__}\n"
                "`authorised_mode` должен быть кортеж с флагами значений.\n"
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
        используйте паттерн декоратор через self.additional_services: list

        Пример:

        ```
        def get_service(self) -> ContestsService:
            contest_service = ContestsService(ContestRepository)
            for mode in self.modes:
                if mode == "metrics":
                    self.additional_services.append(main_metrics_service)
                if mode == "user":
                    self.additional_services.append(main_metrics_service)

            if self.additional_services:
                contest_service = ContestsService(ContestRepository, *self.additional_services)

            return contest_service
        ```
        """
        raise NotImplementedError(f"Вы обязаны реализовать метод get_service в классе наследнике {self.__class__.__name__}")


class ContestFactory(DependsFactory):
    authorised_mode = {"metrics"}

    def get_service(self) -> ContestsService:
        contest_service = ContestsService(ContestRepository)
        if not self.modes:
            return contest_service

        for mode in self.modes:
            if mode == "metrics":
                self.additional_services.append(main_metrics_service)

        if self.additional_services:
            contest_service = ContestsService(ContestRepository, *self.additional_services)

        return contest_service


def contest_service(modes: list[str] | None = None) -> ContestsService:
    """
    Params:
    - "metrics" >> set service MainMetricsService
    """

    factory = ContestFactory(modes)
    return factory.get_service()
