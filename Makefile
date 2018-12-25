BIN=g2html

build:
	go build -o $(BIN)

build-local:
	go build -o $(GOPATH)/bin/$(BIN)

run: build
	./$(BIN) -d $(G2HTML_DIR) -t $(G2HTML_TMP)
