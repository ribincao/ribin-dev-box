from common.config import global_config
from common.logger import logger
from common.utils import aprint
from typing import List
from langchain.chat_models import ChatOpenAI
from langchain.vectorstores import DeepLake
from langchain.schema import Document
from langchain.chains import ConversationalRetrievalChain
from langchain.document_loaders import TextLoader
from langchain.text_splitter import CharacterTextSplitter
from langchain.schema import Document
from langchain.embeddings import OpenAIEmbeddings
import os

global_config.load_config()
logger.init_logger(global_config.log_config)
os.environ["OPENAI_API_KEY"] = global_config.api_keys.openai_api
os.environ["SERPAPI_API_KEY"] = global_config.api_keys.serp_api
os.environ["ACTIVELOOP_TOKEN"] = global_config.api_keys.active_loop_api
DEEP_LAKE_DATASET_PATH = "hub://ribincao/langchain-code"


def filter(x):
    # filter based on source code
    if "something" in x["text"].data()["value"]:
        return False

    # filter based on path e.g. extension
    metadata = x["metadata"].data()["value"]
    return "only_this" in metadata["source"] or "also_that" in metadata["source"]


class CodeAssistant(object):
    def __init__(self):
        pass

    def load_code(self, path: str, file_suffix: str = ".py") -> List[Document]:
        docs: List[Document] = []
        for dir_path, _, file_names in os.walk(path):
            for file_name in file_names:
                if not file_name.endswith(file_suffix):
                    continue
                try:
                    loader = TextLoader(
                        os.path.join(dir_path, file_name), encoding="utf-8"
                    )
                    docs.extend(loader.load_and_split())
                except Exception as error:
                    logger.error(f"load code error {error}")
        return docs

    def split_documents(self, documents: List[Document]) -> List[Document]:
        text_splitter = CharacterTextSplitter(chunk_size=1000, chunk_overlap=0)
        texts = text_splitter.split_documents(documents)
        return texts

    def get_embeddings(self) -> OpenAIEmbeddings:
        return OpenAIEmbeddings()  # type: ignore

    def upload(
        self,
        documents: List[Document],
        embeddings: OpenAIEmbeddings,
        dataset_path: str = DEEP_LAKE_DATASET_PATH,
    ) -> DeepLake:
        db = DeepLake.from_documents(
            documents=documents, embedding=embeddings, dataset_path=dataset_path
        )
        return db

    def get_retriever(
        self,
        embeddings: OpenAIEmbeddings,
        dataset_path: str = DEEP_LAKE_DATASET_PATH,
        is_filter: bool = False,
    ):
        db = DeepLake(
            dataset_path=dataset_path, read_only=True, embedding_function=embeddings
        )
        retriever = db.as_retriever()
        retriever.search_kwargs["distance_metric"] = "cos"
        retriever.search_kwargs["fetch_k"] = 20
        retriever.search_kwargs["maximal_marginal_relevance"] = True
        retriever.search_kwargs["k"] = 20
        if is_filter:
            retriever.search_kwargs["filter"] = filter
        return retriever

    def get_model(self) -> ChatOpenAI:
        return ChatOpenAI(model_name="gpt-3.5-turbo")  # type: ignore

    def get_conversation(self, model: ChatOpenAI, retriever):
        return ConversationalRetrievalChain.from_llm(llm=model, retriever=retriever)

    def run(self, code_root_path: str, dataset_path: str):
        docs = self.load_code(code_root_path)
        logger.info(f"step1. load code finished. {len(docs)}")
        docs = self.split_documents(docs)
        logger.info(f"step2. split code finished. {len(docs)}")
        embeddings = self.get_embeddings()
        self.upload(docs, embeddings, dataset_path)
        logger.info(f"step3. upload finished. {dataset_path}")

        model = self.get_model()
        retriever = self.get_retriever(embeddings, dataset_path, is_filter=False)
        conversion = self.get_conversation(model, retriever)
        logger.info(f"step4. get conversion finished.")
        chat_history = []
        while True:
            question = input("Q: ")
            if not question:
                break
            result = conversion(
                {"question": question, "chat_history": chat_history},
                return_only_outputs=True,
            )
            answer = result.get("answer", "ERROR")
            logger.info(f"Question: {question}, Answer: {answer}")


if __name__ == "__main__":
    code_path = "/Users/ribincao/Desktop/ribin-workspace/ribin-py-2dgame"
    code_ai = CodeAssistant()
    code_ai.run(code_path, DEEP_LAKE_DATASET_PATH)
    pass
