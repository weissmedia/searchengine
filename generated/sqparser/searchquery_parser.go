// Code generated from /app/antlr/SearchQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package sqparser // SearchQuery
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SearchQueryParser struct {
	*antlr.BaseParser
}

var SearchQueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func searchqueryParserInit() {
	staticData := &SearchQueryParserStaticData
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
		"query", "expression", "orExpression", "andExpression", "comparisonExpression",
		"primary", "condition", "comparisonOperator", "value", "rangeExpression",
		"inList", "inValue", "sort_clause", "limit_clause", "offset_clause",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 140, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 3, 0,
		32, 8, 0, 1, 0, 1, 0, 1, 0, 3, 0, 37, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3,
		0, 43, 8, 0, 3, 0, 45, 8, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 52, 8,
		2, 10, 2, 12, 2, 55, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 60, 8, 3, 10, 3, 12,
		3, 63, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 72, 8, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 90, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 3, 8, 99, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 104, 8, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 113, 8, 10, 10, 10, 12, 10, 116,
		9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 124, 8, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 129, 8, 12, 5, 12, 131, 8, 12, 10, 12, 12, 12,
		134, 9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 0, 3, 1, 0, 17, 20, 2, 0, 22, 22,
		25, 26, 1, 0, 10, 11, 144, 0, 31, 1, 0, 0, 0, 2, 46, 1, 0, 0, 0, 4, 48,
		1, 0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10, 71, 1, 0, 0, 0, 12,
		89, 1, 0, 0, 0, 14, 91, 1, 0, 0, 0, 16, 98, 1, 0, 0, 0, 18, 100, 1, 0,
		0, 0, 20, 108, 1, 0, 0, 0, 22, 119, 1, 0, 0, 0, 24, 121, 1, 0, 0, 0, 26,
		135, 1, 0, 0, 0, 28, 137, 1, 0, 0, 0, 30, 32, 3, 2, 1, 0, 31, 30, 1, 0,
		0, 0, 31, 32, 1, 0, 0, 0, 32, 36, 1, 0, 0, 0, 33, 34, 5, 6, 0, 0, 34, 35,
		5, 7, 0, 0, 35, 37, 3, 24, 12, 0, 36, 33, 1, 0, 0, 0, 36, 37, 1, 0, 0,
		0, 37, 44, 1, 0, 0, 0, 38, 39, 5, 8, 0, 0, 39, 42, 3, 26, 13, 0, 40, 41,
		5, 9, 0, 0, 41, 43, 3, 28, 14, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0,
		0, 43, 45, 1, 0, 0, 0, 44, 38, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 1, 1,
		0, 0, 0, 46, 47, 3, 4, 2, 0, 47, 3, 1, 0, 0, 0, 48, 53, 3, 6, 3, 0, 49,
		50, 5, 4, 0, 0, 50, 52, 3, 6, 3, 0, 51, 49, 1, 0, 0, 0, 52, 55, 1, 0, 0,
		0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 5, 1, 0, 0, 0, 55, 53, 1,
		0, 0, 0, 56, 61, 3, 8, 4, 0, 57, 58, 5, 3, 0, 0, 58, 60, 3, 8, 4, 0, 59,
		57, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0,
		0, 62, 7, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 65, 3, 10, 5, 0, 65, 9, 1,
		0, 0, 0, 66, 67, 5, 14, 0, 0, 67, 68, 3, 2, 1, 0, 68, 69, 5, 15, 0, 0,
		69, 72, 1, 0, 0, 0, 70, 72, 3, 12, 6, 0, 71, 66, 1, 0, 0, 0, 71, 70, 1,
		0, 0, 0, 72, 11, 1, 0, 0, 0, 73, 74, 5, 23, 0, 0, 74, 75, 5, 1, 0, 0, 75,
		90, 3, 16, 8, 0, 76, 77, 5, 23, 0, 0, 77, 78, 5, 21, 0, 0, 78, 90, 5, 25,
		0, 0, 79, 80, 5, 23, 0, 0, 80, 81, 3, 14, 7, 0, 81, 82, 3, 16, 8, 0, 82,
		90, 1, 0, 0, 0, 83, 84, 5, 23, 0, 0, 84, 85, 5, 2, 0, 0, 85, 90, 3, 16,
		8, 0, 86, 87, 5, 23, 0, 0, 87, 88, 5, 5, 0, 0, 88, 90, 3, 20, 10, 0, 89,
		73, 1, 0, 0, 0, 89, 76, 1, 0, 0, 0, 89, 79, 1, 0, 0, 0, 89, 83, 1, 0, 0,
		0, 89, 86, 1, 0, 0, 0, 90, 13, 1, 0, 0, 0, 91, 92, 7, 0, 0, 0, 92, 15,
		1, 0, 0, 0, 93, 99, 5, 25, 0, 0, 94, 99, 5, 26, 0, 0, 95, 99, 3, 18, 9,
		0, 96, 99, 5, 24, 0, 0, 97, 99, 5, 22, 0, 0, 98, 93, 1, 0, 0, 0, 98, 94,
		1, 0, 0, 0, 98, 95, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 97, 1, 0, 0, 0,
		99, 17, 1, 0, 0, 0, 100, 101, 5, 12, 0, 0, 101, 103, 5, 22, 0, 0, 102,
		104, 5, 27, 0, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105,
		1, 0, 0, 0, 105, 106, 5, 22, 0, 0, 106, 107, 5, 13, 0, 0, 107, 19, 1, 0,
		0, 0, 108, 109, 5, 14, 0, 0, 109, 114, 3, 22, 11, 0, 110, 111, 5, 16, 0,
		0, 111, 113, 3, 22, 11, 0, 112, 110, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0,
		114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116,
		114, 1, 0, 0, 0, 117, 118, 5, 15, 0, 0, 118, 21, 1, 0, 0, 0, 119, 120,
		7, 1, 0, 0, 120, 23, 1, 0, 0, 0, 121, 123, 5, 23, 0, 0, 122, 124, 7, 2,
		0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 132, 1, 0, 0, 0,
		125, 126, 5, 16, 0, 0, 126, 128, 5, 23, 0, 0, 127, 129, 7, 2, 0, 0, 128,
		127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 125,
		1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 25, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 22, 0, 0,
		136, 27, 1, 0, 0, 0, 137, 138, 5, 22, 0, 0, 138, 29, 1, 0, 0, 0, 14, 31,
		36, 42, 44, 53, 61, 71, 89, 98, 103, 114, 123, 128, 132,
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

// SearchQueryParserInit initializes any static state used to implement SearchQueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSearchQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SearchQueryParserInit() {
	staticData := &SearchQueryParserStaticData
	staticData.once.Do(searchqueryParserInit)
}

// NewSearchQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSearchQueryParser(input antlr.TokenStream) *SearchQueryParser {
	SearchQueryParserInit()
	this := new(SearchQueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SearchQueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SearchQuery.g4"

	return this
}

// SearchQueryParser tokens.
const (
	SearchQueryParserEOF            = antlr.TokenEOF
	SearchQueryParserNOT_EQUALS     = 1
	SearchQueryParserEQUALS         = 2
	SearchQueryParserAND            = 3
	SearchQueryParserOR             = 4
	SearchQueryParserIN             = 5
	SearchQueryParserSORT           = 6
	SearchQueryParserBY             = 7
	SearchQueryParserLIMIT          = 8
	SearchQueryParserOFFSET         = 9
	SearchQueryParserASC            = 10
	SearchQueryParserDESC           = 11
	SearchQueryParserLBRACKET       = 12
	SearchQueryParserRBRACKET       = 13
	SearchQueryParserLPAREN         = 14
	SearchQueryParserRPAREN         = 15
	SearchQueryParserCOMMA          = 16
	SearchQueryParserGREATER        = 17
	SearchQueryParserGREATER_EQUAL  = 18
	SearchQueryParserLESS           = 19
	SearchQueryParserLESS_EQUAL     = 20
	SearchQueryParserFUZZY          = 21
	SearchQueryParserNUMBER         = 22
	SearchQueryParserIDENTIFIER     = 23
	SearchQueryParserWILDCARD       = 24
	SearchQueryParserQUOTED_LITERAL = 25
	SearchQueryParserLITERAL        = 26
	SearchQueryParserWS             = 27
)

