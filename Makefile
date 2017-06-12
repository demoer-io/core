GO=go
GOV=govendor
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GOV) fetch -v

fetch:
	@$(GOGET) github.com/urfave/cli
	@$(GOGET) github.com/gorilla/mux
	@$(GOGET) github.com/gorilla/handlers

build:
	@echo "start building..."
	$(GOINSTALL)
	@echo "Yay! build DONE!"

all: fetch build
