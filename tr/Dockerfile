FROM python:3.7

RUN pip install aiohttp summa nltk pymorphy2

WORKDIR /srv

EXPOSE 80

ENTRYPOINT python -u srv.py