package search

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/weissmedia/searchengine/generated/sqparser"
	"github.com/weissmedia/searchengine/internal/backend"
	"github.com/weissmedia/searchengine/internal/core"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"sort"
	"strconv"
	"strings"
)

type Executor struct {
	*sqparser.BaseSearchQueryVisitor
	ctx       context.Context
	ResultSet []string
	backend   backend.SearchBackend
	log       *zap.Logger
}

func NewExecutor(ctx context.Context, backend backend.SearchBackend, logger *zap.Logger) *Executor {
	logger.Info("Creating new search engine instance...")
	return &Executor{
		BaseSearchQueryVisitor: &sqparser.BaseSearchQueryVisitor{},
		ctx:                    ctx,
		backend:                backend,
		log:                    logger,
	}
}

// Execute executes the processing and converts the final ResultSet to []string
func (r *Executor) Execute(tree antlr.ParseTree) ([]string, error) {
	// Visit the parsed syntax tree (ParseTree) and get the ResultSet as a map[string]struct{}
	resultSet := r.Visit(tree).([]string)
	return resultSet, nil
}

func (r *Executor) Visit(tree antlr.ParseTree) any {
	if tree == nil {
		r.log.Warn("Attempted to visit nil node")
		return nil
	}
	r.log.Debug("Visiting node", zap.String("type", fmt.Sprintf("%T", tree)))
	return tree.Accept(r)
}

func (r *Executor) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	r.log.Error("Visiting ErrorNode")
	return nil
}

func (r *Executor) VisitTerminal(_ antlr.TerminalNode) interface{} {
	r.log.Debug("Visiting TerminalNode")
	return nil
}

func (r *Executor) VisitChildren(tree antlr.RuleNode) any {
	if tree == nil {
		r.log.Error("Tree is nil")
		return nil
	}

	n := tree.GetChildCount()
	r.log.Debug("Visiting children", zap.Int("count", n))

	for i := 0; i < n; i++ {
		c := tree.GetChild(i)
		val, ok := c.(antlr.ParseTree)
		if !ok {
			r.log.Warn("Child is not a ParseTree", zap.Int("index", i))
			continue
		}
		_ = r.Visit(val)
	}
	return 0
}

func (r *Executor) VisitExpression(ctx *sqparser.ExpressionContext) any {
	r.log.Debug("Visiting Expression")
	return r.Visit(ctx.OrExpression())
}

func (r *Executor) VisitQuery(ctx *sqparser.QueryContext) any {
	r.log.Debug("Visiting Query")

	// Visit the place and get the results
	resultSet := r.Visit(ctx.Expression()).(map[string]struct{})
	r.log.Debug("Initial result from expression", zap.Any("resultSet", resultSet))

	r.ResultSet = convertSetToSlice(resultSet)
	r.log.Debug("Result after convertSetToSlice", zap.Strings("resultSet", r.ResultSet))

	// If a sort clause is present, call the sort logic
	if sortCtx, ok := ctx.Sort_clause().(*sqparser.Sort_clauseContext); ok {
		r.ResultSet = r.VisitSort_clause(sortCtx).([]string)
		r.log.Debug("Result after sorting", zap.Strings("resultSet", r.ResultSet))
	}

	// If an OFFSET clause is present, call the OFFSET logic
	if offsetCtx, ok := ctx.Offset_clause().(*sqparser.Offset_clauseContext); ok {
		r.ResultSet = r.VisitOffset_clause(offsetCtx).([]string)
		r.log.Debug("Result after applying OFFSET", zap.Strings("resultSet", r.ResultSet))
	}

	// If a LIMIT clause is present, call the LIMIT logic.
	if limitCtx, ok := ctx.Limit_clause().(*sqparser.Limit_clauseContext); ok {
		r.ResultSet = r.VisitLimit_clause(limitCtx).([]string)
		r.log.Debug("Result after applying LIMIT", zap.Strings("resultSet", r.ResultSet))
	}

	r.log.Debug("Final result after sorting, offset, and limit", zap.Strings("resultSet", r.ResultSet))
	return r.ResultSet
}

