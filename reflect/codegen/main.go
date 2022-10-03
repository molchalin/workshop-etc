package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"reflect"
)

func main() {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "main.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("main_generated.go")
	if err != nil {
		log.Fatal(err)
	}

	w := new(bytes.Buffer)

	fmt.Fprintf(w, "package %v\n\n\n", node.Name.String())
	fmt.Fprintf(w, "import (\n")
	fmt.Fprintf(w, "\"strconv\"\n")
	fmt.Fprintf(w, "\"fmt\"\n")
	fmt.Fprintf(w, ")\n")

	for _, n := range node.Decls {
		g, ok := n.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range g.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := currType.Type.(*ast.StructType)
			if !ok {
				continue
			}
			fmt.Fprintf(w, "func (v *%v) Scan(record []string, m map[string]int) error {\n", currType.Name.Name)
			for _, field := range structType.Fields.List {
				name := field.Names[0].Name
				if field.Tag != nil {
					tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
					if n := tag.Get("csv"); n != "" {
						name = n
					}
				}
				fmt.Fprintf(w, "{\n")
				fmt.Fprintf(w, `idx, ok := m["%v"]`, name)
				fmt.Fprintf(w, "\nif !ok {\n")
				fmt.Fprintf(w, `return fmt.Errorf("can't find field %v")`, name)
				fmt.Fprintf(w, "\n}\n")
				fmt.Fprintf(w, "str := record[idx]\n")
				switch field.Type.(*ast.Ident).Name {
				case "string":
					fmt.Fprintf(w, "v.%v = str\n", field.Names[0].Name)
				case "float64":
					fmt.Fprintf(w, "var err error\n")
					fmt.Fprintf(w, "v.%v, err = strconv.ParseFloat(str, 64)\n", field.Names[0].Name)
					fmt.Fprintf(w, "if err != nil {\n")
					fmt.Fprintf(w, "return err\n")
					fmt.Fprintf(w, "}\n")
				case "int":
					fmt.Fprintf(w, "var err error\n")
					fmt.Fprintf(w, "v.%v, err = strconv.Atoi(str)\n", field.Names[0].Name)
					fmt.Fprintf(w, "if err != nil {\n")
					fmt.Fprintf(w, "return err\n")
					fmt.Fprintf(w, "}\n")
				default:
					panic("unsupported type")
				}
				fmt.Fprintf(w, "}\n")
			}
			fmt.Fprintf(w, "return nil\n}\n")
		}
	}
	b, err := format.Source(w.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(out, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
}
