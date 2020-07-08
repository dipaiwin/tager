import json
import logging
import re
import sys

import cv2
import numpy as np
import pytesseract
from aiohttp import web
from aiohttp.web_response import json_response


def setup_custom_logger(name):
	log = logging.getLogger(name)
	h = logging.StreamHandler(stream=sys.stdout)
	h.setFormatter(logging.Formatter('%(message)s'))
	h.flush = sys.stdout.flush
	log.addHandler(h)
	log.setLevel(logging.INFO)
	return log


async def model(request):
	post = await request.post()
	image = post.get("image")
	buf = image.file.read()
	nparr = np.fromstring(buf, np.uint8)
	img = cv2.imdecode(nparr, cv2.IMREAD_COLOR)
	img = cv2.resize(img, None, fx=3, fy=3, interpolation=cv2.INTER_CUBIC)
	img = cv2.medianBlur(img, 3)
	text = pytesseract.image_to_string(img, lang="rus", config=config)
	text = re.findall(r"\w+|\.", text)
	return json_response(body=json.dumps(' '.join(text), ensure_ascii=False))


if __name__ == '__main__':
	config = "--oem 1 --psm 4"
	logger = setup_custom_logger('root')
	app = web.Application()
	app.add_routes([web.post('/model', model)])
	web.run_app(app, port=80, access_log=logger)
