from typing import List
from langchain.vectorstores import Chroma
from langchain.agents import AgentType, initialize_agent, load_tools, AgentExecutor
from langchain.prompts import PromptTemplate
from langchain.chains import LLMChain, SimpleSequentialChain
from langchain.llms import OpenAI
from langchain.embeddings import OpenAIEmbeddings
from langchain.callbacks import get_openai_callback
from langchain.schema import Document
from langchain.chains import RetrievalQAWithSourcesChain
import os
from common.config import global_config

global_config.load_config()
os.environ["OPENAI_API_KEY"] = global_config.api_keys.openai_api
os.environ["SERPAPI_API_KEY"] = global_config.api_keys.serp_api
os.environ["ACTIVELOOP_TOKEN"] = global_config.api_keys.active_loop_api


def llms_example(temperature: float = 0.9, is_test: bool = False) -> OpenAI:
    llm = OpenAI(temperature=temperature)  # type: ignore
    if is_test:
        with get_openai_callback() as call_back:
            ret = llm.predict("我姓曹，我的儿子起什么名字比较好?")
            print(ret)
            print()
            print(call_back)
    # llm.save("./aigc/models/llm_example.json")
    return llm


def llms_prompt_example(is_test: bool = False) -> PromptTemplate:
    prompt = PromptTemplate.from_template("我姓{last_name},我的儿子起什么名字比较好?")
    if is_test:
        with get_openai_callback() as call_back:
            llm = llms_example()
            message = prompt.format(last_name="曹")
            print(llm.predict(message))
            print()
            print(call_back)
    return prompt


def llms_chain_example() -> LLMChain:
    llm = llms_example()
    prompt = llms_prompt_example()
    chain = LLMChain(llm=llm, prompt=prompt)
    ret = chain.run("曹")
    print(ret)
    # chain.save("./aigc/chains/llm_example.json")
    return chain


def llms_seq_chain_example():
    llm = llms_example(temperature=0.9)
    first_prompt = PromptTemplate(
        template="我姓{last_name}, 生了个儿子，帮我的儿子起个名字", input_variables=["last_name"]
    )
    first_chain = LLMChain(llm=llm, prompt=first_prompt)
    second_prompt = PromptTemplate(
        template="我的儿子名字叫{name}, 给他起个小名", input_variables=["name"]
    )
    second_chain = LLMChain(llm=llm, prompt=second_prompt)
    chain = SimpleSequentialChain(chains=[first_chain, second_chain], verbose=True)
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
    while True:
        question = input("Q: ")
        if not question:
            break
        agent.run(question)
    return agent


def llms_load_document_example(
    file_path: str = "./aigc/data/llm_example2.txt",
) -> List[Document]:
    from langchain.document_loaders import TextLoader

    loader = TextLoader(file_path=file_path)
    data = loader.load()
    return data


def llms_split_example(document: List[Document]) -> List[Document]:
    from langchain.text_splitter import CharacterTextSplitter

    text_splitter = CharacterTextSplitter(chunk_size=1000, chunk_overlap=0)
    text = text_splitter.split_documents(document)
    return text


def llms_embedding_example() -> RetrievalQAWithSourcesChain:
    docs = llms_load_document_example()
    docs = llms_split_example(docs)

    embeddings = OpenAIEmbeddings()  # type: ignore
    db = Chroma.from_documents(
        docs,
        embeddings,
        metadatas=[{"source": str(i)} for i in range(len(docs))],
        # persist_directory="./aigc/chroma/llm_example",
    )

    llm = llms_example()
    retriever = db.as_retriever()
    retriever.search_kwargs["distance_metric"] = "cos"
    retriever.search_kwargs["fetch_k"] = 100
    retriever.search_kwargs["maximal_marginal_relevance"] = True
    retriever.search_kwargs["k"] = 10

    chain = RetrievalQAWithSourcesChain.from_chain_type(
        llm, chain_type="stuff", retriever=retriever
    )
    while True:
        question = input("Q: ")
        if not question:
            break
        answer = chain({"question": question}, return_only_outputs=True)
        print(answer["answer"])
    return chain  # type: ignore


if __name__ == "__main__":
    # llms_example(is_test=True)
    # llms_prompt_example(is_test=True)
    # llms_chain_example()
    # llms_seq_chain_example()
    # llms_agent_example()
    llms_embedding_example()
