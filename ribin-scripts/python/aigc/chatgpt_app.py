from common.config import get_opeai_api_key, get_serp_api_key
from common.utils import aprint
from langchain import OpenAI, ConversationChain, LLMChain, PromptTemplate
from langchain.memory import ConversationBufferWindowMemory
import os
from aigc.prompts import chatgpt_prompt, linux_prompt

os.environ["OPENAI_API_KEY"] = get_opeai_api_key()
os.environ["SERPAPI_API_KEY"] = get_serp_api_key()



def chatgpt(init_context: str = linux_prompt) -> LLMChain:
    prompt = PromptTemplate(
        input_variables=["history", "human_input"], template=chatgpt_prompt
    )
    chatgpt_chain = LLMChain(
        llm=OpenAI(temperature=0.0),  # type: ignore
        prompt=prompt,
        verbose=False,
        memory=ConversationBufferWindowMemory(k=2),
    )
    chatgpt_chain.predict(human_input=init_context)
    while True:
        text = input("Human: ")
        if not text or text == "q":
            break
        answer = chatgpt_chain.predict(human_input=text)
        aprint(answer)

    return chatgpt_chain


def linux_shell():
    chatgpt(linux_prompt)


if __name__ == "__main__":
    linux_shell()

