BIN=g2html

build:
	go build -o $(BIN)

run: build
	./$(BIN) -d $(G2HTML_DIR)
