sql-mask: main.go parser/mysql_lexer.go parser/mysql_parser.go antlr-utils/case_changing_stream.go
	go build -o $@ $<

parser/mysql_lexer.go parser/mysql_parser.go: parser/antlr.jar parser/MySqlLexer.g4 parser/MySqlParser.g4
	java -jar $< -Dlanguage=Go $(@D)/*.g4

parser/antlr.jar:
	curl -o $@ https://www.antlr.org/download/antlr-4.7.2-complete.jar

parser/MySqlLexer.g4:
	curl -o $@ https://raw.githubusercontent.com/antlr/grammars-v4/master/mysql/MySqlLexer.g4

parser/MySqlParser.g4:
	curl -o $@ https://raw.githubusercontent.com/antlr/grammars-v4/master/mysql/MySqlParser.g4
	patch $@ $(@D)/parser.patch

antlr-utils/case_changing_stream.go:
	curl -o $@ https://raw.githubusercontent.com/antlr/antlr4/master/doc/resources/case_changing_stream.go
	patch $@ $(@D)/stream.patch
