FROM tensorflow/tensorflow:1.15.2-py3

RUN pip install deeppavlov

RUN apt-get update
RUN apt-get install -y git

RUN python -m deeppavlov install ner_ontonotes_bert_mult

EXPOSE 10010

ENTRYPOINT python -m deeppavlov riseapi ner_ontonotes_bert_mult -p 10010
