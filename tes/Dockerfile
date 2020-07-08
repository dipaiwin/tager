FROM python:3.7

RUN apt-get update && apt-get install tesseract-ocr -y

RUN pip3 install pytesseract opencv-python

COPY ./model/rus.traineddata /usr/share/tesseract-ocr/4.00/tessdata/

RUN pip install aiohttp

WORKDIR /srv

EXPOSE 80

ENTRYPOINT python -u srv.py