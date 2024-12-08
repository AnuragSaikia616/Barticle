def summarize_text(summarizer, text, max_length=130, min_length=100):
    return summarizer(
        text, max_length=max_length, min_length=min_length, do_sample=False
    )[0]["summary_text"]


def get_answer(qa_pipeline, question, context):
    qa_input = {"question": question, "context": context}
    return qa_pipeline(qa_input)
