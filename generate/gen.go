//go:generate go run gen.go

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"tag-names/model"
	"text/template"
	"time"
)

const tagName = "map"

type MetaStruct struct {
	Package string
	Name    string
	Type    string
	Marshal string
}

func main() {
	println("Generating some awesome go code")
	templates := template.Must(template.New("templates").ParseGlob("../codegen_templates/*"))
	createFile("../marshalling", "trader.go", "marshal.tpl", model.Trade{}, templates)
}

func createFile(outputPath, outputFileName, templateFileName string, object any, templates *template.Template) {
	out, err := os.Create(fmt.Sprintf("%s/%s", outputPath, outputFileName))
	if err != nil {
		log.Printf("%v", err)
	}
	defer out.Close()

	err = templates.ExecuteTemplate(out, templateFileName, createMetaStruct(object))
}

func createMetaStruct(st any) *MetaStruct {
	ms := &MetaStruct{}

	// TypeOf returns the reflection Type that represents the dynamic type of i.
	// If i is a nil interface value, TypeOf returns nil.
	rt := reflect.TypeOf(st)

	ms.Name = rt.Name()
	ms.Package = rt.PkgPath()

	ms.Type = fmt.Sprintf("%s.%s", strings.Split(ms.Package, "/")[1], ms.Name)

	rv := reflect.ValueOf(st)

	if rv.Kind() == reflect.Pointer {
		panic("Cannot marshal pointer")
	}

	buf := new(bytes.Buffer)

	buf.WriteString(`fmt.Sprintf("`)

	var fields []string

	for i := 0; i < rt.NumField(); i++ {
		tf := rt.Field(i)
		vf := rv.Field(i)
		tag := tf.Tag.Get(tagName)

		buf.WriteString(tag)
		buf.WriteString(":")

		switch vf.Interface().(type) {
		case int64:
			buf.WriteString("%d\\n")
			fields = append(fields, fmt.Sprintf("object.%s", tf.Name))
		case time.Time:
			buf.WriteString("%s\\n")
			fields = append(fields, fmt.Sprintf("object.%s.Format(time.RFC3339)", tf.Name))
		case string:
			buf.WriteString("%s\\n")
			fields = append(fields, fmt.Sprintf("object.%s", tf.Name))
		case float64:
			buf.WriteString("%.2f\\n")
			fields = append(fields, fmt.Sprintf("object.%s", tf.Name))
		}
	}

	buf.WriteString(fmt.Sprintf(`", %s)`, strings.Join(fields, ", ")))
	ms.Marshal = buf.String()

	return ms
}
