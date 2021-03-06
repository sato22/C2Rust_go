package parser

import (
	// "fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	//"strings"
)

// 引数の値が変更される場合の処理は？
// return分の処理は？　→　あくまでも変数定義の話だから、形式的な部分は気にしなくていいのかな

type MyListener struct {
	*BaseCListener                 // embedded
	current_arg        CVarDecl    // 引数(CVar)
	current_decl       CVarDecl    // 変数
	current_assignment CAssignment // 代入文
	current_f          CFunction   // CFunction
	body               []interface{}
	stack              Stack // primaryexpressionの値を格納するスタック
	f                  bool  // trueであればCFunctionに代入
	arg                bool  // trueであれば引数として
	decl               bool  // trueであればbody内の変数として扱う
	right              bool  // trueであれば右辺を扱っている
}

// constructor
func NewMyListener() *MyListener {
	return &MyListener{
		BaseCListener: new(BaseCListener),
	}
}

// override

// VisitTerminal is called when a terminal node is visited.
func (s *MyListener) VisitTerminal(node antlr.TerminalNode) {
	log.Printf("VisitTerminal") // "test" を visitterminal という風に関数名にするとどこで呼ばれているかタイミングがわかります
}

// VisitErrorNode is called when an error node is visited.
func (s *MyListener) VisitErrorNode(node antlr.ErrorNode) {
	log.Printf("VisitErrorNode")
}

// EnterEveryRule is called when any rule is entered.
func (s *MyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	log.Printf("EnterEveryRule")
}

// ExitEveryRule is called when any rule is exited.
func (s *MyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	log.Printf("ExitEveryRule")
}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *MyListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {
	log.Printf("EnterPrimaryExpression")
}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *MyListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {
	log.Printf("ExitPrimaryExpression %s", ctx.GetText())
	switch ctx.primarytype {
	case 1:
		log.Printf("primarytype = 1")
		if !s.right {
			s.current_assignment.label = ctx.GetText()
			log.Printf("---s.current_assignment.label = %s---", s.current_assignment.label)
		}
	case 2:
		log.Printf("primarytype = 2")
	case 3:
		log.Printf("primarytype = 3")
	default:
	}

}

// EnterGenericSelection is called when production genericSelection is entered.
func (s *MyListener) EnterGenericSelection(ctx *GenericSelectionContext) {
	log.Printf("EnterGenericSelection")
}

// ExitGenericSelection is called when production genericSelection is exited.
func (s *MyListener) ExitGenericSelection(ctx *GenericSelectionContext) {
	log.Printf("ExitGenericSelection")
}

// EnterGenericAssocList is called when production genericAssocList is entered.
func (s *MyListener) EnterGenericAssocList(ctx *GenericAssocListContext) {
	log.Printf("EnterGenericAssocList")
}

// ExitGenericAssocList is called when production genericAssocList is exited.
func (s *MyListener) ExitGenericAssocList(ctx *GenericAssocListContext) {
	log.Printf("ExitGenericAssocList")
}

// EnterGenericAssociation is called when production genericAssociation is entered.
func (s *MyListener) EnterGenericAssociation(ctx *GenericAssociationContext) {
	log.Printf("EnterGenericAssociation")
}

