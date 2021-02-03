package c2rust

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
	cvar     CVar
	right    string // this should be an expr?
	operator string
}

type CArray struct {
	label        string
	arraytype    string
	size         string
	initializer  string
	label_filled bool
	size_filled  bool
}

type CFunction struct {
	label      string
	args       []CVar
	body       []CStatement
	returntype string
}

type CStatement interface {
	UpdateEnv(*Env) *Env
}
