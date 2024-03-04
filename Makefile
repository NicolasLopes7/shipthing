.PHONY: dev

dev:
	@find ./$(service) ./lib ./config -name "*.go" | entr -r sh -c 'echo "Restarted" && go run ./$(service)'

dev-beta:
	@make -j 2 deployer uploader compose

deployer:
	@bash ./infra/deployer-service.sh

uploader:
	@bash ./infra/uploader-service.sh

compose:
	@docker-compose -f infra/docker-compose.yaml up -d || docker compose -f infra/docker-compose.yaml up -d
