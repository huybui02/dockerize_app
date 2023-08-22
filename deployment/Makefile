SHELL:=/bin/bash

run-fe:
	docker exec -it -uroot go_app sh -c 'cd /app && go mod download && go build -o main . && go run main.go'

run-be:
	docker exec -it -uroot nextjs_app sh -c 'cd /frontend && npm install && npm run dev'