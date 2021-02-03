package c2rust

import (
	"fmt"
	// "log"
)

var vartable map[string]string
var rbody []RWriter
var assignment []RWriter

func init() {
	vartable = make(map[string]string)
	vartable["short int"] = "i16"
	vartable["int"] = "i32"
	vartable["long int"] = "i64"
	vartable["unsigned int"] = "u32"
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
		if initializer == "" {
			v := &Var{
				label:       label,
				vartype:     vt,
				initializer: initializer,
				mutable:     false,
				used:        false,
			}
			env.vars[label] = v
			// fmt.Println("env.NewVar")
		} else {
			v := &Var{
				label:       label,
				vartype:     vt,
				initializer: initializer,
				mutable:     false,
				used:        true,
			}
			env.vars[label] = v
			// fmt.Println("env.NewVar")
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
	// args/cassignment
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
		rbody = append(rbody, s)
	}

	// array
	for _, s := range f.body {
		env = s.UpdateEnv(env)
	}

	// TODO: statements
	return &RFunction{
		label:      f.label,
		args:       args,
		body:       rbody,
		assignment: assignment,
		returntype: vartable[f.returntype],
	}
}

func (s *CVarDecl) UpdateEnv(env *Env) *Env {
	env.NewVar(s.cvar.label, s.cvar.vartype, s.initializer)
	return env
}

func (s *CArray) UpdateEnv(env *Env) *Env {
	env.NewVar(s.label, s.arraytype, s.initializer)
	return env
}

func (s *CAssignment) UpdateEnv(env *Env) *Env {
	v, _ := env.GetVar(s.cvar.label)
	// the right value can be evaluated?
	if s.operator != "==" {
		if !v.used {
			v.initializer = s.right
			v.used = true
		} else {
			v.mutable = true

			if v.label == s.cvar.label {
				rv := RVar{
					label:   v.label,
					vartype: v.vartype,
					mutable: v.mutable,
				}

				rs := &RAssignment{
					rvar:     rv,
					right:    s.right,
					operator: s.operator,
				}

				assignment = append(assignment, rs)
				fmt.Println("assignment = append")
				fmt.Printf("assignment = %v\n", assignment)
			}

		}
	}
	return env
}
