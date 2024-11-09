import time
from fastapi import FastAPI
from pydantic import BaseModel
from transformers import (
    pipeline,
)

qa_pipeline = pipeline("question-answering", model="./Model/qa_model")
summarizer = pipeline("summarization", model="./Model/summ_model")

app = FastAPI()


class SumTextInput(BaseModel):
    text: str


class qaTextInput(BaseModel):
    context: str
    question: str


@app.post("/summarize/")
async def summarize(input: SumTextInput):
    start_time = time.time()
    summary = summarize_text(summarizer, input.text)
    end_time = time.time()
    print(f"SUMMARY_TIME: {end_time - start_time}")
    return {"summary": summary}


@app.post("/answer/")
async def answer(input: qaTextInput):
    answer = get_answer(qa_pipeline, input.question, input.context)
    return {"answer": answer}


def summarize_text(summarizer, text, max_length=130, min_length=30):
    return summarizer(
        text, max_length=max_length, min_length=min_length, do_sample=False
    )[0]["summary_text"]


def get_answer(qa_pipeline, question, context):
    qa_input = {"question": question, "context": context}
    return qa_pipeline(qa_input)
