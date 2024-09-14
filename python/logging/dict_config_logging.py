import json
import logging

json_formatter = logging.Formatter(
    json.dumps(
        {
            "time": "%(asctime)s.%(msecs)03d",
            "level": "%(levelname)-6s",
            "funcName": "%(funcName)s()",
            "line": "%(pathname)s:%(lineno)d",
            "loggerName": "%(name)s",
            "message": "%(message)s",
        },
    ),
    datefmt="%Y-%m-%d %H:%M:%S",
)

ch = logging.StreamHandler()
ch.setFormatter(json_formatter)
logger = logging.getLogger("root")
logger.addHandler(ch)


def test():
    try:
        1 / 0
    except ZeroDivisionError:
        logging.error("Error: ZeroDivisionError", exc_info=True)
    except Exception:
        logging.exception("Error: unexpected error")


if __name__ == "__main__":
    test()
