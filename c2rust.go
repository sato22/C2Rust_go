package c2rust

import (
	"fmt"
)

var vartable map[string]string

func init() {
	vartable = make(map[string]string)
	vartable["int"] = "i32"
	vartable["float"] = "f32"
	vartable["double"] = "f64"
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
		v := &Var{
			label:       label,
			vartype:     vt,
			initializer: initializer,
			mutable:     false,
			used:        false,
		}
		env.vars[label] = v
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
	// args
	args := []RVar{}
	for _, x := range f.args {
		v := RVar{
			label:   x.label,
			vartype: vartable[x.vartype],
		}
		args = append(args, v)
	}
	for _, s := range f.body {
		env = s.UpdateEnv(env)
	}
	body := []RWriter{}
	// decl block
	for k, v := range env.vars {
		fmt.Println(v.vartype)
		rv := RVar{
			label:   k,
			vartype: v.vartype,
			mutable: v.mutable,
		}
		s := &RVarDecl{
			rvar:        rv,
			initializer: v.initializer,
		}
		body = append(body, s)
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
	env.NewVar(s.cvar.label, s.cvar.vartype, s.initializer)
	return env
}

func (s *CAssignment) UpdateEnv(env *Env) *Env {
	v, _ := env.GetVar(s.cvar.label)
	// the right value can be evaluated?
	if !v.used {
		v.initializer = s.right
		v.used = true
	} else {
		v.mutable = true
	}
	return env
}
