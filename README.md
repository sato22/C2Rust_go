# C2Rust_go
### mylistener.go 変更点

```
9～
type MyListener struct {
	*BaseCListener        // embedded
	current_block  []CVar // body内のCVarを格納するスライス
	current_decl   CVar   // body内のCVar
	current_decls  []CVarDecl
	current_args   []CVar    // 引数（CVar）を格納するスライス
	current_arg    CVar      // 引数(CVar)
	current_f      CFunction // CFunction　ここにcurrent_blockとcurrent_argsの情報をいれる
	f              bool      // trueであればCFunctionに代入
	arg            bool      // trueであればargに代入
	decl           bool      // trueであればbody内の変数として扱う
}

946～
// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *MyListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {
	log.Printf("EnterFunctionDefinition")
	s.f = true
	s.arg = false
	s.decl = false
}

824～
// EnterBlockItemList is called when production blockItemList is entered.
func (s *MyListener) EnterBlockItemList(ctx *BlockItemListContext) {
	log.Printf("EnterBlockItemList")
	s.arg = false
	s.decl = true
}

356～
// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *MyListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {
	log.Printf("ExitInitDeclarator %s", ctx.GetText())
	s.current_decl.label = ctx.GetText()
	log.Printf("---s.current_decl.label is %s---", s.current_decl.label)
}

377～
// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *MyListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {
	log.Printf("ExitTypeSpecifier %s", ctx.id.GetText())

	if s.f {
		s.current_f.returntype = ctx.id.GetText()
		log.Printf("---s.current_f.returntype is %s---", s.current_f.returntype)
	}

	if s.arg {
		s.current_arg.vartype = ctx.id.GetText()
		log.Printf("---s.current_arg.vartype is %s---", s.current_arg.vartype)
	}

	if s.decl {
		s.current_decl.vartype = ctx.id.GetText()
		log.Printf("---s.current_decl.vartype is %s---", s.current_decl.vartype)
	}

	/*
		s.current_decl.vartype = ctx.id.GetText()
		log.Printf("---vartype is %s---", s.current_decl.vartype)
	*/
}

567～
// ExitDirectDeclarator is called when production directDeclarator is exited.
func (s *MyListener) ExitDirectDeclarator(ctx *DirectDeclaratorContext) {
	log.Printf("ExitDirectDeclarator %s", ctx.GetText())

	if s.f {
		s.current_f.label = ctx.GetText()
		log.Printf("---s.current_f.label is %s---", s.current_f.label)
		s.f = false
		s.arg = true
	} else if s.arg {
		s.current_arg.label = ctx.GetText()
		log.Printf("---s.current_arg.label is %s---", s.current_arg.label)
	}
}

677～
// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *MyListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {
	log.Printf("ExitParameterDeclaration %s", ctx.GetText())
	s.current_args = append(s.current_args, s.current_arg)
	log.Printf("---current_args is %v---", s.current_args)
}

296～
// ExitDeclaration is called when production declaration is exited.
func (s *MyListener) ExitDeclaration(ctx *DeclarationContext) {
	log.Printf("ExitDeclaration %s", ctx.GetText())
	s.current_block = append(s.current_block, s.current_decl)
	log.Printf("---current_block is %v---", s.current_block)

	s1 := CVarDecl{
		cvar:        s.current_decl,
		initializer: "",
	}
	s.current_decls = append(s.current_decls, s1)

}

956～
// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *MyListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {
	log.Printf("ExitFunctionDefinition")
	log.Printf("---current_block is %v, current_args is %v---", s.current_block, s.current_args)
	/*
		f := CFunction{
			label:      "f",
			args:       []CVar{a, b},
			body:       []CStatement{&s1, &s2},
			returntype: "int",
		}
	*/
	s.current_f.args = s.current_args
	for _, x := range s.current_decls {
		s.current_f.body = append(s.current_f.body, &x)
	}
}

63～
// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *MyListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {
	log.Printf("ExitPrimaryExpression %s", ctx.GetText())
	for _, x := range s.current_decls {
		// log.Printf("%v", x)
		if ctx.GetText() == x.cvar.label {
			s.current_assignment.cvar = x.cvar
			break
		} else {

			s.current_assignment.right = ctx.GetText()
			s.current_assignments = append(s.current_assignments, s.current_assignment)
			break
		}
	}
}

```
