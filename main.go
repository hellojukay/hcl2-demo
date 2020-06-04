package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// Configuration config struct
type Configuration struct {
	Pipeline Pipeline `hcl:"pipeline,block" json:"pipeline"`
}

// Pipeline pipeline struct
type Pipeline struct {
	Env    map[string]string `hcl:"env" json:"env"`
	Stages Task              `hcl:"stages,block" json:"stages"`
}

// Task task struct
type Task struct {
	Stages []Stage `hcl:"stage,block" json:"stage"`
}

// Stage stage struct
type Stage struct {
	Name     string `hcl:"name"`
	Describe string `hcl:"describe,optional"`
}

var pwd = function.New(&function.Spec{
	Type: function.StaticReturnType(cty.String),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		dir, _ := os.Getwd()
		return cty.StringVal(dir), nil
	},
})
var config *string

func init() {
	config = flag.String("f", "./taskfile", "config file path")
	flag.Parse()
}
func main() {
	var parser = hclparse.NewParser()
	file, dg := parser.ParseHCLFile(*config)
	if dg.HasErrors() {
		log.Fatalf("parse error %s", dg.Error())
	}
	file.Body.JustAttributes()
	var config Configuration
	ctx := &hcl.EvalContext{
		Functions: map[string]function.Function{
			"PWD": pwd,
		},
	}
	confDiags := gohcl.DecodeBody(file.Body, ctx, &config)
	if confDiags.HasErrors() {
		log.Fatal(confDiags)
	}
	bf, _ := json.Marshal(config)
	fmt.Printf("%s\n", string(bf))
}
