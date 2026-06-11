import logging
import os

os.environ["CUDA_VISIBLE_DEVICES"] = ""
from datetime import datetime

from fastapi import FastAPI
from pydantic import BaseModel
from transformers import pipeline

from .modelUtils import get_answer, summarize_text
from .scraper import extract


logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S",
)
logger = logging.getLogger(__name__)

# Resolve model directory to an absolute path so transformers doesn't
# mistake "./Model/" for a HuggingFace repo ID.
MODEL_DIR = os.path.abspath("./Model/")

app = FastAPI()

# ---------------------------------------------------------------------------
# Lazy model loading – models are downloaded and cached on first use so the
# server starts immediately (useful when models aren't downloaded yet).
# ---------------------------------------------------------------------------
_summarizer = None
_qa_pipeline = None


def _load_summarizer():
    global _summarizer
    if _summarizer is not None:
        return _summarizer
    path = os.path.join(MODEL_DIR, "summ_model")
    logger.info("Loading summarization model from %s …", path)
    _summarizer = pipeline("summarization", model=path,device=1)
    logger.info("Summarization model loaded")
    return _summarizer


def _load_qa_pipeline():
    global _qa_pipeline
    if _qa_pipeline is not None:
        return _qa_pipeline
    path = os.path.join(MODEL_DIR, "qa_model")
    logger.info("Loading QA model from %s …", path)
    _qa_pipeline = pipeline("question-answering", model=path,device=1)
    logger.info("QA model loaded")
    return _qa_pipeline


# ---------------------------------------------------------------------------
# Request / response models
# ---------------------------------------------------------------------------

class SumTextInput(BaseModel):
    text: str


class qaTextInput(BaseModel):
    context: str
    question: str


class URLinput(BaseModel):
    url: str


# ---------------------------------------------------------------------------
# Endpoints
# ---------------------------------------------------------------------------

@app.post("/summarize/")
def summarize(input: SumTextInput):
    start_time = datetime.now()
    summarizer = _load_summarizer()
    summary = summarize_text(summarizer, input.text)
    elapsed = datetime.now() - start_time
    logger.info("Summary took %s s", elapsed)
    return {"summary": summary}


@app.post("/answer/")
async def answer(input: qaTextInput):
    start_time = datetime.now()
    qa = _load_qa_pipeline()
    result = get_answer(qa, input.question, input.context)
    elapsed = datetime.now() - start_time
    logger.info("Answering took %s s", elapsed)
    return result


@app.post("/scrape/")
async def scrape(input: URLinput):
    start_time = datetime.now()
    scraped_content = extract(input.url)
    elapsed = datetime.now() - start_time
    logger.info("Scraping took %s s", elapsed)
    return {"content": scraped_content}
