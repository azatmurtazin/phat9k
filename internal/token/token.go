package token

type TokenType int

const (
	T_EOF TokenType = iota
	T_OPEN_TAG
	T_CLOSE_TAG
	T_STRING
	T_LNUMBER
	T_DNUMBER
	T_CONSTANT_ENCAPSED_STRING
	T_VARIABLE
	T_ECHO
	T_PRINT
	T_FUNCTION
	T_CLASS
	T_INTERFACE
	T_TRAIT
	T_EXTENDS
	T_IMPLEMENTS
	T_NEW
	T_IF
	T_ELSEIF
	T_ELSE
	T_SWITCH
	T_CASE
	T_DEFAULT
	T_FOR
	T_FOREACH
	T_WHILE
	T_DO
	T_RETURN
	T_BREAK
	T_CONTINUE
	T_EXIT
	T_DIE
	T_TRY
	T_CATCH
	T_THROW
	T_NAMESPACE
	T_USE
	T_PUBLIC
	T_PRIVATE
	T_PROTECTED
	T_STATIC
	T_FINAL
	T_ABSTRACT
	T_CONST
	T_VAR
	T_ARRAY
	T_LIST
	T_GLOBAL
	T_UNSET
	T_EMPTY
	T_ISSET
	T_INCLUDE
	T_INCLUDE_ONCE
	T_REQUIRE
	T_REQUIRE_ONCE
	T_TRUE
	T_FALSE
	T_NULL
	T_OPEN_TAG_ECHO
	T_OPEN_TAG_ECHO_EOF
	T_NS_SEPARATOR
	T_NAME_FULLY_QUALIFIED
	T_NAME_QUALIFIED
	T_NAME_RELATIVE
	T_NAME_VARIABLE
	T_STRING_VARNAME
	T_VARIABLE_VARIABLE
	T_LEFT_BRACE
	T_RIGHT_BRACE
	T_LEFT_BRACKET
	T_RIGHT_BRACKET
	T_LEFT_PAREN
	T_RIGHT_PAREN
	T_LEFT_CURLY_BRACE
	T_RIGHT_CURLY_BRACE
	T_SEMICOLON
	T_COMMA
	T_DOT
	T_COLON
	T_QUESTION
	T_AT
	T_AMPERSAND
	T_BITWISE_AND
	T_MULT
	T_DIV
	T_PLUS
	T_MINUS
	T_CONCAT
	T_MOD
	T_EQUAL
	T_PLUS_EQUAL
	T_MINUS_EQUAL
	T_MULT_EQUAL
	T_DIV_EQUAL
	T_CONCAT_EQUAL
	T_MOD_EQUAL
	T_BITWISE_OR_EQUAL
	T_BITWISE_XOR_EQUAL
	T_BITWISE_AND_EQUAL
	T_SL_EQUAL
	T_SR_EQUAL
	T_COALESCE_EQUAL
	T_BOOLEAN_AND
	T_BOOLEAN_OR
	T_BOOLEAN_NOT
	T_LOGICAL_AND
	T_LOGICAL_OR
	T_LOGICAL_XOR
	T_IS_IDENTICAL
	T_IS_NOT_IDENTICAL
	T_IS_EQUAL
	T_IS_NOT_EQUAL
	T_SMALLER
	T_SMALLER_OR_EQUAL
	T_GREATER
	T_GREATER_OR_EQUAL
	T_SPACESHIP
	T_SL
	T_SR
	T_INC
	T_DEC
	T_DOUBLE_COLON
	T_NS_CODEL
	T_COALESCE
	T_DOUBLE_ARROW
	T_INSTANCEOF
	T_UNSAFE_COALESCE
	T_LOGICAL_OR_DBL
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func (t Token) String() string {
	if t.Literal != "" {
		return t.Literal
	}
	return tokenNames[t.Type]
}

var tokenNames = map[TokenType]string{
	T_EOF:                      "EOF",
	T_OPEN_TAG:                 "<?php",
	T_CLOSE_TAG:                "?>",
	T_STRING:                   "STRING",
	T_LNUMBER:                  "NUMBER",
	T_DNUMBER:                  "DNUMBER",
	T_CONSTANT_ENCAPSED_STRING: "CONSTANT_ENCAPSED_STRING",
	T_VARIABLE:                 "VARIABLE",
	T_ECHO:                     "echo",
	T_PRINT:                    "print",
	T_FUNCTION:                 "function",
	T_CLASS:                    "class",
	T_INTERFACE:                "interface",
	T_TRAIT:                    "trait",
	T_EXTENDS:                  "extends",
	T_IMPLEMENTS:               "implements",
	T_NEW:                      "new",
	T_IF:                       "if",
	T_ELSEIF:                   "elseif",
	T_ELSE:                     "else",
	T_SWITCH:                   "switch",
	T_CASE:                     "case",
	T_DEFAULT:                  "default",
	T_FOR:                      "for",
	T_FOREACH:                  "foreach",
	T_WHILE:                    "while",
	T_DO:                       "do",
	T_RETURN:                   "return",
	T_BREAK:                    "break",
	T_CONTINUE:                 "continue",
	T_EXIT:                     "exit",
	T_DIE:                      "die",
	T_TRY:                      "try",
	T_CATCH:                    "catch",
	T_THROW:                    "throw",
	T_NAMESPACE:                "namespace",
	T_USE:                      "use",
	T_PUBLIC:                   "public",
	T_PRIVATE:                  "private",
	T_PROTECTED:                "protected",
	T_STATIC:                   "static",
	T_FINAL:                    "final",
	T_ABSTRACT:                 "abstract",
	T_CONST:                    "const",
	T_VAR:                      "var",
	T_ARRAY:                    "array",
	T_LIST:                     "list",
	T_GLOBAL:                   "global",
	T_UNSET:                    "unset",
	T_EMPTY:                    "empty",
	T_ISSET:                    "isset",
	T_INCLUDE:                  "include",
	T_INCLUDE_ONCE:             "include_once",
	T_REQUIRE:                  "require",
	T_REQUIRE_ONCE:             "require_once",
	T_TRUE:                     "true",
	T_FALSE:                    "false",
	T_NULL:                     "null",
	T_OPEN_TAG_ECHO:            "<?=",
	T_OPEN_TAG_ECHO_EOF:        "<?=",
	T_NS_SEPARATOR:             "Namespace separator",
	T_NAME_FULLY_QUALIFIED:     "NAME_FULLY_QUALIFIED",
	T_NAME_QUALIFIED:           "NAME_QUALIFIED",
	T_NAME_RELATIVE:            "NAME_RELATIVE",
	T_NAME_VARIABLE:            "NAME_VARIABLE",
	T_STRING_VARNAME:           "STRING_VARNAME",
	T_VARIABLE_VARIABLE:        "Variable variable",
	T_LEFT_BRACE:               "{",
	T_RIGHT_BRACE:              "}",
	T_LEFT_BRACKET:             "[",
	T_RIGHT_BRACKET:            "]",
	T_LEFT_PAREN:               "(",
	T_RIGHT_PAREN:              ")",
	T_LEFT_CURLY_BRACE:         "{",
	T_RIGHT_CURLY_BRACE:        "}",
	T_SEMICOLON:                ";",
	T_COMMA:                    ",",
	T_DOT:                      ".",
	T_COLON:                    ":",
	T_QUESTION:                 "?",
	T_AT:                       "@",
	T_AMPERSAND:                "&",
	T_BITWISE_AND:              "&",
	T_MULT:                     "*",
	T_DIV:                      "/",
	T_PLUS:                     "+",
	T_MINUS:                    "-",
	T_CONCAT:                   ".",
	T_MOD:                      "%",
	T_EQUAL:                    "=",
	T_LOGICAL_AND:              "&&",
	T_LOGICAL_OR:               "||",
	T_BOOLEAN_AND:              "&&",
	T_BOOLEAN_OR:               "||",
	T_LOGICAL_XOR:              "xor",
	T_IS_IDENTICAL:             "===",
	T_IS_NOT_IDENTICAL:         "!==",
	T_IS_EQUAL:                 "==",
	T_IS_NOT_EQUAL:             "!=",
	T_SMALLER:                  "<",
	T_SMALLER_OR_EQUAL:         "<=",
	T_GREATER:                  ">",
	T_GREATER_OR_EQUAL:         ">=",
	T_SPACESHIP:                "<=>",
	T_SL:                       "<<",
	T_SR:                       ">>",
	T_INC:                      "++",
	T_DEC:                      "--",
	T_DOUBLE_COLON:             "::",
	T_NS_CODEL:                 "\\",
	T_COALESCE:                 "??",
	T_DOUBLE_ARROW:             "=>",
	T_INSTANCEOF:               "instanceof",
}
