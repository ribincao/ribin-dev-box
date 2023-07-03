from common.utils import aprint
from typing import List, Callable, Optional
from langchain.chat_models import ChatOpenAI
from langchain.vectorstores import DeepLake
from langchain.schema import Document
from langchain.chains import ConversationalRetrievalChain
from langchain.text_splitter import CharacterTextSplitter
from langchain.schema import Document
from langchain.embeddings import OpenAIEmbeddings


class Assistant(object):
    def __init__(self):
        pass

    def load_documents(self, handler, *args, **kwargs) -> List[Document]:
        documents = handler(*args, **kwargs)
        aprint(f"load documents file_count {len(documents)} finished.")
        return documents

    def split_documents(self, documents: List[Document]) -> List[Document]:
        text_splitter = CharacterTextSplitter(chunk_size=1000, chunk_overlap=0)
        texts = text_splitter.split_documents(documents)
        aprint(f"split code count {len(texts)} finished.")
        return texts

    def get_embeddings(self) -> OpenAIEmbeddings:
        embeddings = OpenAIEmbeddings()  # type: ignore
        aprint(f"get embeddings {embeddings.__class__.__name__} finished.")
        return embeddings

    def upload(
        self,
        documents: List[Document],
        embeddings: OpenAIEmbeddings,
        dataset_path,
    ) -> DeepLake:
        db = DeepLake.from_documents(
            documents=documents, embedding=embeddings, dataset_path=dataset_path
        )
        aprint(f"upload {db.__class__.__name__} finished.")
        return db

    def get_retriever(
        self,
        embeddings: OpenAIEmbeddings,
        dataset_path: str,
        filter: Optional[Callable],
    ):
        db = DeepLake(
            dataset_path=dataset_path, read_only=True, embedding_function=embeddings
        )
        retriever = db.as_retriever()
        retriever.search_kwargs["distance_metric"] = "cos"
        retriever.search_kwargs["fetch_k"] = 20
        retriever.search_kwargs["maximal_marginal_relevance"] = True
        retriever.search_kwargs["k"] = 20
        if filter:
            retriever.search_kwargs["filter"] = filter
        aprint(f"get retriever {retriever.__class__.__name__} finished.")
        return retriever

    def get_model(self) -> ChatOpenAI:
        model = ChatOpenAI(model_name="gpt-3.5-turbo")  # type: ignore
        aprint(f"get model {model.__class__.__name__} finished.")
        return model

    def get_conversation(self, model: ChatOpenAI, retriever):
        conversation = ConversationalRetrievalChain.from_llm(
            llm=model, retriever=retriever
        )
        aprint(f"get conversion {conversation.__class__.__name__} finished.")
        return conversation
