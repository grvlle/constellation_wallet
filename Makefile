bold := $(shell tput bold)
sgr0 := $(shell tput sgr0)

help:
	@printf "\nBelow commands have been made available to ease up Molly Wallet development.\n\n"

	@printf "Molly Wallet App: \n"

	@printf "  $(bold)run_dev_env$(sgr0) - Starts up the dev enviornment on http://localhost:8080\n"
	@printf "  $(bold)build_app$(sgr0) - Compiles the Molly Wallet to the build folder\n"
	@printf "  $(bold)buid_app_production$(sgr0) - Builds a production build for the OS it's running on\n"
	@printf "  $(bold)cross_compile_app$(sgr0) - Compiling for every OS and Platform to the build directory\n\n"

	@printf "Update Module: \n"

	@printf "  $(bold)build_update_module$(sgr0) - Builds the update module to the ~/.dag folder\n"
	@printf "  $(bold)run_update_module$(sgr0) - Builds and runs the update module\n"
	@printf "  $(bold)cross_compile_update_module$(sgr0) - Compiling for every OS and Platform to the build directory\n\n"

	@printf "General: \n"

	@printf "  $(bold)all$(sgr0) - Builds and compiles both the wallet and the Update Module\n"
	@printf "  $(bold)cross_compile_all$(sgr0) - Builds and compiles both the wallet and the Update Module for all Platforms\n\n"

	@printf "  $(bold)clean$(sgr0) - Cleans up old builds from the build directory\n\n"


run_dev_env:
	@echo "Starting up frontend dev env on http://localhost:8080..."
	$(shell wails serve) \
	$(shell cd frontend && npm run serve)

build_app:
	@echo "Building Molly Wallet binary to build folder..."
	wails build

buid_app_production:
	@echo "Building app for production..."
	wails build -f -p

cross_compile_app:
	@echo "Compiling for every OS and Platform..."
	wails build -x darwin/amd64
	wails build -x windows/amd64
	wails build -x linux/amd64

build_update_module:
	@echo "Building update module binary to ~/.dag"
	go build -o ~/.dag/update backend/cmd/update/main.go

run_update_module:
	@echo "Compiling and running the update module"
	go run backend/cmd/update/main.go


cross_compile_update_module:
	@echo "Compiling for every OS and Platform..."
	GOOS=linux GOARCH=amd64 go build -o build/update-linux-amd backend/cmd/update/main.go
	GOOS=windows GOARCH=amd64 go build -o build/update-windows-amd64 backend/cmd/update/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/main-darwin-amd64 backend/cmd/update/main.go

clean:
	@echo "Cleaning up build directory..."
	@rm -rf build

all: build_app build_update_module
cross_compile_all: cross_compile_app cross_compile_update_module