// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package hello

import (
	"demo/api/hello"
)
// ControllerV1 implements the hello.IHelloV1 interface for version 1 of the hello API.
type ControllerV1 struct{}
// NewV1 returns a new instance of ControllerV1 as hello.IHelloV1.
// This function is used by the router to bind the controller to API endpoints.
// Creating route object
func NewV1() hello.IHelloV1 {
	return &ControllerV1{}
}

