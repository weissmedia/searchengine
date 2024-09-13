// service/redis_query_executor.go
package search

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/weissmedia/searchengine/generated/sqparser"
	"github.com/weissmedia/searchengine/internal/backend"
	"github.com/weissmedia/searchengine/internal/core"
	"golang.org/x/net/context"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Executor struct {
	*sqparser.BaseSearchQueryVisitor
	ctx       context.Context
	ResultSet []string
	backend   backend.SearchBackend
}

func NewExecutor(ctx context.Context, backend backend.SearchBackend) *Executor {
	return &Executor{
		BaseSearchQueryVisitor: &sqparser.BaseSearchQueryVisitor{},
		ctx:                    ctx,
		backend:                backend,
	}
}

// Execute führt die Verarbeitung durch und konvertiert das finale ResultSet zu []string
func (r *Executor) Execute(tree antlr.ParseTree) ([]string, error) {
	// Besuche den geparsten Syntaxbaum (ParseTree) und erhalte das ResultSet als map[string]struct{}
	resultSet := r.Visit(tree).([]string)
	return resultSet, nil
}

func (r *Executor) Visit(tree antlr.ParseTree) any {
	if tree == nil {
		log.Println("Attempted to visit nil node")
		return nil
	}
	log.Printf("Visiting node: %T\n", tree)
	return tree.Accept(r)
}
func (r *Executor) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	log.Println("Visiting VisitErrorNode")
	return nil
}
func (r *Executor) VisitTerminal(_ antlr.TerminalNode) interface{} {
	log.Println("Visiting VisitTerminal")
	return nil
}
func (r *Executor) VisitChildren(tree antlr.RuleNode) any {
	if tree == nil {
		log.Println("Error: tree is nil")
		return nil
	}

	n := tree.GetChildCount()
	log.Printf("Visiting %d children\n", n)

	for i := 0; i < n; i++ {
		c := tree.GetChild(i)
		val, ok := c.(antlr.ParseTree)
		if !ok {
			log.Printf("Error: child %d is not a ParseTree\n", i)
			continue
		}
		_ = r.Visit(val)
	}
	return 0
}

func (r *Executor) VisitExpression(ctx *sqparser.ExpressionContext) any {
	log.Println("Visiting Expression")
	return r.Visit(ctx.OrExpression())
}

func (r *Executor) VisitQuery(ctx *sqparser.QueryContext) any {
	log.Println("Visiting Query")

	// Besuche den Ausdruck und hole die Resultate
	resultSet := r.Visit(ctx.Expression()).(map[string]struct{})
	log.Printf("Initial result from expression: %v", resultSet)

	r.ResultSet = convertSetToSlice(resultSet)
	log.Printf("Initial result from expression (convertToSlice): %v", r.ResultSet)

	// Wenn eine Sortierklausel vorhanden ist, rufe die Sortierlogik auf
	if sortCtx, ok := ctx.Sort_clause().(*sqparser.Sort_clauseContext); ok {
		r.ResultSet = r.VisitSort_clause(sortCtx).([]string)
		log.Printf("Result after sorting: %v", r.ResultSet)
	}

	// Wenn eine OFFSET-Klausel vorhanden ist, rufe die OFFSET-Logik auf
	if offsetCtx, ok := ctx.Offset_clause().(*sqparser.Offset_clauseContext); ok {
		r.ResultSet = r.VisitOffset_clause(offsetCtx).([]string)
		log.Printf("Result after applying OFFSET: %v", r.ResultSet)
	}

	// Wenn eine LIMIT-Klausel vorhanden ist, rufe die LIMIT-Logik auf
	if limitCtx, ok := ctx.Limit_clause().(*sqparser.Limit_clauseContext); ok {
		r.ResultSet = r.VisitLimit_clause(limitCtx).([]string)
		log.Printf("Result after applying LIMIT: %v", r.ResultSet)
	}

	log.Printf("Final result after sorting, offset, and limit: %v", r.ResultSet)
	return r.ResultSet
}

