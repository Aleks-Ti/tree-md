"""Фабрика зависимостей для FastApi приложения."""
from collections import defaultdict
from abc import abstractmethod
from typing import Any


class MetricsRepository:
    """from src.metrics.repository import ContestRepository"""
    def increase_value(self) -> None:
        print("I go to the database and increase the value")


class MetricsService:
    """from src.metrics.services import MetricsService"""
    def __init__(self, metrics_repo: MetricsRepository) -> None:
        self.metrics_repo: MetricsRepository = metrics_repo()

    def increase_metric(self) -> None:
        print("I call complex logic for counts and stuff, and then call the repository to access the database")
        self.metrics_repo.increase_value()


class ContestRepository:
    """from src.contests.repository import ContestRepository"""

    def get_db_info(self) -> None:
        print("hello, i am ContestRepo, get data from db!")


class UserRepository:
    """from src.user.repository import UserRepository"""
    def get_user(self) -> None:
        print("I'm a teapot")


def main_metrics_service() -> MetricsService:
    """from src.metrics.dependencies import main_metrics_service"""
    return MetricsService(MetricsRepository)


class ContestsService:
    """from src.contests.service import ContestsService"""
    def __init__(
            self, contest_repo: ContestRepository,
            metrics_service: MetricsService | None = None,
            user_repo: UserRepository | None = None
    ) -> None:
        self.contest_repo: ContestRepository = contest_repo
        self.metrics_service: MetricsService = metrics_service
        self.user_repo: UserRepository = user_repo


