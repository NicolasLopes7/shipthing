.PHONY: dev

dev:
	@find ./$(service) ./lib ./config -name "*.go" | entr -r sh -c 'echo "Restarted" && go run ./$(service)'