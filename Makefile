PHONY: build clean

build:
	@echo "Building the project..."
	go build -v -o pretend 
	@sha256sum pretend

clean:
	@echo "Cleaning up..."
	rm -f pretend

act:
	@echo "Running act..."
	act -W .github/workflows/go.yml
