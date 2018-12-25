package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"strings"
)

const output = "./index.html"

func extract(dir string) (map[string]string, error) {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

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
					continue
				}
			}
		}
	}

	return m, nil
}

func run(dir, tmp string) error {
	m, err := extract(dir)
	if err != nil {
		return err
	}

	t := template.Must(template.ParseFiles(tmp))
	type Data struct {
		Maps map[string]string
	}
	d := Data{Maps: m}

	w, err := os.Create(output)
	if err != nil {
		return err
	}
	defer w.Close()
	if err = t.Execute(w, d); err != nil {
		return err
	}
	return nil
}

func main() {
	dir := flag.String("d", "", "parse directory")
	tmp := flag.String("t", "./template.html", "template html file")
	flag.Parse()

	if err := run(*dir, *tmp); err != nil {
		log.Fatalln(err)
	}
}
