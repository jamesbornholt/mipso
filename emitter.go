package main

func emit_add_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, ADD}
}

func emit_addi_3(rt Register, rs Register, imm Immediate) Instruction {
	return InsnI{ADDI, rs, rt, imm}
}

func emit_addiu_3(rt Register, rs Register, imm Immediate) Instruction {
	return InsnI{ADDIU, rs, rt, imm}
}

func emit_addu_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, ADDU}
}

func emit_and_3(rs Register, rt Register, rd Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, AND}
}

func emit_andi_3(rt Register, rs Register, imm Immediate) Instruction {
	return InsnI{ANDI, rs, rt, imm}
}

func emit_beq_3(rs Register, rt Register, off Immediate) Instruction {
	return InsnI{BEQ, rs, rt, off}
}

func emit_bgez_2(rs Register, off Immediate) Instruction {
	return InsnI{REGIMM, rs, BGEZ, off}
}

func emit_bgezal_2(rs Register, off Immediate) Instruction {
	return InsnI{REGIMM, rs, BGEZAL, off}
}

func emit_bgtz_2(rs Register, off Immediate) Instruction {
	return InsnI{BGTZ, rs, 0, off}
}

func emit_blez_2(rs Register, off Immediate) Instruction {
	return InsnI{BLEZ, rs, 0, off}
}

func emit_bltz_2(rs Register, off Immediate) Instruction {
	return InsnI{REGIMM, rs, BLTZ, off}
}

func emit_bltzal_2(rs Register, off Immediate) Instruction {
	return InsnI{REGIMM, rs, BLTZAL, off}
}

func emit_bne_3(rs Register, rt Register, off Immediate) Instruction {
	return InsnI{BNE, rs, rt, off}
}

func emit_break_0() Instruction {
	return InsnR{SPECIAL, 0, 0, 0, 0, BREAK}
}

func emit_cfc0_2(rt Register, rd Register) Instruction {
	return InsnR{COP0, CF, rt, rd, 0, 0}
}

func emit_cfc1_2(rt Register, rd Register) Instruction {
	return InsnR{COP1, CF, rt, rd, 0, 0}
}

func emit_cfc2_2(rt Register, rd Register) Instruction {
	return InsnR{COP2, CF, rt, rd, 0, 0}
}

func emit_cop0_1(cofun Target) Instruction {
	return InsnJ{COP0, cofun}
}

func emit_cop1_1(cofun Target) Instruction {
	return InsnJ{COP1, cofun}
}

func emit_cop2_1(cofun Target) Instruction {
	return InsnJ{COP2, cofun}
}

func emit_ctc0_2(rt Register, rd Register) Instruction {
	return InsnR{COP0, CT, rt, rd, 0, 0}
}

func emit_ctc1_2(rt Register, rd Register) Instruction {
	return InsnR{COP1, CT, rt, rd, 0, 0}
}

func emit_ctc2_2(rt Register, rd Register) Instruction {
	return InsnR{COP2, CT, rt, rd, 0, 0}
}

func emit_div_2(rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, 0, 0, DIV}
}

func emit_divu_2(rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, 0, 0, DIVU}
}

func emit_j_1(tgt Target) Instruction {
	return InsnJ{J, tgt}
}

func emit_tal_1(tgt Target) Instruction {
	return InsnJ{JAL, tgt}
}

func emit_jalr_1(rs Register) Instruction {
	return InsnR{SPECIAL, rs, 0, 31, 0, JALR}
}

func emit_jalr_2(rd Register, rs Register) Instruction {
	return InsnR{SPECIAL, rs, 0, rd, 0, JALR}
}

func emit_jr_1(rs Register) Instruction {
	return InsnR{SPECIAL, rs, 0, 0, 0, JR}
}

func emit_lb_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LB, base, rt, off}
}

func emit_lbu_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LBU, base, rt, off}
}

func emit_lh_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LH, base, rt, off}
}

func emit_lhu_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LHU, base, rt, off}
}

func emit_lui_2(rt Register, imm Immediate) Instruction {
	return InsnI{LUI, 0, rt, imm}
}

func emit_lw_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LW, base, rt, off}
}

func emit_lwc1_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LWC1, base, rt, off}
}