// ExitGenericAssociation is called when production genericAssociation is exited.
func (s *MyListener) ExitGenericAssociation(ctx *GenericAssociationContext) {
	log.Printf("ExitGenericAssociation")
}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *MyListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {
	log.Printf("EnterPostfixExpression")
}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *MyListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {
	log.Printf("ExitPostfixExpression")
	switch ctx.postfixtype {
	case 1, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		// nothing
	case 2:
		s.current_assignment.size = append(s.current_assignment.size, s.stack.Pop())
	default:
	}
}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *MyListener) EnterArgumentExpressionList(ctx *ArgumentExpressionListContext) {
	log.Printf("EnterArgumentExpressionList")
}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *MyListener) ExitArgumentExpressionList(ctx *ArgumentExpressionListContext) {
	log.Printf("ExitArgumentExpressionList")
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *MyListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {
	log.Printf("EnterUnaryExpression")
}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *MyListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {
	log.Printf("ExitUnaryExpression")
}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *MyListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {
	log.Printf("EnterUnaryOperator")
}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *MyListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {
	log.Printf("ExitUnaryOperator")
}

// EnterCastExpression is called when production castExpression is entered.
func (s *MyListener) EnterCastExpression(ctx *CastExpressionContext) {
	log.Printf("EnterCastExpression")
}

// ExitCastExpression is called when production castExpression is exited.
func (s *MyListener) ExitCastExpression(ctx *CastExpressionContext) {
	log.Printf("ExitCastExpression")
}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *MyListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
	log.Printf("EnterMultiplicativeExpression")
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *MyListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
	log.Printf("ExitMultiplicativeExpression")
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *MyListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {
	log.Printf("EnterAdditiveExpression")
}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *MyListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {
	log.Printf("ExitAdditiveExpression %s", ctx.GetText())
}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *MyListener) EnterShiftExpression(ctx *ShiftExpressionContext) {
	log.Printf("EnterShiftExpression")
}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *MyListener) ExitShiftExpression(ctx *ShiftExpressionContext) {
	log.Printf("ExitShiftExpression")
}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *MyListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {
	log.Printf("EnterRelationalExpression")
}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *MyListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {
	log.Printf("ExitRelationalExpression")
}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *MyListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {
	log.Printf("EnterEqualityExpression %s", ctx.GetText())
}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *MyListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {
	log.Printf("ExitEqualityExpression %s", ctx.GetText())
}

// EnterAndExpression is called when production andExpression is entered.
func (s *MyListener) EnterAndExpression(ctx *AndExpressionContext) {
	log.Printf("EnterAndExpression")
}

// ExitAndExpression is called when production andExpression is exited.
func (s *MyListener) ExitAndExpression(ctx *AndExpressionContext) {
	log.Printf("ExitAndExpression")
}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *MyListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {
	log.Printf("EnterExclusiveOrExpression")
}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *MyListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {
	log.Printf("ExitExclusiveOrExpression")
}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *MyListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {
	log.Printf("EnterInclusiveOrExpression")
}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *MyListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {
	log.Printf("ExitInclusiveOrExpression")
}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *MyListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {
	log.Printf("EnterLogicalAndExpression")
}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *MyListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {
	log.Printf("ExitLogicalAndExpression")
}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *MyListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {
	log.Printf("EnterLogicalOrExpression")
}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *MyListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {
	log.Printf("ExitLogicalOrExpression")
}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *MyListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {
	log.Printf("EnterConditionalExpression")
	s.right = true
}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *MyListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {
	log.Printf("ExitConditionalExpression")
	s.right = false
}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *MyListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {
	log.Printf("EnterAssignmentExpression")
}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *MyListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {
	log.Printf("ExitAssignmentExpression %s", ctx.GetText())
	switch ctx.assignmenttype {
	case 1:
		s.stack.Push(ctx.GetText())
		log.Printf("stack = %v", s.stack)
	case 2:
		s.current_assignment.right = s.stack.Pop()
		log.Printf("---s.current_assignment.right = %s---", ctx.GetText())
		if s.current_assignment.label != "" {
			s.body = append(s.body, s.current_assignment)
			log.Printf("---s.current_assignment is %v---", s.current_assignment)
		}
	case 3:
		// nothing
	default:
	}

}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *MyListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {
	log.Printf("EnterAssignmentOperator")
}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *MyListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {
	log.Printf("ExitAssignmentOperator %s", ctx.GetText())
	s.current_assignment.operator = ctx.GetText()
	log.Printf("---s.current_assignment.operator = %s---", s.current_assignment.operator)
}

// EnterExpression is called when production expression is entered.
func (s *MyListener) EnterExpression(ctx *ExpressionContext) {
	log.Printf("EnterExpression")
}

// ExitExpression is called when production expression is exited.
func (s *MyListener) ExitExpression(ctx *ExpressionContext) {
	log.Printf("ExitExpression")
}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *MyListener) EnterConstantExpression(ctx *ConstantExpressionContext) {
	log.Printf("EnterConstantExpression")
}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *MyListener) ExitConstantExpression(ctx *ConstantExpressionContext) {
	log.Printf("ExitConstantExpression")
}

