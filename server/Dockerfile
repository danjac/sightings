FROM python:3
ENV PYTHONUNBUFFERED 1
RUN apt-get update && apt-get install -y postgresql-client
ADD . /api
WORKDIR /api
RUN pip install -r requirements.txt
