from transformers import (
    AutoModelForQuestionAnswering,
    AutoTokenizer,
    AutoModelForSeq2SeqLM,
    TrainingArguments,
    Trainer,
    DefaultDataCollator,
)
import os
os.environ["CUDA_VISIBLE_DEVICES"] = ""
import torch
from datasets import load_dataset


def load_qa_model(model_name="deepset/roberta-base-squad2"):
    print("INFO: loading model roberta")
    tokenizer = AutoTokenizer.from_pretrained(model_name)
    model = AutoModelForQuestionAnswering.from_pretrained(model_name)
    return model, tokenizer


def load_summarization_model(model_name="facebook/bart-large-cnn"):
    print("INFO: loading model bart")
    tokenizer = AutoTokenizer.from_pretrained(model_name)
    model = AutoModelForSeq2SeqLM.from_pretrained(model_name)
    return model, tokenizer


def preprocess_squad(examples, tokenizer):
    print("INFO: entering preprocess phase")
    questions = [q.strip() for q in examples["question"]]
    inputs = tokenizer(
        questions,
        examples["context"],
        max_length=384,
        truncation="only_second",
        return_offsets_mapping=True,
        padding="max_length",
    )

    offset_mapping = inputs.pop("offset_mapping")
    answers = examples["answers"]
    start_positions = []
    end_positions = []

    for i, offset in enumerate(offset_mapping):
        answer = answers[i]
        start_char = answer["answer_start"][0]
        end_char = answer["answer_start"][0] + len(answer["text"][0])
        sequence_ids = inputs.sequence_ids(i)
        idx = 0
        while sequence_ids[idx] != 1:
            idx += 1
        context_start = idx
        while sequence_ids[idx] == 1:
            idx += 1
        context_end = idx - 1

        if offset[context_start][0] > end_char or offset[context_end][1] < start_char:
            start_positions.append(0)
            end_positions.append(0)
        else:
            idx = context_start
            while idx <= context_end and offset[idx][0] <= start_char:
                idx += 1
            start_positions.append(idx - 1)

            idx = context_end
            while idx >= context_start and offset[idx][1] >= end_char:
                idx -= 1
            end_positions.append(idx + 1)

    inputs["start_positions"] = start_positions
    inputs["end_positions"] = end_positions
    return inputs


def train_qa(model, tokenizer, train_size=10000, eval_size=5000, epochs=1):
    print("INFO: training question-answering model")
    squad = load_dataset("squad")
    tokenized_squad = squad.map(
        lambda examples: preprocess_squad(examples, tokenizer),
        batched=True,
        remove_columns=squad["train"].column_names,
    )

    train_dataset = tokenized_squad["train"].select(range(train_size))
    eval_dataset = tokenized_squad["validation"].select(range(eval_size))

    training_args = TrainingArguments(
        output_dir="./results",
        evaluation_strategy="epoch",
        learning_rate=2e-5,
        per_device_train_batch_size=16,
        per_device_eval_batch_size=16,
        num_train_epochs=epochs,
        weight_decay=0.01,
        fp16=True,
        report_to=[],
    )

    trainer = Trainer(
        model=model,
        args=training_args,
        train_dataset=train_dataset,
        eval_dataset=eval_dataset,
        tokenizer=tokenizer,
        data_collator=DefaultDataCollator(),
    )

    trainer.train()
    return model


def get_answer(qa_pipeline, question, context):
    qa_input = {"question": question, "context": context}
    return qa_pipeline(qa_input)


def summarize_text(summarizer, text, max_length=130, min_length=30):
    return summarizer(
        text, max_length=max_length, min_length=min_length, do_sample=False
    )[0]["summary_text"]


def buildAndSaveModels():
    qa_model, qa_tokenizer = load_qa_model()
    summarization_model, summarization_tokenizer = load_summarization_model()

    qa_model.save_pretrained("./Model/qa_model")
    qa_tokenizer.save_pretrained("./Model/qa_model")
    summarization_model.save_pretrained("./Model/summ_model")
    summarization_tokenizer.save_pretrained("./Model/summ_model")

    if os.path.exists("Model/qa_model"):
        print("INFO: question-answering model saved")
    if os.path.exists("Model/summ_model"):
        print("INFO: summary model saved")


buildAndSaveModels()