// EnterDeclaration is called when production declaration is entered.
func (s *MyListener) EnterDeclaration(ctx *DeclarationContext) {
	log.Printf("EnterDeclaration")
	s.current_decl = CVarDecl{
		label:       "",
		vartype:     "",
		initializer: "",
		dim:         0,
		size:        []string{},
	}
}

// ExitDeclaration is called when production declaration is exited.
func (s *MyListener) ExitDeclaration(ctx *DeclarationContext) {
	log.Printf("ExitDeclaration %s", ctx.GetText())
	s.body = append(s.body, s.current_decl)
	log.Printf("---s.current_decl is %v---", s.current_decl)
}

// EnterDeclarationSpecifiers is called when production declarationSpecifiers is entered.
func (s *MyListener) EnterDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {
	log.Printf("EnterDeclarationSpecifiers")
}

//  ExitDeclarationSpecifiers is called when production declarationSpecifiers is exited.
func (s *MyListener) ExitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {
	log.Printf("ExitDeclarationSpecifiers %s", ctx.GetText())
}

// EnterDeclarationSpecifiers2 is called when production declarationSpecifiers2 is entered.
func (s *MyListener) EnterDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) {
	log.Printf("EnterDeclarationSpecifiers2")
}

// ExitDeclarationSpecifiers2 is called when production declarationSpecifiers2 is exited.
func (s *MyListener) ExitDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) {
	log.Printf("ExitDeclarationSpecifiers2")
}

// EnterDeclarationSpecifier is called when production declarationSpecifier is entered.
func (s *MyListener) EnterDeclarationSpecifier(ctx *DeclarationSpecifierContext) {
	log.Printf("EnterDeclarationSpecifier")
}

// ExitDeclarationSpecifier is called when production declarationSpecifier is exited.
func (s *MyListener) ExitDeclarationSpecifier(ctx *DeclarationSpecifierContext) {
	log.Printf("ExitDeclarationSpecifier %s", ctx.GetText())
}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *MyListener) EnterInitDeclaratorList(ctx *InitDeclaratorListContext) {
	log.Printf("EnterInitDeclaratorList")
}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *MyListener) ExitInitDeclaratorList(ctx *InitDeclaratorListContext) {
	log.Printf("ExitInitDeclaratorList %s", ctx.GetText())
}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *MyListener) EnterInitDeclarator(ctx *InitDeclaratorContext) {
	log.Printf("EnterInitDeclarator")
	s.decl = true
	log.Printf("---s.decl = true---")
}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *MyListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {
	log.Printf("ExitInitDeclarator %s", ctx.GetText())
	s.decl = false
	log.Printf("---s.decl = false---")
}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *MyListener) EnterStorageClassSpecifier(ctx *StorageClassSpecifierContext) {
	log.Printf("EnterStorageClassSpecifier")
}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *MyListener) ExitStorageClassSpecifier(ctx *StorageClassSpecifierContext) {
	log.Printf("ExitStorageClassSpecifier")
}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *MyListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {
	log.Printf("EnterTypeSpecifier")
}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *MyListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {
	log.Printf("ExitTypeSpecifier %s", ctx.id.GetText())

	if s.f {
		s.current_f.returntype += ctx.GetText()
		log.Printf("---s.current_f.returntype is %s---", s.current_f.returntype)
	}

	if s.arg {
		s.current_arg.vartype += ctx.GetText()
		log.Printf("---s.current_arg.vartype is %s---", s.current_arg.vartype)
	}

	if s.decl {
		s.current_decl.vartype += ctx.GetText()
		log.Printf("---s.current_decl.vartype is %s---", s.current_decl.vartype)
	}
}

// EnterStructOrUnionSpecifier is called when production structOrUnionSpecifier is entered.
func (s *MyListener) EnterStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {
	log.Printf("EnterStructOrUnionSpecifier")
}

// ExitStructOrUnionSpecifier is called when production structOrUnionSpecifier is exited.
func (s *MyListener) ExitStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {
	log.Printf("ExitStructOrUnionSpecifier")
}

// EnterStructOrUnion is called when production structOrUnion is entered.
func (s *MyListener) EnterStructOrUnion(ctx *StructOrUnionContext) {
	log.Printf("EnterStructOrUnion")
}

