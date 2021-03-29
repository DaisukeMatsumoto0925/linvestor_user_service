package controllers

import "time"

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

// {
// 	"resource": "/Users/matsumotodaisuke/go/src/github.com/xfpng345/linvestor_user_service/src/app/infrastructure/sqlhandler.go",
// 	"owner": "_generated_diagnostic_collection_name_#22",
// 	"severity": 8,
// 	"message": "cannot use ctx (variable of type interface{}) as context.Context value in argument to handler.Conn.CreateUser: missing method Deadline",
// 	"source": "compiler",
// 	"startLineNumber": 46,
// 	"startColumn": 38,
// 	"endLineNumber": 46,
// 	"endColumn": 41
// }
