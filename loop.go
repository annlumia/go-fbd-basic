package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

// Variable ...
type Variable struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// Program ...
type Program struct {
	ID      string     `json:"id"`
	Method  string     `json:"method"`
	Inputs  []Variable `json:"inputs"`
	Outputs []Variable `json:"outputs"`
}

func parseProgram(p *[]Program) error {
	file, err := ioutil.ReadFile("program.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), p)
	if err != nil {
		return err
	}

	return nil
}

func writeValues(args ...Variable) []interface{} {
	values := []interface{}{}
	for _, arg := range args {
		switch arg.Type {
		case "v":
			values = append(values, readVariableValue(arg.Value.(string)))
		}
	}
	return values
}

func readValues(args ...Variable) []interface{} {
	values := []interface{}{}
	for _, arg := range args {
		switch arg.Type {
		case "v":
			values = append(values, readVariableValue(arg.Value.(string)))
		case "c":
			values = append(values, arg.Value)
		}
	}
	return values
}

func writeVariableValue(v string, val interface{}) interface{} {
	area := v[0:1]
	offset, err := strconv.Atoi(v[1:])
	if err != nil {
		panic(err)
	}
	switch area {
	case "M":
		M[offset] = val.(bool)
	case "R":
		R[offset] = val.(float64)
	}
	return nil
}

func readVariableValue(v string) interface{} {
	area := v[0:1]
	offset, err := strconv.Atoi(v[1:])
	if err != nil {
		panic(err)
	}
	switch area {
	case "M":
		return M[offset]
	case "R":
		return R[offset]
	}
	return nil
}

func writeOutputs(out []Variable, args []interface{}) {
	for index, v := range out {
		switch v.Type {
		case "v":
			writeVariableValue(v.Value.(string), args[index])
		}
	}
}

// Loop ...
func Loop() {
	for {
		programs := []Program{}
		if err := parseProgram(&programs); err != nil {
			log.Println("Error when parsing program. Error :", err.Error())
		}

		for _, p := range programs {
			// Get input values
			inpts := readValues(p.Inputs...)
			outputs := Functions[p.Method](true, inpts...)
			writeOutputs(p.Outputs, outputs)

			time.Sleep(time.Second * 1)
		}
	}
}
