package core

// ComparisonOperator repräsentiert die unterstützten Vergleichsoperatoren.
type ComparisonOperator string

// Definiere die Vergleichsoperatoren als Konstanten.
const (
	OpGreater      ComparisonOperator = ">"
	OpGreaterEqual ComparisonOperator = ">="
	OpLess         ComparisonOperator = "<"
	OpLessEqual    ComparisonOperator = "<="
)

// Comparator definiert den Typ für eine Vergleichsfunktion.
type Comparator func(id1, id2 string) int

// Ordinal vergleicht Werte, die den Vergleichsoperator unterstützen.
type Ordinal interface {
	~int | ~float64 | ~string // Erlaubt nur Typen wie int, float64, und string
}

// Compare vergleicht zwei Werte vom Typ T (int, string, float64).
func Compare[T Ordinal](val1, val2 T, asc bool) int {
	if val1 == val2 {
		return 0
	}
	if asc {
		if val1 < val2 {
			return -1
		}
		return 1
	} else {
		if val1 > val2 {
			return -1
		}
		return 1
	}
}

// MapComparisonOperator gibt eine Vergleichsfunktion für den entsprechenden Operator zurück.
func MapComparisonOperator[T Ordinal](op ComparisonOperator) func(val1, val2 T) int {
	switch op {
	case OpGreater:
		return func(val1, val2 T) int { return Compare(val1, val2, false) }
	case OpGreaterEqual:
		return func(val1, val2 T) int { return Compare(val1, val2, true) }
	case OpLess:
		return func(val1, val2 T) int { return Compare(val1, val2, true) }
	case OpLessEqual:
		return func(val1, val2 T) int { return Compare(val1, val2, true) }
	default:
		return nil
	}
}
