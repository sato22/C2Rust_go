package parser

import (
	"fmt"
	"strings"
)

type RVar struct {
	label   string
	vartype string
	mutable bool
}

type RVarDecl struct {
	rvar        RVar
	initializer string
}

type RAssignment struct {
	rvar     RVar
	right    string
	operator string
}

type RArray struct {
	dim         int
	size        []string
	label       string
	arraytype   string
	initializer string
}

type RFunction struct {
	label      string
	args       []RVar
	body       []RWriter
	returntype string
}

type RWriter interface {
	write()
}

func (f *RFunction) write() {
	argstr := []string{}
	for _, x := range f.args {
		if x.mutable {
			argstr = append(argstr, fmt.Sprintf("mut %s: %s", x.label, x.vartype))
		} else {
			argstr = append(argstr, fmt.Sprintf("%s: %s", x.label, x.vartype))
		}
	}
	if f.returntype != "" {
		fmt.Printf("fn %s(%s) -> %s {\n", f.label, strings.Join(argstr, ", "), f.returntype)
	} else {
		fmt.Printf("fn %s(%s) {\n", f.label, strings.Join(argstr, ", "))
	}
	for _, x := range f.body {
		x.write()
	}
	fmt.Println("}")
}

func (d *RVarDecl) write() {
	if d.rvar.mutable {
		fmt.Printf("let mut %s: %s", d.rvar.label, d.rvar.vartype)
	} else {
		fmt.Printf("let %s: %s", d.rvar.label, d.rvar.vartype)
	}
	if d.initializer != "" {
		fmt.Printf(" = %s", d.initializer)
	}
	fmt.Println()
}

func (s *RAssignment) write() {
	fmt.Printf("%s %s %s", s.rvar.label, s.operator, s.right)
	fmt.Println()
}
