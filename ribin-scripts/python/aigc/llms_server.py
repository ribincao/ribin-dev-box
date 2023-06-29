from common.config import global_config
from langchain.llms import OpenAI
from typing import Optional


def get_llm() -> Optional[OpenAI]:
    if not global_config.service_config:
        return None
    if not global_config.service_config.openai_api:
        return None
    return OpenAI(openai_api_key=global_config.service_config.openai_api)  # type: ignore


if __name__ == "__main__":
    global_config.load_config("./config.yaml")
    llm = get_llm()
    if llm:
        ret = llm.predict(
            "What would be a good company name for a company that makes colorful socks?"
        )
        print(ret)
