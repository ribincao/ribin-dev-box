from common.config import global_config
from langchain.vectorstores import Chroma
from langchain.agents import AgentType, initialize_agent, load_tools, AgentExecutor
from langchain.prompts import PromptTemplate
from langchain.chains import LLMChain, SimpleSequentialChain
from langchain.llms import OpenAI
from langchain import ConversationChain
from langchain.embeddings.sentence_transformer import SentenceTransformerEmbeddings
from common.utils import aprint
from langchain.callbacks.streaming_stdout import StreamingStdOutCallbackHandler
from langchain.llms.loading import load_llm
from langchain.chains import load_chain
from langchain.prompts import load_prompt
from langchain.callbacks import get_openai_callback
import os

global_config.load_config()
os.environ["OPENAI_API_KEY"] = global_config.api_keys.openai_api
os.environ["SERPAPI_API_KEY"] = global_config.api_keys.serp_api
os.environ['ACTIVELOOP_TOKEN'] = global_config.api_keys.active_loop_api


def llms_example(temperature: float = 0.0, is_test: bool = False) -> OpenAI:
    llm = OpenAI(
            temperature=temperature,
            callbacks=[StreamingStdOutCallbackHandler()]
            )  # type: ignore
    if is_test:
        with get_openai_callback() as call_back:
            ret = llm.predict(
                "What would be a good company name for a company that makes colorful socks?"
            )
            aprint(ret)
            aprint(call_back)
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


def llms_load_document_example():
    from langchain.document_loaders import JSONLoader
    from pprint import pprint
    loader = JSONLoader(
        file_path='./aigc/data/llm_example.json',
        jq_schema='.messages[].content'
    )
    data = loader.load()
    aprint(data)
    return data

def llms_split_example(document):
    from langchain.text_splitter import CharacterTextSplitter
    text_splitter = CharacterTextSplitter(
        chunk_size=100, 
        chunk_overlap=0
    )
    text = text_splitter.split_documents(document)
    aprint(text)

def llms_embedding_example() -> SentenceTransformerEmbeddings:
    embedding_function = SentenceTransformerEmbeddings(model_name="all-MiniLM-L6-v2")
    return embedding_function

def llms_chroma_example(docs, embedding):
    db = Chroma.from_documents(docs, embedding, persist_directory="./aigc/chroma/llm_example")
    # query = "What did the president say about Ketanji Brown Jackson"
    # docs = db.similarity_search(query)




if __name__ == "__main__":
    # llms_example(is_test=True)
    # llms_memory_example()
    # llms_prompt_example(True)
    # llms_seq_chain_example()
    # llms_chain_example()
    doc = llms_load_document_example()
    data = llms_split_example(doc)
    embedding = llms_embedding_example()
    llms_chroma_example(data, embedding)
