// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget2.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// import "../NTypes"

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 204,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 5, 21, 156, 10, 21, 3, 22, 3, 22, 3, 22, 3, 23, 6, 23, 162, 10, 23,
	13, 23, 14, 23, 163, 3, 24, 6, 24, 167, 10, 24, 13, 24, 14, 24, 168, 3,
	25, 3, 25, 3, 25, 3, 25, 7, 25, 175, 10, 25, 12, 25, 14, 25, 178, 11, 25,
	3, 25, 3, 25, 3, 26, 6, 26, 183, 10, 26, 13, 26, 14, 26, 184, 3, 26, 3,
	26, 3, 27, 5, 27, 190, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	7, 28, 198, 10, 28, 12, 28, 14, 28, 201, 11, 28, 3, 28, 3, 28, 2, 2, 29,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 3, 2, 8, 4,
	2, 62, 62, 64, 64, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 3, 2, 36, 36, 5,
	2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 2, 213, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2,
	2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 3, 57, 3, 2, 2,
	2, 5, 59, 3, 2, 2, 2, 7, 61, 3, 2, 2, 2, 9, 63, 3, 2, 2, 2, 11, 65, 3,
	2, 2, 2, 13, 70, 3, 2, 2, 2, 15, 76, 3, 2, 2, 2, 17, 81, 3, 2, 2, 2, 19,
	84, 3, 2, 2, 2, 21, 91, 3, 2, 2, 2, 23, 96, 3, 2, 2, 2, 25, 100, 3, 2,
	2, 2, 27, 105, 3, 2, 2, 2, 29, 110, 3, 2, 2, 2, 31, 117, 3, 2, 2, 2, 33,
	122, 3, 2, 2, 2, 35, 131, 3, 2, 2, 2, 37, 139, 3, 2, 2, 2, 39, 146, 3,
	2, 2, 2, 41, 155, 3, 2, 2, 2, 43, 157, 3, 2, 2, 2, 45, 161, 3, 2, 2, 2,
	47, 166, 3, 2, 2, 2, 49, 170, 3, 2, 2, 2, 51, 182, 3, 2, 2, 2, 53, 189,
	3, 2, 2, 2, 55, 193, 3, 2, 2, 2, 57, 58, 7, 63, 2, 2, 58, 4, 3, 2, 2, 2,
	59, 60, 7, 126, 2, 2, 60, 6, 3, 2, 2, 2, 61, 62, 7, 42, 2, 2, 62, 8, 3,
	2, 2, 2, 63, 64, 7, 43, 2, 2, 64, 10, 3, 2, 2, 2, 65, 66, 7, 118, 2, 2,
	66, 67, 7, 123, 2, 2, 67, 68, 7, 114, 2, 2, 68, 69, 7, 103, 2, 2, 69, 12,
	3, 2, 2, 2, 70, 71, 7, 114, 2, 2, 71, 72, 7, 116, 2, 2, 72, 73, 7, 107,
	2, 2, 73, 74, 7, 112, 2, 2, 74, 75, 7, 118, 2, 2, 75, 14, 3, 2, 2, 2, 76,
	77, 7, 117, 2, 2, 77, 78, 7, 107, 2, 2, 78, 79, 7, 124, 2, 2, 79, 80, 7,
	103, 2, 2, 80, 16, 3, 2, 2, 2, 81, 82, 7, 99, 2, 2, 82, 83, 7, 117, 2,
	2, 83, 18, 3, 2, 2, 2, 84, 85, 7, 117, 2, 2, 85, 86, 7, 118, 2, 2, 86,
	87, 7, 116, 2, 2, 87, 88, 7, 107, 2, 2, 88, 89, 7, 112, 2, 2, 89, 90, 7,
	105, 2, 2, 90, 20, 3, 2, 2, 2, 91, 92, 7, 117, 2, 2, 92, 93, 7, 106, 2,
	2, 93, 94, 7, 99, 2, 2, 94, 95, 7, 51, 2, 2, 95, 22, 3, 2, 2, 2, 96, 97,
	7, 111, 2, 2, 97, 98, 7, 102, 2, 2, 98, 99, 7, 55, 2, 2, 99, 24, 3, 2,
	2, 2, 100, 101, 7, 112, 2, 2, 101, 102, 7, 118, 2, 2, 102, 103, 7, 104,
	2, 2, 103, 104, 7, 117, 2, 2, 104, 26, 3, 2, 2, 2, 105, 106, 7, 104, 2,
	2, 106, 107, 7, 107, 2, 2, 107, 108, 7, 110, 2, 2, 108, 109, 7, 103, 2,
	2, 109, 28, 3, 2, 2, 2, 110, 111, 7, 114, 2, 2, 111, 112, 7, 99, 2, 2,
	112, 113, 7, 101, 2, 2, 113, 114, 7, 109, 2, 2, 114, 115, 7, 103, 2, 2,
	115, 116, 7, 118, 2, 2, 116, 30, 3, 2, 2, 2, 117, 118, 7, 114, 2, 2, 118,
	119, 7, 101, 2, 2, 119, 120, 7, 99, 2, 2, 120, 121, 7, 114, 2, 2, 121,
	32, 3, 2, 2, 2, 122, 123, 7, 103, 2, 2, 123, 124, 7, 122, 2, 2, 124, 125,
	7, 107, 2, 2, 125, 126, 7, 104, 2, 2, 126, 127, 7, 107, 2, 2, 127, 128,
	7, 112, 2, 2, 128, 129, 7, 104, 2, 2, 129, 130, 7, 113, 2, 2, 130, 34,
	3, 2, 2, 2, 131, 132, 7, 103, 2, 2, 132, 133, 7, 122, 2, 2, 133, 134, 7,
	118, 2, 2, 134, 135, 7, 116, 2, 2, 135, 136, 7, 99, 2, 2, 136, 137, 7,
	101, 2, 2, 137, 138, 7, 118, 2, 2, 138, 36, 3, 2, 2, 2, 139, 140, 7, 104,
	2, 2, 140, 141, 7, 107, 2, 2, 141, 142, 7, 110, 2, 2, 142, 143, 7, 118,
	2, 2, 143, 144, 7, 103, 2, 2, 144, 145, 7, 116, 2, 2, 145, 38, 3, 2, 2,
	2, 146, 147, 7, 46, 2, 2, 147, 40, 3, 2, 2, 2, 148, 156, 9, 2, 2, 2, 149,
	150, 7, 64, 2, 2, 150, 156, 7, 63, 2, 2, 151, 152, 7, 62, 2, 2, 152, 156,
	7, 63, 2, 2, 153, 154, 7, 63, 2, 2, 154, 156, 7, 63, 2, 2, 155, 148, 3,
	2, 2, 2, 155, 149, 3, 2, 2, 2, 155, 151, 3, 2, 2, 2, 155, 153, 3, 2, 2,
	2, 156, 42, 3, 2, 2, 2, 157, 158, 7, 93, 2, 2, 158, 159, 7, 95, 2, 2, 159,
	44, 3, 2, 2, 2, 160, 162, 9, 3, 2, 2, 161, 160, 3, 2, 2, 2, 162, 163, 3,
	2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 46, 3, 2, 2,
	2, 165, 167, 9, 4, 2, 2, 166, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168,
	166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 48, 3, 2, 2, 2, 170, 176, 7,
	36, 2, 2, 171, 172, 7, 36, 2, 2, 172, 175, 7, 36, 2, 2, 173, 175, 10, 5,
	2, 2, 174, 171, 3, 2, 2, 2, 174, 173, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2,
	176, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 179, 3, 2, 2, 2, 178,
	176, 3, 2, 2, 2, 179, 180, 7, 36, 2, 2, 180, 50, 3, 2, 2, 2, 181, 183,
	9, 6, 2, 2, 182, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 182, 3, 2,
	2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 8, 26, 2, 2,
	187, 52, 3, 2, 2, 2, 188, 190, 7, 15, 2, 2, 189, 188, 3, 2, 2, 2, 189,
	190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 7, 12, 2, 2, 192, 54,
	3, 2, 2, 2, 193, 194, 7, 49, 2, 2, 194, 195, 7, 49, 2, 2, 195, 199, 3,
	2, 2, 2, 196, 198, 10, 7, 2, 2, 197, 196, 3, 2, 2, 2, 198, 201, 3, 2, 2,
	2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 202, 3, 2, 2, 2, 201,
	199, 3, 2, 2, 2, 202, 203, 8, 28, 2, 2, 203, 56, 3, 2, 2, 2, 11, 2, 155,
	163, 168, 174, 176, 184, 189, 199, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'|'", "'('", "')'", "'type'", "'print'", "'size'", "'as'",
	"'string'", "'sha1'", "'md5'", "'ntfs'", "'file'", "'packet'", "'pcap'",
	"'exifinfo'", "'extract'", "'filter'", "','", "", "'[]'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "COMPOP", "LISTOP", "INT", "ID", "STRING", "WS", "NL", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "COMPOP", "LISTOP", "INT", "ID", "STRING", "WS", "NL",
	"LINE_COMMENT",
}

