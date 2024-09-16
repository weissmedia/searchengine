package search

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/weissmedia/searchengine/generated/sqparser"
	"github.com/weissmedia/searchengine/internal/backend"
	"github.com/weissmedia/searchengine/internal/core"
	"github.com/weissmedia/searchengine/internal/profiler"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"sort"
	"strconv"
	"strings"
)

type Executor struct {
	*sqparser.BaseSearchQueryVisitor
	ctx           context.Context
	ResultSet     []string
	backend       backend.SearchBackend
	log           *zap.Logger
	profiler      *profiler.Profiler
	executionLogs []string
}

func NewExecutor(ctx context.Context, backend backend.SearchBackend, logger *zap.Logger, profiler *profiler.Profiler) *Executor {
	logger.Info("Creating new search engine instance...")
	return &Executor{
		BaseSearchQueryVisitor: &sqparser.BaseSearchQueryVisitor{},
		ctx:                    ctx,
		backend:                backend,
		log:                    logger,
		profiler:               profiler,
	}
}

// Execute the query and return results along with timing information
func (r *Executor) Execute(tree antlr.ParseTree) (*core.ExecutionResult, error) {
	// Clear the executionLogs at the start of each execution
	r.executionLogs = []string{}

	// Visit the tree and get the result set
	resultSet := r.Visit(tree).([]string)

	// Return the result along with logs and timings
	return &core.ExecutionResult{
		ResultSet:          resultSet,
		ResultCount:        len(resultSet),
		Timings:            r.profiler.GetTimings(),
		TotalExecutionTime: r.profiler.GetTotalExecutionTime(),
		Log:                r.executionLogs,
	}, nil
}

func (r *Executor) Visit(tree antlr.ParseTree) any {
	if tree == nil {
		r.log.Warn("Attempted to visit nil node")
		return nil
	}
	r.log.Debug("Visiting node", zap.String("type", fmt.Sprintf("%T", tree)))

	// Profiling den Besuch des Baumes
	defer r.profiler.DeferTiming("Visit ParseTree")()

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
	defer r.profiler.DeferTiming("VisitQuery")()

	// Profiling das Besuchen des Ausdrucks
	resultSet := r.Visit(ctx.Expression()).(map[string]struct{})
	r.log.Debug("Initial result from expression", zap.Any("resultSet", resultSet))

	// Konvertiere Set zu Slice
	defer r.profiler.DeferTiming("ConvertSetToSlice")()
	r.ResultSet = r.convertSetToSlice(resultSet)
	r.log.Debug("Result after convertSetToSlice", zap.Strings("resultSet", r.ResultSet))

	// Profiling für Sortierung, falls Sort-Klausel vorhanden ist
	if sortCtx, ok := ctx.Sort_clause().(*sqparser.Sort_clauseContext); ok {
		r.ResultSet = r.VisitSort_clause(sortCtx).([]string)
		r.log.Debug("Result after sorting", zap.Strings("resultSet", r.ResultSet))
	}

	// Profiling für Offset, falls vorhanden
	if offsetCtx, ok := ctx.Offset_clause().(*sqparser.Offset_clauseContext); ok {
		r.ResultSet = r.VisitOffset_clause(offsetCtx).([]string)
		r.log.Debug("Result after applying OFFSET", zap.Strings("resultSet", r.ResultSet))
	}

	// Profiling für Limit, falls vorhanden
	if limitCtx, ok := ctx.Limit_clause().(*sqparser.Limit_clauseContext); ok {
		r.ResultSet = r.VisitLimit_clause(limitCtx).([]string)
		r.log.Debug("Result after applying LIMIT", zap.Strings("resultSet", r.ResultSet))
	}

	r.log.Debug("Final result after sorting, offset, and limit", zap.Strings("resultSet", r.ResultSet))
	return r.ResultSet
}

