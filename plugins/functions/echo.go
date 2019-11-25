package main

import (
	"fmt"
)

type echo struct {
}

func (f *echo) Validate(args []interface{}) error{
	if len(args) != 1{
		return fmt.Errorf("echo function only supports 1 parameter but got %d", len(args))
	}
	return nil
}

func (f *echo) Exec(args []interface{}) (interface{}, bool) {
	result := args[0]
	return result, true
}

var Echo echo
