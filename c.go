package parser

import (
// 	"fmt"
)

type CVarDecl struct {
	label       string
	vartype     string
	initializer string // this should be an expr?
	dim         int
	size        []string
}

type CAssignment struct {
	label    string
	size     []string
	right    string // this should be an expr?
	operator string
}

type CFunction struct {
	label      string
	args       []CVarDecl
	body       []CStatement
	returntype string
}

type CStatement interface {
	UpdateEnv(*Env) *Env
	AddStatement(*Env, []RWriter, []RVarDecl) []RWriter
}