func (r *Executor) VisitAndExpression(ctx *sqparser.AndExpressionContext) any {
	log.Println("Visiting And Expression")

	var finalResult map[string]struct{}

	for i := 0; i < len(ctx.AllComparisonExpression()); i++ {
		set := r.Visit(ctx.ComparisonExpression(i))
		log.Printf("Set for expression %d: %v", i, set)
		resultSet, ok := set.(map[string]struct{})
		log.Printf("ResultSet for expression %d: %v", i, resultSet)
		if !ok || len(resultSet) == 0 {
			log.Printf("Encountered an empty set for expression %s", ctx.ComparisonExpression(i).GetText())
			return map[string]struct{}{} // Leere Menge zurückgeben, wenn eine Bedingung fehlschlägt
		}

		if finalResult == nil {
			finalResult = resultSet
		} else {
			for elem := range finalResult {
				if _, found := resultSet[elem]; !found {
					delete(finalResult, elem)
				}
			}
			log.Printf("Updated final result after intersection: %v", finalResult)
		}
	}

	log.Printf("Final AND intersection result: %v", finalResult)
	return finalResult
}

func (r *Executor) VisitOrExpression(ctx *sqparser.OrExpressionContext) any {
	log.Println("Visiting Or Expression")
	finalResult := make(map[string]struct{})

	for i := 0; i < len(ctx.AllAndExpression()); i++ {
		set := r.Visit(ctx.AndExpression(i))

		resultSet, ok := set.(map[string]struct{})
		log.Printf("Set for OR expression %d: %v", i, resultSet)
		if ok && len(resultSet) > 0 {
			for elem := range resultSet {
				finalResult[elem] = struct{}{}
			}
			log.Printf("Added non-empty set for OR condition: %v", resultSet)
		} else {
			log.Printf("Skipping empty set for OR condition at index %d", i)
		}
	}

	log.Printf("Final OR union result: %v", finalResult)
	return finalResult
}

func (r *Executor) VisitComparisonExpression(ctx *sqparser.ComparisonExpressionContext) any {
	log.Println("Visiting Comparison Expression")
	result := r.Visit(ctx.Primary())
	log.Printf("Result from Primary in Comparison Expression: %v", result)
	return result
}

func (r *Executor) VisitPrimary(ctx *sqparser.PrimaryContext) any {
	log.Println("Visiting Primary")
	if ctx.LPAREN() != nil {
		result := r.Visit(ctx.Expression())
		log.Printf("Processed expression in parentheses: %s, Result: %v", ctx.GetText(), result)
		return result
	}

	if ctx.Condition() != nil {
		result := r.Visit(ctx.Condition())
		log.Printf("Processed condition: %s, Result: %v", ctx.GetText(), result)
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
		// Versuche, den String-Wert von ctx.NUMBER() in einen int zu konvertieren
		numberStr := ctx.NUMBER().GetText()
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Printf("Error converting number: %v", err)
			return nil
		}
		return number
	}
	if ctx.WILDCARD() != nil {
		return strings.Trim(ctx.WILDCARD().GetText(), "'")
	}
	if ctx.RangeExpression() != nil {
		// Hier kannst du entweder den Bereich als String zurückgeben oder den RangeExpression weiterverarbeiten
		return r.VisitRangeExpression(ctx.RangeExpression().(*sqparser.RangeExpressionContext))
	}
	return nil
}

