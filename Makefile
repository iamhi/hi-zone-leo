default: tidy docker-stop docker-up templ run

docker-stop:
	docker-compose stop

docker-up:
	docker-compose up -d

templ:
	templ generate

tidy:
	go mod tidy

build:
	cd cmd/web && go build -o ../../hi-zone-leo.go main.go

run:
	air

run-internal:
	cd cmd/internal && air

stop: docker-stop
