.PHONY:db-migrate
db-migrate:
	export ENV=local && go run main.go db:migrate

.PHONY:db-rollback
db-rollback:
	export ENV=local && go run main.go db:rollback

.PHONY:serve
serve:
	export ENV=local && go run main.go serve

.PHONY:db-up
db-up:
	docker-compose up -d postgres

.PHONY:api-up
api-up:
	docker build -t tiger-kittens . && docker-compose up -d tiger-kittens

.PHONY:api-down
api-down:
	docker-compose rm -s tiger-kittens

.PHONY:down
down:
	docker-compose down