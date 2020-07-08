FROM python:3.5

RUN pip3 install --upgrade pip setuptools

RUN pip3 install numpy scipy opencv-python pillow matplotlib h5py keras jsonpickle requests https://github.com/OlafenwaMoses/ImageAI/releases/download/2.0.2/imageai-2.0.2-py3-none-any.whl

RUN pip install tensorflow==1.13.1

RUN pip install aiohttp

RUN pip uninstall -y pillow && pip install pillow==6.0

RUN pip install googletrans

WORKDIR /srv

EXPOSE 80

ENTRYPOINT python -u srv.py