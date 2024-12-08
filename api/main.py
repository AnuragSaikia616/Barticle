from datetime import datetime
from fastapi import FastAPI
from pydantic import BaseModel
from transformers import (
    pipeline,
)
from .modelUtils import get_answer, summarize_text
from .scraper import extract

# Model directory
model_dir = "./Model/"

# Loading the summary model
summarizer = pipeline("summarization", model=model_dir + "summ_model")
# Loading the question anwering model
qa_pipeline = pipeline("question-answering", model=model_dir + "qa_model")


app = FastAPI()


# This class is for /summarize handler input
class SumTextInput(BaseModel):
    text: str


# This handler takes in text content and returns a summary as response
@app.post("/summarize/")
def summarize(input: SumTextInput):
    start_time = datetime.now()
    # getting summary
    summary = summarize_text(summarizer, input.text)
    end_time = datetime.now()
    time_delta = end_time - start_time
    print("INFO: ", "Summary took ", time_delta, " s")
    return {"summary": summary}


# This class is for /answer handler input
class qaTextInput(BaseModel):
    context: str
    question: str


# This handler texts in qaTextInput
# it returns answer as a response to the question based on the
# context provided
@app.post("/answer/")
async def answer(input: qaTextInput):
    start_time = datetime.now()
    # getting answer
    answer = get_answer(qa_pipeline, input.question, input.context)
    end_time = datetime.now()
    time_delta = end_time - start_time

    print("INFO: ", "answering took ", time_delta, " s")
    return answer


class URLinput(BaseModel):
    url: str


@app.post("/scrape/")
async def scrape(input: URLinput):
    url = input.url
    start_time = datetime.now()
    # Scratping website content
    scrapedContent = extract(url)
    end_time = datetime.now()
    time_delta = end_time - start_time
    print("INFO: ", "Scraping took ", time_delta, " s")
    return {"content": scrapedContent}
