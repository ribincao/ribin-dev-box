from common.singleton import Singleton
from typing import Optional


class ServiceConfig:
    openai_api: str = ""
    env: str = ""
    redis_addr: str = ""
    redis_username: str = ""
    redis_password: str = ""


class LogConfig:
    log_level: str = ""
    log_path: str = ""
    log_mode: str = ""


class GlobalConfig(Singleton):
    def __init__(self):
        self.service_config: Optional[ServiceConfig] = None
        self.log_config: Optional[LogConfig] = None

    def load_config(self, path: str = "./config.yaml"):
        from common.utils import data_util

        data = data_util.load_from_yaml(path)
        if not data:
            return
        for k, d in data.items():
            if k == "ServiceConfig":
                if not self.service_config:
                    self.service_config = ServiceConfig()
                self.service_config.__dict__.update(d)
            elif k == "LogConfig":
                if not self.log_config:
                    self.log_config = LogConfig()
                self.log_config.__dict__.update(d)


global_config = GlobalConfig()


if __name__ == "__main__":
    global_config.load_config("../config.yaml")
    if global_config.service_config:
        print(global_config.service_config.openai_api)
    if global_config.log_config:
        print(global_config.log_config.log_path)
