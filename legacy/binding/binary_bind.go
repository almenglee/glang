package binding

import (
	"main/legacy/binding/boundNode"
	"main/legacy/expression"
	"main/legacy/syntax"
)

func (b *Binder) BindBinaryExpression(exp syntax.ExpressionSyntax) boundNode.BoundExpression {
	biExp := exp.(*expression.BinaryExpressionSyntax)
	left := b.Bind(biExp.Left)
	right := b.Bind(biExp.Right)
	operKind := BindBinaryOperator(biExp.Kind(), left.Type(), right.Type())
	if operKind == IllegalBinaryOperator {
		oper := biExp.OperatorToken
		Tleft, Tright := left.Type().String(), right.Type().String()
		b.Diag.UndefinedBinaryOperator(oper.Span, oper.Text, Tleft, Tright)
		return left
	}
	return NewBoundBinaryExpression(left, operKind, right)
}
