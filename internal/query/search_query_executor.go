package query

import (
	"fmt"
	"github.com/weissmedia/searchengine/generated/sqparser"
	"github.com/weissmedia/searchengine/internal/backend"
	"log"
	"sort"
	"strconv"
	"strings"
)

type SearchQueryExecutor struct {
	*sqparser.BaseSearchQueryVisitor
	backend   backend.SearchBackend
	ResultSet []string
}

func NewExecutor(searchBackend backend.SearchBackend) *SearchQueryExecutor {
	return &SearchQueryExecutor{
		BaseSearchQueryVisitor: &sqparser.BaseSearchQueryVisitor{},
		backend:                searchBackend,
	}
}

// Execute führt die Verarbeitung durch
func (r *SearchQueryExecutor) Execute(tree antlr.ParseTree) ([]string, error) {
	// Besuche den geparsten Syntaxbaum (ParseTree)
	result := r.Visit(tree).([]string) // Die resultierende Menge
	return result, nil
}

func (r *SearchQueryExecutor) Visit(tree antlr.ParseTree) any {
	if tree == nil {
		log.Println("Attempted to visit nil node")
		return nil
	}
	log.Printf("Visiting node: %T\n", tree)
	return tree.Accept(r)
}
func (r *SearchQueryExecutor) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	log.Println("Visiting VisitErrorNode")
	return nil
}
func (r *SearchQueryExecutor) VisitTerminal(_ antlr.TerminalNode) interface{} {
	log.Println("Visiting VisitTerminal")
	return nil
}
func (r *SearchQueryExecutor) VisitChildren(tree antlr.RuleNode) any {
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

func (r *SearchQueryExecutor) _VisitQuery(ctx *sqparser.QueryContext) any {
	log.Println("Visiting Query")
	result := r.Visit(ctx.Expression())
	log.Printf("Result after visiting Query: %v\n", result)
	return result
}

func (r *SearchQueryExecutor) VisitExpression(ctx *sqparser.ExpressionContext) any {
	log.Println("Visiting Expression")
	return r.Visit(ctx.OrExpression())
}

func (r *SearchQueryExecutor) VisitQuery(ctx *sqparser.QueryContext) any {
	log.Println("Visiting Query")

	// Besuche den Ausdruck und hole die Resultate
	r.ResultSet = r.Visit(ctx.Expression()).([]string)
	log.Printf("Initial result from expression: %v", r.ResultSet)

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

	// Gib das finale ResultSet nach Sortierung, Offset und Limit zurück
	log.Printf("Final result after sorting, offset, and limit: %v", r.ResultSet)
	return r.ResultSet
}

func (r *SearchQueryExecutor) VisitOrExpression(ctx *sqparser.OrExpressionContext) any {
	log.Println("Visiting Or Expression")
	var results [][]string

	for i := 0; i < len(ctx.AllAndExpression()); i++ {
		setList := r.Visit(ctx.AndExpression(i))

		if setList != nil {
			resultList := convertToSet(setList)
			if len(resultList) > 0 {
				log.Printf("Adding non-empty set for OR condition: %v", resultList)
				results = append(results, resultList)
			} else {
				log.Printf("Skipping empty set for OR condition at index %d", i)
			}
		}
	}

	finalResult := unionSets(results)
	log.Printf("Final OR union result: %v", finalResult)
	return finalResult
}

func (r *SearchQueryExecutor) VisitAndExpression(ctx *sqparser.AndExpressionContext) any {
	log.Println("Visiting And Expression")
	var results [][]string

	for i := 0; i < len(ctx.AllComparisonExpression()); i++ {
		currentResult := r.Visit(ctx.ComparisonExpression(i))

		resultList, ok := currentResult.([]string)
		if !ok || resultList == nil || len(resultList) == 0 {
			log.Printf("Encountered an empty set for expression %s", ctx.ComparisonExpression(i).GetText())
			return []string{} // Leere Menge zurückgeben, wenn eine Bedingung fehlschlägt
		}

		results = append(results, resultList)
	}

	finalResult := intersectSets(results)
	log.Printf("Final AND intersection result: %v", finalResult)
	return finalResult
}

func (r *SearchQueryExecutor) VisitComparisonExpression(ctx *sqparser.ComparisonExpressionContext) any {
	log.Println("Visiting Comparison Expression")
	return r.Visit(ctx.Primary())
}

func (r *SearchQueryExecutor) _VisitPrimary(ctx *sqparser.PrimaryContext) any {
	log.Println("Visiting Primary")
	if ctx.Condition() != nil {
		return r.Visit(ctx.Condition())
	}
	return nil
}

func (r *SearchQueryExecutor) VisitPrimary(ctx *sqparser.PrimaryContext) any {
	log.Println("Visiting Primary")
	if ctx.LPAREN() != nil {
		result := r.Visit(ctx.Expression())
		log.Printf("Processed expression in parentheses: %s, Result: %v", ctx.GetText(), result)
		return result
	}

	if ctx.Condition() != nil {
		return r.Visit(ctx.Condition())
	}
	return nil
}

func (r *SearchQueryExecutor) _VisitCondition(ctx *sqparser.ConditionContext) any {
	identifier := ctx.IDENTIFIER().GetText()
	var value string

	// Verarbeite die '!=' Bedingung
	if ctx.NOT_EQUALS() != nil {
		value = trimQuotes(ctx.Value().QUOTED_LITERAL().GetText())
		redisKey := fmt.Sprintf("%s:%s", identifier, value)
		log.Printf("Visiting '!=' Condition: %s != %s (redisKey: %s)", identifier, value, redisKey)

		allKeys := r.getAllKeysForIdentifier(identifier)
		matchingSet := r.RedisData[redisKey]

		resultSet := subtractSets(allKeys, matchingSet)
		log.Printf("Final result for '!=' condition: %v", resultSet)
		return resultSet
	}

	// Verarbeite die '=' Bedingung
	if ctx.EQUALS() != nil {
		value = trimQuotes(ctx.Value().QUOTED_LITERAL().GetText())
		redisKey := fmt.Sprintf("%s:%s", identifier, value)
		log.Printf("Visiting '=' Condition: %s = %s (redisKey: %s)", identifier, value, redisKey)

		resultSet := r.RedisData[redisKey]
		log.Printf("Found data for %s: %v", redisKey, resultSet)
		return resultSet
	}

	// Verarbeite die 'IN' Bedingung
	if ctx.IN() != nil {
		setList := make(map[string]struct{})
		for _, inValueCtx := range ctx.InList().AllInValue() {
			value := trimQuotes(inValueCtx.GetText())
			redisKey := fmt.Sprintf("%s:%s", identifier, value)

			if resultSet, found := r.RedisData[redisKey]; found {
				log.Printf("Found data for %s: %v", redisKey, resultSet)
				for _, item := range resultSet {
					setList[item] = struct{}{}
				}
			}
		}

		finalResult := make([]string, 0, len(setList))
		for item := range setList {
			finalResult = append(finalResult, item)
		}
		log.Printf("Final result for IN condition: %v", finalResult)
		return finalResult
	}

	if ctx.Value() != nil {
		log.Printf("Value is not nil: %s", ctx.Value().GetText())

		if ctx.Value().RangeExpression() != nil {
			log.Printf("Range expression detected: %s", ctx.Value().GetText())

			// Typ assertion von IRangeExpressionContext zu *RangeExpressionContext
			if rangeExprCtx, ok := ctx.Value().RangeExpression().(*sqparser.RangeExpressionContext); ok {
				rangeExpr := r.VisitRangeExpression(rangeExprCtx)

				// Verarbeite den Bereichsausdruck
				redisExpression := fmt.Sprintf("@%s:%s", identifier, rangeExpr)
				log.Printf("Processing range expression: %s", redisExpression)

				// Hole die Resultate aus Redis für den Bereichsausdruck
				if resultSet, found := r.RedisData[redisExpression]; found {
					log.Printf("Retrieved set for %s: %v", redisExpression, resultSet)
					return resultSet
				} else {
					log.Printf("No data found for range expression: %s", redisExpression)
				}
			} else {
				log.Println("Error: Could not assert RangeExpressionContext type")
			}
		} else {
			log.Println("No RangeExpression detected")
		}
	} else {
		log.Println("Error: ctx.Value() is nil")
	}

	if ctx.ComparisonOperator() != nil {
		log.Println("Visiting Comparison Operator")
		value := ctx.Value().GetText()
		op := ctx.ComparisonOperator().GetText()
		var redisExpression string
		switch op {
		case ">=":
			redisExpression = fmt.Sprintf("@%s:[%s +inf]", identifier, value)
		case ">":
			redisExpression = fmt.Sprintf("@%s:[(%s +inf]", identifier, value)
		case "<=":
			redisExpression = fmt.Sprintf("@%s:[-inf %s]", identifier, value)
		case "<":
			redisExpression = fmt.Sprintf("@%s:[-inf (%s]", identifier, value)
		}
		log.Printf("Redis Range expression: %v", redisExpression)
		// Ruft die Daten aus der Redis-Datenbank ab
		resultSet := r.RedisData[redisExpression]
		log.Printf("Retrieved set for %s: %v", redisExpression, resultSet)
		return resultSet
	}

	log.Printf("No matching condition found for: %s", identifier)
	return nil
}

func (r *SearchQueryExecutor) VisitValue(ctx *sqparser.ValueContext) any {
	if ctx.QUOTED_LITERAL() != nil {
		return strings.Trim(ctx.QUOTED_LITERAL().GetText(), "'")
	}
	if ctx.LITERAL() != nil {
		return ctx.LITERAL().GetText()
	}
	if ctx.NUMBER() != nil {
		return ctx.NUMBER().GetText()
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

func (r *SearchQueryExecutor) VisitCondition(ctx *sqparser.ConditionContext) any {
	identifier := ctx.IDENTIFIER().GetText()

	if ctx.IN() != nil {
		// Verwende VisitInList, um die Werte zu extrahieren
		inList := r.Visit(ctx.InList()) // Verwende das Interface IInListContext hier
		inValues, ok := inList.([]string)
		if !ok || len(inValues) == 0 {
			log.Printf("No valid values found for 'IN' clause with identifier %s", identifier)
			return nil
		}

		setList := make(map[string]struct{})
		for _, value := range inValues {
			redisKey := fmt.Sprintf("%s:%s", identifier, value)

			if resultSet, found := r.RedisData[redisKey]; found {
				log.Printf("Found data for %s: %v", redisKey, resultSet)
				for _, item := range resultSet {
					setList[item] = struct{}{}
				}
			}
		}

		// Erzeuge das finale Ergebnis aus setList
		finalResult := make([]string, 0, len(setList))
		for item := range setList {
			finalResult = append(finalResult, item)
		}

		log.Printf("Final result for IN condition: %v", finalResult)
		return finalResult
	}

	// Fuzzy-Abfrage
	if ctx.FUZZY() != nil {
		log.Println("Fuzzy comparison")
		fuzzyValue := strings.Trim(ctx.QUOTED_LITERAL().GetText(), "'")
		if fuzzyValue == "" {
			log.Printf("Fuzzy value missing for identifier %s", identifier)
			return nil
		}
		redisExpression := fmt.Sprintf("@%s:%%%%%s%%%%", identifier, fuzzyValue)
		log.Printf("Fuzzy comparison detected: %s", redisExpression)
		resultSet := r.RedisData[redisExpression]
		log.Printf("Found data for %s: %v", redisExpression, resultSet)
		return resultSet
	}

	// Cast von IValueContext zu *ValueContext
	valueCtx, ok := ctx.Value().(*sqparser.ValueContext)
	if !ok {
		log.Printf("Error: Could not cast ctx.Value() to *ValueContext for identifier %s", identifier)
		return nil
	}

	// Verwende VisitValue, um den Wert zu erhalten
	value, ok := r.VisitValue(valueCtx).(string)
	if !ok || value == "" {
		log.Printf("No valid value found for identifier %s", identifier)
		return nil
	}

	// Wildcard-Abfrage
	if ctx.Value().WILDCARD() != nil {
		redisExpression := fmt.Sprintf("@%s:%s", identifier, value)
		log.Printf("Wildcard search detected: %s", redisExpression)
		resultSet := r.RedisData[redisExpression]
		log.Printf("Found data for %s: %v", redisExpression, resultSet)
		return resultSet
	}

	// Range Expression
	if ctx.Value().RangeExpression() != nil {
		log.Printf("Range expression detected: %s", ctx.Value().GetText())
		redisExpression := fmt.Sprintf("@%s:%s", identifier, value)
		log.Printf("Processing range expression: %s", redisExpression)
		resultSet := r.RedisData[redisExpression]
		log.Printf("Found data for %s: %v", redisExpression, resultSet)
		return resultSet
	}

	// NOT EQUALS-Abfrage
	if ctx.NOT_EQUALS() != nil && value != "" {
		allKeys := r.getAllKeysForIdentifier(identifier)
		matchingSet := r.RedisData[fmt.Sprintf("%s:%s", identifier, value)]
		result := subtractSets(allKeys, matchingSet)
		log.Printf("Processed '!=' condition: %s, Result: %v", ctx.GetText(), result)
		return result
	}

	// EQUALS-Abfrage
	if ctx.EQUALS() != nil && value != "" {
		redisExpression := fmt.Sprintf("%s:%s", identifier, value)
		log.Printf("Processed '=' condition: %s", redisExpression)
		resultSet := r.RedisData[redisExpression]
		log.Printf("Found data for %s: %v", redisExpression, resultSet)
		return resultSet
	}

	// Vergleichsoperatoren
	if ctx.ComparisonOperator() != nil && value != "" {
		op := ctx.ComparisonOperator().GetText()
		var redisExpression string
		switch op {
		case ">=":
			redisExpression = fmt.Sprintf("@%s:[%s +inf]", identifier, value)
		case ">":
			redisExpression = fmt.Sprintf("@%s:[(%s +inf]", identifier, value)
		case "<=":
			redisExpression = fmt.Sprintf("@%s:[-inf %s]", identifier, value)
		case "<":
			redisExpression = fmt.Sprintf("@%s:[-inf (%s]", identifier, value)
		}
		log.Printf("Processed comparison operator: %s", redisExpression)
		resultSet := r.RedisData[redisExpression]
		log.Printf("Found data for %s: %v", redisExpression, resultSet)
		return resultSet
	}

	// IN-Abfrage
	if ctx.IN() != nil {
		inList := make([]string, 0)
		for _, inValue := range ctx.InList().AllInValue() {
			inList = append(inList, fmt.Sprintf("%s:%s", identifier, strings.Trim(inValue.GetText(), "'")))
		}

		setList := make(map[string]struct{})
		for _, expr := range inList {
			resultSet := r.RedisData[expr]
			log.Printf("Processing OR condition (IN transformation): %s, Retrieved Set: %v", expr, resultSet)
			for _, item := range resultSet {
				setList[item] = struct{}{}
			}
		}
		finalResult := make([]string, 0, len(setList))
		for item := range setList {
			finalResult = append(finalResult, item)
		}
		return finalResult
	}

	log.Printf("Unsupported or incomplete condition: %s", ctx.GetText())
	return nil
}

func (r *SearchQueryExecutor) VisitInList(ctx *sqparser.InListContext) any {
	var values []string
	for _, inValueCtx := range ctx.AllInValue() {
		value := r.Visit(inValueCtx).(string) // Expecting a string from VisitInValue
		values = append(values, value)
	}
	return values
}

func (r *SearchQueryExecutor) VisitInValue(ctx *sqparser.InValueContext) any {
	if ctx.QUOTED_LITERAL() != nil {
		return strings.Trim(ctx.QUOTED_LITERAL().GetText(), "'")
	}
	if ctx.LITERAL() != nil {
		return ctx.LITERAL().GetText()
	}
	return ""
}

func (r *SearchQueryExecutor) VisitSort_clause(ctx *sqparser.Sort_clauseContext) any {
	// Erstellt die orderMaps
	orderMaps := r.createOrderMaps(ctx.AllIDENTIFIER())

	// Array zur Speicherung von Comparator-Funktionen
	comparators := make([]func(id1, id2 string) int, 0, len(ctx.AllIDENTIFIER()))

	// Iteriere über alle Sortierbedingungen und erstelle die entsprechenden Vergleichsfunktionen
	for i, identifierCtx := range ctx.AllIDENTIFIER() {
		identifier := identifierCtx.GetText()
		direction := "ASC"
		if ctx.ASC(i) != nil {
			direction = "ASC"
		} else if ctx.DESC(i) != nil {
			direction = "DESC"
		}

		// Überprüfen, ob das Sortierfeld in den orderMaps verfügbar ist
		orderMap, ok := orderMaps[identifier]
		if !ok || len(orderMap) == 0 {
			// Wenn das Sortierfeld nicht verfügbar ist, logge eine Warnung
			log.Printf("Warning: Sort field %s not available in Redis data", identifier)
			continue
		}

		asc := direction == "ASC"

		// Bestimme den Typ des Wertes, um den passenden Comparator zu erstellen
		for _, val := range orderMap {
			switch val.(type) {
			case int:
				comparators = append(comparators, func(id1, id2 string) int {
					val1, ok1 := orderMap[id1]
					val2, ok2 := orderMap[id2]

					// Überprüfen, ob die Werte existieren
					if !ok1 || !ok2 {
						log.Printf("Warning: Missing value for ids %s or %s in field %s", id1, id2, identifier)
						return 0 // Ignoriere fehlende Werte, könnte auch eine andere Logik sein
					}

					intVal1, intVal2 := val1.(int), val2.(int)
					if intVal1 == intVal2 {
						return 0
					}
					if asc {
						if intVal1 < intVal2 {
							return -1
						}
						return 1
					}
					if intVal1 > intVal2 {
						return -1
					}
					return 1
				})
			case string:
				comparators = append(comparators, func(id1, id2 string) int {
					val1, ok1 := orderMap[id1]
					val2, ok2 := orderMap[id2]

					// Überprüfen, ob die Werte existieren
					if !ok1 || !ok2 {
						log.Printf("Warning: Missing value for ids %s or %s in field %s", id1, id2, identifier)
						return 0 // Ignoriere fehlende Werte, könnte auch eine andere Logik sein
					}

					strVal1, strVal2 := val1.(string), val2.(string)
					if strVal1 == strVal2 {
						return 0
					}
					if asc {
						if strVal1 < strVal2 {
							return -1
						}
						return 1
					}
					if strVal1 > strVal2 {
						return -1
					}
					return 1
				})
			}
			break
		}
	}

	// Wenn keine Comparatoren vorhanden sind, gib die ursprüngliche Liste zurück
	if len(comparators) == 0 {
		log.Println("Warning: No valid sort fields provided, returning original order")
		return r.ResultSet
	}

	// Sortiere das Ergebnis basierend auf den Comparatoren
	sort.SliceStable(r.ResultSet, func(i, j int) bool {
		id1, id2 := r.ResultSet[i], r.ResultSet[j]
		for _, comparator := range comparators {
			if result := comparator(id1, id2); result != 0 {
				return result < 0
			}
		}
		return false
	})

	// Gib das sortierte Ergebnis zurück
	log.Printf("Final result after sorting: %v", r.ResultSet)
	return r.ResultSet
}

func (r *SearchQueryExecutor) VisitOffset_clause(ctx *sqparser.Offset_clauseContext) any {
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

func (r *SearchQueryExecutor) VisitLimit_clause(ctx *sqparser.Limit_clauseContext) any {
	log.Println("Visiting Limit Clause")

	limit, _ := strconv.Atoi(ctx.NUMBER().GetText()) // Konvertiere die LIMIT-Zahl
	if limit < len(r.ResultSet) {
		r.ResultSet = r.ResultSet[:limit]
	}

	log.Printf("Result after applying LIMIT: %v", r.ResultSet)
	return r.ResultSet
}

func (r *SearchQueryExecutor) VisitRangeExpression(ctx *sqparser.RangeExpressionContext) any {
	log.Println("VisitRangeExpression called")

	// Extrahiere die Nummern aus dem Kontext
	startNumber := ctx.NUMBER(0).GetText() // Erste Zahl
	endNumber := ctx.NUMBER(1).GetText()   // Zweite Zahl

	// Debug-Ausgaben
	log.Printf("Start number: %s", startNumber)
	log.Printf("End number: %s", endNumber)

	// Erstelle den Bereichsausdruck
	rangeExpr := fmt.Sprintf("[%s %s]", startNumber, endNumber)

	// Debug-Ausgabe des finalen Bereichsausdrucks
	log.Printf("Range expression: %s", rangeExpr)

	return rangeExpr
}

func (r *SearchQueryExecutor) createOrderMaps(identifiers []antlr.TerminalNode) map[string]map[string]interface{} {
	orderMaps := make(map[string]map[string]interface{})

	// Iteriere über die Identifiers, die im Query verwendet werden
	for _, identifierCtx := range identifiers {
		field := identifierCtx.GetText()

		// Überprüfe, ob das Feld in den Redis-Daten vorhanden ist
		entries, ok := r.RedisData["sorting:"+field]
		if !ok || len(entries) == 0 {
			log.Printf("Warning: Field %s not available in Redis data", field)
			continue
		}

		orderMap := make(map[string]interface{})

		// Verarbeite die Einträge und konvertiere sie in die passende Struktur
		for _, entry := range entries {
			parts := strings.Split(entry, ":")
			if len(parts) != 2 {
				continue
			}

			id := parts[1]    // Der erste Teil ist die ID
			value := parts[0] // Der zweite Teil ist der Wert

			// Versuche, die Werte zu Zahlen zu konvertieren, wenn möglich
			if numericValue, err := strconv.Atoi(value); err == nil {
				orderMap[id] = numericValue // Wenn der Wert numerisch ist
			} else {
				orderMap[id] = value // Wenn der Wert ein String ist
			}
		}

		// Speichere das orderMap für das aktuelle Feld in orderMaps
		orderMaps[field] = orderMap
	}

	return orderMaps
}

// Hilfsfunktion, um alle möglichen Werte für einen bestimmten identifier zu holen
func (r *SearchQueryExecutor) getAllKeysForIdentifier(identifier string) []string {
	allValues := make([]string, 0)
	for key, values := range r.RedisData {
		if keyHasIdentifier(key, identifier) {
			allValues = append(allValues, values...)
		}
	}
	return allValues
}

// Überprüfe, ob der Schlüssel den gesuchten identifier enthält
func keyHasIdentifier(key, identifier string) bool {
	return len(key) > len(identifier) && key[:len(identifier)] == identifier
}

// subtractSets subtrahiert die Werte von der Menge der allKeys
func subtractSets(allKeys, excludeSet []string) []string {
	resultSet := make(map[string]struct{})

	// Füge alle Werte von allKeys in das Set ein
	for _, value := range allKeys {
		resultSet[value] = struct{}{}
	}

	// Entferne alle Werte, die im excludeSet enthalten sind
	for _, value := range excludeSet {
		delete(resultSet, value)
	}

	// Konvertiere das resultSet zurück in ein Slice
	finalResult := make([]string, 0, len(resultSet))
	for item := range resultSet {
		finalResult = append(finalResult, item)
	}

	return finalResult
}

// trimQuotes entfernt die umschließenden Apostrophe, wenn sie vorhanden sind
func trimQuotes(s string) string {
	if len(s) > 1 && s[0] == '\'' && s[len(s)-1] == '\'' {
		return s[1 : len(s)-1]
	}
	return s
}

// intersectSets berechnet die Schnittmenge mehrerer Mengen (Slices von Strings)
func intersectSets(sets [][]string) []string {
	if len(sets) == 0 {
		return []string{}
	}

	// Beginne mit der ersten Menge
	resultSet := make(map[string]struct{})
	for _, item := range sets[0] {
		resultSet[item] = struct{}{}
	}

	// Berechne die Schnittmenge mit den übrigen Mengen
	for _, set := range sets[1:] {
		tempSet := make(map[string]struct{})
		for _, item := range set {
			if _, found := resultSet[item]; found {
				tempSet[item] = struct{}{}
			}
		}
		resultSet = tempSet // Aktualisiere die Schnittmenge
	}

	// Konvertiere das resultSet zurück in ein Slice
	finalResult := make([]string, 0, len(resultSet))
	for item := range resultSet {
		finalResult = append(finalResult, item)
	}

	return finalResult
}

// unionSets kombiniert mehrere Mengen von Ergebnissen zu einer einzigen Menge (ähnlich wie union_sets im Python-Code)
func unionSets(sets [][]string) []string {
	unionSet := make(map[string]struct{})

	// Füge jedes Set zur Vereinigung hinzu
	for _, set := range sets {
		for _, item := range set {
			unionSet[item] = struct{}{}
		}
	}

	// Wandle die Menge in eine Liste um
	finalResult := make([]string, 0, len(unionSet))
	for item := range unionSet {
		finalResult = append(finalResult, item)
	}

	return finalResult
}

// convertToSet konvertiert ein Ergebnis in ein Set von Strings (hier als Hilfsfunktion verwendet)
func convertToSet(result any) []string {
	if result == nil {
		return []string{}
	}
	if res, ok := result.([]string); ok {
		return res
	}
	return []string{}
}
