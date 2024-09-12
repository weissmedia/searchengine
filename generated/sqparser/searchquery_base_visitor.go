// Code generated from SearchQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package sqparser // SearchQuery

import "github.com/antlr4-go/antlr/v4"

type BaseSearchQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSearchQueryVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitOrExpression(ctx *OrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitRangeExpression(ctx *RangeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitInList(ctx *InListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitInValue(ctx *InValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitSort_clause(ctx *Sort_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitLimit_clause(ctx *Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitOffset_clause(ctx *Offset_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}
