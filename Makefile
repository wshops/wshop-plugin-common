requirements:
	@echo "Installing requirements..."
	go mod download
	go mod tidy

clean-packages:
	go clean -modcache