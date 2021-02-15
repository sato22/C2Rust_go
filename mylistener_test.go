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

func TestMyListener1(t *testing.T) {
	is := antlr.NewInputStream(`
	int f(float a,int b){
		int d = 1;
		d = 2;
		d = 1+3;
		unsigned int c[3] = {1,2,3};
		double s = b+1;
		c[0][1] = 1;
		c[1][2] = 2;
		float f[5] = {1,2,3,4,5};
		f[0] = 1;
		b = a+1;
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
func TestMyListener2(t *testing.T) {
	is := antlr.NewInputStream(`
	int f(int a,int b){
		int d = a+1;
		a = 1;
	}
	`)
	lexer := NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCParser(stream)
	listener := NewMyListener()
	p.AddParseListener(listener)
	tree := p.FunctionDefinition()
	fmt.Println(tree.ToStringTree([]string{}, p))
	cfunction := listener.current_f

	fmt.Println("")
	fmt.Println(cfunction) // current_fを確認
	// antlrで作成したfunc構造体を変換して表示
	rf := cfunction.toRust()
	fmt.Println("---convert c to rust---")
	rf.write()
}
*/

/*
func TestMyListener3(t *testing.T) {
	fmt.Println("TestMyListener3")
}
*/
