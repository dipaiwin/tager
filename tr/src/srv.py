import json
import logging
import sys

import pymorphy2
from aiohttp import web
from aiohttp.web_response import json_response
from summa import keywords

import nltk
from nltk.corpus import stopwords

nltk.download("stopwords")
russian_stopwords = stopwords.words("russian")


def setup_custom_logger(name):
	log = logging.getLogger(name)
	h = logging.StreamHandler(stream=sys.stdout)
	h.setFormatter(logging.Formatter('%(message)s'))
	h.flush = sys.stdout.flush
	log.addHandler(h)
	log.setLevel(logging.INFO)
	return log


async def model(request):
	inp = await request.json()
	text = inp["x"]
	words = keywords.keywords(text[0], language="russian", additional_stopwords=stopwords.words("russian"), split=False)
	res = set()
	for word in words.split():
		morph_word = morph.parse(word)[0]
		if morph_word.tag.POS in ('NOUN', 'ADJF', 'ADJS', 'VERB', 'INFN', 'PRTF', 'PRTS', 'GRND', 'NUMR'):
			res.add(morph_word.normal_form)
	return json_response(body=json.dumps(list(res), ensure_ascii=False))


if __name__ == '__main__':
	morph = pymorphy2.MorphAnalyzer()
	logger = setup_custom_logger('root')
	app = web.Application()
	app.add_routes([web.post('/model', model)])
	web.run_app(app, port=80, access_log=logger)