func emit_lwc2_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LWC2, base, rt, off}
}

func emit_lwc3_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LWC3, base, rt, off}
}

func emit_lwl_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LWL, base, rt, off}
}

func emit_lwr_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{LWR, base, rt, off}
}

func emit_mfc0_2(rt Register, rd Register) Instruction {
	return InsnR{COP0, MF, rt, rd, 0, 0}
}

func emit_mfc1_2(rt Register, rd Register) Instruction {
	return InsnR{COP1, MF, rt, rd, 0, 0}
}

func emit_mfc2_2(rt Register, rd Register) Instruction {
	return InsnR{COP2, MF, rt, rd, 0, 0}
}

func emit_mfhi_1(rd Register) Instruction {
	return InsnR{SPECIAL, 0, 0, rd, 0, MFHI}
}

func emit_mflo_1(rd Register) Instruction {
	return InsnR{SPECIAL, 0, 0, rd, 0, MFLO}
}

func emit_mtc0_2(rt Register, rd Register) Instruction {
	return InsnR{COP0, MT, rt, rd, 0, 0}
}

func emit_mtc1_2(rt Register, rd Register) Instruction {
	return InsnR{COP1, MT, rt, rd, 0, 0}
}

func emit_mtc2_2(rt Register, rd Register) Instruction {
	return InsnR{COP2, MT, rt, rd, 0, 0}
}

func emit_mthi_1(rs Register) Instruction {
	return InsnR{SPECIAL, rs, 0, 0, 0, MTHI}
}

func emit_mtlo_1(rs Register) Instruction {
	return InsnR{SPECIAL, rs, 0, 0, 0, MTLO}
}

func emit_mult_2(rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, 0, 0, MULT}
}

func emit_multu_2(rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, 0, 0, MULTU}
}

func emit_nor_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, NOR}
}

func emit_or_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, OR}
}

func emit_ori_3(rt Register, rs Register, imm Immediate) Instruction {
	return InsnI{ORI, rs, rt, imm}
}

func emit_sb_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{SB, base, rt, off}
}

func emit_sh_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{SH, base, rt, off}
}

func emit_sll_3(rd Register, rt Register, sa ShiftAmt) Instruction {
	return InsnR{SPECIAL, 0, rt, rd, sa, SLL}
}

func emit_sllv_3(rd Register, rt Register, rs Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, SLLV}
}

func emit_slt_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, SLT}
}

func emit_slti_3(rt Register, rs Register, imm Immediate) Instruction {
	return InsnI{SLTI, rs, rt, imm}
}

func emit_sltiu_3(rt Register, rs Register, imm Immediate) Instruction {
	return InsnI{SLTIU, rs, rt, imm}
}

func emit_sltu_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, SLTU}
}

func emit_sra_3(rd Register, rt Register, sa ShiftAmt) Instruction {
	return InsnR{SPECIAL, 0, rt, rd, sa, SRA}
}

func emit_srav_3(rd Register, rt Register, rs Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, SRAV}
}

func emit_srl_3(rd Register, rt Register, sa ShiftAmt) Instruction {
	return InsnR{SPECIAL, 0, rt, rd, sa, SRL}
}

func emit_srlv_3(rd Register, rt Register, rs Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, SRLV}
}

func emit_sub_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, SUB}
}

func emit_subu_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, SUBU}
}

func emit_sw_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{SW, base, rt, off}
}

func emit_swc1_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{SWC1, base, rt, off}
}

func emit_swc2_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{SWC2, base, rt, off}
}

func emit_swc3_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{SWC3, base, rt, off}
}

func emit_swl_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{SWL, base, rt, off}
}

func emit_swr_3(rt Register, off Immediate, base Register) Instruction {
	return InsnI{SWR, base, rt, off}
}

func emit_syscall_0() Instruction {
	return InsnR{SPECIAL, 0, 0, 0, 0, SYSCALL}
}

func emit_xor_3(rd Register, rs Register, rt Register) Instruction {
	return InsnR{SPECIAL, rs, rt, rd, 0, XOR}
}

func emit_xori_3(rt Register, rs Register, imm Immediate) Instruction {
	return InsnI{XORI, rs, rt, imm}
}
