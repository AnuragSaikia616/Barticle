FROM python:3.9

WORKDIR /app

COPY ./requirements.txt .

RUN pip3 install --no-cache-dir -r requirements.txt

COPY . .

RUN python3 modelBuilder.py

EXPOSE 8000

CMD [ "fastapi","run","main.py", "--proxy-headers" ]
