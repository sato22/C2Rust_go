package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestMyListener1(t *testing.T) {
	is := antlr.NewInputStream(`
	f(int a, int b) {
	       int c;
	       int d;
	       c = 10;
	       c = 20;
	}
	`)
	lexer := NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCParser(stream)
	listener := NewMyListener()
	p.AddParseListener(listener)
	tree := p.FunctionDefinition()
	fmt.Println(tree.ToStringTree([]string{}, p))

	fmt.Println(listener.current_f) // current_fを確認
	// antlrで作成したfunc構造体を変換して表示
	rf := listener.current_f.toRust()
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
