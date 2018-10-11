package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

func run(dir string) error {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	fmt.Println(pkgs)

	return nil
}

func main() {
	dir := flag.String("d", "", "parse directory")
	flag.Parse()

	if err := run(*dir); err != nil {
		log.Fatalln(err)
	}
}
