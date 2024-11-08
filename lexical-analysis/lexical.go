package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Token represents a lexical token.
type LexicalToken struct {
	Type    string `json:"type"`
	Value  string `json:"value"`
}

// ASTNode represents a node in the Abstract Syntax Tree
type ASTNode struct {
    Type     string    `json:"type"`
    Value    string    `json:"value,omitempty"`
    Children []ASTNode `json:"children,omitempty"`
}


// TokenizeHandler should send tokens to parser
func LexicalTokenizeHandler(w http.ResponseWriter, r *http.Request) {
    sourceCode := r.URL.Query().Get("code")
    tokens := tokenize(sourceCode)

    // Forward tokens to the parsing service
    SendTokensToParser(tokens)
    
    // Respond with tokens for verification
    json.NewEncoder(w).Encode(tokens)
}


// tokenize is a basic tokenizing function (stub)
func tokenize(code string) []LexicalToken {
    var tokens []LexicalToken
    var currentWord string

    // Define keywords, operators, and delimiters
    keywords := map[string]bool{"alignas": true, "alignof": true, "and":true,	"and_eq":true,	"asm":true,	"auto":true,	"bitand":true,	"bitor":true,	"bool":true,	"break":true,	"case":true,	"catch":true,	"char":true,	"char8_t":true,	"char16_t":true,	"char32_t":true,	"class":true, "compl":true, "concept":true, "const":true, "const_cast":true, "consteval":true, "constexpr":true, "constinit":true, "continue":true, "co_await":true, "co_return":true, "co_yield":true, "decltype":true, "default":true, "delete":true, "do":true, "double":true, "dynamic_cast":true, "else":true, "enum":true, "explicit":true, "export":true, "extern":true, "false":true, "float":true, "for":true, "friend":true, "goto":true, "if":true, "inline":true, "int":true, "long":true, "mutable":true, "namespace":true, "new":true, "noexcept":true, "not":true, "not_eq":true, "nullptr":true, "operator":true, "or":true, "or_eq":true, "private":true, "protected":true, "public":true, "register":true, "reinterpret_cast":true, "requires":true, "return":true, "short":true, "signed":true, "sizeof":true, "static":true, "static_assert":true, "static_cast":true, "struct":true, "switch":true, "template":true, "this":true, "thread_local":true, "throw":true, "true":true, "try":true, "typedef":true, "typeid":true, "typename":true, "union":true, "unsigned":true, "using":true, "virtual":true, "void":true, "volatile":true, "wchar_t":true, "while":true, "xor":true, "xor_eq":true}

    operators := map[string]bool{"=": true, "+": true, "-": true, "*": true, "/": true, "%": true, "++": true, "--": true, "==": true, "!=": true, ">": true, "<": true, ">=": true, "<=": true, "&&": true, "||": true, "!": true, "&": true, "|": true, "^": true, "~": true, "<<": true, ">>": true, "+=": true, "-=": true, "*=": true, "/=": true, "%=": true, "&=": true, "|=": true, "^=": true, "<<=": true, ">>=": true, "=>": true, "->": true}

    delimiters := map[string]bool{";": true, "(": true, ")": true, "{": true, "}": true, "[": true, "]": true, ",": true, ".": true, ":": true}

    for i := 0; i < len(code); i++ {
        char := string(code[i])

        // Check if character is whitespace to separate tokens
        if char == " " || char == "\n" || char == "\t" {
            if currentWord != "" {
                tokens = append(tokens, classifyToken(currentWord, keywords, operators))
                currentWord = ""
            }
            continue
        }

        // Check if character is an operator
        if operators[char] {
            if currentWord != "" {
                tokens = append(tokens, classifyToken(currentWord, keywords, operators))
                currentWord = ""
            }
            tokens = append(tokens, LexicalToken{Type: "OPERATOR", Value: char})
            continue
        }

        // Check if character is a delimiter
        if delimiters[char] {
            if currentWord != "" {
                tokens = append(tokens, classifyToken(currentWord, keywords, operators))
                currentWord = ""
            }
            tokens = append(tokens, LexicalToken{Type: "DELIMITER", Value: char})  // Add delimiter token here
            continue
        }

        // If none of the above, accumulate the character in currentWord
        currentWord += char
    }

    // Final token for any remaining word
    if currentWord != "" {
        tokens = append(tokens, classifyToken(currentWord, keywords, operators))
    }

    return tokens
}

// classifyToken classifies a word as a keyword, identifier, or number
func classifyToken(word string, keywords, operators map[string]bool) LexicalToken {
    if keywords[word] {
        return LexicalToken{Type: "KEYWORD", Value: word}
    }
    if _, err := strconv.Atoi(word); err == nil {
        return LexicalToken{Type: "NUMBER", Value: word}
    }
	if operators[word] {
		return LexicalToken{Type: "OPERATOR", Value: word}
	}
    return LexicalToken{Type: "IDENTIFIER", Value: word}
}


// main function to start the server
func main() {
	http.HandleFunc("/tokenize", LexicalTokenizeHandler)
	fmt.Println("Lexical Analysis service is running on port 8081.")
	// Example URLs to test the service:

	// int x = 10;
	fmt.Println("http://localhost:8081/tokenize?code=int%20x%20=%2010%3B")
	http.ListenAndServe(":8081", nil)
}

/* 
run in terminal: go run lexical.go

When you start the server, it listens on port 8081.
To test it, you can send a GET request with the source code in the URL:

http://localhost:8081/tokenize?code=int%20x%20=%205;

The `?code=int%20x%20=%205;` part is: 
A query parameter with:
	code as the Key.
	int x = 5; as the Value.

(spaces are URL-encoded as %20).

In this setup, you don't need a separate source code file. 
Instead, the code is passed directly as text in the URL. 
For production, you might accept a POST request with the source code in the request body if you want to send longer or more complex code snippets.
*/

// SendTokensToParser sends the tokens to the parsing service
func SendTokensToParser(tokens []LexicalToken) {
    jsonData, _ := json.Marshal(tokens)
    resp, err := http.Post("http://parsing:8082/parse", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error sending tokens to parser:", err)
        return
    }
    defer resp.Body.Close()

    var ast ASTNode
    json.NewDecoder(resp.Body).Decode(&ast)
    fmt.Println("Received AST from parser:", ast)
}

// Call SendTokensToParser within TokenizeHandler after tokenization


// Modify TokenizeHandler to send tokens to parser after tokenization
func TokenizeHandler(w http.ResponseWriter, r *http.Request) {
    sourceCode := r.URL.Query().Get("code")
    tokens := tokenize(sourceCode)

    // Send tokens to parsing service
    SendTokensToParser(tokens)

    // Respond with tokens for local testing
    json.NewEncoder(w).Encode(tokens)
}

