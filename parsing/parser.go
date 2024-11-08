// parser.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Token represents a lexical token received from lexical analysis
type ParseToken struct {
    Type  string `json:"type"`
    Value string `json:"value"`
}

// ASTNode represents a node in the Abstract Syntax Tree
type ASTNode struct {
    Type     string    `json:"type"`
    Value    string    `json:"value,omitempty"`
    Children []ASTNode `json:"children,omitempty"`
}

// MachineCode represents a line of machine code or assembly instruction
type MachineCode struct {
    Instruction string `json:"instruction"`
}


// SendASTToCodeGen sends the AST to the code generation service
func SendASTToCodeGen(ast ASTNode) {
    jsonData, _ := json.Marshal(ast)
    resp, err := http.Post("http://codegen:8083/generate", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error sending AST to code generator:", err)
        return
    }
    defer resp.Body.Close()

    var machineCode []MachineCode
    json.NewDecoder(resp.Body).Decode(&machineCode)
    fmt.Println("Received machine code from code generator:", machineCode)
}


// ParseHandler should send AST to code generation
func ParseHandler(w http.ResponseWriter, r *http.Request) {
    var tokens []ParseToken
    if err := json.NewDecoder(r.Body).Decode(&tokens); err != nil {
        http.Error(w, "Invalid tokens received", http.StatusBadRequest)
        fmt.Println("Error decoding tokens:", err)
        return
    }
    fmt.Println("Tokens received in parser:", tokens)

    ast := parse(tokens)

    // Forward AST to code generation
    SendASTToCodeGen(ast)

    // Respond with AST for verification
    json.NewEncoder(w).Encode(ast)
}


// parse is a simple function to build an AST from tokens (stub)
func parse(tokens []ParseToken ) ASTNode {
    // TODO: Implement detailed parsing logic
    return ASTNode{
        Type: "Program",
        Children: []ASTNode{
            {
                Type:  "Declaration",
                Value: "int x = 10;",
            },
        },
    }
}

func main() {
    http.HandleFunc("/parse", ParseHandler)
    fmt.Println("Parsing service running on port 8082")
    http.ListenAndServe(":8082", nil)
}
