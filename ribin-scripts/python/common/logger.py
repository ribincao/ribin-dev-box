from common.singleton import Singleton
from common.config import LogConfig
from typing import Optional
import logging
import sys
import os

LOG_NAME = "ribin"
LOG_FORMAT = "%(levelname)s | %(asctime)s | %(message)s"


class LogMode:
    file: str = "file"
    console: str = "console"


class Logger(Singleton):
    def __init__(self):
        self._logger: logging.Logger = logging.getLogger(LOG_NAME)
        self.log_mode: str = "console"
        self.log_level: str = "DEBUG"
        self.log_path: str = "server.log"

    def init_logger(self, log_config: Optional[LogConfig]):
        if log_config:
            self.log_path = log_config.log_path
            self.log_mode = log_config.log_mode
            self.log_level = log_config.log_level

        self._logger.setLevel(self.log_level)
        handler = None
        if self.log_mode == LogMode.file:
            up_path = os.path.abspath(os.path.dirname(os.path.dirname(__file__)))
            handler = logging.FileHandler(
                up_path + f"/{os.getpid()}-" + self.log_path, encoding="utf-8"
            )

        if self.log_mode == LogMode.console:
            handler = logging.StreamHandler(sys.stdout)

        if not handler:
            return

        handler.setLevel(level=self.log_level)
        log_format = logging.Formatter(LOG_FORMAT)
        handler.setFormatter(log_format)

        self._logger.addHandler(handler)

    @staticmethod
    def format_msg(msg):
        try:
            caller = sys._getframe(2)
            file_name = "/".join(caller.f_code.co_filename.split("/")[-2:])
            call_name, file_no = caller.f_code.co_name, caller.f_lineno
            return ":".join([file_name, call_name, str(file_no)]) + f" | {msg}"
        except Exception as e:
            return msg

    def debug(self, msg: str):
        msg = self.format_msg(msg)
        if self.log_mode == LogMode.console:
            msg = f"\033[34m{msg}\033[0m"
        self._logger.debug(msg)

    def info(self, msg: str):
        msg = self.format_msg(msg)
        if self.log_mode == LogMode.console:
            msg = f"\033[32m{msg}\033[0m"
        self._logger.info(msg)

    def warning(self, msg: str):
        msg = self.format_msg(msg)
        if self.log_mode == LogMode.console:
            msg = f"\033[33m{msg}\033[0m"
        self._logger.warning(msg)

    def error(self, msg: str):
        msg = self.format_msg(msg)
        if self.log_mode == LogMode.console:
            msg = f"\033[31m{msg}\033[0m"
        self._logger.error(msg, stack_info=True)


logger = Logger()


if __name__ == "__main__":
    logger.init_logger(None)
    logger.debug("debug")
    logger.info("info")
    logger.warning("warning")
    logger.error("error")
