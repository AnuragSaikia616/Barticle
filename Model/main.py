from fastapi import FastAPI
from pydantic import BaseModel
from transformers import (
    pipeline,
)

qa_pipeline = pipeline("question-answering", model="qa_model")
summarizer = pipeline("summarization", model="summ_model")

app = FastAPI()


class SumTextInput(BaseModel):
    text: str


class qaTextInput(BaseModel):
    context: str
    question: str


@app.post("/summarize/")
def summarize(input: SumTextInput):
    summary = summarize_text(summarizer, input.text)
    return {"summary": summary}


@app.post("/answer/")
async def answer(input: qaTextInput):
    answer = get_answer(qa_pipeline, input.question, input.context)
    return {"answer": answer}


def summarize_text(summarizer, text, max_length=130, min_length=100):
    return summarizer(
        text, max_length=max_length, min_length=min_length, do_sample=False
    )[0]["summary_text"]


def get_answer(qa_pipeline, question, context):
    qa_input = {"question": question, "context": context}
    return qa_pipeline(qa_input)
