FROM python:3.9-slip

RUN apt update && apt install    -y \
    chromium-driver \ 
    && rm -rf /var/lib/apt/lists/*

RUN pip install --no-cache-dir pipenv

WORKDIR /app

COPY . /app

RUN pipenv install --deploy --system

CMD ["python", "bot.py"]