// SearchQueryParser rules.
const (
	SearchQueryParserRULE_query                = 0
	SearchQueryParserRULE_expression           = 1
	SearchQueryParserRULE_orExpression         = 2
	SearchQueryParserRULE_andExpression        = 3
	SearchQueryParserRULE_comparisonExpression = 4
	SearchQueryParserRULE_primary              = 5
	SearchQueryParserRULE_condition            = 6
	SearchQueryParserRULE_comparisonOperator   = 7
	SearchQueryParserRULE_value                = 8
	SearchQueryParserRULE_rangeExpression      = 9
	SearchQueryParserRULE_inList               = 10
	SearchQueryParserRULE_inValue              = 11
	SearchQueryParserRULE_sort_clause          = 12
	SearchQueryParserRULE_limit_clause         = 13
	SearchQueryParserRULE_offset_clause        = 14
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	SORT() antlr.TerminalNode
	BY() antlr.TerminalNode
	Sort_clause() ISort_clauseContext
	LIMIT() antlr.TerminalNode
	Limit_clause() ILimit_clauseContext
	OFFSET() antlr.TerminalNode
	Offset_clause() IOffset_clauseContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *QueryContext) SORT() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserSORT, 0)
}

func (s *QueryContext) BY() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserBY, 0)
}

func (s *QueryContext) Sort_clause() ISort_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISort_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISort_clauseContext)
}

func (s *QueryContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserLIMIT, 0)
}

func (s *QueryContext) Limit_clause() ILimit_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimit_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimit_clauseContext)
}

func (s *QueryContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserOFFSET, 0)
}

