from common.config import get_opeai_api_key
from langchain.chat_models import ChatOpenAI
from langchain.schema import AIMessage, HumanMessage, SystemMessage
from langchain.prompts.chat import (
    ChatPromptTemplate,
    SystemMessagePromptTemplate,
    HumanMessagePromptTemplate,
)
from langchain import LLMChain


def chat_model_example(is_test: bool = False) -> ChatOpenAI:
    chat = ChatOpenAI(temperature=0, openai_api_key=get_opeai_api_key())  # type: ignore
    if is_test:
        ret = chat.predict_messages(
            [
                HumanMessage(
                    content="Translate this sentence from English to Chinese. I love programming."
                )
            ]
        )
        print(ret.content)
    return chat


def chat_model_prompt_example(is_test: bool = False) -> ChatPromptTemplate:
    template = "You are a helpful assistant that translates {input_language} to {output_language}."
    system_message_tmeplate = SystemMessagePromptTemplate.from_template(template)
    human_template = "{text}"
    human_message_template = HumanMessagePromptTemplate.from_template(human_template)
    chat_prompt = ChatPromptTemplate.from_messages(
        [system_message_tmeplate, human_message_template]
    )
    if is_test:
        chat_prompt.format_messages(
            input_language="English",
            output_language="Chinese",
            text="I love programming.",
        )
    return chat_prompt


def chat_model_chain_example() -> LLMChain:
    chat = chat_model_example()
    prompt = chat_model_prompt_example()
    chain = LLMChain(llm=chat, prompt=prompt)
    ret = chain.run(
        input_language="English", output_language="Chinese", text="I love programming."
    )
    print(ret)
    return chain


if __name__ == "__main__":
    chat_model_chain_example()
