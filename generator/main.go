package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

type ParamsFieldsStruct struct {
	Name    string `json:"name"`
	ApiName string `json:"api_name"`
	Type    string `json:"type"`
}

type ParamsStruct struct {
	Name   string               `json:"name"`
	Fields []ParamsFieldsStruct `json:"fields"`
}

type TerraformArguments struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
	Computed    bool   `json:"computed,omitempty"`
	Optional    bool   `json:"optional,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

type Terraform struct {
	Name      string               `json:"name"`
	Arguments []TerraformArguments `json:"arguments,omitempty"`
}

type ApiArguments struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

type ApiFunction struct {
	Name            string         `json:"name"`
	ResultIsPointer bool           `json:"isPointer"`
	Arguments       []ApiArguments `json:"arguments"`
	ApiParams       ParamsStruct   `json:"params"`
}

type TerraformResult struct {
	ApiToTerraformFunction string `json:"api_to_terraform"`
	MessageField           string `json:"msg_field"`
	ResultField            string `json:"result_field"`
	ResultWrapperFunction  string `json:"result_wrapper"`
	TerraformField         string `json:"terraform_field"`
	TerraformId            string `json:"terraform_id"`
	ApiToIdFn              string `json:"api_to_terraform_id"`
	InputToIdFn            string `json:"input_to_terraform_id"`
}

type Operation struct {
	ApiFunction              ApiFunction     `json:"api_function"`
	Result                   TerraformResult `json:"result"`
	InputToTerraformFunction string          `json:"input_to_terraform"`
	Name                     string
	ElementName              string
	CRUD                     string
}

type Schema struct {
	SchemaFunction          string `json:"schema_function"`
	SchemaFunctionArguments string `json:"schema_function_arguments"`
}

type Component struct {
	Name          string               `json:"name"`
	Terraform     Terraform            `json:"terraform"`
	ElementName   string               `json:"element_name"`
	ResultSchemas map[string]Schema    `json:"result_schemas"`
	CRUD          map[string]Operation `json:"crud"`
}

type InputData struct {
	ApiAlias    string      `json:"api_alias"`
	ApiPackage  string      `json:"api_package"`
	Package     string      `json:"package"`
	DataSources []Component `json:"data_sources"`
	Resources   []Component `json:"resources"`
}

// Main function
func main() {
	flag.Parse()

	if *input == unset || *generate == unset {
		flag.CommandLine.Usage()
		os.Exit(0)
	}

	data := InputData{}
	inputs := strings.Split(*input, ",")
	for i, input := range inputs {
		inputFile, err := ioutil.ReadFile(input)
		if err != nil {
			panic(err)
		}

		parsed := InputData{}
		err = json.Unmarshal([]byte(inputFile), &parsed)
		if err != nil {
			panic(err)
		}
		if i == 0 {
			data.ApiAlias = parsed.ApiAlias
			data.ApiPackage = parsed.ApiPackage
			data.Package = parsed.Package
			data.DataSources = parsed.DataSources
			data.Resources = parsed.Resources
		} else {
			data.DataSources = append(data.DataSources, parsed.DataSources...)
			data.Resources = append(data.Resources, parsed.Resources...)
		}
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
	caser := cases.Title(language.Und, cases.NoLower)
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"isNotEmpty": func(s string) bool {
			return len(s) > 0
		},
		"isPointer": func(p bool) string {
			if p {
				return "*"
			} else {
				return ""
			}
		},
		"contains": func(k string, m map[string]interface{}) bool {
			_, ok := m[k]
			return ok
		},
		"default": func(d string, v string) string {
			if v != "" {
				return v
			} else {
				return d
			}
		},
		"ToLower": strings.ToLower,
		"toTitle": caser.String,
		"enrichOperation": func(op Operation, n string, c string, en string) Operation {
			op.Name = n
			op.CRUD = c
			op.ElementName = en
			return op
		},
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
