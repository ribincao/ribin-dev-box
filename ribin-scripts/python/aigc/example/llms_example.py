from common.config import global_config
from langchain.agents import AgentType, initialize_agent, load_tools, AgentExecutor
from langchain.prompts import PromptTemplate
from langchain.chains import LLMChain, SimpleSequentialChain
from langchain.llms import OpenAI
from langchain import ConversationChain
import os
from common.utils import aprint
from langchain.callbacks.streaming_stdout import StreamingStdOutCallbackHandler
from langchain.llms.loading import load_llm
from langchain.chains import load_chain
from langchain.prompts import load_prompt

global_config.load_config()
os.environ["OPENAI_API_KEY"] = global_config.api_keys.openai_api
os.environ["SERPAPI_API_KEY"] = global_config.api_keys.serp_api


def llms_example(temperature: float = 0.0, is_test: bool = False) -> OpenAI:
    llm = OpenAI(
            temperature=temperature,
            callbacks=[StreamingStdOutCallbackHandler()]
            )  # type: ignore
    if is_test:
        ret = llm.predict(
            "What would be a good company name for a company that makes colorful socks?"
        )
        aprint(ret)
    llm.save("./aigc/models/llm_example.json")
    return llm


def llms_prompt_example(is_test: bool = False) -> PromptTemplate:
    prompt = PromptTemplate.from_template(
        "what is a good name for a company that makes {product}?"
    )
    if is_test:
        prompt: PromptTemplate = load_prompt("./aigc/prompts/llm_example.json")  # type: ignore
        llm = llms_example()
        aprint(llm(prompt.format(product="colorful socks")))
    return prompt  


def llms_chain_example() -> LLMChain:
    llm = llms_example()
    prompt = llms_prompt_example()
    chain = LLMChain(llm=llm, prompt=prompt)
    ret = chain.run("colorful socks")
    aprint(ret)
    chain.save("./aigc/chains/llm_example.json")
    return chain

def llms_seq_chain_example():
    llm = llms_example(temperature=0.9)
    first_prompt = PromptTemplate(
            template="我姓{last_name}, 生了个儿子，帮我的儿子起个名字",
            input_variables=["last_name"]
            )
    first_chain = LLMChain(llm=llm, prompt=first_prompt)
    second_prompt = PromptTemplate(
            template="我的儿子名字叫{name}, 给他起个小名",
            input_variables=["name"]
            )
    secone_chain = LLMChain(llm=llm, prompt=second_prompt)
    chain = SimpleSequentialChain(chains=[first_chain, secone_chain], verbose=True)
    while True:
        last_name = input("我姓: ")
        if not last_name:
            break
        chain.run(last_name)


def llms_agent_example() -> AgentExecutor:
    llm = llms_example()
    tools = load_tools(["serpapi", "llm-math"], llm=llm)
    agent = initialize_agent(
        tools, llm, agent=AgentType.ZERO_SHOT_REACT_DESCRIPTION, verbose=True
    )
    agent.run(
        "What was the high temperature in SF yesterday in Fahrenheit? What is that number rasied to the .023 power?"
    )
    return agent


def llms_memory_example(is_test: bool = False) -> ConversationChain:
    llm = llms_example()
    conversion = ConversationChain(llm=llm, verbose=False)
    if is_test:
        aprint(conversion.run("Hi there!"))
        aprint(conversion.run("I'm doing well! Just having a conversion with an AI."))
        aprint(conversion.run("Bye!"))
    else:
        while True:
            text = input("Human: ")
            if not text or text == "q":
                break
            answer = conversion.run(text)
            aprint("AI: " + answer)

    return conversion


if __name__ == "__main__":
    # llms_memory_example()
    llms_prompt_example(True)
    # llms_seq_chain_example()
    # llms_chain_example()
