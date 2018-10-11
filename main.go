package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func run(dir string) error {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	fmt.Println(pkgs)

	m := map[string]string{}

	for _, v := range pkgs {
		for fk, fv := range v.Files {
			if !strings.HasSuffix(fk, ".go") || strings.HasSuffix(fk, "_test.go") {
				continue
			}
			commentMap := ast.NewCommentMap(fset, fv, fv.Comments)
			for ck, _ := range commentMap {
				switch tv := ck.(type) {
				case *ast.GenDecl:
					if tv.Tok != token.TYPE {
						break
					}
					var typeName string
					for _, sv := range tv.Specs {
						ssv := sv.(*ast.TypeSpec)
						typeName = ssv.Name.Name
					}
					m[typeName] = tv.Doc.Text()
				default:
					fmt.Println("skip")
				}
			}
		}
	}

	fmt.Print(m)
	return nil
}

func main() {
	dir := flag.String("d", "", "parse directory")
	flag.Parse()

	if err := run(*dir); err != nil {
		log.Fatalln(err)
	}
}
