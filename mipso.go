package main

import "fmt"
import "github.com/jamesbornholt/mipso/assembler"

func main() {
	// v := emit_add_3(0, 1, 2)
	// fmt.Printf("%v\n", v)

    // bar := make([]Instruction, 10)
    // fmt.Printf("%v\n", bar)

    foo, _ := assembler.ParseFile("/Users/bornholt/code/mipso/test.s")
    fmt.Printf("foo: %v\n", foo)
    // for {
    //     switch t := foo.NextToken(); {
    //     case t.Typ == assembler.TokEOF:
    //         break
    //     default:
    //         fmt.Printf("%v\n", t);
    //     }
    // }
}