func (s *QueryContext) Offset_clause() IOffset_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffset_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffset_clauseContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SearchQueryParserRULE_query)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SearchQueryParserLPAREN || _la == SearchQueryParserIDENTIFIER {
		{
			p.SetState(30)
			p.Expression()
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SearchQueryParserSORT {
		{
			p.SetState(33)
			p.Match(SearchQueryParserSORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(SearchQueryParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Sort_clause()
		}

	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SearchQueryParserLIMIT {
		{
			p.SetState(38)
			p.Match(SearchQueryParserLIMIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Limit_clause()
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SearchQueryParserOFFSET {
			{
				p.SetState(40)
				p.Match(SearchQueryParserOFFSET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(41)
				p.Offset_clause()
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OrExpression() IOrExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) OrExpression() IOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SearchQueryParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.OrExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAndExpression() []IAndExpressionContext
	AndExpression(i int) IAndExpressionContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_orExpression
	return p
}

func InitEmptyOrExpressionContext(p *OrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_orExpression
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AllAndExpression() []IAndExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAndExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndExpressionContext); ok {
			tst[i] = t.(IAndExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExpressionContext) AndExpression(i int) IAndExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(SearchQueryParserOR)
}

func (s *OrExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(SearchQueryParserOR, i)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) OrExpression() (localctx IOrExpressionContext) {
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SearchQueryParserRULE_orExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.AndExpression()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SearchQueryParserOR {
		{
			p.SetState(49)
			p.Match(SearchQueryParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.AndExpression()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparisonExpression() []IComparisonExpressionContext
	ComparisonExpression(i int) IComparisonExpressionContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_andExpression
	return p
}

func InitEmptyAndExpressionContext(p *AndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_andExpression
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) AllComparisonExpression() []IComparisonExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonExpressionContext); ok {
			len++
		}
	}

	tst := make([]IComparisonExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonExpressionContext); ok {
			tst[i] = t.(IComparisonExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExpressionContext) ComparisonExpression(i int) IComparisonExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionContext)
}

func (s *AndExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(SearchQueryParserAND)
}

func (s *AndExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(SearchQueryParserAND, i)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) AndExpression() (localctx IAndExpressionContext) {
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SearchQueryParserRULE_andExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.ComparisonExpression()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SearchQueryParserAND {
		{
			p.SetState(57)
			p.Match(SearchQueryParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.ComparisonExpression()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonExpressionContext is an interface to support dynamic dispatch.
type IComparisonExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext

	// IsComparisonExpressionContext differentiates from other interfaces.
	IsComparisonExpressionContext()
}

type ComparisonExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonExpressionContext() *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_comparisonExpression
	return p
}

func InitEmptyComparisonExpressionContext(p *ComparisonExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_comparisonExpression
}

func (*ComparisonExpressionContext) IsComparisonExpressionContext() {}

func NewComparisonExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_comparisonExpression

	return p
}

func (s *ComparisonExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExpressionContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitComparisonExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) ComparisonExpression() (localctx IComparisonExpressionContext) {
	localctx = NewComparisonExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SearchQueryParserRULE_comparisonExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Primary()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Condition() IConditionContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserLPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserRPAREN, 0)
}

func (s *PrimaryContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SearchQueryParserRULE_primary)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SearchQueryParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(SearchQueryParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Expression()
		}
		{
			p.SetState(68)
			p.Match(SearchQueryParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SearchQueryParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Condition()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	NOT_EQUALS() antlr.TerminalNode
	Value() IValueContext
	FUZZY() antlr.TerminalNode
	QUOTED_LITERAL() antlr.TerminalNode
	ComparisonOperator() IComparisonOperatorContext
	EQUALS() antlr.TerminalNode
	IN() antlr.TerminalNode
	InList() IInListContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserIDENTIFIER, 0)
}

func (s *ConditionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserNOT_EQUALS, 0)
}

func (s *ConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ConditionContext) FUZZY() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserFUZZY, 0)
}

func (s *ConditionContext) QUOTED_LITERAL() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserQUOTED_LITERAL, 0)
}

func (s *ConditionContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ConditionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserEQUALS, 0)
}

func (s *ConditionContext) IN() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserIN, 0)
}

func (s *ConditionContext) InList() IInListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInListContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SearchQueryParserRULE_condition)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Match(SearchQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(SearchQueryParserNOT_EQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(SearchQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(SearchQueryParserFUZZY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(SearchQueryParserQUOTED_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			p.Match(SearchQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.ComparisonOperator()
		}
		{
			p.SetState(81)
			p.Value()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Match(SearchQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(SearchQueryParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Value()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(86)
			p.Match(SearchQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(SearchQueryParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.InList()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GREATER() antlr.TerminalNode
	GREATER_EQUAL() antlr.TerminalNode
	LESS() antlr.TerminalNode
	LESS_EQUAL() antlr.TerminalNode

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_comparisonOperator
	return p
}

func InitEmptyComparisonOperatorContext(p *ComparisonOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_comparisonOperator
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GREATER() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserGREATER, 0)
}

func (s *ComparisonOperatorContext) GREATER_EQUAL() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserGREATER_EQUAL, 0)
}

func (s *ComparisonOperatorContext) LESS() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserLESS, 0)
}

func (s *ComparisonOperatorContext) LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserLESS_EQUAL, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SearchQueryParserRULE_comparisonOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1966080) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED_LITERAL() antlr.TerminalNode
	LITERAL() antlr.TerminalNode
	RangeExpression() IRangeExpressionContext
	WILDCARD() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) QUOTED_LITERAL() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserQUOTED_LITERAL, 0)
}

func (s *ValueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserLITERAL, 0)
}

func (s *ValueContext) RangeExpression() IRangeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeExpressionContext)
}

