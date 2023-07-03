from common.config import global_config
from common.utils import aprint
from aigc.understanding import Assistant
from langchain.schema import Document
from typing import List
from langchain.document_loaders import TextLoader
import os

global_config.load_config()
os.environ["OPENAI_API_KEY"] = global_config.api_keys.openai_api
os.environ["SERPAPI_API_KEY"] = global_config.api_keys.serp_api
os.environ["ACTIVELOOP_TOKEN"] = global_config.api_keys.active_loop_api
DEEP_LAKE_DATASET_PATH = "hub://ribincao/langchain-code"


def load_code(path: str, file_suffix: str = ".py") -> List[Document]:
    docs: List[Document] = []
    for dir_path, _, file_names in os.walk(path):
        for file_name in file_names:
            if not file_name.endswith(file_suffix):
                continue
            try:
                loader = TextLoader(os.path.join(dir_path, file_name), encoding="utf-8")
                docs.extend(loader.load_and_split())
            except Exception as error:
                aprint(f"[ERROR] load code error {error}")
    return docs


def filter(x):
    # filter based on source code
    if "something" in x["text"].data()["value"]:
        return False

    # filter based on path e.g. extension
    metadata = x["metadata"].data()["value"]
    return "only_this" in metadata["source"] or "also_that" in metadata["source"]


class CodeAssistant(Assistant):
    def __init__(self):
        pass

    def run(self, code_root_path: str, dataset_path: str):
        docs = self.load_documents(load_code, code_root_path, ".py")
        docs = self.split_documents(docs)
        embeddings = self.get_embeddings()

        self.upload(docs, embeddings, dataset_path)

        model = self.get_model()
        retriever = self.get_retriever(embeddings, dataset_path, filter)
        conversion = self.get_conversation(model, retriever)
        chat_history = []
        while True:
            question = input("Question: ")
            if not question:
                break
            result = conversion(
                {"question": question, "chat_history": chat_history},
                return_only_outputs=True,
            )
            answer = result.get("answer", "ERROR")
            aprint(f"{answer}")


if __name__ == "__main__":
    code_path = "/Users/ribincao/Desktop/ribin-workspace/ribin-py-2dgame"
    code_ai = CodeAssistant()
    code_ai.run(code_path, DEEP_LAKE_DATASET_PATH)
