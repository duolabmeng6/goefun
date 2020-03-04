package coreUtil

import (
	"github.com/robertkrimen/otto"
)

func RunJs(content string,functionName string, args... interface{}) string{
	var err error
	vm := otto.New()
	_, err = vm.Run(content)
	if err!=nil {
		panic(err)
	}
	value, err := vm.Call(functionName, nil, args...)
	if err != nil {
		panic(err)
	}
	return value.String()
}