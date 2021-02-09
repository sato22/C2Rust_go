package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

// 配列
// 原始的な型の対応
// ポインタ

// let a: [i32,3] = {4,5,6}

/*
func TestMyListener1(t *testing.T) {
	is := antlr.NewInputStream(`
	int f(){
		int a[3][2] = {4,5,6};
		int b = 7;
		b = 8;
	}
	`)
	lexer := NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCParser(stream)
	listener := NewMyListener()
	p.AddParseListener(listener)
	tree := p.FunctionDefinition()
	fmt.Println(tree.ToStringTree([]string{}, p))

	fmt.Println("")
	fmt.Println(listener.current_f) // current_fを確認
	// antlrで作成したfunc構造体を変換して表示
	rf := listener.current_f.toRust()
	fmt.Println("convert c to rust")
	rf.write()
}
*/

func TestMyListener2(t *testing.T) {
	is := antlr.NewInputStream(`
	int f(int c){
		int a = 6;
		a = 5+6;
		int b = a+4;
		b = a + 7;
		int c = b;
	}
	`)
	lexer := NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCParser(stream)
	listener := NewMyListener()
	p.AddParseListener(listener)
	tree := p.FunctionDefinition()
	fmt.Println(tree.ToStringTree([]string{}, p))

	fmt.Println("")
	fmt.Println(listener.current_f) // current_fを確認
	// antlrで作成したfunc構造体を変換して表示
	rf := listener.current_f.toRust()
	fmt.Println("convert c to rust")
	rf.write()
}

/*
func TestMyListener3(t *testing.T) {
	fmt.Println("TestMyListener3")
}
*/