func (r *Executor) VisitCondition(ctx *sqparser.ConditionContext) any {
	identifier := ctx.IDENTIFIER().GetText()

	// IN-Abfrage
	if ctx.IN() != nil {
		inList := r.Visit(ctx.InList()) // Verwende das Interface IInListContext hier
		inValues, ok := inList.([]string)
		if !ok || len(inValues) == 0 {
			log.Printf("No valid values found for 'IN' clause with identifier %s", identifier)
			return nil
		}
		result, err := r.backend.GetMap(r.ctx, identifier, inValues)
		if err != nil {
			log.Printf("Error getting map for 'IN' clause with identifier %s: %v", identifier, err)
			return nil
		}
		log.Printf("Returning result for IN condition: %v", result)
		return result
	}

	// Fuzzy-Abfrage
	if ctx.FUZZY() != nil {
		// Extrahiere den Fuzzy-Wert aus dem Kontext
		value := strings.Trim(ctx.QUOTED_LITERAL().GetText(), "'")

		// Überprüfen, ob der Fuzzy-Wert leer ist
		if value == "" {
			log.Printf("Fuzzy value missing for identifier %s", identifier)
			return nil
		}

		// Suche in Redis nach dem Ausdruck
		resultSet, err := r.backend.SearchFuzzyMap(identifier, value)
		if err != nil {
			log.Printf("Error during fuzzy search for identifier %s: %v", identifier, err)
			return nil
		}

		// Protokolliere und gib das Ergebnis zurück
		log.Printf("Found data for fuzzy search: %v", resultSet)
		return resultSet
	}

	// Cast von IValueContext zu *ValueContext
	valueCtx, ok := ctx.Value().(*sqparser.ValueContext)
	if !ok {
		log.Printf("Error: Could not cast ctx.Value() to *ValueContext for identifier %s", identifier)
		return nil
	}

	// Wildcard-Abfrage
	if ctx.Value().WILDCARD() != nil {
		if value, ok := r.VisitValue(valueCtx).(string); ok {
			// Führe eine RedisSearch-Abfrage durch, anstatt direkt auf r.RedisData zuzugreifen.
			resultSet, err := r.backend.SearchWildcardMap(identifier, value)
			if err != nil {
				log.Printf("Error executing wildcard search: %v", err)
				return nil
			}

			log.Printf("WILDCARD Processed '!=' condition: %s, Result: %v", identifier, resultSet)
			return resultSet
		}
	}

	// Range-Abfrage
	if ctx.Value().RangeExpression() != nil {
		rangeValues, ok := r.VisitValue(valueCtx).([]int)
		if ok && len(rangeValues) == 2 {
			log.Printf("Range expression detected: %s", identifier)

			startValue := rangeValues[0]
			endValue := rangeValues[1]

			// Führe die Range-Map-Suche aus
			resultSet, err := r.backend.SearchRangeMap(identifier, startValue, endValue)
			if err != nil {
				log.Printf("Error searching range for %s: %v", identifier, err)
				return nil
			}

			log.Printf("Found data for %s: in range %d - %d: %v", identifier, startValue, endValue, resultSet)
			return resultSet
		} else {
			log.Printf("Invalid range values: %v", rangeValues)
		}
	}

	// NOT EQUALS-Abfrage
	if ctx.NOT_EQUALS() != nil {
		if value, ok := r.VisitValue(valueCtx).(string); ok && value != "" {
			resultSet, err := r.backend.GetMapExcluding(r.ctx, identifier, value)
			if err != nil {
				log.Printf("Error getting values for %s: %v", identifier, err)
				return nil
			}
			log.Printf("Processed '!=' condition: %s, Result: %v", identifier, resultSet)
			return resultSet
		}
	}

	// EQUALS-Abfrage
	if ctx.EQUALS() != nil {
		if value, ok := r.VisitValue(valueCtx).(string); ok && value != "" {
			resultSet, err := r.backend.GetMap(r.ctx, identifier, value)
			if err != nil {
				log.Printf("Error processing value %s: %v", value, err)
				return nil
			}
			log.Printf("Processed '=' condition: %s, Result: %v", identifier, resultSet)
			return resultSet
		}
	}

	// Vergleichsoperatoren
	if ctx.ComparisonOperator() != nil {
		if comparisonCtx, ok := ctx.ComparisonOperator().(sqparser.IComparisonOperatorContext); ok {

			// Bestimme den Operator
			operator, err := sqparser.DetermineComparisonOperator(comparisonCtx)
			if err != nil {
				log.Printf("Error determining comparison operator: %v", err)
				return nil
			}

			// Verarbeite den Wert (z.B. int)
			if value, ok := r.VisitValue(valueCtx).(int); ok {
				// Verarbeite den Operator mit dem Backend
				resultSet, err := r.backend.SearchComparisonMap(identifier, operator, value)
				if err != nil {
					log.Printf("Error searching comparison for %s: %v", identifier, err)
					return nil
				}
				log.Printf("Found data for comparison %s: %v", identifier, resultSet)
				return resultSet
			} else {
				log.Printf("Invalid value type for comparison: %v", valueCtx.GetText())
				return nil
			}
		}
	}

	log.Printf("Unsupported or incomplete condition: %s", ctx.GetText())
	return nil
}

