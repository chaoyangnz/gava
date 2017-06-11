package gvm

const MAIN_METHOD = "main([Ljava/lang/String;)V"

type Instruction struct {
	mnemonic        string
	interpreter     func(uint8, *StackFrame, *Thread, *ClassType, *Method)
	switchFrame     bool
}

var instructions = [JVM_OPC_MAX+1]Instruction{
	/* ----- CONSTANTS -----------*/
	/*00 (0x00)*/ Instruction{"nop", NOP, false},
	/*01 (0x01)*/ Instruction{"aconst_null", ACONST_NULL, false},
	/*02 (0x02)*/ Instruction{"iconst_m1", ICONST_M1, false},
	/*03 (0x03)*/ Instruction{"iconst_0", ICONST_0, false},
	/*04 (0x04)*/ Instruction{"iconst_1", ICONST_1, false},
	/*05 (0x05)*/ Instruction{"iconst_2", ICONST_2, false},
	/*06 (0x06)*/ Instruction{"iconst_3", ICONST_3, false},
	/*07 (0x07)*/ Instruction{"iconst_4", ICONST_4, false},
	/*08 (0x08)*/ Instruction{"iconst_5", ICONST_5, false},
	/*09 (0x09)*/ Instruction{"lconst_0", LCONST_0, false},
	/*10 (0x0A)*/ Instruction{"lconst_1", LCONST_1, false},
	/*11 (0x0B)*/ Instruction{"fconst_0", FCONST_0, false},
	/*12 (0x0C)*/ Instruction{"fconst_1", FCONST_1, false},
	/*13 (0x0D)*/ Instruction{"fconst_2", FCONST_2, false},
	/*14 (0x0E)*/ Instruction{"dconst_0", DCONST_0, false},
	/*15 (0x0F)*/ Instruction{"dconst_1", DCONST_1, false},
	/*16 (0x10)*/ Instruction{"bipush", BIPUSH, false},
	/*17 (0x11)*/ Instruction{"sipush", SIPUSH, false},
	/*18 (0x12)*/ Instruction{"ldc", LDC, false},
	/*19 (0x13)*/ Instruction{"ldc_w", LDC_W, false},
	/*20 (0x14)*/ Instruction{"ldc2_w", LDC2_W, false},
	/*--------LOADS ----------------*/
	/*21 (0x15)*/ Instruction{"iload", ILOAD, false},
	/*22 (0x16)*/ Instruction{"lload", LLOAD, false},
	/*23 (0x17)*/ Instruction{"fload", FLOAD, false},
	/*24 (0x18)*/ Instruction{"dload", DLOAD, false},
	/*25 (0x19)*/ Instruction{"aload", ALOAD, false},
	/*26 (0x1A)*/ Instruction{"iload_0", ILOAD_0, false},
	/*27 (0x1B)*/ Instruction{"iload_1", ILOAD_1, false},
	/*28 (0x1C)*/ Instruction{"iload_2", ILOAD_2, false},
	/*29 (0x1D)*/ Instruction{"iload_3", ILOAD_3, false},
	/*30 (0x1E)*/ Instruction{"lload_0", LLOAD_0, false},
	/*31 (0x1F)*/ Instruction{"lload_1", LLOAD_1, false},
	/*32 (0x20)*/ Instruction{"lload_2", LLOAD_2, false},
	/*33 (0x21)*/ Instruction{"lload_3", LLOAD_3, false},
	/*34 (0x22)*/ Instruction{"fload_0", FLOAD_0, false},
	/*35 (0x23)*/ Instruction{"fload_1", FLOAD_1, false},
	/*36 (0x24)*/ Instruction{"fload_2", FLOAD_2, false},
	/*37 (0x25)*/ Instruction{"fload_3", FLOAD_3, false},
	/*38 (0x26)*/ Instruction{"dload_0", DLOAD_0, false},
	/*39 (0x27)*/ Instruction{"dload_1", DLOAD_1, false},
	/*40 (0x28)*/ Instruction{"dload_2", DLOAD_2, false},
	/*41 (0x29)*/ Instruction{"dload_3", DLOAD_3, false},
	/*42 (0x2A)*/ Instruction{"aload_0", ALOAD_0, false},
	/*43 (0x2B)*/ Instruction{"aload_1", ALOAD_1, false},
	/*44 (0x2C)*/ Instruction{"aload_2", ALOAD_2, false},
	/*45 (0x2D)*/ Instruction{"aload_3", ALOAD_3, false},
	/*46 (0x2E)*/ Instruction{"iaload", IALOAD, false},
	/*47 (0x2F)*/ Instruction{"laload", LALOAD, false},
	/*48 (0x30)*/ Instruction{"faload", FALOAD, false},
	/*49 (0x31)*/ Instruction{"daload", DALOAD, false},
	/*50 (0x32)*/ Instruction{"aaload", AALOAD, false},
	/*51 (0x33)*/ Instruction{"baload", BALOAD, false},
	/*52 (0x34)*/ Instruction{"caload", CALOAD, false},
	/*53 (0x35)*/ Instruction{"saload", SALOAD, false},
	/*--------STORES ----------------*/
	/*54 (0x36)*/ Instruction{"istore", ISTORE, false},
	/*55 (0x37)*/ Instruction{"lstore", LSTORE, false},
	/*56 (0x38)*/ Instruction{"fstore", FSTORE, false},
	/*57 (0x39)*/ Instruction{"dstore", DSTORE, false},
	/*58 (0x3A)*/ Instruction{"astore", ASTORE, false},
	/*59 (0x3B)*/ Instruction{"istore_0", ISTORE_0, false},
	/*60 (0x3C)*/ Instruction{"istore_1", ISTORE_1, false},
	/*61 (0x3D)*/ Instruction{"istore_2", ISTORE_2, false},
	/*62 (0x3E)*/ Instruction{"istore_3", ISTORE_3, false},
	/*63 (0x3F)*/ Instruction{"lstore_0", LSTORE_0, false},
	/*64 (0x40)*/ Instruction{"lstore_1", LSTORE_1, false},
	/*65 (0x41)*/ Instruction{"lstore_2", LSTORE_2, false},
	/*66 (0x42)*/ Instruction{"lstore_3", LSTORE_3, false},
	/*67 (0x43)*/ Instruction{"fstore_0", FSTORE_0, false},
	/*68 (0x44)*/ Instruction{"fstore_1", FSTORE_1, false},
	/*69 (0x45)*/ Instruction{"fstore_2", FSTORE_2, false},
	/*70 (0x46)*/ Instruction{"fstore_3", FSTORE_3, false},
	/*71 (0x47)*/ Instruction{"dstore_0", DSTORE_0, false},
	/*72 (0x48)*/ Instruction{"dstore_1", DSTORE_1, false},
	/*73 (0x49)*/ Instruction{"dstore_2", DSTORE_2, false},
	/*74 (0x4A)*/ Instruction{"dstore_3", DSTORE_3, false},
	/*75 (0x4B)*/ Instruction{"astore_0", ASTORE_0, false},
	/*76 (0x4C)*/ Instruction{"astore_1", ASTORE_1, false},
	/*77 (0x4D)*/ Instruction{"astore_2", ASTORE_2, false},
	/*78 (0x4E)*/ Instruction{"astore_3", ASTORE_3, false},
	/*79 (0x4F)*/ Instruction{"iastore", IASTORE, false},
	/*80 (0x50)*/ Instruction{"lastore", LASTORE, false},
	/*81 (0x51)*/ Instruction{"fastore", FASTORE, false},
	/*82 (0x52)*/ Instruction{"dastore", DASTORE, false},
	/*83 (0x53)*/ Instruction{"aastore", AASTORE, false},
	/*84 (0x54)*/ Instruction{"bastore", BASTORE, false},
	/*85 (0x55)*/ Instruction{"castore", CASTORE, false},
	/*86 (0x56)*/ Instruction{"sastore", SASTORE, false},
	/*--------STACK---------------*/
	/*87 (0x57)*/ Instruction{"pop", POP, false},
	/*88 (0x58)*/ Instruction{"pop2", POP2, false},
	/*89 (0x59)*/ Instruction{"dup", DUP, false},
	/*90 (0x5A)*/ Instruction{"dup_x1", DUP_X1, false},
	/*91 (0x5B)*/ Instruction{"dup_x2", DUP_X2, false},
	/*92 (0x5C)*/ Instruction{"dup2", DUP2, false},
	/*93 (0x5D)*/ Instruction{"dup2_x1", DUP2_X1, false},
	/*94 (0x5E)*/ Instruction{"dup2_x2", DUP2_X2, false},
	/*95 (0x5F)*/ Instruction{"swap", SWAP, false},
	/*---------MATH -------------*/
	/*96 (0x60)*/  Instruction{"iadd", IADD, false},
	/*97 (0x61)*/  Instruction{"ladd", LADD, false},
	/*98 (0x62)*/  Instruction{"fadd", FADD, false},
	/*99 (0x63)*/  Instruction{"dadd", DADD, false},
	/*100 (0x64)*/ Instruction{"isub", ISUB, false},
	/*101 (0x65)*/ Instruction{"lsub", LSUB, false},
	/*102 (0x66)*/ Instruction{"fsub", FSUB, false},
	/*103 (0x67)*/ Instruction{"dsub", DSUB, false},
	/*104 (0x68)*/ Instruction{"imul", IMUL, false},
	/*105 (0x69)*/ Instruction{"lmul", LMUL, false},
	/*106 (0x6A)*/ Instruction{"fmul", FMUL, false},
	/*107 (0x6B)*/ Instruction{"dmul", DMUL, false},
	/*108 (0x6C)*/ Instruction{"idiv", IDIV, false},
	/*109 (0x6D)*/ Instruction{"ldiv", LDIV, false},
	/*110 (0x6E)*/ Instruction{"fdiv", FDIV, false},
	/*111 (0x6F)*/ Instruction{"ddiv", DDIV, false},
	/*112 (0x70)*/ Instruction{"irem", IREM, false},
	/*113 (0x71)*/ Instruction{"lrem", LREM, false},
	/*114 (0x72)*/ Instruction{"frem", FREM, false},
	/*115 (0x73)*/ Instruction{"drem", DREM, false},
	/*116 (0x74)*/ Instruction{"ineg", INEG, false},
	/*117 (0x75)*/ Instruction{"lneg", LNEG, false},
	/*118 (0x76)*/ Instruction{"fneg", FNEG, false},
	/*119 (0x77)*/ Instruction{"dneg", DNEG, false},
	/*120 (0x78)*/ Instruction{"ishl", ISHL, false},
	/*121 (0x79)*/ Instruction{"lshl", LSHL, false},
	/*122 (0x7A)*/ Instruction{"ishr", ISHR, false},
	/*123 (0x7B)*/ Instruction{"lshr", LSHR, false},
	/*124 (0x7C)*/ Instruction{"iushr", IUSHR, false},
	/*125 (0x7D)*/ Instruction{"lushr", LUSHR, false},
	/*126 (0x7E)*/ Instruction{"iand", IAND, false},
	/*127 (0x7F)*/ Instruction{"land", LAND, false},
	/*128 (0x80)*/ Instruction{"ior", IOR, false},
	/*129 (0x81)*/ Instruction{"lor", LOR, false},
	/*130 (0x82)*/ Instruction{"ixor", IXOR, false},
	/*131 (0x83)*/ Instruction{"lxor", LXOR, false},
	/*132 (0x84)*/ Instruction{"iinc", IINC, false},
	/*--------CONVERSIONS-----------*/
	/*133 (0x85)*/ Instruction{"i2l", I2L, false},
	/*134 (0x86)*/ Instruction{"i2f", I2F, false},
	/*135 (0x87)*/ Instruction{"i2d", I2D, false},
	/*136 (0x88)*/ Instruction{"l2i", L2I, false},
	/*137 (0x89)*/ Instruction{"l2f", L2F, false},
	/*138 (0x8A)*/ Instruction{"l2d", L2D, false},
	/*139 (0x8B)*/ Instruction{"f2i", F2I, false},
	/*140 (0x8C)*/ Instruction{"f2l", F2L, false},
	/*141 (0x8D)*/ Instruction{"f2d", F2D, false},
	/*142 (0x8E)*/ Instruction{"d2i", D2I, false},
	/*143 (0x8F)*/ Instruction{"d2l", D2L, false},
	/*144 (0x90)*/ Instruction{"d2f", D2F, false},
	/*145 (0x91)*/ Instruction{"i2b", I2B, false},
	/*146 (0x92)*/ Instruction{"i2c", I2C, false},
	/*147 (0x93)*/ Instruction{"i2s", I2S, false},
	/*-----------COMPARASON -----------*/
	/*148 (0x94)*/ Instruction{"lcmp", LCMP, false},
	/*149 (0x95)*/ Instruction{"fcmpl", FCMPL, false},
	/*150 (0x96)*/ Instruction{"fcmpg", FCMPG, false},
	/*151 (0x97)*/ Instruction{"dcmpl", DCMPL, false},
	/*152 (0x98)*/ Instruction{"dcmpg", DCMPG, false},
	/*153 (0x99)*/ Instruction{"ifeq", IFEQ, false},
	/*154 (0x9A)*/ Instruction{"ifne", IFNE, false},
	/*155 (0x9B)*/ Instruction{"iflt", IFLT, false},
	/*156 (0x9C)*/ Instruction{"ifge", IFGE, false},
	/*157 (0x9D)*/ Instruction{"ifgt", IFGT, false},
	/*158 (0x9E)*/ Instruction{"ifle", IFLE, false},
	/*159 (0x9F)*/ Instruction{"if_icmpeq", IF_ICMPEQ, false},
	/*160 (0xA0)*/ Instruction{"if_icmpne", IF_ICMPNE, false},
	/*161 (0xA1)*/ Instruction{"if_icmplt", IF_ICMPLT, false},
	/*162 (0xA2)*/ Instruction{"if_icmpge", IF_ICMPGE, false},
	/*163 (0xA3)*/ Instruction{"if_icmpgt", IF_ICMPGT, false},
	/*164 (0xA4)*/ Instruction{"if_icmple", IF_ICMPLE, false},
	/*165 (0xA5)*/ Instruction{"if_acmpeq", IF_ACMPEQ, false},
	/*166 (0xA6)*/ Instruction{"if_acmpne", IF_ACMPNE, false},
	/*---------REFERENCES -------------*/
	/*167 (0xA7)*/ Instruction{"goto", GOTO, false},
	/*168 (0xA8)*/ Instruction{"jsr", JSR, false},
	/*169 (0xA9)*/ Instruction{"ret", RET, true},
	/*170 (0xAA)*/ Instruction{"tableswitch", TABLESWITCH, false}, // variable bytecode length
	/*171 (0xAB)*/ Instruction{"lookupswitch", LOOKUPSWITCH, false},
	/*172 (0xAC)*/ Instruction{"ireturn", IRETURN, true},
	/*173 (0xAD)*/ Instruction{"lreturn", LRETURN, true},
	/*174 (0xAE)*/ Instruction{"freturn", FRETURN, true},
	/*175 (0xAF)*/ Instruction{"dreturn", DRETURN, true},
	/*176 (0xB0)*/ Instruction{"areturn", ARETURN, true},
	/*177 (0xB1)*/ Instruction{"return", RETURN, true},
	/*-------CONTROLS------------------*/
	/*178 (0xB2)*/ Instruction{"getstatic", GETSTATIC, false},
	/*179 (0xB3)*/ Instruction{"putstatic", PUTSTATIC, false},
	/*180 (0xB4)*/ Instruction{"getfield", GETFIELD, false},
	/*181 (0xB5)*/ Instruction{"putfield", PUTFIELD, false},
	/*182 (0xB6)*/ Instruction{"invokevirtual", INVOKEVIRTUAL, true},
	/*183 (0xB7)*/ Instruction{"invokespecial", INVOKESPECIAL, true},
	/*184 (0xB8)*/ Instruction{"invokestatic", INVOKESTATIC, true},
	/*185 (0xB9)*/ Instruction{"invokeinterface", INVOKEINTERFACE, true},
	/*186 (0xBA)*/ Instruction{"invokedynamic", INVOKEDYNAMIC, true},
	/*187 (0xBB)*/ Instruction{"newObject", NEW, false},
	/*188 (0xBC)*/ Instruction{"newarray", NEWARRAY, false},
	/*189 (0xBD)*/ Instruction{"anewarray", ANEWARRAY, false},
	/*190 (0xBE)*/ Instruction{"arraylength", ARRAYLENGTH, false},
	/*191 (0xBF)*/ Instruction{"athrow", ATHROW, false},
	/*192 (0xC0)*/ Instruction{"checkcast", CHECKCAST, false},
	/*193 (0xC1)*/ Instruction{"instanceof", INSTANCEOF, false},
	/*194 (0xC2)*/ Instruction{"monitorenter", MONITORENTER, false},
	/*195 (0xC3)*/ Instruction{"monitorexit", MONITOREXIT, false},
	/*--------EXTENDED-----------------*/
	/*196 (0xC4)*/ Instruction{"wide", WIDE, false},
	/*197 (0xC5)*/ Instruction{"multianewarray", MULTIANEWARRAY, false},
	/*198 (0xC6)*/ Instruction{"ifnull", IFNULL, false},
	/*199 (0xC7)*/ Instruction{"ifnonnull", IFNONNULL, false},
	/*200 (0xC8)*/ Instruction{"goto_w", GOTO_W, false},
	/*201 (0xC9)*/ Instruction{"jsr_w", JSR_W, false},
	/*----------RESERVED ---------------*/
	///*202 (0xCA)*/ Instruction{"breakpoint", BREAKPOINT, false},
	///*254 (0xFE)*/    Instruction{"impdep1", IMPDEP1, false},
	///*255 (0xFF)*/    Instruction{"impdep2", IMPDEP2, false},
}

func run(mainClass *ClassType)  {
	mainMethod := mainClass.findMethod(MAIN_METHOD)
	thread := newThread("main")
	thread.pushFrame(NewStackFrame(mainMethod))
	for len(thread.vmStack) != 0 { // per stack frame
		f := thread.peekFrame()
		bytecode := f.method.code
		trace("⤮ %s.%s%s", f.method.class.thisClassName, f.method.name, f.method.descriptor)
		for f.pc < uint32(len(f.method.code)) {
			pc := f.pc
			opcode := bytecode[pc]
			instruction := instructions[opcode]
			trace("   %04d ➢ %s", int(pc), instruction.mnemonic)
			instruction.interpreter(opcode, f, thread, f.method.class, f.method)
			// jump instruction can operate pc
			// some instruction also have variable length: tableswitch...
			// these intructions will control pc themselves
			instruction_length := JVM_OPCODE_LENGTH_INITIALIZER[opcode]
			if f.pc == pc && instruction_length > 0 {
				f.pc += uint32(instruction_length)
			}
			if instruction.switchFrame {
				break
			}
		}
	}
}



