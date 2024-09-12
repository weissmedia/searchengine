// Code generated from SearchQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package sqparser // SearchQuery

import "github.com/antlr4-go/antlr/v4"

// SearchQueryListener is a complete listener for a parse tree produced by SearchQueryParser.
type SearchQueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOrExpression is called when entering the orExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterRangeExpression is called when entering the rangeExpression production.
	EnterRangeExpression(c *RangeExpressionContext)

	// EnterInList is called when entering the inList production.
	EnterInList(c *InListContext)

	// EnterInValue is called when entering the inValue production.
	EnterInValue(c *InValueContext)

	// EnterSort_clause is called when entering the sort_clause production.
	EnterSort_clause(c *Sort_clauseContext)

	// EnterLimit_clause is called when entering the limit_clause production.
	EnterLimit_clause(c *Limit_clauseContext)

	// EnterOffset_clause is called when entering the offset_clause production.
	EnterOffset_clause(c *Offset_clauseContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOrExpression is called when exiting the orExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitRangeExpression is called when exiting the rangeExpression production.
	ExitRangeExpression(c *RangeExpressionContext)

	// ExitInList is called when exiting the inList production.
	ExitInList(c *InListContext)

	// ExitInValue is called when exiting the inValue production.
	ExitInValue(c *InValueContext)

	// ExitSort_clause is called when exiting the sort_clause production.
	ExitSort_clause(c *Sort_clauseContext)

	// ExitLimit_clause is called when exiting the limit_clause production.
	ExitLimit_clause(c *Limit_clauseContext)

	// ExitOffset_clause is called when exiting the offset_clause production.
	ExitOffset_clause(c *Offset_clauseContext)
}
