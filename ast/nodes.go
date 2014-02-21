package ast

import (
	"fmt"
)

// Node encapsulates every AST node.
type Node interface {

// An Identifier is specifically a variable in code.
type Identifier struct {
	Name string
	Type Type
}

// EvaluatesTo returns the known type of the variable.
func (i Identifier) EvaluatesTo() Type {
	return i.Type
}

// NewIdentifier intializes an identifier node with its type set to AnyType.
func NewIdentifier(name string) Identifier {
	return Identifier{name, AnyType}
}

// A statement is an executable piece of code. It may be as simple as
// a function call or a variable assignment. It also includes things like
// "if".
type Statement interface {
	Node
}

// An Expression is a snippet of code that evaluates to a single value when run
// and does not represent a program instruction.
type Expression interface {
	EvaluatesTo() Type
}

// AnyType is a bitmask of all the valid types.
const AnyType = String | Integer | Float | Boolean | Null | Resource | Array | Object

// OperatorExpression is an expression that applies an operator to one, two, or three
// operands. The operator determines how many operands it should contain.
type OperatorExpression struct {
	Operand1 Expression
	Operand2 Expression
	Operand3 Expression
	Type     Type
	Operator string
}

func (o OperatorExpression) String() string {
	if o.Operand2 == nil {
		return fmt.Sprintf("(%s%v~%v)", o.Operand1, o.Operator, o.Type)
	}
	if o.Operand3 == nil {
		return fmt.Sprintf("(%s %v %s~%v)", o.Operand1, o.Operator, o.Operand2, o.Type)
	}
	return fmt.Sprintf("(%s ? %s : %s~%v)", o.Operand1, o.Operand2, o.Operand3, o.Type)
}

func (o OperatorExpression) EvaluatesTo() Type {
	return o.Type
}

// Echo returns a new echo statement.
func Echo(expr Expression) EchoStmt {
	return EchoStmt{Expression: expr}
}

// Echo represents an echo statement. It may be either a literal statement
// or it may be from data outside PHP-mode, such as "here" in: <? not here ?> here <? not here ?>
type EchoStmt struct {
	Expression Expression
}

// ReturnStmt represents a function return.
type ReturnStmt struct {
	Expression
}

// AssignmentStmt represents an assignment.
type AssignmentStmt struct {
	Assignee Identifier
	Value    Expression
}

type FunctionCallStmt struct {
	FunctionCallExpression
}

type FunctionCallExpression struct {
	FunctionName string
	Arguments    []Expression
}

func (f FunctionCallExpression) EvaluatesTo() Type {
	return String | Integer | Float | Boolean | Null | Resource | Array | Object
}

type Block struct {
	Statements []Statement
	Scope      Scope
}

type FunctionStmt struct {
	FunctionDefinition
	Body Block
}

type FunctionDefinition struct {
	Name      string
	Arguments []FunctionArgument
}

type FunctionArgument struct {
	TypeHint   string
	Identifier Identifier
}

type Class struct {
	Name       string
	Extends    *Class
	Implements []*Interface
	Methods    []Method
}

type Constant struct {
	Identifier
	Value interface{}
}

type Interface struct {
	Methods []FunctionDefinition
}

type Method struct {
	*FunctionStmt
	Visibility Visibility
}

type Visibility int

const (
	Private Visibility = iota
	Protected
	Public
)

type IfStmt struct {
	Condition  Expression
	TrueBlock  Statement
	FalseBlock Statement
}

type ForStmt struct {
	Initialization Expression
	Termination    Expression
	Iteration      Expression
	LoopBlock      Block
}

type WhileStmt struct {
	Termination Expression
	LoopBlock   Block
}

type DoWhileStmt struct {
	Termination Expression
	LoopBlock   Block
}

type TryStmt struct {
	TryBlock     *Block
	FinallyBlock *Block
	CatchStmts   []*CatchStmt
}

type CatchStmt struct {
	CatchBlock Block
	CatchType  Type
}

type Literal struct {
	Type Type
}

func (l Literal) EvaluatesTo() Type {
	return l.Type
}

type ForeachStmt struct {
	Source    Expression
	Key       *Identifier
	Value     Identifier
	LoopBlock Block
}
