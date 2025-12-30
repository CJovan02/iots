FROM python:3.13

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY app ./app
COPY model_artifacts ./model_artifacts

EXPOSE 8080

CMD [ "fastapi", "run", "./app/main.py", "--port", "8080" ]
