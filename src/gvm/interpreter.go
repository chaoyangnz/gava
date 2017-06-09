package gvm

const MAIN_METHOD = "main([Ljava/lang/String;)V"

type Instruction struct {
	mnemonic        string
	interpreter     func(uint8, *StackFrame, *Thread, *JavaClass, *JavaMethod)
	length          int
	switchFrame     bool
}

var instructions = []Instruction{
	/* ----- CONSTANTS -----------*/
	/*00 (0x00)*/ Instruction{"nop", NOP, 1, false},
	/*01 (0x01)*/ Instruction{"aconst_null", ACONST_NULL, 1, false},
	/*02 (0x02)*/ Instruction{"iconst_m1", ICONST_M1, 1, false},
	/*03 (0x03)*/ Instruction{"iconst_0", ICONST_0, 1, false},
	/*04 (0x04)*/ Instruction{"iconst_1", ICONST_1, 1, false},
	/*05 (0x05)*/ Instruction{"iconst_2", ICONST_2, 1, false},
	/*06 (0x06)*/ Instruction{"iconst_3", ICONST_3, 1, false},
	/*07 (0x07)*/ Instruction{"iconst_4", ICONST_4, 1, false},
	/*08 (0x08)*/ Instruction{"iconst_5", ICONST_5, 1, false},
	/*09 (0x09)*/ Instruction{"lconst_0", LCONST_0, 1, false},
	/*10 (0x0A)*/ Instruction{"lconst_1", LCONST_1, 1, false},
	/*11 (0x0B)*/ Instruction{"fconst_0", FCONST_0, 1, false},
	/*12 (0x0C)*/ Instruction{"fconst_1", FCONST_1, 1, false},
	/*13 (0x0D)*/ Instruction{"fconst_2", FCONST_2, 1, false},
	/*14 (0x0E)*/ Instruction{"dconst_0", DCONST_0, 1, false},
	/*15 (0x0F)*/ Instruction{"dconst_1", DCONST_1, 1, false},
	/*16 (0x10)*/ Instruction{"bipush", BIPUSH, 2, false},
	/*17 (0x11)*/ Instruction{"sipush", SIPUSH, 3, false},
	/*18 (0x12)*/ Instruction{"ldc", LDC, 2, false},
	/*19 (0x13)*/ Instruction{"ldc_w", LDC_W, 3, false},
	/*20 (0x14)*/ Instruction{"ldc2_w", LDC2_W, 3, false},
	/*--------LOADS ----------------*/
	/*21 (0x15)*/ Instruction{"iload", ILOAD, 2, false},
	/*22 (0x16)*/ Instruction{"lload", LLOAD, 2, false},
	/*23 (0x17)*/ Instruction{"fload", FLOAD, 2, false},
	/*24 (0x18)*/ Instruction{"dload", DLOAD, 2, false},
	/*25 (0x19)*/ Instruction{"aload", ALOAD, 2, false},
	/*26 (0x1A)*/ Instruction{"iload_0", ILOAD_0, 1, false},
	/*27 (0x1B)*/ Instruction{"iload_1", ILOAD_1, 1, false},
	/*28 (0x1C)*/ Instruction{"iload_2", ILOAD_2, 1, false},
	/*29 (0x1D)*/ Instruction{"iload_3", ILOAD_3, 1, false},
	/*30 (0x1E)*/ Instruction{"lload_0", LLOAD_0, 1, false},
	/*31 (0x1F)*/ Instruction{"lload_1", LLOAD_1, 1, false},
	/*32 (0x20)*/ Instruction{"lload_2", LLOAD_2, 1, false},
	/*33 (0x21)*/ Instruction{"lload_3", LLOAD_3, 1, false},
	/*34 (0x22)*/ Instruction{"fload_0", FLOAD_0, 1, false},
	/*35 (0x23)*/ Instruction{"fload_1", FLOAD_1, 1, false},
	/*36 (0x24)*/ Instruction{"fload_2", FLOAD_2, 1, false},
	/*37 (0x25)*/ Instruction{"fload_3", FLOAD_3, 1, false},
	/*38 (0x26)*/ Instruction{"dload_0", DLOAD_0, 1, false},
	/*39 (0x27)*/ Instruction{"dload_1", DLOAD_1, 1, false},
	/*40 (0x28)*/ Instruction{"dload_2", DLOAD_2, 1, false},
	/*41 (0x29)*/ Instruction{"dload_3", DLOAD_3, 1, false},
	/*42 (0x2A)*/ Instruction{"aload_0", ALOAD_0, 1, false},
	/*43 (0x2B)*/ Instruction{"aload_1", ALOAD_1, 1, false},
	/*44 (0x2C)*/ Instruction{"aload_2", ALOAD_2, 1, false},
	/*45 (0x2D)*/ Instruction{"aload_3", ALOAD_3, 1, false},
	/*46 (0x2E)*/ Instruction{"iaload", IALOAD, 1, false},
	/*47 (0x2F)*/ Instruction{"laload", LALOAD, 1, false},
	/*48 (0x30)*/ Instruction{"faload", FALOAD, 1, false},
	/*49 (0x31)*/ Instruction{"daload", DALOAD, 1, false},
	/*50 (0x32)*/ Instruction{"aaload", AALOAD, 1, false},
	/*51 (0x33)*/ Instruction{"baload", BALOAD, 1, false},
	/*52 (0x34)*/ Instruction{"caload", CALOAD, 1, false},
	/*53 (0x35)*/ Instruction{"saload", SALOAD, 1, false},
	/*--------STORES ----------------*/
	/*54 (0x36)*/ Instruction{"istore", ISTORE, 2, false},
	/*55 (0x37)*/ Instruction{"lstore", LSTORE, 2, false},
	/*56 (0x38)*/ Instruction{"fstore", FSTORE, 2, false},
	/*57 (0x39)*/ Instruction{"dstore", DSTORE, 2, false},
	/*58 (0x3A)*/ Instruction{"astore", ASTORE, 2, false},
	/*59 (0x3B)*/ Instruction{"istore_0", ISTORE_0, 1, false},
	/*60 (0x3C)*/ Instruction{"istore_1", ISTORE_1, 1, false},
	/*61 (0x3D)*/ Instruction{"istore_2", ISTORE_2, 1, false},
	/*62 (0x3E)*/ Instruction{"istore_3", ISTORE_3, 1, false},
	/*63 (0x3F)*/ Instruction{"lstore_0", LSTORE_0, 1, false},
	/*64 (0x40)*/ Instruction{"lstore_1", LSTORE_1, 1, false},
	/*65 (0x41)*/ Instruction{"lstore_2", LSTORE_2, 1, false},
	/*66 (0x42)*/ Instruction{"lstore_3", LSTORE_3, 1, false},
	/*67 (0x43)*/ Instruction{"fstore_0", FSTORE_0, 1, false},
	/*68 (0x44)*/ Instruction{"fstore_1", FSTORE_1, 1, false},
	/*69 (0x45)*/ Instruction{"fstore_2", FSTORE_2, 1, false},
	/*70 (0x46)*/ Instruction{"fstore_3", FSTORE_3, 1, false},
	/*71 (0x47)*/ Instruction{"dstore_0", DSTORE_0, 1, false},
	/*72 (0x48)*/ Instruction{"dstore_1", DSTORE_1, 1, false},
	/*73 (0x49)*/ Instruction{"dstore_2", DSTORE_2, 1, false},
	/*74 (0x4A)*/ Instruction{"dstore_3", DSTORE_3, 1, false},
	/*75 (0x4B)*/ Instruction{"astore_0", ASTORE_0, 1, false},
	/*76 (0x4C)*/ Instruction{"astore_1", ASTORE_1, 1, false},
	/*77 (0x4D)*/ Instruction{"astore_2", ASTORE_2, 1, false},
	/*78 (0x4E)*/ Instruction{"astore_3", ASTORE_3, 1, false},
	/*79 (0x4F)*/ Instruction{"iastore", IASTORE, 1, false},
	/*80 (0x50)*/ Instruction{"lastore", LASTORE, 1, false},
	/*81 (0x51)*/ Instruction{"fastore", FASTORE, 1, false},
	/*82 (0x52)*/ Instruction{"dastore", DASTORE, 1, false},
	/*83 (0x53)*/ Instruction{"aastore", AASTORE, 1, false},
	/*84 (0x54)*/ Instruction{"bastore", BASTORE, 1, false},
	/*85 (0x55)*/ Instruction{"castore", CASTORE, 1, false},
	/*86 (0x56)*/ Instruction{"sastore", SASTORE, 1, false},
	/*--------STACK---------------*/
	/*87 (0x57)*/ Instruction{"pop", POP, 1, false},
	/*88 (0x58)*/ Instruction{"pop2", POP2, 1, false},
	/*89 (0x59)*/ Instruction{"dup", DUP, 1, false},
	/*90 (0x5A)*/ Instruction{"dup_x1", DUP_X1, 1, false},
	/*91 (0x5B)*/ Instruction{"dup_x2", DUP_X2, 1, false},
	/*92 (0x5C)*/ Instruction{"dup2", DUP2, 1, false},
	/*93 (0x5D)*/ Instruction{"dup2_x1", DUP2_X1, 1, false},
	/*94 (0x5E)*/ Instruction{"dup2_x2", DUP2_X2, 1, false},
	/*95 (0x5F)*/ Instruction{"swap", SWAP, 1, false},
	/*---------MATH -------------*/
	/*96 (0x60)*/  Instruction{"iadd", IADD, 1, false},
	/*97 (0x61)*/  Instruction{"ladd", LADD, 1, false},
	/*98 (0x62)*/  Instruction{"fadd", FADD, 1, false},
	/*99 (0x63)*/  Instruction{"dadd", DADD, 1, false},
	/*100 (0x64)*/ Instruction{"isub", ISUB, 1, false},
	/*101 (0x65)*/ Instruction{"lsub", LSUB, 1, false},
	/*102 (0x66)*/ Instruction{"fsub", FSUB, 1, false},
	/*103 (0x67)*/ Instruction{"dsub", DSUB, 1, false},
	/*104 (0x68)*/ Instruction{"imul", IMUL, 1, false},
	/*105 (0x69)*/ Instruction{"lmul", LMUL, 1, false},
	/*106 (0x6A)*/ Instruction{"fmul", FMUL, 1, false},
	/*107 (0x6B)*/ Instruction{"dmul", DMUL, 1, false},
	/*108 (0x6C)*/ Instruction{"idiv", IDIV, 1, false},
	/*109 (0x6D)*/ Instruction{"ldiv", LDIV, 1, false},
	/*110 (0x6E)*/ Instruction{"fdiv", FDIV, 1, false},
	/*111 (0x6F)*/ Instruction{"ddiv", DDIV, 1, false},
	/*112 (0x70)*/ Instruction{"irem", IREM, 1, false},
	/*113 (0x71)*/ Instruction{"lrem", LREM, 1, false},
	/*114 (0x72)*/ Instruction{"frem", FREM, 1, false},
	/*115 (0x73)*/ Instruction{"drem", DREM, 1, false},
	/*116 (0x74)*/ Instruction{"ineg", INEG, 1, false},
	/*117 (0x75)*/ Instruction{"lneg", LNEG, 1, false},
	/*118 (0x76)*/ Instruction{"fneg", FNEG, 1, false},
	/*119 (0x77)*/ Instruction{"dneg", DNEG, 1, false},
	/*120 (0x78)*/ Instruction{"ishl", ISHL, 1, false},
	/*121 (0x79)*/ Instruction{"lshl", LSHL, 1, false},
	/*122 (0x7A)*/ Instruction{"ishr", ISHR, 1, false},
	/*123 (0x7B)*/ Instruction{"lshr", LSHR, 1, false},
	/*124 (0x7C)*/ Instruction{"iushr", IUSHR, 1, false},
	/*125 (0x7D)*/ Instruction{"lushr", LUSHR, 1, false},
	/*126 (0x7E)*/ Instruction{"iand", IAND, 1, false},
	/*127 (0x7F)*/ Instruction{"land", LAND, 1, false},
	/*128 (0x80)*/ Instruction{"ior", IOR, 1, false},
	/*129 (0x81)*/ Instruction{"lor", LOR, 1, false},
	/*130 (0x82)*/ Instruction{"ixor", IXOR, 1, false},
	/*131 (0x83)*/ Instruction{"lxor", LXOR, 1, false},
	/*132 (0x84)*/ Instruction{"iinc", IINC, 3, false},
	/*--------CONVERSIONS-----------*/
	/*133 (0x85)*/ Instruction{"i2l", I2L, 1, false},
	/*134 (0x86)*/ Instruction{"i2f", I2F, 1, false},
	/*135 (0x87)*/ Instruction{"i2d", I2D, 1, false},
	/*136 (0x88)*/ Instruction{"l2i", L2I, 1, false},
	/*137 (0x89)*/ Instruction{"l2f", L2F, 1, false},
	/*138 (0x8A)*/ Instruction{"l2d", L2D, 1, false},
	/*139 (0x8B)*/ Instruction{"f2i", F2I, 1, false},
	/*140 (0x8C)*/ Instruction{"f2l", F2L, 1, false},
	/*141 (0x8D)*/ Instruction{"f2d", F2D, 1, false},
	/*142 (0x8E)*/ Instruction{"d2i", D2I, 1, false},
	/*143 (0x8F)*/ Instruction{"d2l", D2L, 1, false},
	/*144 (0x90)*/ Instruction{"d2f", D2F, 1, false},
	/*145 (0x91)*/ Instruction{"i2b", I2B, 1, false},
	/*146 (0x92)*/ Instruction{"i2c", I2C, 1, false},
	/*147 (0x93)*/ Instruction{"i2s", I2S, 1, false},
	/*-----------COMPARASON -----------*/
	/*148 (0x94)*/ Instruction{"lcmp", LCMP, 1, false},
	/*149 (0x95)*/ Instruction{"fcmpl", FCMPL, 1, false},
	/*150 (0x96)*/ Instruction{"fcmpg", FCMPG, 1, false},
	/*151 (0x97)*/ Instruction{"dcmpl", DCMPL, 1, false},
	/*152 (0x98)*/ Instruction{"dcmpg", DCMPG, 1, false},
	/*153 (0x99)*/ Instruction{"ifeq", IFEQ, 3, false},
	/*154 (0x9A)*/ Instruction{"ifne", IFNE, 3, false},
	/*155 (0x9B)*/ Instruction{"iflt", IFLT, 3, false},
	/*156 (0x9C)*/ Instruction{"ifge", IFGE, 3, false},
	/*157 (0x9D)*/ Instruction{"ifgt", IFGT, 3, false},
	/*158 (0x9E)*/ Instruction{"ifle", IFLE, 3, false},
	/*159 (0x9F)*/ Instruction{"if_icmpeq", IF_ICMPEQ, 3, false},
	/*160 (0xA0)*/ Instruction{"if_icmpne", IF_ICMPNE, 3, false},
	/*161 (0xA1)*/ Instruction{"if_icmplt", IF_ICMPLT, 3, false},
	/*162 (0xA2)*/ Instruction{"if_icmpge", IF_ICMPGE, 3, false},
	/*163 (0xA3)*/ Instruction{"if_icmpgt", IF_ICMPGT, 3, false},
	/*164 (0xA4)*/ Instruction{"if_icmple", IF_ICMPLE, 3, false},
	/*165 (0xA5)*/ Instruction{"if_acmpeq", IF_ACMPEQ, 3, false},
	/*166 (0xA6)*/ Instruction{"if_acmpne", IF_ACMPNE, 3, false},
	/*---------REFERENCES -------------*/
	/*167 (0xA7)*/ Instruction{"goto", GOTO, 3, false},
	/*168 (0xA8)*/ Instruction{"jsr", JSR, 3, false},
	/*169 (0xA9)*/ Instruction{"ret", RET, 2, true},
	/*170 (0xAA)*/ Instruction{"tableswitch", TABLESWITCH, -1, false}, // variable bytecode length
	/*171 (0xAB)*/ Instruction{"lookupswitch", LOOKUPSWITCH, -1, false},
	/*172 (0xAC)*/ Instruction{"ireturn", IRETURN, 1, true},
	/*173 (0xAD)*/ Instruction{"lreturn", LRETURN, 1, true},
	/*174 (0xAE)*/ Instruction{"freturn", FRETURN, 1, true},
	/*175 (0xAF)*/ Instruction{"dreturn", DRETURN, 1, true},
	/*176 (0xB0)*/ Instruction{"areturn", ARETURN, 1, true},
	/*177 (0xB1)*/ Instruction{"return", RETURN, 1, true},
	/*-------CONTROLS------------------*/
	/*178 (0xB2)*/ Instruction{"getstatic", GETSTATIC, 3, false},
	/*179 (0xB3)*/ Instruction{"putstatic", PUTSTATIC, 3, false},
	/*180 (0xB4)*/ Instruction{"getfield", GETFIELD, 3, false},
	/*181 (0xB5)*/ Instruction{"putfield", PUTFIELD, 3, false},
	/*182 (0xB6)*/ Instruction{"invokevirtual", INVOKEVIRTUAL, 3, true},
	/*183 (0xB7)*/ Instruction{"invokespecial", INVOKESPECIAL, 3, true},
	/*184 (0xB8)*/ Instruction{"invokestatic", INVOKESTATIC, 3, true},
	/*185 (0xB9)*/ Instruction{"invokeinterface", INVOKEINTERFACE, 5, true},
	/*186 (0xBA)*/ Instruction{"invokedynamic", INVOKEDYNAMIC, 5, true},
	/*187 (0xBB)*/ Instruction{"new", NEW, 3, false},
	/*188 (0xBC)*/ Instruction{"newarray", NEWARRAY, 2, false},
	/*189 (0xBD)*/ Instruction{"anewarray", ANEWARRAY, 3, false},
	/*190 (0xBE)*/ Instruction{"arraylength", ARRAYLENGTH, 1, false},
	/*191 (0xBF)*/ Instruction{"athrow", ATHROW, 1, false},
	/*192 (0xC0)*/ Instruction{"checkcast", CHECKCAST, 3, false},
	/*193 (0xC1)*/ Instruction{"instanceof", INSTANCEOF, 3, false},
	/*194 (0xC2)*/ Instruction{"monitorenter", MONITORENTER, 1, false},
	/*195 (0xC3)*/ Instruction{"monitorexit", MONITOREXIT, 1, false},
	/*--------EXTENDED-----------------*/
	/*196 (0xC4)*/ Instruction{"wide", WIDE, -1, false},
	/*197 (0xC5)*/ Instruction{"multianewarray", MULTIANEWARRAY, 4, false},
	/*198 (0xC6)*/ Instruction{"ifnull", IFNULL, 3, false},
	/*199 (0xC7)*/ Instruction{"ifnonnull", IFNONNULL, 3, false},
	/*200 (0xC8)*/ Instruction{"goto_w", GOTO_W, 5, false},
	/*201 (0xC9)*/ Instruction{"jsr_w", JSR_W, 5, false},
	/*----------RESERVED ---------------*/
	///*202 (0xCA)*/ Instruction{"breakpoint", BREAKPOINT, 1, false},
	///*254 (0xFE)*/    Instruction{"impdep1", IMPDEP1, 1, false},
	///*255 (0xFF)*/    Instruction{"impdep2", IMPDEP2, 1, false},
}

func run(mainClass *JavaClass)  {
	mainMethod := mainClass.findMethod(MAIN_METHOD)
	thread := newThread("main")
	thread.pushFrame(NewStackFrame(mainMethod))
	for len(thread.vmStack) != 0 { // per stack frame
		f := thread.peekFrame()
		bytecode := f.method.code
		trace("⤮ %s.%s%s \n", f.method.class.thisClassName, f.method.name, f.method.descriptor)
		for f.pc < uint32(len(f.method.code)) {
			pc := f.pc
			opcode := bytecode[pc]
			instruction := instructions[opcode]
			trace("%04d ➢ %s \n", int(pc), instruction.mnemonic)
			instruction.interpreter(opcode, f, thread, f.method.class, f.method)
			// jump instruction can operate pc
			// some instruction also have variable length: tableswitch...
			// these intructions will control pc themselves
			if f.pc == pc && instruction.length > 0 {
				f.pc += uint32(instruction.length)
			}
			if instruction.switchFrame {
				break
			}
		}
	}
}