// ExitStructOrUnion is called when production structOrUnion is exited.
func (s *MyListener) ExitStructOrUnion(ctx *StructOrUnionContext) {
	log.Printf("ExitStructOrUnion")
}

// EnterStructDeclarationList is called when production structDeclarationList is entered.
func (s *MyListener) EnterStructDeclarationList(ctx *StructDeclarationListContext) {
	log.Printf("EnterStructDeclarationList")
}

// ExitStructDeclarationList is called when production structDeclarationList is exited.
func (s *MyListener) ExitStructDeclarationList(ctx *StructDeclarationListContext) {
	log.Printf("ExitStructDeclarationList")
}

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *MyListener) EnterStructDeclaration(ctx *StructDeclarationContext) {
	log.Printf("EnterStructDeclaration")
}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *MyListener) ExitStructDeclaration(ctx *StructDeclarationContext) {
	log.Printf("ExitStructDeclaration")
}

// EnterSpecifierQualifierList is called when production specifierQualifierList is entered.
func (s *MyListener) EnterSpecifierQualifierList(ctx *SpecifierQualifierListContext) {
	log.Printf("EnterSpecifierQualifierList")
}

// ExitSpecifierQualifierList is called when production specifierQualifierList is exited.
func (s *MyListener) ExitSpecifierQualifierList(ctx *SpecifierQualifierListContext) {
	log.Printf("ExitSpecifierQualifierList")
}

// EnterStructDeclaratorList is called when production structDeclaratorList is entered.
func (s *MyListener) EnterStructDeclaratorList(ctx *StructDeclaratorListContext) {
	log.Printf("EnterStructDeclaratorList")
}

// ExitStructDeclaratorList is called when production structDeclaratorList is exited.
func (s *MyListener) ExitStructDeclaratorList(ctx *StructDeclaratorListContext) {
	log.Printf("ExitStructDeclaratorList")
}

// EnterStructDeclarator is called when production structDeclarator is entered.
func (s *MyListener) EnterStructDeclarator(ctx *StructDeclaratorContext) {
	log.Printf("EnterStructDeclarator")
}

// ExitStructDeclarator is called when production structDeclarator is exited.
func (s *MyListener) ExitStructDeclarator(ctx *StructDeclaratorContext) {
	log.Printf("ExitStructDeclarator")
}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *MyListener) EnterEnumSpecifier(ctx *EnumSpecifierContext) {
	log.Printf("EnterEnumSpecifier")
}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *MyListener) ExitEnumSpecifier(ctx *EnumSpecifierContext) {
	log.Printf("ExitEnumSpecifier")
}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *MyListener) EnterEnumeratorList(ctx *EnumeratorListContext) {
	log.Printf("EnterEnumeratorList")
}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *MyListener) ExitEnumeratorList(ctx *EnumeratorListContext) {
	log.Printf("ExitEnumeratorList")
}

// EnterEnumerator is called when production enumerator is entered.
func (s *MyListener) EnterEnumerator(ctx *EnumeratorContext) {
	log.Printf("EnterEnumerator")
}

// ExitEnumerator is called when production enumerator is exited.
func (s *MyListener) ExitEnumerator(ctx *EnumeratorContext) {
	log.Printf("ExitAlignmentSpecifier")
}

// EnterEnumerationConstant is called when production enumerationConstant is entered.
func (s *MyListener) EnterEnumerationConstant(ctx *EnumerationConstantContext) {
	log.Printf("EnterEnumerationConstant")
}

// ExitEnumerationConstant is called when production enumerationConstant is exited.
func (s *MyListener) ExitEnumerationConstant(ctx *EnumerationConstantContext) {
	log.Printf("ExitEnumerationConstant")
}

// EnterAtomicTypeSpecifier is called when production atomicTypeSpecifier is entered.
func (s *MyListener) EnterAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) {
	log.Printf("EnterAtomicTypeSpecifier")
}

// ExitAtomicTypeSpecifier is called when production atomicTypeSpecifier is exited.
func (s *MyListener) ExitAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) {
	log.Printf("ExitAtomicTypeSpecifier")
}