func (s *ValueContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserWILDCARD, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserNUMBER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SearchQueryParserRULE_value)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SearchQueryParserQUOTED_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(SearchQueryParserQUOTED_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SearchQueryParserLITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(SearchQueryParserLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SearchQueryParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.RangeExpression()
		}

	case SearchQueryParserWILDCARD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)
			p.Match(SearchQueryParserWILDCARD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SearchQueryParserNUMBER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)
			p.Match(SearchQueryParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeExpressionContext is an interface to support dynamic dispatch.
type IRangeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	WS() antlr.TerminalNode

	// IsRangeExpressionContext differentiates from other interfaces.
	IsRangeExpressionContext()
}

type RangeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeExpressionContext() *RangeExpressionContext {
	var p = new(RangeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_rangeExpression
	return p
}

func InitEmptyRangeExpressionContext(p *RangeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_rangeExpression
}

func (*RangeExpressionContext) IsRangeExpressionContext() {}

func NewRangeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeExpressionContext {
	var p = new(RangeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_rangeExpression

	return p
}

func (s *RangeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeExpressionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserLBRACKET, 0)
}

func (s *RangeExpressionContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(SearchQueryParserNUMBER)
}

func (s *RangeExpressionContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(SearchQueryParserNUMBER, i)
}

func (s *RangeExpressionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserRBRACKET, 0)
}

func (s *RangeExpressionContext) WS() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserWS, 0)
}

func (s *RangeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitRangeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) RangeExpression() (localctx IRangeExpressionContext) {
	localctx = NewRangeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SearchQueryParserRULE_rangeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(SearchQueryParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(SearchQueryParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SearchQueryParserWS {
		{
			p.SetState(102)
			p.Match(SearchQueryParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(105)
		p.Match(SearchQueryParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(SearchQueryParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInListContext is an interface to support dynamic dispatch.
type IInListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllInValue() []IInValueContext
	InValue(i int) IInValueContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsInListContext differentiates from other interfaces.
	IsInListContext()
}

type InListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInListContext() *InListContext {
	var p = new(InListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_inList
	return p
}

func InitEmptyInListContext(p *InListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_inList
}

func (*InListContext) IsInListContext() {}

func NewInListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InListContext {
	var p = new(InListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_inList

	return p
}

func (s *InListContext) GetParser() antlr.Parser { return s.parser }

func (s *InListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserLPAREN, 0)
}

func (s *InListContext) AllInValue() []IInValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInValueContext); ok {
			len++
		}
	}

	tst := make([]IInValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInValueContext); ok {
			tst[i] = t.(IInValueContext)
			i++
		}
	}

	return tst
}

func (s *InListContext) InValue(i int) IInValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInValueContext)
}

func (s *InListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserRPAREN, 0)
}

func (s *InListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SearchQueryParserCOMMA)
}

func (s *InListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SearchQueryParserCOMMA, i)
}

