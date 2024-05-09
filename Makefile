.PHONY:db-migrate
db-migrate:
	export ENV=local && go run main.go db:migrate

.PHONY:db-rollback
db-rollback:
	export ENV=local && go run main.go db:rollback

.PHONY:serve
serve:
	export ENV=local && go run main.go serve