func (r *Executor) VisitAndExpression(ctx *sqparser.AndExpressionContext) any {
	r.log.Debug("Visiting And Expression")

	var finalResult map[string]struct{}

	for i := 0; i < len(ctx.AllComparisonExpression()); i++ {
		set := r.Visit(ctx.ComparisonExpression(i))
		r.log.Debug("Set for expression", zap.Int("index", i), zap.Any("set", set))
		resultSet, ok := set.(map[string]struct{})
		if !ok || len(resultSet) == 0 {
			r.log.Warn("Empty set for expression", zap.String("expression", ctx.ComparisonExpression(i).GetText()))
			return map[string]struct{}{}
		}

		if finalResult == nil {
			finalResult = resultSet
		} else {
			for elem := range finalResult {
				if _, found := resultSet[elem]; !found {
					delete(finalResult, elem)
				}
			}
			r.log.Debug("Updated final result after intersection", zap.Any("finalResult", finalResult))
		}
	}

	r.log.Debug("Final AND intersection result", zap.Any("finalResult", finalResult))
	return finalResult
}

func (r *Executor) VisitOrExpression(ctx *sqparser.OrExpressionContext) any {
	r.log.Debug("Visiting Or Expression")
	finalResult := make(map[string]struct{})

	for i := 0; i < len(ctx.AllAndExpression()); i++ {
		set := r.Visit(ctx.AndExpression(i))

		resultSet, ok := set.(map[string]struct{})
		if ok && len(resultSet) > 0 {
			for elem := range resultSet {
				finalResult[elem] = struct{}{}
			}
			r.log.Debug("Added non-empty set for OR condition", zap.Any("resultSet", resultSet))
		} else {
			r.log.Warn("Skipping empty set for OR condition", zap.Int("index", i))
		}
	}

	r.log.Debug("Final OR union result", zap.Any("finalResult", finalResult))
	return finalResult
}

func (r *Executor) VisitComparisonExpression(ctx *sqparser.ComparisonExpressionContext) any {
	r.log.Debug("Visiting Comparison Expression")
	result := r.Visit(ctx.Primary())
	r.log.Debug("Result from Primary in Comparison Expression", zap.Any("result", result))
	return result
}

func (r *Executor) VisitPrimary(ctx *sqparser.PrimaryContext) any {
	r.log.Debug("Visiting Primary")
	if ctx.LPAREN() != nil {
		result := r.Visit(ctx.Expression())
		r.log.Debug("Processed expression in parentheses", zap.String("expression", ctx.GetText()), zap.Any("result", result))
		return result
	}

	if ctx.Condition() != nil {
		result := r.Visit(ctx.Condition())
		r.log.Debug("Processed condition", zap.String("condition", ctx.GetText()), zap.Any("result", result))
		return result
	}

	return nil
}

func (r *Executor) VisitValue(ctx *sqparser.ValueContext) any {
	if ctx.QUOTED_LITERAL() != nil {
		return strings.Trim(ctx.QUOTED_LITERAL().GetText(), "'")
	}
	if ctx.LITERAL() != nil {
		return ctx.LITERAL().GetText()
	}
	if ctx.NUMBER() != nil {
		// Attempts to convert the string value of ctx.NUMBER() to an int
		numberStr := ctx.NUMBER().GetText()
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			r.log.Error("Error converting number", zap.Error(err))
			return nil
		}
		return number
	}
	if ctx.WILDCARD() != nil {
		return strings.Trim(ctx.WILDCARD().GetText(), "'")
	}
	if ctx.RangeExpression() != nil {
		// Here you can either return the range as a string or process the RangeExpression further
		return r.VisitRangeExpression(ctx.RangeExpression().(*sqparser.RangeExpressionContext))
	}
	return nil
}

