//go:build ignore
// +build ignore

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

// The following directive is necessary to make the package coherent:

// This program generates contributors.go. It can be invoked by running
// go generate

type FieldDef struct {
	Name string
	T    string
}
type StructDef struct {
	Name   string
	Fields []FieldDef
}

func main() {
	fileName := os.Getenv("GOFILE")
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fullPath := cwd + "/" + fileName
	log.Println(fullPath)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fullPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}
	for _, node := range f.Decls {
		str := StructDef{}
		switch node.(type) {
		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)

					str.Name = typeSpec.Name.Name

					switch typeSpec.Type.(type) {
					case *ast.StructType:
						structType := typeSpec.Type.(*ast.StructType)
						for _, field := range structType.Fields.List {
							switch field.Type.(type) {
							case *ast.Ident:
								i := field.Type.(*ast.Ident)
								str.Fields = append(str.Fields, FieldDef{
									Name: field.Names[0].Name,
									T:    i.Name,
								})
							case *ast.SelectorExpr:
								i := field.Type.(*ast.SelectorExpr)
								str.Fields = append(str.Fields, FieldDef{
									Name: field.Names[0].Name,
									T:    i.Sel.Name,
								})
							default:
								log.Println("default")
							}
						}
					}
				}
			}
		}
		log.Println(str)
	}
}