// EnterTypeQualifier is called when production typeQualifier is entered.
func (s *MyListener) EnterTypeQualifier(ctx *TypeQualifierContext) {
	log.Printf("EnterTypeQualifier")
}

// ExitTypeQualifier is called when production typeQualifier is exited.
func (s *MyListener) ExitTypeQualifier(ctx *TypeQualifierContext) {
	log.Printf("ExitTypeQualifier")
}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *MyListener) EnterFunctionSpecifier(ctx *FunctionSpecifierContext) {
	log.Printf("EnterFunctionSpecifier")
}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *MyListener) ExitFunctionSpecifier(ctx *FunctionSpecifierContext) {
	log.Printf("ExitFunctionSpecifier")
}

// EnterAlignmentSpecifier is called when production alignmentSpecifier is entered.
func (s *MyListener) EnterAlignmentSpecifier(ctx *AlignmentSpecifierContext) {
	log.Printf("EnterAlignmentSpecifier")
}

// ExitAlignmentSpecifier is called when production alignmentSpecifier is exited.
func (s *MyListener) ExitAlignmentSpecifier(ctx *AlignmentSpecifierContext) {
	log.Printf("ExitAlignmentSpecifier")
}

// EnterDeclarator is called when production declarator is entered.
func (s *MyListener) EnterDeclarator(ctx *DeclaratorContext) {
	log.Printf("EnterDeclarator")
}

// ExitDeclarator is called when production declarator is exited.
func (s *MyListener) ExitDeclarator(ctx *DeclaratorContext) {
	log.Printf("ExitDeclarator %s", ctx.GetText())
}

// EnterDirectDeclarator is called when production directDeclarator is entered.
func (s *MyListener) EnterDirectDeclarator(ctx *DirectDeclaratorContext) {
	log.Printf("EnterDirectDeclarator")
}

// ctDeclarator is called when production directDeclarator is exited.
func (s *MyListener) ExitDirectDeclarator(ctx *DirectDeclaratorContext) {
	log.Printf("ExitDirectDeclarator %s", ctx.GetText())

	switch ctx.declaratortype {
	case 1:
		if s.decl {
			s.current_decl.label = ctx.GetText()
			log.Printf("---s.current_decl.label is %s---", s.current_decl.label)
		}
	case 2, 4, 5, 6:
		// nothing
	case 3:
		// array
		// log.Printf("directdeclarator type = 3")
		s.current_decl.dim++
		s.current_decl.size = append(s.current_decl.size, s.stack.Pop())
		log.Printf("---stack = %v---", s.stack)
	default:
	}

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

// EnterGccDeclaratorExtension is called when production gccDeclaratorExtension is entered.
func (s *MyListener) EnterGccDeclaratorExtension(ctx *GccDeclaratorExtensionContext) {
	log.Printf("EnterGccDeclaratorExtension")
}

// ExitGccDeclaratorExtension is called when production gccDeclaratorExtension is exited.
func (s *MyListener) ExitGccDeclaratorExtension(ctx *GccDeclaratorExtensionContext) {
	log.Printf("ExitGccDeclaratorExtension")
}

// EnterGccAttributeSpecifier is called when production gccAttributeSpecifier is entered.
func (s *MyListener) EnterGccAttributeSpecifier(ctx *GccAttributeSpecifierContext) {
	log.Printf("EnterGccAttributeSpecifier")
}

// ExitGccAttributeSpecifier is called when production gccAttributeSpecifier is exited.
func (s *MyListener) ExitGccAttributeSpecifier(ctx *GccAttributeSpecifierContext) {
	log.Printf("ExitGccAttributeSpecifier")
}

// EnterGccAttributeList is called when production gccAttributeList is entered.
func (s *MyListener) EnterGccAttributeList(ctx *GccAttributeListContext) {
	log.Printf("EnterGccAttributeList")
}

// ExitGccAttributeList is called when production gccAttributeList is exited.
func (s *MyListener) ExitGccAttributeList(ctx *GccAttributeListContext) {
	log.Printf("ExitGccAttributeList")
}

// EnterGccAttribute is called when production gccAttribute is entered.
func (s *MyListener) EnterGccAttribute(ctx *GccAttributeContext) {
	log.Printf("EnterGccAttribute")
}

// ExitGccAttribute is called when production gccAttribute is exited.
func (s *MyListener) ExitGccAttribute(ctx *GccAttributeContext) {
	log.Printf("ExitGccAttribute")
}

// EnterNestedParenthesesBlock is called when production nestedParenthesesBlock is entered.
func (s *MyListener) EnterNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) {
	log.Printf("EnterNestedParenthesesBlock")
}