func (r *Executor) VisitCondition(ctx *sqparser.ConditionContext) any {
	identifier := ctx.IDENTIFIER().GetText()

	// IN query
	if ctx.IN() != nil {
		inList := r.Visit(ctx.InList()) // Use the IInListContext interface here
		inValues, ok := inList.([]string)
		if !ok || len(inValues) == 0 {
			r.log.Warn("No valid values found for 'IN' clause", zap.String("identifier", identifier))
			return nil
		}
		result, err := r.backend.GetMap(r.ctx, identifier, inValues)
		if err != nil {
			r.log.Error("Error getting map for 'IN' clause", zap.String("identifier", identifier), zap.Error(err))
			return nil
		}
		r.log.Debug("Returning result for IN condition", zap.Any("result", result))
		return result
	}

	// Fuzzy query
	if ctx.FUZZY() != nil {
		value := strings.Trim(ctx.QUOTED_LITERAL().GetText(), "'")

		if value == "" {
			r.log.Warn("Fuzzy value missing", zap.String("identifier", identifier))
			return nil
		}

		resultSet, err := r.backend.SearchFuzzyMap(identifier, value)
		if err != nil {
			r.log.Error("Error during fuzzy search", zap.String("identifier", identifier), zap.Error(err))
			return nil
		}

		r.log.Debug("Found data for fuzzy search", zap.Any("resultSet", resultSet))
		return resultSet
	}

	valueCtx, ok := ctx.Value().(*sqparser.ValueContext)
	if !ok {
		r.log.Error("Error: Could not cast ctx.Value() to *ValueContext", zap.String("identifier", identifier))
		return nil
	}

	// Wildcard query
	if ctx.Value().WILDCARD() != nil {
		if value, ok := r.VisitValue(valueCtx).(string); ok {
			resultSet, err := r.backend.SearchWildcardMap(identifier, value)
			if err != nil {
				r.log.Error("Error executing wildcard search", zap.Error(err))
				return nil
			}

			r.log.Debug("Processed '!=' condition", zap.String("identifier", identifier), zap.Any("resultSet", resultSet))
			return resultSet
		}
	}

	// Range query
	if ctx.Value().RangeExpression() != nil {
		rangeValues, ok := r.VisitValue(valueCtx).([]int)
		if ok && len(rangeValues) == 2 {
			startValue := rangeValues[0]
			endValue := rangeValues[1]

			resultSet, err := r.backend.SearchRangeMap(identifier, startValue, endValue)
			fmt.Println("ZZZZZZZZ", resultSet)
			if err != nil {
				r.log.Error("Error searching range", zap.String("identifier", identifier), zap.Error(err))
				return nil
			}

			r.log.Debug("Found data for range", zap.String("identifier", identifier), zap.Int("start", startValue), zap.Int("end", endValue), zap.Any("resultSet", resultSet))
			return resultSet
		} else {
			r.log.Warn("Invalid range values", zap.Any("rangeValues", rangeValues))
		}
	}

	// NOT EQUALS query
	if ctx.NOT_EQUALS() != nil {
		if value, ok := r.VisitValue(valueCtx).(string); ok && value != "" {
			resultSet, err := r.backend.GetMapExcluding(r.ctx, identifier, value)
			if err != nil {
				r.log.Error("Error getting values for '!=' condition", zap.String("identifier", identifier), zap.Error(err))
				return nil
			}
			r.log.Debug("Processed '!=' condition", zap.String("identifier", identifier), zap.Any("resultSet", resultSet))
			return resultSet
		}
	}

	// EQUALS query
	if ctx.EQUALS() != nil {
		if value, ok := r.VisitValue(valueCtx).(string); ok && value != "" {
			resultSet, err := r.backend.GetMap(r.ctx, identifier, value)
			if err != nil {
				r.log.Error("Error processing value", zap.String("value", value), zap.Error(err))
				return nil
			}
			r.log.Debug("Processed '=' condition", zap.String("identifier", identifier), zap.Any("resultSet", resultSet))
			return resultSet
		}
	}

	// relational operators
	if ctx.ComparisonOperator() != nil {
		if comparisonCtx, ok := ctx.ComparisonOperator().(sqparser.IComparisonOperatorContext); ok {
			operator, err := sqparser.DetermineComparisonOperator(comparisonCtx)
			if err != nil {
				r.log.Error("Error determining comparison operator", zap.Error(err))
				return nil
			}

			if value, ok := r.VisitValue(valueCtx).(int); ok {
				resultSet, err := r.backend.SearchComparisonMap(identifier, operator, value)
				if err != nil {
					r.log.Error("Error searching comparison", zap.String("identifier", identifier), zap.Error(err))
					return nil
				}
				r.log.Debug("Found data for comparison", zap.String("identifier", identifier), zap.Any("resultSet", resultSet))
				return resultSet
			} else {
				r.log.Warn("Invalid value type for comparison", zap.String("value", valueCtx.GetText()))
				return nil
			}
		}
	}

	r.log.Warn("Unsupported or incomplete condition", zap.String("condition", ctx.GetText()))
	return nil
}

func (r *Executor) VisitInList(ctx *sqparser.InListContext) any {
	var values []string
	for _, inValueCtx := range ctx.AllInValue() {
		value := r.Visit(inValueCtx).(string)
		values = append(values, value)
	}
	return values
}

func (r *Executor) VisitInValue(ctx *sqparser.InValueContext) any {
	if ctx.QUOTED_LITERAL() != nil {
		return strings.Trim(ctx.QUOTED_LITERAL().GetText(), "'")
	}
	if ctx.LITERAL() != nil {
		return ctx.LITERAL().GetText()
	}
	return ""
}

