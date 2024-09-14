import logging


def config_logging(level):
    logging.basicConfig(
        level=level,
        datefmt="%Y-%m-%d %H:%M:%S",
        format="[%(asctime)s.%(msecs)03d] %(module)s:%(lineno)d %(levelname)-7s - %(pathname)s:%(lineno)d - %(message)s"
    )


logger = logging.getLogger(__name__)


def test1():
    logger.error("ERROR")
    logger.warning("WARNING")
    logger.info("Info")
    logger.debug("Debug")
    try:
        1 / 0
    except ZeroDivisionError:
        logging.error("Error: ZeroDivisionError", exc_info=True)
    except Exception:
        logging.exception("Error: unexpected error")


def main():
    config_logging(logging.INFO)
    test1()


if __name__ == "__main__":
    main()
