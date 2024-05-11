.PHONY:dev-db-migrate
dev-db-migrate:
	export ENV=local && go run main.go db:migrate

.PHONY:dev-db-rollback
dev-db-rollback:
	export ENV=local && go run main.go db:rollback

.PHONY:dev-db-seed
dev-db-seed:
	export ENV=local && go run main.go db:seed

.PHONY:serve
serve:
	export ENV=local && go run main.go serve

.PHONY:db-up
db-up:
	docker-compose up -d postgres

.PHONY:api-up
api-up:
	docker build -t tiger-kittens . && docker-compose up -d tiger-kittens

.PHONY:db-migrate
db-migrate:
	docker build -t tiger-kittens . && docker-compose up -d tiger-kittens-migrate

.PHONY:db-seed
db-seed:
	docker build -t tiger-kittens . && docker-compose up -d tiger-kittens-seed

.PHONY:api-down
api-down:
	docker-compose rm -s tiger-kittens

.PHONY:down
down:
	docker-compose down