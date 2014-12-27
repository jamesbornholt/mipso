package main

type Opcode int
type Register int
type Immediate int
type Target int
type ShiftAmt int
type Funct int

type Instruction interface {
	isInstruction()
}

type InsnI struct {
	op     Opcode
	rs, rt Register
	imm    Immediate
}

func (i InsnI) isInstruction() {}

type InsnJ struct {
	op  Opcode
	tgt Target
}

func (i InsnJ) isInstruction() {}

type InsnR struct {
	op         Opcode
	rs, rt, rd Register
	shamt      ShiftAmt
	funct      Funct
}

func (i InsnR) isInstruction() {}