func (r *Executor) VisitAndExpression(ctx *sqparser.AndExpressionContext) any {
	defer r.profiler.DeferTiming("VisitAndExpression")()

	var finalResult map[string]struct{}
	var expressions []string

	for i := 0; i < len(ctx.AllComparisonExpression()); i++ {
		// Measure the time for each condition
		expression := ctx.ComparisonExpression(i).GetText()
		r.profiler.TimeOperation(fmt.Sprintf("AND Condition: %s", expression), func() {
			set := r.Visit(ctx.ComparisonExpression(i))
			resultSet, ok := set.(map[string]struct{})
			if !ok || len(resultSet) == 0 {
				r.log.Warn("Empty set for expression", zap.String("expression", expression))
				r.executionLogs = append(r.executionLogs, fmt.Sprintf("Empty set for expression: %s", expression))
				finalResult = map[string]struct{}{}
				return
			}

			if finalResult == nil {
				finalResult = resultSet
			} else {
				for elem := range finalResult {
					if _, found := resultSet[elem]; !found {
						delete(finalResult, elem)
					}
				}
			}
		})

		// Capture the expression for the output
		expressions = append(expressions, expression)
	}

	// Output of the AND conditions
	r.log.Debug("AND Expression",
		zap.String("expression", strings.Join(expressions, " AND")),
	)

	return finalResult
}

func (r *Executor) VisitOrExpression(ctx *sqparser.OrExpressionContext) any {
	defer r.profiler.DeferTiming("VisitOrExpression")()

	finalResult := make(map[string]struct{})
	var expressions []string

	for i := 0; i < len(ctx.AllAndExpression()); i++ {
		expression := ctx.AndExpression(i).GetText()
		r.profiler.TimeOperation(fmt.Sprintf("OR Condition: %s", expression), func() {
			set := r.Visit(ctx.AndExpression(i))
			resultSet, ok := set.(map[string]struct{})
			if ok && len(resultSet) > 0 {
				for elem := range resultSet {
					finalResult[elem] = struct{}{}
				}
			}
		})

		expressions = append(expressions, expression)
	}

	// Output of the OR conditions
	r.log.Debug("OR Expression",
		zap.String("expression", strings.Join(expressions, " OR")),
	)

	return finalResult
}

func (r *Executor) VisitComparisonExpression(ctx *sqparser.ComparisonExpressionContext) any {
	defer r.profiler.DeferTiming("VisitComparisonExpression")()

	result := r.Visit(ctx.Primary())
	r.log.Debug("Result from Primary in Comparison Expression", zap.Any("result", result))
	return result
}

func (r *Executor) VisitPrimary(ctx *sqparser.PrimaryContext) any {
	defer r.profiler.DeferTiming("VisitPrimary")()

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
	defer r.profiler.DeferTiming("VisitCondition")()

	identifier := ctx.IDENTIFIER().GetText()

	// IN query Profiling
	if ctx.IN() != nil {
		inList := r.Visit(ctx.InList())
		inValues, ok := inList.([]string)
		if !ok || len(inValues) == 0 {
			r.log.Warn("No valid values found for 'IN' clause", zap.String("identifier", identifier))
			return nil
		}

		// Profiling for Redis IN query
		result := r.profiler.TimeOperationWithReturn("GetMap for IN query", func() interface{} {
			result, err := r.backend.GetMap(r.ctx, identifier, inValues)
			if err != nil {
				r.log.Error("Error getting map for 'IN' clause", zap.String("identifier", identifier), zap.Error(err))
				return nil
			}
			r.log.Debug("Returning result for IN condition", zap.Any("result", result))
			return result
		}).(map[string]struct{})

		return result
	}

	// Profiling for Fuzzy query
	if ctx.FUZZY() != nil {
		value := strings.Trim(ctx.QUOTED_LITERAL().GetText(), "'")

		if value == "" {
			r.log.Warn("Fuzzy value missing", zap.String("identifier", identifier))
			return nil
		}

		// Profiling for a Fuzzy query
		resultSet := r.profiler.TimeOperationWithReturn("SearchFuzzyMap", func() interface{} {
			resultSet, err := r.backend.SearchFuzzyMap(identifier, value)
			if err != nil {
				r.log.Error("Error during fuzzy search", zap.String("identifier", identifier), zap.Error(err))
				return nil
			}
			return resultSet
		}).(map[string]struct{})

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
	defer r.profiler.DeferTiming("VisitSort_clause")()

	fields := core.NewSortFieldList(r.log)

	// Profiling for adding sort fields
	r.profiler.TimeOperation("AddSortFields", func() {
		for i, identifierCtx := range ctx.AllIDENTIFIER() {
			order := "ASC"
			if ctx.ASC(i) != nil {
				order = "ASC"
			} else if ctx.DESC(i) != nil {
				order = "DESC"
			}
			fields.AddSortField(identifierCtx.GetText(), order)
		}
	})

	// Time taken to query the sorting values from Redis
	resultChan := r.profiler.TimeOperationWithReturn("GetSortedFieldValuesMap", func() interface{} {
		resultChan, err := r.backend.GetSortedFieldValuesMap(r.ctx, fields)
		if err != nil {
			r.log.Error("Error getting sorted field values", zap.Error(err))
			return nil
		}
		return resultChan // return the receive-only channel ←chan core.SortResult
	}).(<-chan core.SortResult) // Correctly assert as a receive-only channel

	results := make([]core.SortResult, fields.Len())

	// Record time for mapping the sorting results
	r.profiler.TimeOperation("MapSortResults", func() {
		for result := range resultChan {
			results[result.Index] = result
		}
	})

	comparators := make([]func(id1, id2 string) int, 0, fields.Len())

	// define comparison functions
	r.profiler.TimeOperation("SetupComparators", func() {
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
					if asc {
						return compareIntsAsc(val1, val2)
					}
					return compareIntsDesc(val1, val2)
				})
			case core.StringType:
				comparators = append(comparators, func(id1, id2 string) int {
					val1, val2 := orderMap[id1].(string), orderMap[id2].(string)
					if asc {
						return compareStringsAsc(val1, val2)
					}
					return compareStringsDesc(val1, val2)
				})
			}
		}
	})

	// Sort the ResultSet
	r.profiler.TimeOperation("SortResults", func() {
		sort.SliceStable(r.ResultSet, func(i, j int) bool {
			id1, id2 := r.ResultSet[i], r.ResultSet[j]
			for _, comparator := range comparators {
				if result := comparator(id1, id2); result != 0 {
					return result < 0
				}
			}
			return false
		})
	})

	return r.ResultSet
}

