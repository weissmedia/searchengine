// Code generated from /app/antlr/SearchQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package sqparser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SearchQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SearchQueryLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func searchquerylexerLexerInit() {
	staticData := &SearchQueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'!='", "'='", "'AND'", "'OR'", "'IN'", "'SORT'", "'BY'", "'LIMIT'",
		"'OFFSET'", "'ASC'", "'DESC'", "'['", "']'", "'('", "')'", "','", "'>'",
		"'>='", "'<'", "'<='", "'~'",
	}
	staticData.SymbolicNames = []string{
		"", "NOT_EQUALS", "EQUALS", "AND", "OR", "IN", "SORT", "BY", "LIMIT",
		"OFFSET", "ASC", "DESC", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN",
		"COMMA", "GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL", "FUZZY",
		"NUMBER", "IDENTIFIER", "WILDCARD", "QUOTED_LITERAL", "LITERAL", "WS",
	}
	staticData.RuleNames = []string{
		"NOT_EQUALS", "EQUALS", "AND", "OR", "IN", "SORT", "BY", "LIMIT", "OFFSET",
		"ASC", "DESC", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN", "COMMA",
		"GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL", "FUZZY", "NUMBER",
		"IDENTIFIER", "WILDCARD", "QUOTED_LITERAL", "LITERAL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 175, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 21, 4, 21, 124, 8, 21, 11, 21, 12, 21, 125, 1, 22, 1, 22, 5,
		22, 130, 8, 22, 10, 22, 12, 22, 133, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 3, 23, 151, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 157, 8,
		24, 10, 24, 12, 24, 160, 9, 24, 1, 24, 1, 24, 1, 25, 4, 25, 165, 8, 25,
		11, 25, 12, 25, 166, 1, 26, 4, 26, 170, 8, 26, 11, 26, 12, 26, 171, 1,
		26, 1, 26, 0, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 1, 0, 6, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 2, 0, 39, 39, 92, 92, 5, 0, 45, 45, 48, 57,
		65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 182, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 1, 55, 1, 0,
		0, 0, 3, 58, 1, 0, 0, 0, 5, 60, 1, 0, 0, 0, 7, 64, 1, 0, 0, 0, 9, 67, 1,
		0, 0, 0, 11, 70, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0, 15, 78, 1, 0, 0, 0, 17,
		84, 1, 0, 0, 0, 19, 91, 1, 0, 0, 0, 21, 95, 1, 0, 0, 0, 23, 100, 1, 0,
		0, 0, 25, 102, 1, 0, 0, 0, 27, 104, 1, 0, 0, 0, 29, 106, 1, 0, 0, 0, 31,
		108, 1, 0, 0, 0, 33, 110, 1, 0, 0, 0, 35, 112, 1, 0, 0, 0, 37, 115, 1,
		0, 0, 0, 39, 117, 1, 0, 0, 0, 41, 120, 1, 0, 0, 0, 43, 123, 1, 0, 0, 0,
		45, 127, 1, 0, 0, 0, 47, 150, 1, 0, 0, 0, 49, 152, 1, 0, 0, 0, 51, 164,
		1, 0, 0, 0, 53, 169, 1, 0, 0, 0, 55, 56, 5, 33, 0, 0, 56, 57, 5, 61, 0,
		0, 57, 2, 1, 0, 0, 0, 58, 59, 5, 61, 0, 0, 59, 4, 1, 0, 0, 0, 60, 61, 5,
		65, 0, 0, 61, 62, 5, 78, 0, 0, 62, 63, 5, 68, 0, 0, 63, 6, 1, 0, 0, 0,
		64, 65, 5, 79, 0, 0, 65, 66, 5, 82, 0, 0, 66, 8, 1, 0, 0, 0, 67, 68, 5,
		73, 0, 0, 68, 69, 5, 78, 0, 0, 69, 10, 1, 0, 0, 0, 70, 71, 5, 83, 0, 0,
		71, 72, 5, 79, 0, 0, 72, 73, 5, 82, 0, 0, 73, 74, 5, 84, 0, 0, 74, 12,
		1, 0, 0, 0, 75, 76, 5, 66, 0, 0, 76, 77, 5, 89, 0, 0, 77, 14, 1, 0, 0,
		0, 78, 79, 5, 76, 0, 0, 79, 80, 5, 73, 0, 0, 80, 81, 5, 77, 0, 0, 81, 82,
		5, 73, 0, 0, 82, 83, 5, 84, 0, 0, 83, 16, 1, 0, 0, 0, 84, 85, 5, 79, 0,
		0, 85, 86, 5, 70, 0, 0, 86, 87, 5, 70, 0, 0, 87, 88, 5, 83, 0, 0, 88, 89,
		5, 69, 0, 0, 89, 90, 5, 84, 0, 0, 90, 18, 1, 0, 0, 0, 91, 92, 5, 65, 0,
		0, 92, 93, 5, 83, 0, 0, 93, 94, 5, 67, 0, 0, 94, 20, 1, 0, 0, 0, 95, 96,
		5, 68, 0, 0, 96, 97, 5, 69, 0, 0, 97, 98, 5, 83, 0, 0, 98, 99, 5, 67, 0,
		0, 99, 22, 1, 0, 0, 0, 100, 101, 5, 91, 0, 0, 101, 24, 1, 0, 0, 0, 102,
		103, 5, 93, 0, 0, 103, 26, 1, 0, 0, 0, 104, 105, 5, 40, 0, 0, 105, 28,
		1, 0, 0, 0, 106, 107, 5, 41, 0, 0, 107, 30, 1, 0, 0, 0, 108, 109, 5, 44,
		0, 0, 109, 32, 1, 0, 0, 0, 110, 111, 5, 62, 0, 0, 111, 34, 1, 0, 0, 0,
		112, 113, 5, 62, 0, 0, 113, 114, 5, 61, 0, 0, 114, 36, 1, 0, 0, 0, 115,
		116, 5, 60, 0, 0, 116, 38, 1, 0, 0, 0, 117, 118, 5, 60, 0, 0, 118, 119,
		5, 61, 0, 0, 119, 40, 1, 0, 0, 0, 120, 121, 5, 126, 0, 0, 121, 42, 1, 0,
		0, 0, 122, 124, 7, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 44, 1, 0, 0, 0, 127, 131,
		7, 1, 0, 0, 128, 130, 7, 2, 0, 0, 129, 128, 1, 0, 0, 0, 130, 133, 1, 0,
		0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 46, 1, 0, 0, 0,
		133, 131, 1, 0, 0, 0, 134, 135, 5, 39, 0, 0, 135, 136, 5, 42, 0, 0, 136,
		137, 3, 51, 25, 0, 137, 138, 5, 39, 0, 0, 138, 151, 1, 0, 0, 0, 139, 140,
		5, 39, 0, 0, 140, 141, 3, 51, 25, 0, 141, 142, 5, 42, 0, 0, 142, 143, 5,
		39, 0, 0, 143, 151, 1, 0, 0, 0, 144, 145, 5, 39, 0, 0, 145, 146, 5, 42,
		0, 0, 146, 147, 3, 51, 25, 0, 147, 148, 5, 42, 0, 0, 148, 149, 5, 39, 0,
		0, 149, 151, 1, 0, 0, 0, 150, 134, 1, 0, 0, 0, 150, 139, 1, 0, 0, 0, 150,
		144, 1, 0, 0, 0, 151, 48, 1, 0, 0, 0, 152, 158, 5, 39, 0, 0, 153, 157,
		8, 3, 0, 0, 154, 155, 5, 92, 0, 0, 155, 157, 9, 0, 0, 0, 156, 153, 1, 0,
		0, 0, 156, 154, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0,
		158, 159, 1, 0, 0, 0, 159, 161, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 161,
		162, 5, 39, 0, 0, 162, 50, 1, 0, 0, 0, 163, 165, 7, 4, 0, 0, 164, 163,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0,
		0, 0, 167, 52, 1, 0, 0, 0, 168, 170, 7, 5, 0, 0, 169, 168, 1, 0, 0, 0,
		170, 171, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172,
		173, 1, 0, 0, 0, 173, 174, 6, 26, 0, 0, 174, 54, 1, 0, 0, 0, 8, 0, 125,
		131, 150, 156, 158, 166, 171, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SearchQueryLexerInit initializes any static state used to implement SearchQueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSearchQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SearchQueryLexerInit() {
	staticData := &SearchQueryLexerLexerStaticData
	staticData.once.Do(searchquerylexerLexerInit)
}

// NewSearchQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSearchQueryLexer(input antlr.CharStream) *SearchQueryLexer {
	SearchQueryLexerInit()
	l := new(SearchQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SearchQueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SearchQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SearchQueryLexer tokens.
const (
	SearchQueryLexerNOT_EQUALS     = 1
	SearchQueryLexerEQUALS         = 2
	SearchQueryLexerAND            = 3
	SearchQueryLexerOR             = 4
	SearchQueryLexerIN             = 5
	SearchQueryLexerSORT           = 6
	SearchQueryLexerBY             = 7
	SearchQueryLexerLIMIT          = 8
	SearchQueryLexerOFFSET         = 9
	SearchQueryLexerASC            = 10
	SearchQueryLexerDESC           = 11
	SearchQueryLexerLBRACKET       = 12
	SearchQueryLexerRBRACKET       = 13
	SearchQueryLexerLPAREN         = 14
	SearchQueryLexerRPAREN         = 15
	SearchQueryLexerCOMMA          = 16
	SearchQueryLexerGREATER        = 17
	SearchQueryLexerGREATER_EQUAL  = 18
	SearchQueryLexerLESS           = 19
	SearchQueryLexerLESS_EQUAL     = 20
	SearchQueryLexerFUZZY          = 21
	SearchQueryLexerNUMBER         = 22
	SearchQueryLexerIDENTIFIER     = 23
	SearchQueryLexerWILDCARD       = 24
	SearchQueryLexerQUOTED_LITERAL = 25
	SearchQueryLexerLITERAL        = 26
	SearchQueryLexerWS             = 27
)
