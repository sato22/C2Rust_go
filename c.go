package parser

import (
// 	"fmt"
)

type CVar struct {
	label   string
	vartype string
}

type CVarDecl struct {
	cvar        CVar
	initializer string // this should be an expr?
}

type CAssignment struct {
	label    string
	right    string // this should be an expr?
	operator string
}

type CArray struct {
	dim         int      // arraydim exitdirectdeclarator で case:3の際に++
	size        []string // a[3][4] の数字部分
	label       string   // arraylabel
	arraytype   string
	initializer string
}

type CFunction struct {
	label      string
	args       []CVar
	body       []CStatement
	returntype string
}

type CStatement interface {
	UpdateEnv(*Env) *Env
	AddStatement(*Env, []RWriter) []RWriter
}