// ExitNestedParenthesesBlock is called when production nestedParenthesesBlock is exited.
func (s *MyListener) ExitNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) {
	log.Printf("ExitNestedParenthesesBlock")
}

// EnterPointer is called when production pointer is entered.
func (s *MyListener) EnterPointer(ctx *PointerContext) {
	log.Printf("EnterPointer")
}

// ExitPointer is called when production pointer is exited.
func (s *MyListener) ExitPointer(ctx *PointerContext) {
	log.Printf("ExitPointer")
}

// EnterTypeQualifierList is called when production typeQualifierList is entered.
func (s *MyListener) EnterTypeQualifierList(ctx *TypeQualifierListContext) {
	log.Printf("EnterTypeQualifierList")
}

// ExitTypeQualifierList is called when production typeQualifierList is exited.
func (s *MyListener) ExitTypeQualifierList(ctx *TypeQualifierListContext) {
	log.Printf("ExitTypeQualifierList")
}

// EnterParameterTypeList is called when production parameterTypeList is entered.
func (s *MyListener) EnterParameterTypeList(ctx *ParameterTypeListContext) {
	log.Printf("EnterParameterTypeList")
}

// ExitParameterTypeList is called when production parameterTypeList is exited.
func (s *MyListener) ExitParameterTypeList(ctx *ParameterTypeListContext) {
	log.Printf("ExitParameterTypeList")
}

// EnterParameterList is called when production parameterList is entered.
func (s *MyListener) EnterParameterList(ctx *ParameterListContext) {
	log.Printf("EnterParameterList")
}

// ExitParameterList is called when production parameterList is exited.
func (s *MyListener) ExitParameterList(ctx *ParameterListContext) {
	log.Printf("ExitParameterList")
}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *MyListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {
	log.Printf("EnterParameterDeclaration")
	s.current_arg = CVarDecl{
		label:       "",
		vartype:     "",
		initializer: "",
	}
}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *MyListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {
	log.Printf("ExitParameterDeclaration %s", ctx.GetText())
	s.current_f.args = append(s.current_f.args, s.current_arg)
	log.Printf("---current_args is %v---", s.current_f.args)
}

// EnterIdentifierList is called when production identifierList is entered.
func (s *MyListener) EnterIdentifierList(ctx *IdentifierListContext) {
	log.Printf("EnterIdentifierList")
}

// ExitIdentifierList is called when production identifierList is exited.
func (s *MyListener) ExitIdentifierList(ctx *IdentifierListContext) {
	log.Printf("ExitIdentifierList")
}

// EnterTypeName is called when production typeName is entered.
func (s *MyListener) EnterTypeName(ctx *TypeNameContext) {
	log.Printf("EnterTypeName")
}

// ExitTypeName is called when production typeName is exited.
func (s *MyListener) ExitTypeName(ctx *TypeNameContext) {
	log.Printf("ExitTypeName")
}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *MyListener) EnterAbstractDeclarator(ctx *AbstractDeclaratorContext) {
	log.Printf("EnterAbstractDeclarator")
}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *MyListener) ExitAbstractDeclarator(ctx *AbstractDeclaratorContext) {
	log.Printf("ExitAbstractDeclarator")
}

// EnterDirectAbstractDeclarator is called when production directAbstractDeclarator is entered.
func (s *MyListener) EnterDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) {
	log.Printf("EnterDirectAbstractDeclarator")
}

// ExitDirectAbstractDeclarator is called when production directAbstractDeclarator is exited.
func (s *MyListener) ExitDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) {
	log.Printf("ExitDirectAbstractDeclarator")
}

// EnterTypedefName is called when production typedefName is entered.
func (s *MyListener) EnterTypedefName(ctx *TypedefNameContext) {
	log.Printf("EnterTypedefName")
}

