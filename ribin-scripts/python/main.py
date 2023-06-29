from common.config import global_config
from common.logger import logger


def init_environment():
    global_config.load_config("./config.yaml")
    logger.init_logger(global_config.log_config)


def main():
    init_environment()
    logger.info("StartServer")


if __name__ == "__main__":
    main()