func (r *Executor) VisitOffset_clause(ctx *sqparser.Offset_clauseContext) any {
	defer r.profiler.DeferTiming("VisitOffset_clause")()

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
	defer r.profiler.DeferTiming("VisitLimit_clause")()

	limit, _ := strconv.Atoi(ctx.NUMBER().GetText()) // Convert the LIMIT number
	if limit < len(r.ResultSet) {
		r.ResultSet = r.ResultSet[:limit]
	}

	r.log.Debug("Result after applying LIMIT", zap.Strings("resultSet", r.ResultSet))
	return r.ResultSet
}

func (r *Executor) VisitRangeExpression(ctx *sqparser.RangeExpressionContext) any {
	defer r.profiler.DeferTiming(fmt.Sprintf("VisitRangeExpression: %s", ctx.GetText()))()

	// Extract start and end values from the range condition
	startNumberStr := ctx.NUMBER(0).GetText()
	endNumberStr := ctx.NUMBER(1).GetText()

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

	// Log and output the range
	r.log.Debug("Range Expression",
		zap.String("expression", ctx.GetText()), // Output of the expression as a string
		zap.Int("start", startNumber),           // Start value as an integer
		zap.Int("end", endNumber),               // End value as an intege
	)

	return []int{startNumber, endNumber}
}

// Auxiliary function for converting map[string]struct{} to []string with profiling
func (r *Executor) convertSetToSlice(set map[string]struct{}) []string {
	defer r.profiler.DeferTiming("ConvertSetToSlice")()

	result := make([]string, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	return result
}

// compareIntsAsc compares two integers in ascending order
func compareIntsAsc(val1, val2 int) int {
	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

// compareIntsDesc compares two integers in descending order
func compareIntsDesc(val1, val2 int) int {
	if val1 > val2 {
		return -1
	} else if val1 < val2 {
		return 1
	}
	return 0
}

// compareStringsAsc compares two strings in ascending order
func compareStringsAsc(val1, val2 string) int {
	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

// compareStringsDesc compares two strings in descending order
func compareStringsDesc(val1, val2 string) int {
	if val1 > val2 {
		return -1
	} else if val1 < val2 {
		return 1
	}
	return 0
}
