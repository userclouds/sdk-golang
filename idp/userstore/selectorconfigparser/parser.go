// Code generated by goyacc -o idp/userstore/selectorconfigparser/parser.go idp/userstore/selectorconfigparser/parser.y. DO NOT EDIT.

package selectorconfigparser

import __yyfmt__ "fmt"

type yySymType struct {
	yys int
}

const ABS_OPERATOR = 57346
const ANY = 57347
const ARRAY_OPERATOR = 57348
const BOOL_VALUE = 57349
const COLUMN_IDENTIFIER = 57350
const COLUMN_OPERATOR = 57351
const COMMA = 57352
const CONJUNCTION = 57353
const DATE_ARGUMENT = 57354
const DATE_OPERATOR = 57355
const INT_VALUE = 57356
const IS = 57357
const LEFT_BRACKET = 57358
const LEFT_PARENTHESIS = 57359
const NOT = 57360
const NULL = 57361
const NUMBER_PART_OPERATOR = 57362
const QUOTED_VALUE = 57363
const RIGHT_BRACKET = 57364
const RIGHT_PARENTHESIS = 57365
const OPERATOR = 57366
const UNKNOWN = 57367
const VALUE_PLACEHOLDER = 57368

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ABS_OPERATOR",
	"ANY",
	"ARRAY_OPERATOR",
	"BOOL_VALUE",
	"COLUMN_IDENTIFIER",
	"COLUMN_OPERATOR",
	"COMMA",
	"CONJUNCTION",
	"DATE_ARGUMENT",
	"DATE_OPERATOR",
	"INT_VALUE",
	"IS",
	"LEFT_BRACKET",
	"LEFT_PARENTHESIS",
	"NOT",
	"NULL",
	"NUMBER_PART_OPERATOR",
	"QUOTED_VALUE",
	"RIGHT_BRACKET",
	"RIGHT_PARENTHESIS",
	"OPERATOR",
	"UNKNOWN",
	"VALUE_PLACEHOLDER",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 71

var yyAct = [...]int8{
	43, 48, 3, 30, 18, 24, 21, 24, 21, 50,
	58, 55, 51, 22, 54, 22, 25, 29, 25, 34,
	23, 49, 23, 44, 32, 20, 12, 20, 47, 33,
	45, 39, 28, 16, 19, 10, 52, 41, 31, 5,
	6, 38, 35, 46, 7, 5, 6, 15, 4, 37,
	7, 8, 14, 56, 57, 36, 9, 8, 27, 26,
	1, 53, 42, 40, 11, 13, 2, 0, 0, 0,
	17,
}

var yyPact = [...]int16{
	31, -1000, 45, 11, 31, -1000, 35, 30, 16, 31,
	-1, -1000, 40, 9, 37, 12, 37, -1000, 1, -1000,
	-1000, -1000, -1000, -1000, 39, 1, -1000, 22, -1000, 8,
	53, -1000, -1000, 12, 52, -1000, 1, 7, -1000, -1000,
	37, 5, -5, 14, 51, -1000, -9, -1000, -12, -1000,
	-1000, -5, -1000, 1, -1000, -1000, -13, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 60, 66, 2, 3, 1, 23, 64, 0,
}

var yyR1 = [...]int8{
	0, 1, 1, 3, 3, 3, 3, 2, 2, 2,
	2, 6, 6, 6, 6, 6, 6, 8, 8, 4,
	4, 4, 5, 5, 5, 7, 7,
}

var yyR2 = [...]int8{
	0, 1, 3, 1, 4, 6, 6, 4, 3, 2,
	3, 1, 1, 1, 1, 4, 3, 1, 3, 1,
	1, 3, 1, 1, 3, 2, 3,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, 17, 8, 9, 13, 20, 11,
	24, -7, 15, -1, 17, 17, 17, -1, 5, -6,
	26, 7, 14, 21, 6, 17, 19, 18, 23, -3,
	-4, 26, 12, 17, -3, -6, 16, -6, 19, 23,
	10, -4, 10, -8, -6, 23, -3, 23, -5, 26,
	14, 17, 22, 10, 23, 23, -5, -8, 23,
}

var yyDef = [...]int8{
	0, -2, 1, 0, 0, 3, 0, 0, 0, 0,
	0, 9, 0, 0, 0, 0, 0, 2, 0, 8,
	11, 12, 13, 14, 0, 0, 25, 0, 10, 0,
	0, 19, 20, 0, 0, 7, 0, 0, 26, 4,
	0, 0, 0, 0, 17, 16, 0, 21, 0, 22,
	23, 0, 15, 0, 5, 6, 0, 18, 24,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = true
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1 /* scope-safe */

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	}
	goto yystack /* stack new state and value */
}