// ExitTypedefName is called when production typedefName is exited.
func (s *MyListener) ExitTypedefName(ctx *TypedefNameContext) {
	log.Printf("ExitTypedefName")
}

// EnterInitializer is called when production initializer is entered.
func (s *MyListener) EnterInitializer(ctx *InitializerContext) {
	log.Printf("EnterInitializer")
}

// ExitInitializer is called when production initializer is exited.
func (s *MyListener) ExitInitializer(ctx *InitializerContext) {
	log.Printf("ExitInitializer %s", ctx.GetText())
	switch ctx.initializertype {
	case 1:
		s.current_decl.initializer = s.stack.Pop()
		log.Printf("---s.current_decl.initializer = %s---", ctx.GetText())
	case 2:
		s.current_decl.initializer = ctx.GetText()
		log.Printf("---s.current_decl.initializer = %s---", ctx.GetText())
	default:
	}
}

// EnterInitializerList is called when production initializerList is entered.
func (s *MyListener) EnterInitializerList(ctx *InitializerListContext) {
	log.Printf("EnterInitializerList")
}

// ExitInitializerList is called when production initializerList is exited.
func (s *MyListener) ExitInitializerList(ctx *InitializerListContext) {
	log.Printf("ExitInitializerList")
}

// EnterDesignation is called when production designation is entered.
func (s *MyListener) EnterDesignation(ctx *DesignationContext) {
	log.Printf("EnterDesignation")
}

// ExitDesignation is called when production designation is exited.
func (s *MyListener) ExitDesignation(ctx *DesignationContext) {
	log.Printf("ExitDesignation")
}

// EnterDesignatorList is called when production designatorList is entered.
func (s *MyListener) EnterDesignatorList(ctx *DesignatorListContext) {
	log.Printf("EnterDesignatorList")
}

// ExitDesignatorList is called when production designatorList is exited.
func (s *MyListener) ExitDesignatorList(ctx *DesignatorListContext) {
	log.Printf("ExitDesignatorList")
}

// EnterDesignator is called when production designator is entered.
func (s *MyListener) EnterDesignator(ctx *DesignatorContext) {
	log.Printf("EnterDesignator")
}

// ExitDesignator is called when production designator is exited.
func (s *MyListener) ExitDesignator(ctx *DesignatorContext) {
	log.Printf("ExitDesignator")
}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *MyListener) EnterStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {
	log.Printf("EnterStaticAssertDeclaration")
}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *MyListener) ExitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {
	log.Printf("ExitStaticAssertDeclaration")
}

// EnterStatement is called when production statement is entered.
func (s *MyListener) EnterStatement(ctx *StatementContext) {
	log.Printf("EnterStatement")
}

// ExitStatement is called when production statement is exited.
func (s *MyListener) ExitStatement(ctx *StatementContext) {
	log.Printf("ExitStatement")
}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *MyListener) EnterLabeledStatement(ctx *LabeledStatementContext) {
	log.Printf("EnterLabeledStatement")
}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *MyListener) ExitLabeledStatement(ctx *LabeledStatementContext) {
	log.Printf("ExitLabeledStatement")
}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *MyListener) EnterCompoundStatement(ctx *CompoundStatementContext) {
	log.Printf("EnterCompoundStatement")
}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *MyListener) ExitCompoundStatement(ctx *CompoundStatementContext) {
	log.Printf("ExitCompoundStatement")
}

// EnterBlockItemList is called when production blockItemList is entered.
func (s *MyListener) EnterBlockItemList(ctx *BlockItemListContext) {
	log.Printf("EnterBlockItemList")
	s.arg = false
	s.decl = true
}

// ExitBlockItemList is called when production blockItemList is exited.
func (s *MyListener) ExitBlockItemList(ctx *BlockItemListContext) {
	log.Printf("ExitBlockItemList")
}

// EnterBlockItem is called when production blockItem is entered.
func (s *MyListener) EnterBlockItem(ctx *BlockItemContext) {
	log.Printf("EnterBlockItem")
}

