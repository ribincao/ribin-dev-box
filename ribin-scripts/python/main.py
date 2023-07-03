from common.config import global_config
from common.logger import logger


def init_environment():
    global_config.load_config("./config.yaml")
    logger.init_logger(global_config.log_config)

    import os

    os.environ["OPENAI_API_KEY"] = global_config.api_keys.openai_api
    os.environ["SERPAPI_API_KEY"] = global_config.api_keys.serp_api
    os.environ["ACTIVELOOP_TOKEN"] = global_config.api_keys.active_loop_api


def main():
    init_environment()
    logger.info("StartServer")


if __name__ == "__main__":
    main()
