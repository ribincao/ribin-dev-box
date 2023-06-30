from common.config import get_opeai_api_key, get_serp_api_key
from langchain.agents import AgentType, initialize_agent, load_tools
from langchain.prompts import PromptTemplate
from langchain.chains import LLMChain
from langchain.llms import OpenAI
import os

os.environ["OPENAI_API_KEY"] = get_opeai_api_key()
os.environ["SERPAPI_API_KEY"] = get_serp_api_key()


def llms_example(temperature: float = 0.0, is_test: bool = False) -> OpenAI:
    llm = OpenAI(temperature=temperature)  # type: ignore
    if is_test:
        ret = llm.predict(
            "What would be a good company name for a company that makes colorful socks?"
        )
        print(ret)
    return llm


def llms_prompt_example(is_test: bool = False) -> PromptTemplate:
    prompt = PromptTemplate.from_template(
        "what is a good name for a company that makes {product}?"
    )
    if is_test:
        prompt.format(product="colorful socks")
    return prompt


def llms_chain_example() -> LLMChain:
    llm = llms_example()
    prompt = llms_prompt_example()
    chain = LLMChain(llm=llm, prompt=prompt)
    ret = chain.run("colorful socks")
    print(ret)
    return chain


def llms_agent_example():
    llm = llms_example()
    tools = load_tools(["serpapi", "llm-math"], llm=llm)
    agent = initialize_agent(
        tools, llm, agent=AgentType.ZERO_SHOT_REACT_DESCRIPTION, verbose=True
    )
    agent.run(
        "What was the high temperature in SF yesterday in Fahrenheit? What is that number rasied to the .023 power?"
    )


if __name__ == "__main__":
    llms_agent_example()
