package assembler

import "fmt"
import "io/ioutil"
import "unicode"
import "unicode/utf8"

type token struct {
	Typ tokenType
	Val string
}

type tokenType int

const eof = -1

const (
	TokEOF     tokenType = iota
	tokStmtSep           // newline or ;
	tokIdentifier
	tokColon
	tokDecimal
	tokComma
	tokRegister
	tokPlus
	tokMinus
	tokDot
	tokEquals
)

type stateFn func(*lexer) stateFn

type lexer struct {
	input  string     // input string TODO: could stream it...
	start  int        // start of the current token
	pos    int        // current location in the input
	width  int        // width of last rune
	tokens chan token // channel for scanned tokens
	state  stateFn    // current state
}

func LexFromFile(path string) *lexer {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err));
	}
	return Lex(string(bytes))
}

func Lex(input string) *lexer {
	l := &lexer{
		input:  input,
		tokens: make(chan token),
		state:  lexStatement,
	}
	go l.run()
	return l
}

func (l *lexer) run() {
	for l.state = lexStatement; l.state != nil; {
		l.state = l.state(l)
	}
}

func (l *lexer) NextToken() token {
	tok := <-l.tokens
	return tok
}

func (l *lexer) emit(t tokenType) {
	l.tokens <- token{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) next() (r rune) {
	fmt.Printf("next(); start=%v, pos=%v, width=%v\n", l.start, l.pos, l.width)
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	fmt.Printf("read rune %d, pos=%d\n", r, l.pos)
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func lexStatement(l *lexer) stateFn {
	fmt.Printf("lexStatement")
Loop:
	for {
		switch r := l.next(); {
		case r == eof:
			fmt.Printf("eof\n")
			l.emit(TokEOF)
			break Loop
		case isSpace(r):
			fmt.Printf("isSpace\n")
			l.ignore()
		case isEndOfLine(r) || r == ';':
			fmt.Printf("isEndOfLine\n")
			return lexStatementSep
		case r == ':':
			fmt.Printf(":\n")
			l.emit(tokColon)
		case r == ',':
			fmt.Printf(",\n")
			l.emit(tokComma)
		case r == '$':
			fmt.Printf("$\n")
			return lexRegister
		case r == '+':
			fmt.Printf("+\n")
			l.emit(tokPlus)
		case r == '-':
			fmt.Printf("-\n")
			l.emit(tokMinus)
		case r == '.':
			fmt.Printf(".\n")
			return lexDirective
		case isAlphaNumeric(r):
			fmt.Printf("isAlphaNumeric\n")
			return lexIdentifier
		default:
			fmt.Printf("??????\n")
			panic(fmt.Sprintf("unexpected character %v", r))
		}
	}
	return nil
}

func lexStatementSep(l *lexer) stateFn {
	fmt.Printf("lexStatementSep\n")
Loop:
	for {
		switch r := l.next(); {
		case isSpace(r) || isEndOfLine(r) || r == ';':
			// eat it
		default:
			l.backup()
			l.emit(tokStmtSep)
			break Loop
		}
	}
	return lexStatement
}

func lexIdentifier(l *lexer) stateFn {
	fmt.Printf("lexIdentifier\n")
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// eat it
		default:
			l.backup()
			l.emit(tokIdentifier)
			break Loop
		}
	}
	fmt.Printf("lexIdentifier done\n")
	return lexStatement
}

// Lex a register name $blah
// The $ has already been lexed
func lexRegister(l *lexer) stateFn {
	l.ignore() // don't want the $ as part of the register name
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// eat it
		default:
			l.backup()
			l.emit(tokRegister)
			break Loop
		}
	}
	return lexStatement
}

// Lex a directive that starts with .
func lexDirective(l *lexer) stateFn {
	l.emit(tokDot)
	return lexStatement
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}
