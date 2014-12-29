package assembler

import "bufio"
import "fmt"
import "os"
import "strings"

type Statement struct {
    Typ statementType
    Insn Instruction
    Label string
    // TODO: data
    // TODO: symdef
    Directive string
}

type Instruction struct {
    Opcode string
    Operands []string
}

type statementType int
const (
    stmtInstruction statementType = iota
    stmtLabel
    // TODO: data
    // TODO: symdef
    stmtDirective
)

func ParseFile(path string) ([]Statement, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    scan := bufio.NewScanner(f)

    stmts := make([]Statement, 0)
    stmt_chan := make(chan Statement, 2)
    done_chan := make(chan bool)

    go func() {
        for {
            s, more := <-stmt_chan
            if more {
                stmts = append(stmts, s)
            } else {
                done_chan <- true
                return
            }
        }
    }()

    for scan.Scan() {
        line := scan.Text()
        parseStatement(line, stmt_chan)
        fmt.Printf("%s\n", line)
    }
    close(stmt_chan)

    <-done_chan

    return stmts, nil
}

func parseStatement(line string, stmt_chan chan Statement) {
    line = strings.TrimSpace(line)
    switch {
    case len(line) == 0:
        // nothing
    case line[0] == '.':
        parseDirective(line, stmt_chan)
    case strings.ContainsRune(line, ':'):
        i := strings.IndexRune(line, ':')
        parseLabel(line[:i], stmt_chan)
        if len(line) > i+1 {
            parseInstruction(line[i+1:], stmt_chan)
        }
    default:
        parseInstruction(line, stmt_chan)
    }
    return
}

func parseDirective(line string, stmt_chan chan Statement) {
    s := Statement{Typ:stmtDirective, Directive:line}
    stmt_chan <- s
}

func parseLabel(line string, stmt_chan chan Statement) {
    s := Statement{Typ:stmtLabel, Label:line}
    stmt_chan <- s
}

func parseInstruction(line string, stmt_chan chan Statement) {
    i := strings.IndexAny(line, " \t")
    if i == -1 {
        opcode := line
        stmt_chan <- Statement{Typ:stmtInstruction, 
                               Insn: Instruction{Opcode: opcode}}
        return
    }

    opcode := line[:i]
    line = eatSpace(line)
    operands := strings.Split(line, ",")
    for n := range(operands) {
        operands[n] = strings.TrimSpace(operands[n])
    }

    insn := Instruction{Opcode: opcode, Operands: operands}
    stmt_chan <- Statement{Typ:stmtInstruction, Insn:insn}
}

func eatSpace(s string) string {
    for n := range(s) {
        if s[n] == ' ' || s[n] == '\t' {
            return s[n:]
        }
    }
    return ""
}