func (r *Executor) VisitInList(ctx *sqparser.InListContext) any {
	var values []string
	for _, inValueCtx := range ctx.AllInValue() {
		value := r.Visit(inValueCtx).(string) // Expecting a string from VisitInValue
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
	var fields core.SortFieldList
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
		log.Printf("Error getting sorted field values: %v", err)
	}

	results := make([]core.SortResult, fields.Len())
	for result := range resultChan {
		// Setze die Ergebnisse basierend auf dem Index
		results[result.Index] = result
	}

	comparators := make([]func(id1, id2 string) int, 0, fields.Len())

	for _, result := range results {
		field := fields.GetSortField(result.Field)
		orderMap := result.OrderMap
		orderMapType := result.OrderMapType
		if len(orderMap) == 0 {
			// Log a warning message and skip this sort condition
			log.Printf("Warning: Sort field %s is not available for sorting", field)
			continue
		}

		asc := field.Order == core.Asc

		// Assign comparator function based on the attribute type and order.
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
		log.Println("Warning: No valid sort fields provided, returning original order")
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
	log.Println("Visiting Offset Clause")

	offset, _ := strconv.Atoi(ctx.NUMBER().GetText()) // Konvertiere die OFFSET-Zahl
	if offset < len(r.ResultSet) {
		r.ResultSet = r.ResultSet[offset:]
	} else {
		r.ResultSet = []string{} // Falls der Offset größer ist als die Länge des Ergebnisses
	}

	log.Printf("Result after applying OFFSET: %v", r.ResultSet)
	return r.ResultSet
}

func (r *Executor) VisitLimit_clause(ctx *sqparser.Limit_clauseContext) any {
	log.Println("Visiting Limit Clause")

	limit, _ := strconv.Atoi(ctx.NUMBER().GetText()) // Konvertiere die LIMIT-Zahl
	if limit < len(r.ResultSet) {
		r.ResultSet = r.ResultSet[:limit]
	}

	log.Printf("Result after applying LIMIT: %v", r.ResultSet)
	return r.ResultSet
}

func (r *Executor) VisitRangeExpression(ctx *sqparser.RangeExpressionContext) any {
	log.Println("VisitRangeExpression called")

	// Extrahiere die Nummern aus dem Kontext
	startNumberStr := ctx.NUMBER(0).GetText() // Erste Zahl
	endNumberStr := ctx.NUMBER(1).GetText()   // Zweite Zahl

	// Debug-Ausgaben
	log.Printf("Start number: %s", startNumberStr)
	log.Printf("End number: %s", endNumberStr)

	// Konvertiere die extrahierten Strings in int
	startNumber, err := strconv.Atoi(startNumberStr)
	if err != nil {
		log.Printf("Error converting start number: %v", err)
		return nil
	}

	endNumber, err := strconv.Atoi(endNumberStr)
	if err != nil {
		log.Printf("Error converting end number: %v", err)
		return nil
	}

	// Erstelle das Range-Array
	rangeValues := []int{startNumber, endNumber}

	// Debug-Ausgabe des finalen Bereichsausdrucks
	rangeExpr := fmt.Sprintf("[%d %d]", startNumber, endNumber)
	log.Printf("Range expression: %s", rangeExpr)

	return rangeValues
}

// Hilfsfunktion zur Konvertierung von map[string]struct{} zu []string
func convertSetToSlice(set map[string]struct{}) []string {
	result := make([]string, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	return result
}
