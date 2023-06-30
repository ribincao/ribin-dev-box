from common.singleton import Singleton
from typing import Optional
from rich.console import Console


class DataUtil(Singleton):
    def __init__(self):
        pass

    @staticmethod
    def load_from_yaml(path: str) -> Optional[dict]:
        with open(path, "r", encoding="utf-8") as f:
            import yaml

            data = yaml.safe_load(f)
        return data

    @staticmethod
    def load_from_json(path: str) -> Optional[dict]:
        with open(path, "r", encoding="utf-8") as f:
            import json

            data = json.load(f)
        return data


data_util = DataUtil()
console = Console()


def aprint(msg: object, color: str = "green"):
    console.log(f"[{color}] {msg}")
