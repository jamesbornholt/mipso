package assembler

import "bufio"
import "fmt"
import "os"
import "strings"

import "github.com/jamesbornholt/mipso/isa"

type Statement struct {
    Typ statementType
    Insn isa.Instruction
    Label string
    // TODO: data
    // TODO: symdef
    Directive string
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
    for scan.Scan() {
        line := scan.Text()
        s, err := parseStatement(line)
        if err != nil {
            return stmts, err
        }
        stmts = append(stmts, s)
        fmt.Printf("%s\n", line)
    }

    return stmts, nil
}

func parseStatement(line string) (stmt Statement, err error) {
    line = strings.TrimSpace(line)
    switch {
    case len(line) == 0:
        stmt, err = Statement{}, nil
    case line[0] == '.':
        stmt, err = parseDirective(line)
    case strings.ContainsRune(line, ':'):
        stmt, err = parseLabel(line)
    default:
        stmt, err = parseInstruction(line)
    }
    return
}

func parseDirective(line string) (stmt Statement, err error) {
    s := Statement{Typ:stmtDirective, Directive:line}
    return s, nil
}

func parseLabel(line string) (stmt Statement, err error) {
    s := Statement{Typ:stmtLabel, Label:line}
    return s, nil
}

func parseInstruction(line string) (stmt Statement, err error) {
    s := Statement{Typ:stmtInstruction, Insn: isa.InsnJ{}}
    return s, nil
}
