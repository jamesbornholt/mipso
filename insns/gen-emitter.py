#!/usr/bin/env python

import os
import sys

def parse_operand(op):
    op = op.strip()
    i = op.find("(")
    if i > -1:
        assert op[-1] == ")"
        b = op[:i]
        o = op[i+1:-1]
        return [b, o]
    else:
        return [op]

folder = os.path.dirname(os.path.realpath(sys.argv[0]))
constants_path = os.path.join(folder, "constants.tsv")
constants_go_path = os.path.join(folder, "constants.go")
insns_path = os.path.join(folder, "insns.txt")
emitter_go_path = os.path.join(folder, "emitter.go")

FORMAT_TYPES = {
    "i": ["Opcode", "Register", "Register", "Immediate"],
    "j": ["Opcode", "Target"],
    "r": ["Opcode", "Register", "Register", "Register", "ShiftAmt", "Funct"]
}
FORMAT_STRUCTS = {
    "i": "InsnI",
    "j": "InsnJ",
    "r": "InsnR"
}

# constants
constants = {}
current_type = "Opcode"
constants_file = open(constants_path, "r")

constants_go_file = open(constants_go_path, "w")
constants_go_file.write("package isa\n\n")
constants_go_file.write("const (\n")

for line in constants_file:
    line = line.strip()
    if not line or line[:2] == "##":
        continue
    if line[0] == "#":
        current_type = line[1:].strip()
        continue
    name, val = line.split()
    constants[name] = val
    constants_go_file.write("\t%s %s = %s\n" % (name, current_type, val))
constants_go_file.write(")\n")
constants_go_file.close()
constants_file.close()

# read insns
insns_file = open(insns_path, "r")
current_insn = None

emitter_go_file = open(emitter_go_path, "w")
emitter_go_file.write("package isa\n\n")

for line in insns_file:
    line = line.strip()
    if not line or line[0] == '#':
        continue
    if current_insn is None:
        i = line.find(" ")
        if i > -1:
            opcode = line[:i]
            operands = line[i+1:].split(",")
            operands = [o for op in operands for o in parse_operand(op)]
        else:
            opcode = line
            operands = []
        current_insn = (opcode, operands)
    else:
        opcode, operands = current_insn

        assert line[0] in FORMAT_TYPES
        assert line[1] == ":"

        fmt = line[0]
        variables = line[2:].strip().split()
        assert len(variables) == len(FORMAT_TYPES[fmt])

        # figure out the types of each operand according to the insn fmt
        operand_types = {}
        for v, ty in zip(variables, FORMAT_TYPES[fmt]):
            if v in constants:
                continue
            if v.isdigit():
                continue
            if v not in operands:
                print "unknown variable %s for %s %s" % (v, opcode, operands)
            operand_types[v] = ty
        
        fn_name = "emit_%s_%s" % (opcode.lower(), len(operands))
        arg_list = ", ".join("%s %s" % (o, operand_types[o]) for o in operands)
        signature = "func %s(%s) Instruction" % (fn_name, arg_list)

        insn_struct = FORMAT_STRUCTS[fmt]
        constructor_args = ", ".join(variables)
        constructor = "%s{%s}" % (insn_struct, constructor_args)

        emitter_go_file.write("%s {\n" % signature)
        emitter_go_file.write("\treturn %s\n" % constructor)
        emitter_go_file.write("}\n\n")

        current_insn = None
emitter_go_file.close()
insns_file.close()
