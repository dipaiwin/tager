import json
import logging
import os
import sys

from aiohttp import web
from aiohttp.web_response import json_response
from imageai.Detection import ObjectDetection
from googletrans import Translator
import numpy as np
import cv2

execution_path = os.getcwd()


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
	buf = post.get("image").file.read()
	np_arr = np.frombuffer(buf, np.uint8)
	img = cv2.imdecode(np_arr, cv2.IMREAD_COLOR)
	cv2.imwrite('/tmp/tmp.jpg', img)
	detections = detector.detectObjectsFromImage(input_image="/tmp/tmp.jpg", output_type="array")[1]
	res = list({translator.translate(eachObject["name"], dest='ru').text.lower() for eachObject in detections})
	return json_response(body=json.dumps(res, ensure_ascii=False))


if __name__ == '__main__':
	logger = setup_custom_logger('root')
	detector = ObjectDetection()
	detector.setModelTypeAsYOLOv3()
	detector.setModelPath(os.path.join(execution_path, "/model/yolo.h5"))
	detector.loadModel()
	translator = Translator()
	app = web.Application()
	app.add_routes([web.post('/model', model)])
	web.run_app(app, port=80, access_log=logger)
