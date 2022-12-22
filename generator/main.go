package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

var (
	//go:embed templates/*.tmpl
	templates embed.FS

	input      *string
	output     *string
	outputFile *os.File
	generate   *string
)

const (
	unset string = "<UNSET>"
)

func init() {
	input = flag.String("i", unset, "The input file name")
	output = flag.String("o", unset, "The output file name")
	generate = flag.String("generate", unset, "The type of code to generate datasources|converters")
}

type Datasource struct {
	ParamsStructName       string `json:"params_struct_name"`
	ApiFunction            string `json:"api_function"`
	ApiToTerraformFunction string `json:"api_to_terraform"`
	SchemaFunction         string `json:"schema_function"`
	ElementName            string `json:"element_name"`
	ResultField            string `json:"result_field"`
	ResultId               string `json:"result_id"`
}

type InputData struct {
	ApiAlias    string       `json:"api_alias"`
	ApiPackage  string       `json:"api_package"`
	Package     string       `json:"package"`
	DataSources []Datasource `json:"data_sources"`
}

// Main function
func main() {
	flag.Parse()

	if *input == unset || *generate == unset {
		flag.CommandLine.Usage()
		os.Exit(0)
	}

	inputFile, err := ioutil.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	data := InputData{}
	err = json.Unmarshal([]byte(inputFile), &data)
	if err != nil {
		panic(err)
	}

	if *output == unset {
		outputFile = os.Stdout
	} else {
		file, err := os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		outputFile = file
	}

	tmplName := fmt.Sprintf("%s.go.tmpl", *generate)
	tmplFile := fmt.Sprintf("templates/%s", tmplName)
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"apiToTerraformType": func(t string) string {
			switch t {
			case "int":
				return "schema.TypeInt"
			case "string":
				return "schema.TypeString"
			case "boolean":
				return "schema.TypeBool"
			default:
				panic(fmt.Sprintf("Unsupported type: %s", t))
			}
		},
	}).ParseFS(templates, tmplFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(tmpl.Name())

	if err := tmpl.Execute(outputFile, data); err != nil {
		panic(err)
	}
}
