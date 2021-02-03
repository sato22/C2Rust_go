package c2rust

import (
	"fmt"
	"testing"
)

func TestVar(t *testing.T) {
	env := NewEnv(nil)         // enter a scope
	env.NewVar("a", "int", "") // int a;
	fmt.Println(env)
	// a = 1
	val := "1" // literal 1
	if x, ok := env.GetVar("a"); ok {
		x.initializer = val
		fmt.Println(x)
	}
	fmt.Println(env)
	// ....
	if x, ok := env.GetVar("a"); ok {
		x.mutable = false
		fmt.Println(x)
	}
	fmt.Println(env)
}

func TestC(t *testing.T) {
	/*
	   int f(int a, int b) {
	       int c;
	       c = 10;
	       // ....
	   }
	*/
	a := CVar{
		label:   "a",
		vartype: "int",
	}
	b := CVar{
		label:   "b",
		vartype: "int",
	}
	s1 := CVarDecl{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		initializer: "",
	}
	s2 := CAssignment{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		right: "10",
	}

	f := CFunction{
		label:      "f",
		args:       []CVar{a, b},
		body:       []CStatement{&s1, &s2},
		returntype: "int",
	}

	fmt.Println(f)
}

func TestRust(t *testing.T) {
	a := RVar{
		label:   "a",
		vartype: "i32",
	}
	b := RVar{
		label:   "b",
		vartype: "i32",
	}
	s1 := RVarDecl{
		rvar: RVar{
			label:   "c",
			vartype: "i32",
		},
		initializer: "1",
	}
	f := RFunction{
		label:      "f",
		args:       []RVar{a, b},
		body:       []RWriter{&s1},
		returntype: "i32",
	}

	f.write()
}

func TestC2R(t *testing.T) {
	/*
	   int f(int a, int b) {
	       int c;
	       c = 10;
	       // ....
	   }
	*/
	a := CVar{
		label:   "a",
		vartype: "int",
	}
	b := CVar{
		label:   "b",
		vartype: "int",
	}
	s1 := CVarDecl{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		initializer: "",
	}
	s2 := CAssignment{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		right: "10",
	}
	s3 := CVarDecl{
		cvar: CVar{
			label:   "d",
			vartype: "int",
		},
		initializer: "",
	}

	f := CFunction{
		label:      "f",
		args:       []CVar{a, b},
		body:       []CStatement{&s1, &s2, &s3},
		returntype: "int",
	}

	fmt.Println(f)

	rf := f.toRust()
	fmt.Println("convert c to rust")
	rf.write()
}

func TestC2R2(t *testing.T) {
	/*
	   int f(int a, int b) {
	       int c;
	       c = 10;
	       c = 20;
	       // ....
	   }
	*/
	a := CVar{
		label:   "a",
		vartype: "int",
	}
	b := CVar{
		label:   "b",
		vartype: "int",
	}
	s1 := CVarDecl{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		initializer: "",
	}
	s2 := CAssignment{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		right: "10",
	}
	s3 := CAssignment{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		right: "20",
	}

	f := CFunction{
		label:      "f",
		args:       []CVar{a, b},
		body:       []CStatement{&s1, &s2, &s3},
		returntype: "int",
	}

	fmt.Println(f)

	rf := f.toRust()
	fmt.Println("convert c to rust")
	rf.write()
}

func TestC2R3(t *testing.T) {
	/*
	   int f(int a, int b) {
	       int c;
	       c = 10;
	       c = 20;
	       int d;
	       d = 10;
	       d = 20;
	       // ....
	   }
	*/
	a := CVar{
		label:   "a",
		vartype: "int",
	}
	b := CVar{
		label:   "b",
		vartype: "int",
	}
	s1 := CVarDecl{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		initializer: "",
	}
	s2 := CAssignment{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		right: "10",
	}
	s3 := CAssignment{
		cvar: CVar{
			label:   "c",
			vartype: "int",
		},
		right: "20",
	}
	s4 := CVarDecl{
		cvar: CVar{
			label:   "d",
			vartype: "int",
		},
		initializer: "",
	}
	s5 := CAssignment{
		cvar: CVar{
			label:   "d",
			vartype: "int",
		},
		right: "10",
	}
	s6 := CAssignment{
		cvar: CVar{
			label:   "d",
			vartype: "int",
		},
		right: "20",
	}

	f := CFunction{
		label:      "f",
		args:       []CVar{a, b},
		body:       []CStatement{&s1, &s2, &s3, &s4, &s5, &s6},
		returntype: "int",
	}

	fmt.Println(f)

	rf := f.toRust()
	fmt.Println("convert c to rust")
	rf.write()
}

func TestC2R4(t *testing.T) {
	/*
	   int f(int a, int b) {
	       int c;
	       c = 10;
	       c = 20;
	       // ....
	   }
	*/
	a := CVar{
		label:   "a",
		vartype: "int",
	}
	b := CVar{
		label:   "b",
		vartype: "int",
	}
	s1 := CArray{
		label:        "a",
		arraytype:    "int",
		size:         "3",
		initializer:  "{4,5,6}",
		label_filled: true,
		size_filled:  true,
	}
	s2 := CAssignment{
		cvar: CVar{
			label:   "a",
			vartype: "int",
		},
		right:    "10",
		operator: "=",
	}

	f := CFunction{
		label:      "f",
		args:       []CVar{a, b},
		body:       []CStatement{&s1, &s2},
		returntype: "int",
	}

	fmt.Println(f)

	rf := f.toRust()
	fmt.Println("convert c to rust")
	rf.write()
}
