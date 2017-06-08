GO=go
GOV=govendor
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GOV) fetch -v

deps:
	@$(GOGET) github.com/urfave/cli

fetch:
	@$(GOGET) github.com/urfave/cli

build:
	@echo "start building..."
	$(GOINSTALL)
	@echo "Yay! build DONE!"

all: deps build
