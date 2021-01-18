package parser

import (
	// "./c2rust"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestMyListener1(t *testing.T) {
	is := antlr.NewInputStream(`
	int f(int a, int b) {
	       int c;
	       int d;
	}
	`)
	lexer := NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCParser(stream)
	listener := NewMyListener()
	p.AddParseListener(listener)
	tree := p.CompilationUnit()
	fmt.Println(tree.ToStringTree([]string{}, p))

	// antlrで作成したfunc構造体を変換して表示
	rf := MyListener.current_f.toRust()
	fmt.Println("convert c to rust")
	rf.write()
}

/*
func TestMyListener2(t *testing.T) {
	is := antlr.NewInputStream(`
	#include <stdio.h>
	main() {
	    printf("Hello world");
	}
	`)
	lexer := NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(NewMyListener(), p.Declaration())
}

func TestMyListener3(t *testing.T) {
	fmt.Println("TestMyListener3")
}
*/
