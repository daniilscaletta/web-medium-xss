FROM python:3.9-slim

RUN apt update && apt install    -y \
    chromium \
    chromium-driver \
    libglib2.0-0 \
    libnss3 \
    libgdk-pixbuf2.0-0 \
    libxss1 \
    libxtst6 \
    libappindicator3-1 \
    libasound2 \
    fonts-liberation \
    xdg-utils \
    && rm -rf /var/lib/apt/lists/*

RUN pip install --no-cache-dir pipenv

WORKDIR /app

COPY . /app


RUN pipenv install --deploy --system

CMD ["python", "bot.py"]