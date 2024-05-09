.PHONE:migrate
migrate:
	export ENV=local && go run main.go db:migrate