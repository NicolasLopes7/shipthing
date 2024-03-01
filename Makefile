.PHONY: dev

dev:
	@find ./$(service) -name "*.go" | entr -r sh -c 'echo "Restarted" && go run ./$(service)'