package sqparser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type syntaxError struct {
	line, col int
	msg       string
	query     string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("Syntax error at line %d:%d - %s in query: %s", e.line, e.col, e.msg, e.query)
}

type queryErrorListener struct {
	*antlr.DefaultErrorListener
	query string
}

func newLynxErrorListener(query string) *queryErrorListener {
	return &queryErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		query:                query,
	}
}

func (l *queryErrorListener) SyntaxError(
	_ antlr.Recognizer,
	_ any,
	line, col int,
	msg string,
	ex antlr.RecognitionException,
) {
	panic(&syntaxError{line, col, msg, l.query})
}

func Parse(input string) (tree IQueryContext, err error) {
	defer func() {
		if r := recover(); r != nil {
			if rErr, ok := r.(*syntaxError); ok {
				tree, err = nil, rErr
			} else {
				panic(r)
			}
		}
	}()

	// Erstelle den InputStream und den Lexer mit dem Query
	stream := antlr.NewInputStream(input)
	lexer := NewSearchQueryLexer(stream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(newLynxErrorListener(input))

	// Token-Stream und Parser vorbereiten
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewSearchQueryParser(tokens)
	parser.BuildParseTrees = true
	parser.RemoveErrorListeners()
	parser.AddErrorListener(newLynxErrorListener(input))

	// Parsen des Queries
	return parser.Query(), nil
}
