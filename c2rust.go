package parser

import (
	"fmt"
	// "log"
)

var vartable map[string]string
var body []RWriter

func init() {
	vartable = make(map[string]string)
	vartable["shortint"] = "i16"
	vartable["int"] = "i32"
	vartable["longint"] = "i64"
	vartable["unsignedint"] = "u32"
	vartable["float"] = "f32"
	vartable["double"] = "f64"
	vartable["char"] = "char"
}

type Var struct {
	label       string
	vartype     string
	initializer string
	mutable     bool
	used        bool
}

type Env struct {
	parent *Env
	vars   map[string]*Var
}

func NewEnv(parent *Env) *Env {
	return &Env{
		parent: parent,
		vars:   make(map[string]*Var),
	}
}

func (env *Env) NewVar(label string, vartype string, initializer string) {
	if vt, ok := vartable[vartype]; ok {
		if initializer == "" {
			v := &Var{
				label:       label,
				vartype:     vt,
				initializer: initializer,
				mutable:     false,
				used:        false,
			}
			env.vars[label] = v
		} else {
			v := &Var{
				label:       label,
				vartype:     vt,
				initializer: initializer,
				mutable:     false,
				used:        true,
			}
			env.vars[label] = v
		}
	} else {
		fmt.Println("error")
	}
}

func (env *Env) GetVar(varlabel string) (*Var, bool) {
	current := env
	for current != nil {
		if v, ok := env.vars[varlabel]; ok {
			return v, true
		} else {
			current = env.parent
		}
	}
	return nil, false
}

func (f *CFunction) toRust() *RFunction {
	env := NewEnv(nil)

	// update env
	for _, x := range f.args {
		env = x.UpdateEnv(env)
	}
	for _, s := range f.body {
		env = s.UpdateEnv(env)
	}

	// args
	args := []RVarDecl{}
	for _, x := range f.args {
		v, _ := env.GetVar(x.label)
		arg := RVarDecl{
			label:   v.label,
			vartype: v.vartype,
			mutable: false,
		}
		if v.used {
			arg.mutable = true
		}
		args = append(args, arg)
	}

	// decl block/assignment
	for _, s := range f.body {
		body = s.AddStatement(env, body, args)
	}

	// TODO: statements
	return &RFunction{
		label:      f.label,
		args:       args,
		body:       body,
		returntype: vartable[f.returntype],
	}
}

func (s *CVarDecl) UpdateEnv(env *Env) *Env {
	env.NewVar(s.label, s.vartype, s.initializer)
	return env
}

func (s *CAssignment) UpdateEnv(env *Env) *Env {
	v, _ := env.GetVar(s.label)
	// the right value can be evaluated?

	if !v.used {
		v.initializer = s.right
		v.used = true
	} else {
		v.mutable = true
	}
	return env
}

func (s *CVarDecl) AddStatement(env *Env, body []RWriter, args []RVarDecl) []RWriter {
	v, _ := env.GetVar(s.label)

	fmt.Println(v.vartype)

	rvdecl := &RVarDecl{
		label:       s.label,
		vartype:     v.vartype,
		mutable:     v.mutable,
		initializer: v.initializer,
		dim:         s.dim,
		size:        s.size,
	}
	body = append(body, rvdecl)

	return body
}

func (s *CAssignment) AddStatement(env *Env, body []RWriter, args []RVarDecl) []RWriter {
	v, _ := env.GetVar(s.label)

	if s.operator != "" {
		rs := &RAssignment{
			label:    v.label,
			vartype:  v.vartype,
			size:     s.size,
			right:    s.right,
			operator: s.operator,
		}

		body = append(body, rs)
	}
	return body
}