class DependsFactory:
    """
    Фабрика по сборке зависимостей(доп параметров определенных в конструкторе __init__) к сервису других сервисов \
          или репозиториев (могут быть как и другие модули, вспомогательные функции, интерфейсы или структуры).<br>

    Params:
        authorised_mode (dict[str, Any]): Обязательный параметр authorised_mode.
            Нужно объявлять заранее, какие флаги существуют для вашей фабрики.
            Пример:
            ```
            def __init__(self, modes) -> None:
                super().__init__(modes, "metrics", "user")
            ...
            ```
        modes (tuple[str]): флаги для присоединения зависимостей к сервису.
            Пример:
            ```
            # получаем modes и инициализируем с ним экземпляр класса.
            factory = ContestFactory(modes)

            # далее передаем флаги в конструктор __init__ в классе потомке.
            class ContestFactory(DependsFactory):
                def __init__(self, modes) -> None:
                    super().__init__(modes, "metrics", "user")
            ...
            ```
    """

    def __init__(self, modes: tuple[str | None], *authorised_mode: str) -> None:
        self.authorised_mode: defaultdict = self.set_authorised_mode(*authorised_mode)
        self.modes: tuple[str | None] = modes
        self.validate()

    def validate(self) -> None:
        self.validate_authorised_mode()
        self.validate_modes_type_hint()
        self.validate_values_modes()

    def validate_modes_type_hint(self) -> None:
        if self.modes:
            if not isinstance(self.modes, tuple):
                raise ValueError(
                    f"Было передано {type(self.modes)}, когда ожидает кортеж строк >> tuple(str) | None\n"
                    "`mode` передаваемые поля, должны быть строкового типа.\n"
                    'Пример: `service = name_service("metrics", "user", ...))`\n'
                    "Пример в FastApi:\n"
                    '`service: ContestService: Annotated[ContestService, Depends(lambda: *_service("metrics", "user", ...))]`\n'
                )

    def validate_values_modes(self) -> None:
        if self.modes:
            if not all(isinstance(mode, str) for mode in self.modes):
                raise ValueError(
                    "`mode` передаваемые поля, должны быть строкового типа.\n"
                    'Пример: `service = name_service("metrics", "user", ...))`\n'
                    "Пример в FastApi:\n"
                    '`service: ContestService: Annotated[ContestService, Depends(lambda: *_service("metrics", "user", ...))]`\n'
                )
            for mode in self.modes:
                if mode not in self.authorised_mode:
                    raise ValueError(
                        f"mode >> {mode}, не объявлен в разрешенных значениях `authorised_mode`, если mode `{mode}` "
                        "является валидным флагом для фабрики зависимостей, просто объявите его в authorised_mode,"
                        f"либо устраните опечатку флага или не валидный флаг {mode}\n"
                        "Пример:\n"
                        "class ContestFactory(DependsFactory):\n"
                        "    def __init__(self, modes: tuple) -> None:\n"
                        '        super().__init__(modes, "metrics", "user")'
                    )

    def validate_authorised_mode(self) -> None:
        if not self.authorised_mode:
            raise ValueError(
                f"Error {self.__class__.__name__}\n"
                "Отсутствует флаги authorised_mode.\n"
                "Определите к authorised_mode ключи, в классе потомке.\n"
                "Пример:\n"
                "class ContestFactory(DependsFactory):\n"
                "    def __init__(self, modes: tuple) -> None:\n"
                '        super().__init__(modes, "metrics", "user")'
            )
        if not isinstance(self.authorised_mode, defaultdict):
            raise ValueError(
                f"Error {self.__class__.__name__}\n"
                f"Нельзя переопределять: `authorised_mode`, тогда как ожидалось authorised_mode: defaultdict\n"
                "`authorised_mode` должен быть  defaultdict со строковыми ключами и дефолтными значениями None.\n"
                "Пример:\n"
                "class ContestFactory(DependsFactory):\n"
                "    def __init__(self, modes: tuple) -> None:\n"
                '        super().__init__(modes, "metrics", "user")'
            )

    def set_authorised_mode(self, *args: str) -> defaultdict:
        default_dict = defaultdict(None)
        for flag in args:
            if not isinstance(flag, str):
                raise ValueError(
                    f"Error {self.__class__.__name__}\n"
                    f"Было передано для `authorised_mode({type(flag)})`,"
                    f' тогда как ожидалось authorised_mode("string", "string", ...)\n'
                    "Пример:\n"
                    "class ContestFactory(DependsFactory):\n"
                    "    def __init__(self, modes: tuple) -> None:\n"
                    '        super().__init__(modes, "metrics", "user")'
                )
            default_dict.setdefault(flag)
        return default_dict

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
            def __init__(self, modes) -> None:
                super().__init__(modes, "metrics", "user")

            def get_service(self) -> ContestsService:
                if self.modes:
                    for mode in self.modes:
                        if mode == "metrics":
                            self.authorised_mode["metrics"] = main_metrics_service()
                        if mode == "user":
                            self.authorised_mode["user"] = UserRepository()

                contest_service = ContestsService(
                    contest_repo=ContestRepository(),
                    user_repo=self.authorised_mode["user"],
                    metrics_service=self.authorised_mode["metrics"]
                )
                return contest_service
        ```
        """
        raise NotImplementedError(f"Вы обязаны реализовать метод get_service в классе наследнике {self.__class__.__name__}")


class ContestFactory(DependsFactory):
    def __init__(self, modes) -> None:
        super().__init__(modes, "metrics", "user")

    def get_service(self) -> ContestsService:
        if self.modes:
            for mode in self.modes:
                if mode == "metrics":
                    self.authorised_mode["metrics"] = main_metrics_service()
                if mode == "user":
                    self.authorised_mode["user"] = UserRepository()

        contest_service = ContestsService(
            contest_repo=ContestRepository(),
            user_repo=self.authorised_mode["user"],
            metrics_service=self.authorised_mode["metrics"],
        )

        return contest_service


def contest_service(*modes: str) -> ContestsService:
    """

    Flags:
        user (str): Добавляет репозиторий User к сервису.
        metrics (str): Добавляет сервис метрик к сервису.

    Всегда объявлять как callable, даже без параметров.

    Пример без параметров для FastApi:
    ```
    async def create_contest(
        ...,
        contest_service: Annotated[ContestsService, Depends(lambda: _contest_service())],
        ...,
    ):
    ```

    Пример с параметрами для FastApi:
    ```
    async def create_contest(
        ...,
        contest_service: Annotated[ContestsService, Depends(lambda: _contest_service("metrics", "users", ..., ...,))],
        ...,
    ):
    ```
    Пример вне контекста FastApi, в остальных случаях:
    ```
    contest_service = _contest_service("metrics", "users", ..., ...,))]
    ```

    Params:
    - "metrics" >> set service MainMetricsService
    - "user" >> set service UserRepository
    """

    factory = ContestFactory(modes)
    return factory.get_service()


def main() -> None:
    example_service: ContestsService = contest_service()
    example_service.contest_repo.get_db_info()

    example_service_1: ContestsService = contest_service("metrics")
    example_service_1.metrics_service.increase_metric()

    example_service_2: ContestsService = contest_service("user")
    example_service_2.user_repo.get_user()

    example_service_3: ContestsService = contest_service("metrics", "user")
    example_service_3.metrics_service.metrics_repo.increase_value()
    example_service_3.user_repo.get_user()


if __name__ == "__main__":
    main()
