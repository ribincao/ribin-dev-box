from common.config import get_opeai_api_key
from langchain.llms import OpenAI
from langchain.prompts import PromptTemplate
from langchain.chains import LLMChain


def llms_example(is_test: bool = False) -> OpenAI:
    llm = OpenAI(openai_api_key=get_opeai_api_key(), temperature=0.9)  # type: ignore
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


if __name__ == "__main__":
    llms_chain_example()
