grammar ECMAScript;

// 11: Lexical Grammar

SourceCharacter
    : '\u0000'..'\u10FFFF'
    ;

InputElementDiv
    : WhiteSpace
    | LineTerminator
    | Comment
    | CommonToken
    | DivPunctuator
    | RightBracePunctuator
    ;

InputElementRegExp
    : WhiteSpace
    | LineTerminator
    | Comment
    | CommonToken
    | RightBracePunctuator
    | RegularExpressionLiteral
    ;

InputElementRegExpOrTemplateTail
    : WhiteSpace
    | LineTerminator
    | Comment
    | CommonToken
    | RegularExpressionLiteral
    | TemplateSubstitutionTail
    ;

InputElementTemplateTail
    : WhiteSpace
    | LineTerminator
    | Comment
    | CommonToken
    | DivPunctuator
    | TemplateSubstitutionTail
    ;

// 11.1: Unicode Format-Control Characters

fragment ZWNJ   : '\u200C'; // Zero Width Non-Joiner
fragment ZWJ    : '\u200D'; // Zero Width Joiner
fragment ZWNBSP : '\uFEFF'; // Zero Width No-Break Space

// 11.2: White Space

fragment TAB  : '\u0009';
fragment VT   : '\u0009';
fragment FF   : '\u0009';
fragment SP   : '\u0009';
fragment NBSP : '\u0009';
// TODO: Any other Unicode "Space-Separator" code point <USP>

WhiteSpace
    : TAB
    | VT
    | FF
    | SP
    | NBSP
    | ZWNBSP
//  | USP
    ;

// 11.3: Line Terminators

fragment CRLF : '\u000D\u000A'; // does this work?
fragment LF   : '\u000A';
fragment CR   : '\u000D';
fragment LS   : '\u2028';
fragment PS   : '\u2029';

LineTerminator
    : LF
    | CR
    | LS
    | PS
    ;

LineTerminatorSequence
    : LF
    | CRLF
    | CR
    | LS
    | PS
    ;

// 11.4: Comments

Comment
    : MultiLineComment
    | SingleLineComment
    ;

MultiLineComment
    : '/*' MultiLineCommentChars* '*/'
    ;

MultiLineCommentChars
    : MultiLineNotAsteriskChar MultiLineCommentChars*
    | '*' PostAsteriskCommentChars
    ;

PostAsteriskChars
    : MultilineNotForwardSlashOrAsteriskChar MultiLineCommentChars*
    | '*' PostAsteriskChars
    ;

MultiLineNotAsteriskChar
    : SourceCharacter
    ;