type Nugget2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewNugget2Lexer(input antlr.CharStream) *Nugget2Lexer {

	l := new(Nugget2Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Nugget2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Nugget2Lexer tokens.
const (
	Nugget2LexerT__0         = 1
	Nugget2LexerT__1         = 2
	Nugget2LexerT__2         = 3
	Nugget2LexerT__3         = 4
	Nugget2LexerT__4         = 5
	Nugget2LexerT__5         = 6
	Nugget2LexerT__6         = 7
	Nugget2LexerT__7         = 8
	Nugget2LexerT__8         = 9
	Nugget2LexerT__9         = 10
	Nugget2LexerT__10        = 11
	Nugget2LexerT__11        = 12
	Nugget2LexerT__12        = 13
	Nugget2LexerT__13        = 14
	Nugget2LexerT__14        = 15
	Nugget2LexerT__15        = 16
	Nugget2LexerT__16        = 17
	Nugget2LexerT__17        = 18
	Nugget2LexerT__18        = 19
	Nugget2LexerCOMPOP       = 20
	Nugget2LexerLISTOP       = 21
	Nugget2LexerINT          = 22
	Nugget2LexerID           = 23
	Nugget2LexerSTRING       = 24
	Nugget2LexerWS           = 25
	Nugget2LexerNL           = 26
	Nugget2LexerLINE_COMMENT = 27
)
