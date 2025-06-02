# Default target: run with `make`
.DEFAULT_GOAL := run

# Run the CLI with ARGS, e.g. `make run ARGS="hi"`
run:
	go run ./cmd/rpd $(ARGS)

# Build the rpd binary into ~/bin
build:
	go build -o ~/bin/rpd ./cmd/rpd

# Alias for build (semantically clearer if you use 'make install')
install: build

update:
	go run ./cmd/rpd update