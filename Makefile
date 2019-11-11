ANTLR_VERSION=4.7.2
GRAMMAR_VERSION=33e2791d5fe4d92f6842f2ed1f8305fb4bbbb0c6

sql-mask: main.go parser/mysql_lexer.go parser/mysql_parser.go
	go build -o $@ $<

libsql_mask.%: main.go parser/mysql_lexer.go parser/mysql_parser.go
	CGO_ENABLED=1 GOOS=`tr -d . <<< $(suffix $@)` go build -ldflags="-w -s" -buildmode=c-shared -o $@

parser/mysql_lexer.go parser/mysql_parser.go: parser/antlr.jar parser/MySqlLexer.g4 parser/MySqlParser.g4
	java -jar $< -Dlanguage=Go $(@D)/*.g4

parser/antlr.jar:
	curl -o $@ https://www.antlr.org/download/antlr-$(ANTLR_VERSION)-complete.jar

parser/MySqlLexer.g4:
	curl -o $@ https://raw.githubusercontent.com/antlr/grammars-v4/$(GRAMMAR_VERSION)/mysql/Positive-Technologies/MySqlLexer.g4

parser/MySqlParser.g4:
	curl -o $@ https://raw.githubusercontent.com/antlr/grammars-v4/$(GRAMMAR_VERSION)/mysql/Positive-Technologies/MySqlParser.g4
