include .env

test_data=`pwd`/test_data
dir_iai=$(test_data)/iai
dir_tes=$(test_data)/tes

down:
	docker-compose $@
logs:
	docker-compose $@ -f
S=
up:
	docker-compose $@ -d $(S)
restart:
	docker-compose $@ $(S)
test_iai:
	echo "IAI TEST"
	for f in $(shell ls ${dir_iai}); do \
		curl -i -X POST -H "Content-Type: multipart/form-data" -F "image=@$(dir_iai)/$${f}" http://localhost:$(PORT_IAI)/model; \
	done

test_ner:
	@curl -d '{"x":["Санкт-Петербург – русский портовый город на побережье Балтийского моря, который в течение двух веков служил столицей Российской империи."]}'\
	 -H "Content-Type: application/json" -X POST http://localhost:$(PORT_NER)/model

test_tes:
	echo "TEST TES"
	for f in $(shell ls ${dir_tes}); do \
			curl -i -X POST -H "Content-Type: multipart/form-data" -F "image=@$(dir_tes)/$${f}" http://localhost:$(PORT_TES)/model; \
		done

test_tr:
	echo "TEST TR"
	curl -d '{"x":["Делюсь с вами рецептом очень быстрого,вкусного и полезного завтрака🤤 «ТОСТЫ С ГУАКАМОЛЕ И ЯЙЦОМ-ПАШОТ»😋 Ингредиенты:Для гуакамоле:Спелый авокадо-1шт Помидор-1шт Зелень (кинза) Лимонный сок.Соль,перец,паприка и сушенный чеснок(можно свежий)Приготовление:Авокадо очистить и размять вилкой,помидор нарезать кубиком и кинзу мелко нарезать. Все смешать,добавить соль,перец,паприку,чеснок и лимонный сок. Все хорошо перемешать и выложить на поджаренные тосты.Готовим яйцо-пашот:В кипящую воду добавляем 1 ст.л. уксуса,делаем ложкой воронку и в эту воронку быстро,но очень аккуратно вливаем яйцо,уменьшаем огонь до минимума и варим 2 минуты. Достаём шумовкой и аккуратно выкладываем на наши тосты с гуакамоле.Вкуснейший и нежный завтрак готов.Обязательно приготовьте.Приятного аппетита"]}' -H "Content-Type: application/json" -X POST http://localhost:$(PORT_TR)/model

test: test_iai test_ner test_tes test_tr


#download:
#	# python -m deeppavlov download ner_ontonotes_bert_mult
#	docker run --rm -ti -v `pwd`/src:/root/.deeppavlov --entrypoint bash $(SRV)