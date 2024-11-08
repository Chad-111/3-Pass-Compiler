// codegen.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ASTNode represents a node in the Abstract Syntax Tree
type CodegenASTNode struct {
    Type     string    `json:"type"`
    Value    string    `json:"value,omitempty"`
    Children []CodegenASTNode  `json:"children,omitempty"`
}

// MachineCode represents a single line of machine code
type MachineCode struct {
    Instruction string `json:"instruction"`
}

// GenerateHandler handles converting AST to machine code
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
    var ast CodegenASTNode
    if err := json.NewDecoder(r.Body).Decode(&ast); err != nil {
        http.Error(w, "Invalid AST received", http.StatusBadRequest)
        fmt.Println("Error decoding AST:", err)
        return
    }
    fmt.Println("AST received in code generation:", ast)

    machineCode := generateCode(ast)
    json.NewEncoder(w).Encode(machineCode)
}

// generateCode is a basic code generation function (stub)
func generateCode(ast CodegenASTNode) []MachineCode {
    var machineCode []MachineCode

    // Example: Handle a simple assignment (e.g., "int x = 10;")
    if ast.Type == "Program" && len(ast.Children) > 0 {
        declaration := ast.Children[0]
        if declaration.Type == "Declaration" && declaration.Value == "int x = 10;" {
            machineCode = append(machineCode, MachineCode{Instruction: "LOAD 10, R1"})
            machineCode = append(machineCode, MachineCode{Instruction: "STORE R1, x"})
        }
    }

    return machineCode
}



func main() {
    http.HandleFunc("/generate", GenerateHandler)
    fmt.Println("Code Generation service running on port 8083")
    http.ListenAndServe(":8083", nil)
}