func (s *InListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitInList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) InList() (localctx IInListContext) {
	localctx = NewInListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SearchQueryParserRULE_inList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(SearchQueryParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.InValue()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SearchQueryParserCOMMA {
		{
			p.SetState(110)
			p.Match(SearchQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.InValue()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(SearchQueryParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInValueContext is an interface to support dynamic dispatch.
type IInValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED_LITERAL() antlr.TerminalNode
	LITERAL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsInValueContext differentiates from other interfaces.
	IsInValueContext()
}

type InValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInValueContext() *InValueContext {
	var p = new(InValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_inValue
	return p
}

func InitEmptyInValueContext(p *InValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_inValue
}

func (*InValueContext) IsInValueContext() {}

func NewInValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InValueContext {
	var p = new(InValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_inValue

	return p
}

func (s *InValueContext) GetParser() antlr.Parser { return s.parser }

func (s *InValueContext) QUOTED_LITERAL() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserQUOTED_LITERAL, 0)
}

func (s *InValueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserLITERAL, 0)
}

func (s *InValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserNUMBER, 0)
}

func (s *InValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitInValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) InValue() (localctx IInValueContext) {
	localctx = NewInValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SearchQueryParserRULE_inValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&104857600) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISort_clauseContext is an interface to support dynamic dispatch.
type ISort_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllASC() []antlr.TerminalNode
	ASC(i int) antlr.TerminalNode
	AllDESC() []antlr.TerminalNode
	DESC(i int) antlr.TerminalNode

	// IsSort_clauseContext differentiates from other interfaces.
	IsSort_clauseContext()
}

type Sort_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySort_clauseContext() *Sort_clauseContext {
	var p = new(Sort_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_sort_clause
	return p
}

func InitEmptySort_clauseContext(p *Sort_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_sort_clause
}

func (*Sort_clauseContext) IsSort_clauseContext() {}

func NewSort_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sort_clauseContext {
	var p = new(Sort_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_sort_clause

	return p
}

func (s *Sort_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Sort_clauseContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(SearchQueryParserIDENTIFIER)
}

func (s *Sort_clauseContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(SearchQueryParserIDENTIFIER, i)
}

func (s *Sort_clauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SearchQueryParserCOMMA)
}

func (s *Sort_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SearchQueryParserCOMMA, i)
}

func (s *Sort_clauseContext) AllASC() []antlr.TerminalNode {
	return s.GetTokens(SearchQueryParserASC)
}

func (s *Sort_clauseContext) ASC(i int) antlr.TerminalNode {
	return s.GetToken(SearchQueryParserASC, i)
}

func (s *Sort_clauseContext) AllDESC() []antlr.TerminalNode {
	return s.GetTokens(SearchQueryParserDESC)
}

func (s *Sort_clauseContext) DESC(i int) antlr.TerminalNode {
	return s.GetToken(SearchQueryParserDESC, i)
}

func (s *Sort_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sort_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sort_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitSort_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) Sort_clause() (localctx ISort_clauseContext) {
	localctx = NewSort_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SearchQueryParserRULE_sort_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(SearchQueryParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SearchQueryParserASC || _la == SearchQueryParserDESC {
		{
			p.SetState(122)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SearchQueryParserASC || _la == SearchQueryParserDESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SearchQueryParserCOMMA {
		{
			p.SetState(125)
			p.Match(SearchQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(SearchQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SearchQueryParserASC || _la == SearchQueryParserDESC {
			{
				p.SetState(127)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SearchQueryParserASC || _la == SearchQueryParserDESC) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimit_clauseContext is an interface to support dynamic dispatch.
type ILimit_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsLimit_clauseContext differentiates from other interfaces.
	IsLimit_clauseContext()
}

type Limit_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimit_clauseContext() *Limit_clauseContext {
	var p = new(Limit_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_limit_clause
	return p
}

func InitEmptyLimit_clauseContext(p *Limit_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_limit_clause
}

func (*Limit_clauseContext) IsLimit_clauseContext() {}

func NewLimit_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Limit_clauseContext {
	var p = new(Limit_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_limit_clause

	return p
}

func (s *Limit_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Limit_clauseContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserNUMBER, 0)
}

func (s *Limit_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Limit_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Limit_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitLimit_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) Limit_clause() (localctx ILimit_clauseContext) {
	localctx = NewLimit_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SearchQueryParserRULE_limit_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(SearchQueryParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOffset_clauseContext is an interface to support dynamic dispatch.
type IOffset_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsOffset_clauseContext differentiates from other interfaces.
	IsOffset_clauseContext()
}

type Offset_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffset_clauseContext() *Offset_clauseContext {
	var p = new(Offset_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_offset_clause
	return p
}

func InitEmptyOffset_clauseContext(p *Offset_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SearchQueryParserRULE_offset_clause
}

func (*Offset_clauseContext) IsOffset_clauseContext() {}

func NewOffset_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Offset_clauseContext {
	var p = new(Offset_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchQueryParserRULE_offset_clause

	return p
}

func (s *Offset_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Offset_clauseContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SearchQueryParserNUMBER, 0)
}

func (s *Offset_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Offset_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Offset_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchQueryVisitor:
		return t.VisitOffset_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchQueryParser) Offset_clause() (localctx IOffset_clauseContext) {
	localctx = NewOffset_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SearchQueryParserRULE_offset_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(SearchQueryParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