func (r *Executor) VisitSort_clause(ctx *sqparser.Sort_clauseContext) any {
	fields := core.NewSortFieldList(r.log)
	for i, identifierCtx := range ctx.AllIDENTIFIER() {
		order := "ASC"
		if ctx.ASC(i) != nil {
			order = "ASC"
		} else if ctx.DESC(i) != nil {
			order = "DESC"
		}
		fields.AddSortField(identifierCtx.GetText(), order)
	}

	resultChan, err := r.backend.GetSortedFieldValuesMap(r.ctx, fields.SortFields())
	if err != nil {
		r.log.Error("Error getting sorted field values", zap.Error(err))
	}
	results := make([]core.SortResult, fields.Len())
	for result := range resultChan {
		results[result.Index] = result
	}

	comparators := make([]func(id1, id2 string) int, 0, fields.Len())

	for _, result := range results {
		field := fields.GetSortField(result.Field)
		orderMap := result.OrderMap
		orderMapType := result.OrderMapType
		if len(orderMap) == 0 {
			r.log.Warn("Sort field not available for sorting", zap.String("field", field.Name))
			continue
		}

		asc := field.Order == core.Asc

		switch orderMapType {
		case core.IntType:
			comparators = append(comparators, func(id1, id2 string) int {
				val1, val2 := orderMap[id1].(int), orderMap[id2].(int)
				if val1 == val2 {
					return 0
				}
				if asc {
					if val1 < val2 {
						return -1
					}
					return 1
				}
				if val1 > val2 {
					return -1
				}
				return 1
			})
		case core.StringType:
			comparators = append(comparators, func(id1, id2 string) int {
				val1, val2 := orderMap[id1].(string), orderMap[id2].(string)
				if val1 == val2 {
					return 0
				}
				if asc {
					if val1 < val2 {
						return -1
					}
					return 1
				}
				if val1 > val2 {
					return -1
				}
				return 1
			})
		}
	}

	if len(comparators) == 0 {
		r.log.Warn("No valid sort fields provided, returning original order")
		return r.ResultSet
	}

	sort.SliceStable(r.ResultSet, func(i, j int) bool {
		id1, id2 := r.ResultSet[i], r.ResultSet[j]
		for _, comparator := range comparators {
			if result := comparator(id1, id2); result != 0 {
				return result < 0
			}
		}
		return false
	})

	return r.ResultSet
}

func (r *Executor) VisitOffset_clause(ctx *sqparser.Offset_clauseContext) any {
	r.log.Debug("Visiting Offset Clause")

	offset, _ := strconv.Atoi(ctx.NUMBER().GetText()) // Convert the OFFSET number
	if offset < len(r.ResultSet) {
		r.ResultSet = r.ResultSet[offset:]
	} else {
		r.ResultSet = []string{} // If the offset is greater than the length of the result
	}

	r.log.Debug("Result after applying OFFSET", zap.Strings("resultSet", r.ResultSet))
	return r.ResultSet
}

func (r *Executor) VisitLimit_clause(ctx *sqparser.Limit_clauseContext) any {
	r.log.Debug("Visiting Limit Clause")

	limit, _ := strconv.Atoi(ctx.NUMBER().GetText()) // Convert the LIMIT number
	if limit < len(r.ResultSet) {
		r.ResultSet = r.ResultSet[:limit]
	}

	r.log.Debug("Result after applying LIMIT", zap.Strings("resultSet", r.ResultSet))
	return r.ResultSet
}

func (r *Executor) VisitRangeExpression(ctx *sqparser.RangeExpressionContext) any {
	r.log.Debug("Visiting RangeExpression")

	// Extract the numbers from the context
	startNumberStr := ctx.NUMBER(0).GetText() // first number
	endNumberStr := ctx.NUMBER(1).GetText()   // Second number

	startNumber, err := strconv.Atoi(startNumberStr)
	if err != nil {
		r.log.Error("Error converting start number", zap.Error(err))
		return nil
	}

	endNumber, err := strconv.Atoi(endNumberStr)
	if err != nil {
		r.log.Error("Error converting end number", zap.Error(err))
		return nil
	}

	rangeValues := []int{startNumber, endNumber}

	r.log.Debug("Range expression", zap.Int("start", startNumber), zap.Int("end", endNumber))

	return rangeValues
}

// Auxiliary function for converting map[string]struct{} to []string
func convertSetToSlice(set map[string]struct{}) []string {
	result := make([]string, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	return result
}