// ExitBlockItem is called when production blockItem is exited.
func (s *MyListener) ExitBlockItem(ctx *BlockItemContext) {
	log.Printf("ExitBlockItem")
}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *MyListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {
	log.Printf("EnterExpressionStatement")
	s.current_assignment = CAssignment{
		label:    "",
		size:     []string{},
		right:    "",
		operator: "",
	}
}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *MyListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {
	log.Printf("ExitExpressionStatement")
}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *MyListener) EnterSelectionStatement(ctx *SelectionStatementContext) {
	log.Printf("EnterSelectionStatement")
}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *MyListener) ExitSelectionStatement(ctx *SelectionStatementContext) {
	log.Printf("ExitSelectionStatement")
}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *MyListener) EnterIterationStatement(ctx *IterationStatementContext) {
	log.Printf("EnterIterationStatement")
}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *MyListener) ExitIterationStatement(ctx *IterationStatementContext) {
	log.Printf("ExitIterationStatement")
}

// EnterForCondition is called when production forCondition is entered.
func (s *MyListener) EnterForCondition(ctx *ForConditionContext) {
	log.Printf("EnterForCondition")
}

// ExitForCondition is called when production forCondition is exited.
func (s *MyListener) ExitForCondition(ctx *ForConditionContext) {
	log.Printf("ExitForCondition")
}

// EnterForDeclaration is called when production forDeclaration is entered.
func (s *MyListener) EnterForDeclaration(ctx *ForDeclarationContext) {
	log.Printf("EnterForDeclaration")
}

// ExitForDeclaration is called when production forDeclaration is exited.
func (s *MyListener) ExitForDeclaration(ctx *ForDeclarationContext) {
	log.Printf("ExitForDeclaration")
}

// EnterForExpression is called when production forExpression is entered.
func (s *MyListener) EnterForExpression(ctx *ForExpressionContext) {
	log.Printf("EnterForExpression")
}

// ExitForExpression is called when production forExpression is exited.
func (s *MyListener) ExitForExpression(ctx *ForExpressionContext) {
	log.Printf("ExitForExpression")
}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *MyListener) EnterJumpStatement(ctx *JumpStatementContext) {
	log.Printf("EnterJumpStatement")
}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *MyListener) ExitJumpStatement(ctx *JumpStatementContext) {
	log.Printf("ExitJumpStatement")
}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *MyListener) EnterCompilationUnit(ctx *CompilationUnitContext) {
	log.Printf("EnterCompilationUnit")
}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *MyListener) ExitCompilationUnit(ctx *CompilationUnitContext) {
	log.Printf("ExitCompilationUnit")
}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *MyListener) EnterTranslationUnit(ctx *TranslationUnitContext) {
	log.Printf("EnterTranslationUnit")
}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *MyListener) ExitTranslationUnit(ctx *TranslationUnitContext) {
	log.Printf("ExitTranslationUnit")
}

// EnterExternalDeclaration is called when production externalDeclaration is entered.
func (s *MyListener) EnterExternalDeclaration(ctx *ExternalDeclarationContext) {
	log.Printf("EnterExternalDeclaration")
}

// ExitExternalDeclaration is called when production externalDeclaration is exited.
func (s *MyListener) ExitExternalDeclaration(ctx *ExternalDeclarationContext) {
	log.Printf("ExitExternalDeclaration")
}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *MyListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {
	log.Printf("EnterFunctionDefinition %s", ctx.GetText())

	s.current_f = CFunction{}

	s.f = true
	s.arg = false
	s.decl = false
	s.stack.top = -1

}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *MyListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {
	log.Printf("ExitFunctionDefinition")
	log.Printf("---s.current_f.body is %v---", s.current_f.body)
	for _, v := range s.body {
		switch x := v.(type) {
		case CVarDecl:
			s.current_f.body = append(s.current_f.body, &x)
		case CAssignment:
			s.current_f.body = append(s.current_f.body, &x)
		default:
		}
	}
}

// EnterDeclarationList is called when production declarationList is entered.
func (s *MyListener) EnterDeclarationList(ctx *DeclarationListContext) {
	log.Printf("EnterDeclarationList")
}

// ExitDeclarationList is called when production declarationList is exited.
func (s *MyListener) ExitDeclarationList(ctx *DeclarationListContext) {
	log.Printf("ExitDeclarationList")
}
