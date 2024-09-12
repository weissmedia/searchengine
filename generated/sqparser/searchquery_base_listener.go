// Code generated from SearchQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package sqparser // SearchQuery

import "github.com/antlr4-go/antlr/v4"

// BaseSearchQueryListener is a complete listener for a parse tree produced by SearchQueryParser.
type BaseSearchQueryListener struct{}

var _ SearchQueryListener = &BaseSearchQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSearchQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSearchQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSearchQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSearchQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSearchQueryListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSearchQueryListener) ExitQuery(ctx *QueryContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSearchQueryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSearchQueryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseSearchQueryListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseSearchQueryListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseSearchQueryListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseSearchQueryListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseSearchQueryListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseSearchQueryListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseSearchQueryListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseSearchQueryListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseSearchQueryListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseSearchQueryListener) ExitCondition(ctx *ConditionContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSearchQueryListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSearchQueryListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSearchQueryListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSearchQueryListener) ExitValue(ctx *ValueContext) {}

// EnterRangeExpression is called when production rangeExpression is entered.
func (s *BaseSearchQueryListener) EnterRangeExpression(ctx *RangeExpressionContext) {}

// ExitRangeExpression is called when production rangeExpression is exited.
func (s *BaseSearchQueryListener) ExitRangeExpression(ctx *RangeExpressionContext) {}

// EnterInList is called when production inList is entered.
func (s *BaseSearchQueryListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production inList is exited.
func (s *BaseSearchQueryListener) ExitInList(ctx *InListContext) {}

// EnterInValue is called when production inValue is entered.
func (s *BaseSearchQueryListener) EnterInValue(ctx *InValueContext) {}

// ExitInValue is called when production inValue is exited.
func (s *BaseSearchQueryListener) ExitInValue(ctx *InValueContext) {}

// EnterSort_clause is called when production sort_clause is entered.
func (s *BaseSearchQueryListener) EnterSort_clause(ctx *Sort_clauseContext) {}

// ExitSort_clause is called when production sort_clause is exited.
func (s *BaseSearchQueryListener) ExitSort_clause(ctx *Sort_clauseContext) {}

// EnterLimit_clause is called when production limit_clause is entered.
func (s *BaseSearchQueryListener) EnterLimit_clause(ctx *Limit_clauseContext) {}

// ExitLimit_clause is called when production limit_clause is exited.
func (s *BaseSearchQueryListener) ExitLimit_clause(ctx *Limit_clauseContext) {}

// EnterOffset_clause is called when production offset_clause is entered.
func (s *BaseSearchQueryListener) EnterOffset_clause(ctx *Offset_clauseContext) {}

// ExitOffset_clause is called when production offset_clause is exited.
func (s *BaseSearchQueryListener) ExitOffset_clause(ctx *Offset_clauseContext) {}
