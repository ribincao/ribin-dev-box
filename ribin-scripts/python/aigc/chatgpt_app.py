from common.config import get_opeai_api_key, get_serp_api_key
from common.utils import aprint
from langchain import OpenAI, ConversationChain, LLMChain, PromptTemplate
from langchain.memory import ConversationBufferWindowMemory
import os

os.environ["OPENAI_API_KEY"] = get_opeai_api_key()
os.environ["SERPAPI_API_KEY"] = get_serp_api_key()

chatgpt_template = """Assistant is a large language model trained by OpenAI.
Assistant is designed to be able to assist with a wide range of tasks, from answering simple questions to providing in-depth explanations and discussions on a wide range of topics. As a language model, Assistant is able to generate human-like text based on the input it receives, allowing it to engage in natural-sounding conversations and provide responses that are coherent and relevant to the topic at hand.
Assistant is constantly learning and improving, and its capabilities are constantly evolving. It is able to process and understand large amounts of text, and can use this knowledge to provide accurate and informative responses to a wide range of questions. Additionally, Assistant is able to generate its own text based on the input it receives, allowing it to engage in discussions and provide explanations and descriptions on a wide range of topics.
Overall, Assistant is a powerful tool that can help with a wide range of tasks and provide valuable insights and information on a wide range of topics. Whether you need help with a specific question or just want to have a conversation about a particular topic, Assistant is here to assist.
{history}
Human: {human_input}
Assistant:"""


def chatgpt(text, temperature: float = 0.0) -> LLMChain:
    prompt = PromptTemplate(
        input_variables=["history", "human_input"], 
        template=chatgpt_template
    )
    chatgpt_chain = LLMChain(
        llm=OpenAI(temperature=temperature),  # type: ignore
        prompt=prompt,
        verbose=False,
        memory=ConversationBufferWindowMemory(k=2),
    )
    chatgpt_chain.predict(human_input=text)
    while True:
        text = input("Human: ")
        if not text or text == "q":
            break
        answer = chatgpt_chain.predict(human_input=text)
        aprint(answer)

    return chatgpt_chain


def linux_shell():
    template = "I want you to act as a Linux terminal. I will type commands and you will reply with what the terminal should show. I want you to only reply with the terminal output inside one unique code block, and nothing else. Do not write explanations. Do not type commands unless I instruct you to do so. When I need to tell you something in English I will do so by putting text inside curly brackets {like this}. My first command is pwd."
    chatgpt(template)


def suanming_app():
    template = "你是一名精通八卦、占星和玄学的命理大师,我将称呼你大师,我会问你一些命理相关的问题,希望你基于自己的专业回答我"
    chatgpt(template, temperature=0.9)

if __name__ == "__main__":
    linux_shell()
    # suanming_app()

