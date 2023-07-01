from common.singleton import Singleton


class ApiKeys:
    openai_api: str = ""
    serp_api: str = ""
    active_loop_api: str = ""


class ServiceConfig:
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
        self.service_config: ServiceConfig = ServiceConfig()
        self.log_config: LogConfig = LogConfig()
        self.api_keys: ApiKeys = ApiKeys()

    def load_config(self, path: str = "./config.yaml"):
        from common.utils import data_util

        data = data_util.load_from_yaml(path)
        if not data:
            return
        for k, d in data.items():
            if k == "ServiceConfig":
                self.service_config.__dict__.update(d)
            elif k == "LogConfig":
                self.log_config.__dict__.update(d)
            elif k == "ApiKeys":
                self.api_keys.__dict__.update(d)


global_config = GlobalConfig()


if __name__ == "__main__":
    global_config.load_config("../config.yaml")
    print(global_config.api_keys.openai_api)
    print(global_config.log_config.log_path)
