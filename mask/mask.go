package mask

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	utils "github.com/kaz/sql-mask/antlr-utils"
	"github.com/kaz/sql-mask/parser"
)

type (
	errorListener struct {
		*antlr.DefaultErrorListener

		err error
	}
	parserListener struct {
		*parser.BaseMySqlParserListener

		inConstant bool
		hasError   bool
		positions  []int
	}
)

func (el *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.err = fmt.Errorf("line %d:%d %s", line, column, msg)
}

func (pl *parserListener) EnterConstant(ctx *parser.ConstantContext) {
	pl.inConstant = true
}
func (pl *parserListener) ExitConstant(ctx *parser.ConstantContext) {
	pl.inConstant = false
}
func (pl *parserListener) VisitTerminal(term antlr.TerminalNode) {
	if pl.inConstant {
		sym := term.GetSymbol()
		pl.positions = append(pl.positions, sym.GetStart(), sym.GetStop()+1)
	}
}

func Mask(sql string) (string, error) {
	psr := parser.NewMySqlParser(antlr.NewCommonTokenStream(parser.NewMySqlLexer(utils.NewCaseChangingStream(antlr.NewInputStream(sql), true)), 0))
	psr.RemoveErrorListeners()

	eListener := &errorListener{}
	pListener := &parserListener{}

	psr.AddErrorListener(eListener)
	antlr.ParseTreeWalkerDefault.Walk(pListener, psr.Root())

	if eListener.err != nil {
		return "", eListener.err
	}

	b := 0
	chunks := []string{}
	for _, i := range pListener.positions {
		chunks = append(chunks, sql[b:i])
		b = i
	}
	chunks = append(chunks, sql[b:])

	for i := 1; i < len(chunks); i += 2 {
		chunks[i] = "?"
	}

	return strings.Join(chunks, ""), nil
}
