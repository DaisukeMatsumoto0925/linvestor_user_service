package controllers

import "time"

// Context is context.Context
type Context interface {
	Param(string) string
	Bind(interface{}) error
	BindJSON(interface{}) error
	JSON(int, interface{})
	Deadline() (time.Time, bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
