package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/yunkon-kim/data-model-generation/pkg/data-model/infra"
	"gopkg.in/yaml.v3"
)

func main() {
	vpc := &infra.VPC{
		NameId:   "vpc-01",
		IPv4CIDR: "",
		Subnets: []struct {
			NameId   string "json:\"nameId\"  yaml:\"nameId\"  validate:\"required\""
			IPv4CIDR string "json:\"iPv4CIDR\"  yaml:\"iPv4CIDR\"  validate:\"cidrv4\""
		}{
			{"subnet-01", "192.168.0.0/24"},
			{"subnet-02", "192.168.1.0/24"},
		},
	}

	fmt.Println("== (Start) Validate a struct ==")
	validate := validator.New()
	valErr := validate.Struct(vpc)
	if valErr != nil {
		fmt.Println(valErr)
		return
	}
	fmt.Println("== (End) Validate a struct ==")

	fmt.Println("== (Start) Struct to JSON Bytes ==")
	jsonBytes, err := json.Marshal(vpc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))
	fmt.Println("== (End) Struct to JSON Bytes ==")

	fmt.Println("== (Start) Struct to JSON Bytes (pretty) ==")
	jsonPrettyBytes, err2 := json.MarshalIndent(vpc, "", "\t")
	if err != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(string(jsonPrettyBytes))

	fmt.Println("== (End) Struct to JSON Bytes (pretty) ==")

	fmt.Println("== (Start) Struct to YAML Bytes ==")
	yamlBytes, err3 := yaml.Marshal(vpc)
	if err2 != nil {
		fmt.Println(err3)
		return
	}

	fmt.Println(string(yamlBytes))
	fmt.Println("== (End) Struct to YAML Bytes ==")

}
