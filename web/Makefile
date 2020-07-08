include .env
main: ver ci push comp

log:
	docker-compose logs -f ci

R=docker-compose run --rm

remove-old:
	docker image rm -f $(SRV)

ci:
	time docker-compose build $@

prune:
	docker image prune -f

run-ci: ci prune
	$R ci

ID=`docker-compose ps -q`