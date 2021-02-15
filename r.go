package parser

import (
	"fmt"
	"strings"
)

type RVarDecl struct {
	label       string
	vartype     string
	mutable     bool
	initializer string
	dim         int
	size        []string
}

type RAssignment struct {
	label    string
	vartype  string
	mutable  bool
	size     []string
	right    string
	operator string
}

type RFunction struct {
	label      string
	args       []RVarDecl
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
	if d.mutable {
		fmt.Printf("let mut %s: ", d.label)
	} else {
		fmt.Printf("let %s: ", d.label)
	}
	if d.dim > 0 {
		if len(d.size) == 2 {
			fmt.Printf("[[%s;%s]%s]", d.vartype, d.size[0], d.size[1])
		} else {
			fmt.Printf("[%s;%s]", d.vartype, d.size[0])
		}
	} else {
		fmt.Printf("%s", d.vartype)
	}

	if d.initializer != "" {
		d.initializer = strings.Replace(d.initializer, "{", "[", -1)
		d.initializer = strings.Replace(d.initializer, "}", "]", -1)
		fmt.Printf(" = %s", d.initializer)
	}
	fmt.Println()
}

func (s *RAssignment) write() {
	if len(s.size) == 0 || strings.Contains(s.right, "[") {
		fmt.Printf("%s %s %s", s.label, s.operator, s.right)
	} else {
		fmt.Printf("%s", s.label)
		if len(s.size) == 2 {
			fmt.Printf("[%s][%s]", s.size[0], s.size[1])
		} else {
			fmt.Printf("[%s]", s.size[0])
		}
		fmt.Printf(" %s %s", s.operator, s.right)
	}
	fmt.Println()
}
