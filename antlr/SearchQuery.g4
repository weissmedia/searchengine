grammar SearchQuery;


// ==================== Parser Rules ====================
// Parser rules describe the structure of the language
// These define how the syntax of the query language is structured

// Start rule - the main entry point for a query
query
    : expression?                          // Optional expression for filtering
      (SORT BY sort_clause)?               // Optional sorting clause
      (LIMIT limit_clause (OFFSET offset_clause)?)? // Optional limit and offset for pagination
    ;

// Logical expressions: OR, AND, parentheses, and conditions
expression
    : orExpression                         // Start with OR expressions
    ;

orExpression
    : andExpression (OR andExpression)*    // An OR expression can be a series of AND expressions joined by OR
    ;

andExpression
    : comparisonExpression (AND comparisonExpression)*  // AND expressions are comparisons joined by AND
    ;

comparisonExpression
    : primary                             // A comparison expression is reduced to a primary condition
    ;

primary
    : LPAREN expression RPAREN            // Allows grouping with parentheses
    | condition                           // A primary expression can also be a condition
    ;

// Condition expressions - define the different types of comparisons in the query
condition
    : IDENTIFIER NOT_EQUALS value         // Condition: field != value
    | IDENTIFIER FUZZY QUOTED_LITERAL     // Fuzzy search: field ~ "value"
    | IDENTIFIER comparisonOperator value // Condition: field with comparison operator
    | IDENTIFIER EQUALS value             // Condition: field = value
    | IDENTIFIER IN inList                // Condition: field IN (list of values)
    ;

// Comparison operators
comparisonOperator
    : GREATER                             // Greater than: >
    | GREATER_EQUAL                       // Greater or equal: >=
    | LESS                                // Less than: <
    | LESS_EQUAL                          // Less or equal: <=
    ;

// Values: literals, ranges, wildcards, fuzzy matching
value
    : QUOTED_LITERAL                      // String values enclosed in quotes
    | LITERAL                             // Simple literals
    | rangeExpression                     // A range of values
    | WILDCARD                            // Wildcard matching (* prefix/suffix/infix)
    | NUMBER                              // Numeric values
    ;

// Range expressions: [n m]
rangeExpression
    : LBRACKET NUMBER WS? NUMBER RBRACKET // Range expression: [n m] defines a range between two numbers
    ;

// IN lists, only allowing literals and simple values
inList
    : LPAREN inValue (COMMA inValue)* RPAREN // A list of values enclosed in parentheses, separated by commas
    ;

inValue
    : QUOTED_LITERAL                      // A quoted literal value
    | LITERAL                             // A simple literal value
    | NUMBER                              // A numeric value
    ;

// Sorting clause: defines how the results should be sorted
sort_clause
    : IDENTIFIER (ASC | DESC)?            // Sort by a field, optionally specify ascending or descending
      (COMMA IDENTIFIER (ASC | DESC)?)*
    ;

// Limit and offset for pagination
limit_clause
    : NUMBER                              // Number of results to return (limit)
    ;

offset_clause
    : NUMBER                              // Number of results to skip (offset)
    ;

// ==================== Lexer Rules ====================
// Lexer rules describe how tokens are recognized
// These rules define how different parts of the query are identified

NOT_EQUALS        : '!=';                 // Not equals operator
EQUALS            : '=';                  // Equals operator
AND               : 'AND';                // AND logical operator
OR                : 'OR';                 // OR logical operator
IN                : 'IN';                 // IN operator for lists
SORT              : 'SORT';               // SORT keyword for ordering
BY                : 'BY';                 // BY keyword for sorting
LIMIT             : 'LIMIT';              // LIMIT keyword for pagination
OFFSET            : 'OFFSET';             // OFFSET keyword for pagination
ASC               : 'ASC';                // Ascending sort order
DESC              : 'DESC';               // Descending sort order
LBRACKET          : '[';                  // Left bracket for ranges
RBRACKET          : ']';                  // Right bracket for ranges
LPAREN            : '(';                  // Left parenthesis
RPAREN            : ')';                  // Right parenthesis
COMMA             : ',';                  // Comma for separating values in lists

// Comparison operators
GREATER           : '>';                  // Greater than operator
GREATER_EQUAL     : '>=';                 // Greater or equal operator
LESS              : '<';                  // Less than operator
LESS_EQUAL        : '<=';                 // Less or equal operator

// Fuzzy matching
FUZZY             : '~';                  // Fuzzy matching operator

// Numbers
NUMBER            : [0-9]+;               // A sequence of digits representing a number

// Identifiers for field names
IDENTIFIER        : [a-zA-Z_][a-zA-Z0-9_]*; // Identifiers are alphanumeric sequences starting with a letter or underscore

// Wildcard matching, only allowed with EQUALS
WILDCARD
    : '\'' '*' LITERAL '\''               // Suffix wildcard: *value
    | '\'' LITERAL '*' '\''               // Prefix wildcard: value*
    | '\'' '*' LITERAL '*' '\''           // Infix wildcard: *value*
    ;

// Quoted literals: strings enclosed in single quotes
QUOTED_LITERAL    : '\'' (~['\\] | '\\' .)* '\'';

// Literals: alphanumeric values, including hyphens and underscores
LITERAL           : [a-zA-Z0-9_-]+;

// Skip whitespace characters
WS                : [ \t\r\n]+ -> skip;   // Ignore whitespace (spaces, tabs, newlines)
