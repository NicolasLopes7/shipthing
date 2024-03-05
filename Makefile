.PHONY: dev

compose:
	@docker-compose up -d || docker compose up -d

deployer:
	@make watch ; go run ./deployer-service

uploader:
	@make watch ; go run ./uploader-service

watch:
	@find . -name "*.go" | entr -r sh -c 'echo "Restarted"'

dev:
	@make compose ; make deployer ; make uploader
