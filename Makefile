.PHONY: dev

dev:
	@find . -name "*.go" | entr -r sh -c 'echo "Restarted" && go run .'