// Code generated from SearchQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package sqparser // SearchQuery

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SearchQueryParser.
type SearchQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SearchQueryParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#rangeExpression.
	VisitRangeExpression(ctx *RangeExpressionContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#inList.
	VisitInList(ctx *InListContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#inValue.
	VisitInValue(ctx *InValueContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#sort_clause.
	VisitSort_clause(ctx *Sort_clauseContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#limit_clause.
	VisitLimit_clause(ctx *Limit_clauseContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#offset_clause.
	VisitOffset_clause(ctx *Offset_clauseContext) interface{}
}
