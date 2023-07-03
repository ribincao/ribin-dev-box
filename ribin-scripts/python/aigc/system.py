from common.config import global_config
from langchain.chat_models import ChatOpenAI
from langchain.prompts.chat import (
    ChatPromptTemplate,
    SystemMessagePromptTemplate,
    HumanMessagePromptTemplate,
)
from langchain.chains import LLMChain
import os
from common.utils import aprint
from typing import List
from langchain.callbacks.streaming_stdout import StreamingStdOutCallbackHandler

global_config.load_config()
os.environ["OPENAI_API_KEY"] = global_config.api_keys.openai_api
os.environ["SERPAPI_API_KEY"] = global_config.api_keys.serp_api


class ChatSystem(object):
    def __init__(self, system_template: str, human_template: str):
        self.system_template: str = system_template
        self.system_variables: List[str] = []
        self.human_template: str = human_template
        self.human_variables: List[str] = []

    def get_system_prompt(self) -> SystemMessagePromptTemplate:
        system_message_tmeplate = SystemMessagePromptTemplate.from_template(
            self.system_template
        )
        self.system_variables = system_message_tmeplate.input_variables
        return system_message_tmeplate

    def get_human_prompt(self) -> HumanMessagePromptTemplate:
        human_message_template = HumanMessagePromptTemplate.from_template(
            self.human_template
        )
        self.human_variables = human_message_template.input_variables
        return human_message_template

    def get_prompt(self) -> ChatPromptTemplate:
        system_message_tmeplate = self.get_system_prompt()
        human_message_template = self.get_human_prompt()
        chat_prompt = ChatPromptTemplate.from_messages(
            [system_message_tmeplate, human_message_template]
        )
        return chat_prompt

    def get_model(self, temperature: float = 0.0) -> ChatOpenAI:
        chat = ChatOpenAI(
            temperature=temperature, callbacks=[StreamingStdOutCallbackHandler()]
        )  # type: ignore
        return chat

    def get_coversation_chain(
        self, model: ChatOpenAI, prompt: ChatPromptTemplate
    ) -> LLMChain:
        conversation = LLMChain(llm=model, prompt=prompt)
        return conversation

    def run(self):
        prompt = self.get_prompt()
        kwargs = {}
        for k in self.system_variables:
            v = input(f"Please enter value for system variable [{k}]: ")
            kwargs[k] = v

        model = self.get_model()
        conversation = self.get_coversation_chain(model, prompt)
        while True:
            for k in self.human_variables:
                v = input(f"Please enter value for human variable [{k}]: ")
                kwargs[k] = v
            answer = conversation.run(**kwargs)
            aprint(f"{answer}")


if __name__ == "__main__":
    system_template = "You are a helpful assistant that translates {input_language} to {output_language}."
    human_template = "{text}"
    translater = ChatSystem(system_template, human_template)
    translater.run()
