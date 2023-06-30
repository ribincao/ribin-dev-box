from common.config import get_opeai_api_key, get_serp_api_key
from langchain.chat_models import ChatOpenAI
from langchain.schema import AIMessage, HumanMessage, SystemMessage
from langchain.agents import AgentType, initialize_agent, load_tools
from langchain.prompts.chat import (
    ChatPromptTemplate,
    SystemMessagePromptTemplate,
    HumanMessagePromptTemplate,
)
from langchain import LLMChain
import os

os.environ["OPENAI_API_KEY"] = get_opeai_api_key()
os.environ["SERPAPI_API_KEY"] = get_serp_api_key()


def chat_model_example(temperature: float = 0.0, is_test: bool = False) -> ChatOpenAI:
    chat = ChatOpenAI(temperature=temperature)  # type: ignore
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


def chat_model_agent_example():
    chat = chat_model_example(temperature=0.9)
    from aigc.llms_server import llms_example

    llm = llms_example()
    tools = load_tools(["serpapi", "llm-math"], llm=llm)
    agent = initialize_agent(
        tools, chat, agent=AgentType.CHAT_ZERO_SHOT_REACT_DESCRIPTION, verbose=True
    )
    agent.run(
        "Who is Olivia Wilde's boyfriend? What is his current age raised to the 0.23 power?"
    )


if __name__ == "__main__":
    chat_model_agent_example()
