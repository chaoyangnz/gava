package jago

type InstructionRegistry []Instruction

func (this InstructionRegistry) GetInstruction(opcode uint8) Instruction {
	return this[opcode]
}

func (this InstructionRegistry) RegisterInstructions() {
	instructions := []Instruction{// JVM_OPC_MAX+1
		/* ----- CONSTANTS -----------*/
		/*00 (0x00)*/ {"nop", NOP},
		/*01 (0x01)*/ {"aconst_null", ACONST_NULL},
		/*02 (0x02)*/ {"iconst_m1", ICONST_M1},
		/*03 (0x03)*/ {"iconst_0", ICONST_0},
		/*04 (0x04)*/ {"iconst_1", ICONST_1},
		/*05 (0x05)*/ {"iconst_2", ICONST_2},
		/*06 (0x06)*/ {"iconst_3", ICONST_3},
		/*07 (0x07)*/ {"iconst_4", ICONST_4},
		/*08 (0x08)*/ {"iconst_5", ICONST_5},
		/*09 (0x09)*/ {"lconst_0", LCONST_0},
		/*10 (0x0A)*/ {"lconst_1", LCONST_1},
		/*11 (0x0B)*/ {"fconst_0", FCONST_0},
		/*12 (0x0C)*/ {"fconst_1", FCONST_1},
		/*13 (0x0D)*/ {"fconst_2", FCONST_2},
		/*14 (0x0E)*/ {"dconst_0", DCONST_0},
		/*15 (0x0F)*/ {"dconst_1", DCONST_1},
		/*16 (0x10)*/ {"bipush", BIPUSH},
		/*17 (0x11)*/ {"sipush", SIPUSH},
		/*18 (0x12)*/ {"ldc", LDC},
		/*19 (0x13)*/ {"ldc_w", LDC_W},
		/*20 (0x14)*/ {"ldc2_w", LDC2_W},
		/*--------LOADS ----------------*/
		/*21 (0x15)*/ {"iload", ILOAD},
		/*22 (0x16)*/ {"lload", LLOAD},
		/*23 (0x17)*/ {"fload", FLOAD},
		/*24 (0x18)*/ {"dload", DLOAD},
		/*25 (0x19)*/ {"aload", ALOAD},
		/*26 (0x1A)*/ {"iload_0", ILOAD_0},
		/*27 (0x1B)*/ {"iload_1", ILOAD_1},
		/*28 (0x1C)*/ {"iload_2", ILOAD_2},
		/*29 (0x1D)*/ {"iload_3", ILOAD_3},
		/*30 (0x1E)*/ {"lload_0", LLOAD_0},
		/*31 (0x1F)*/ {"lload_1", LLOAD_1},
		/*32 (0x20)*/ {"lload_2", LLOAD_2},
		/*33 (0x21)*/ {"lload_3", LLOAD_3},
		/*34 (0x22)*/ {"fload_0", FLOAD_0},
		/*35 (0x23)*/ {"fload_1", FLOAD_1},
		/*36 (0x24)*/ {"fload_2", FLOAD_2},
		/*37 (0x25)*/ {"fload_3", FLOAD_3},
		/*38 (0x26)*/ {"dload_0", DLOAD_0},
		/*39 (0x27)*/ {"dload_1", DLOAD_1},
		/*40 (0x28)*/ {"dload_2", DLOAD_2},
		/*41 (0x29)*/ {"dload_3", DLOAD_3},
		/*42 (0x2A)*/ {"aload_0", ALOAD_0},
		/*43 (0x2B)*/ {"aload_1", ALOAD_1},
		/*44 (0x2C)*/ {"aload_2", ALOAD_2},
		/*45 (0x2D)*/ {"aload_3", ALOAD_3},
		/*46 (0x2E)*/ {"iaload", IALOAD},
		/*47 (0x2F)*/ {"laload", LALOAD},
		/*48 (0x30)*/ {"faload", FALOAD},
		/*49 (0x31)*/ {"daload", DALOAD},
		/*50 (0x32)*/ {"aaload", AALOAD},
		/*51 (0x33)*/ {"baload", BALOAD},
		/*52 (0x34)*/ {"caload", CALOAD},
		/*53 (0x35)*/ {"saload", SALOAD},
		/*--------STORES ----------------*/
		/*54 (0x36)*/ {"istore", ISTORE},
		/*55 (0x37)*/ {"lstore", LSTORE},
		/*56 (0x38)*/ {"fstore", FSTORE},
		/*57 (0x39)*/ {"dstore", DSTORE},
		/*58 (0x3A)*/ {"astore", ASTORE},
		/*59 (0x3B)*/ {"istore_0", ISTORE_0},
		/*60 (0x3C)*/ {"istore_1", ISTORE_1},
		/*61 (0x3D)*/ {"istore_2", ISTORE_2},
		/*62 (0x3E)*/ {"istore_3", ISTORE_3},
		/*63 (0x3F)*/ {"lstore_0", LSTORE_0},
		/*64 (0x40)*/ {"lstore_1", LSTORE_1},
		/*65 (0x41)*/ {"lstore_2", LSTORE_2},
		/*66 (0x42)*/ {"lstore_3", LSTORE_3},
		/*67 (0x43)*/ {"fstore_0", FSTORE_0},
		/*68 (0x44)*/ {"fstore_1", FSTORE_1},
		/*69 (0x45)*/ {"fstore_2", FSTORE_2},
		/*70 (0x46)*/ {"fstore_3", FSTORE_3},
		/*71 (0x47)*/ {"dstore_0", DSTORE_0},
		/*72 (0x48)*/ {"dstore_1", DSTORE_1},
		/*73 (0x49)*/ {"dstore_2", DSTORE_2},
		/*74 (0x4A)*/ {"dstore_3", DSTORE_3},
		/*75 (0x4B)*/ {"astore_0", ASTORE_0},
		/*76 (0x4C)*/ {"astore_1", ASTORE_1},
		/*77 (0x4D)*/ {"astore_2", ASTORE_2},
		/*78 (0x4E)*/ {"astore_3", ASTORE_3},
		/*79 (0x4F)*/ {"iastore", IASTORE},
		/*80 (0x50)*/ {"lastore", LASTORE},
		/*81 (0x51)*/ {"fastore", FASTORE},
		/*82 (0x52)*/ {"dastore", DASTORE},
		/*83 (0x53)*/ {"aastore", AASTORE},
		/*84 (0x54)*/ {"bastore", BASTORE},
		/*85 (0x55)*/ {"castore", CASTORE},
		/*86 (0x56)*/ {"sastore", SASTORE},
		/*--------STACK---------------*/
		/*87 (0x57)*/ {"pop", POP},
		/*88 (0x58)*/ {"pop2", POP2},
		/*89 (0x59)*/ {"dup", DUP},
		/*90 (0x5A)*/ {"dup_x1", DUP_X1},
		/*91 (0x5B)*/ {"dup_x2", DUP_X2},
		/*92 (0x5C)*/ {"dup2", DUP2},
		/*93 (0x5D)*/ {"dup2_x1", DUP2_X1},
		/*94 (0x5E)*/ {"dup2_x2", DUP2_X2},
		/*95 (0x5F)*/ {"swap", SWAP},
		/*---------MATH -------------*/
		/*96 (0x60)*/ {"iadd", IADD},
		/*97 (0x61)*/ {"ladd", LADD},
		/*98 (0x62)*/ {"fadd", FADD},
		/*99 (0x63)*/ {"dadd", DADD},
		/*100 (0x64)*/ {"isub", ISUB},
		/*101 (0x65)*/ {"lsub", LSUB},
		/*102 (0x66)*/ {"fsub", FSUB},
		/*103 (0x67)*/ {"dsub", DSUB},
		/*104 (0x68)*/ {"imul", IMUL},
		/*105 (0x69)*/ {"lmul", LMUL},
		/*106 (0x6A)*/ {"fmul", FMUL},
		/*107 (0x6B)*/ {"dmul", DMUL},
		/*108 (0x6C)*/ {"idiv", IDIV},
		/*109 (0x6D)*/ {"ldiv", LDIV},
		/*110 (0x6E)*/ {"fdiv", FDIV},
		/*111 (0x6F)*/ {"ddiv", DDIV},
		/*112 (0x70)*/ {"irem", IREM},
		/*113 (0x71)*/ {"lrem", LREM},
		/*114 (0x72)*/ {"frem", FREM},
		/*115 (0x73)*/ {"drem", DREM},
		/*116 (0x74)*/ {"ineg", INEG},
		/*117 (0x75)*/ {"lneg", LNEG},
		/*118 (0x76)*/ {"fneg", FNEG},
		/*119 (0x77)*/ {"dneg", DNEG},
		/*120 (0x78)*/ {"ishl", ISHL},
		/*121 (0x79)*/ {"lshl", LSHL},
		/*122 (0x7A)*/ {"ishr", ISHR},
		/*123 (0x7B)*/ {"lshr", LSHR},
		/*124 (0x7C)*/ {"iushr", IUSHR},
		/*125 (0x7D)*/ {"lushr", LUSHR},
		/*126 (0x7E)*/ {"iand", IAND},
		/*127 (0x7F)*/ {"land", LAND},
		/*128 (0x80)*/ {"ior", IOR},
		/*129 (0x81)*/ {"lor", LOR},
		/*130 (0x82)*/ {"ixor", IXOR},
		/*131 (0x83)*/ {"lxor", LXOR},
		/*132 (0x84)*/ {"iinc", IINC},
		/*--------CONVERSIONS-----------*/
		/*133 (0x85)*/ {"i2l", I2L},
		/*134 (0x86)*/ {"i2f", I2F},
		/*135 (0x87)*/ {"i2d", I2D},
		/*136 (0x88)*/ {"l2i", L2I},
		/*137 (0x89)*/ {"l2f", L2F},
		/*138 (0x8A)*/ {"l2d", L2D},
		/*139 (0x8B)*/ {"f2i", F2I},
		/*140 (0x8C)*/ {"f2l", F2L},
		/*141 (0x8D)*/ {"f2d", F2D},
		/*142 (0x8E)*/ {"d2i", D2I},
		/*143 (0x8F)*/ {"d2l", D2L},
		/*144 (0x90)*/ {"d2f", D2F},
		/*145 (0x91)*/ {"i2b", I2B},
		/*146 (0x92)*/ {"i2c", I2C},
		/*147 (0x93)*/ {"i2s", I2S},
		/*-----------COMPARASON -----------*/
		/*148 (0x94)*/ {"lcmp", LCMP},
		/*149 (0x95)*/ {"fcmpl", FCMPL},
		/*150 (0x96)*/ {"fcmpg", FCMPG},
		/*151 (0x97)*/ {"dcmpl", DCMPL},
		/*152 (0x98)*/ {"dcmpg", DCMPG},
		/*153 (0x99)*/ {"ifeq", IFEQ},
		/*154 (0x9A)*/ {"ifne", IFNE},
		/*155 (0x9B)*/ {"iflt", IFLT},
		/*156 (0x9C)*/ {"ifge", IFGE},
		/*157 (0x9D)*/ {"ifgt", IFGT},
		/*158 (0x9E)*/ {"ifle", IFLE},
		/*159 (0x9F)*/ {"if_icmpeq", IF_ICMPEQ},
		/*160 (0xA0)*/ {"if_icmpne", IF_ICMPNE},
		/*161 (0xA1)*/ {"if_icmplt", IF_ICMPLT},
		/*162 (0xA2)*/ {"if_icmpge", IF_ICMPGE},
		/*163 (0xA3)*/ {"if_icmpgt", IF_ICMPGT},
		/*164 (0xA4)*/ {"if_icmple", IF_ICMPLE},
		/*165 (0xA5)*/ {"if_acmpeq", IF_ACMPEQ},
		/*166 (0xA6)*/ {"if_acmpne", IF_ACMPNE},
		/*---------REFERENCES -------------*/
		/*167 (0xA7)*/ {"goto", GOTO},
		/*168 (0xA8)*/ {"jsr", JSR},
		/*169 (0xA9)*/ {"ret", RET},
		/*170 (0xAA)*/ {"tableswitch", TABLESWITCH}, // variable bytecode length
		/*171 (0xAB)*/ {"lookupswitch", LOOKUPSWITCH},
		/*172 (0xAC)*/ {"ireturn", IRETURN},
		/*173 (0xAD)*/ {"lreturn", LRETURN},
		/*174 (0xAE)*/ {"freturn", FRETURN},
		/*175 (0xAF)*/ {"dreturn", DRETURN},
		/*176 (0xB0)*/ {"areturn", ARETURN},
		/*177 (0xB1)*/ {"return", RETURN},
		/*-------CONTROLS------------------*/
		/*178 (0xB2)*/ {"getstatic", GETSTATIC},
		/*179 (0xB3)*/ {"putstatic", PUTSTATIC},
		/*180 (0xB4)*/ {"getfield", GETFIELD},
		/*181 (0xB5)*/ {"putfield", PUTFIELD},
		/*182 (0xB6)*/ {"invokevirtual", INVOKEVIRTUAL},
		/*183 (0xB7)*/ {"invokespecial", INVOKESPECIAL},
		/*184 (0xB8)*/ {"invokestatic", INVOKESTATIC},
		/*185 (0xB9)*/ {"invokeinterface", INVOKEINTERFACE},
		/*186 (0xBA)*/ {"invokedynamic", INVOKEDYNAMIC},
		/*187 (0xBB)*/ {"new", NEW},
		/*188 (0xBC)*/ {"newarray", NEWARRAY},
		/*189 (0xBD)*/ {"anewarray", ANEWARRAY},
		/*190 (0xBE)*/ {"arraylength", ARRAYLENGTH},
		/*191 (0xBF)*/ {"athrow", ATHROW},
		/*192 (0xC0)*/ {"checkcast", CHECKCAST},
		/*193 (0xC1)*/ {"instanceof", INSTANCEOF},
		/*194 (0xC2)*/ {"monitorenter", MONITORENTER},
		/*195 (0xC3)*/ {"monitorexit", MONITOREXIT},
		/*--------EXTENDED-----------------*/
		/*196 (0xC4)*/ {"wide", WIDE},
		/*197 (0xC5)*/ {"multianewarray", MULTIANEWARRAY},
		/*198 (0xC6)*/ {"ifnull", IFNULL},
		/*199 (0xC7)*/ {"ifnonnull", IFNONNULL},
		/*200 (0xC8)*/ {"goto_w", GOTO_W},
		/*201 (0xC9)*/ {"jsr_w", JSR_W},
		/*----------RESERVED ---------------*/
		/*202 (0xCA)*/ {"breakpoint", BREAKPOINT},
		///*254 (0xFE)*/    Instruction{"impdep1", IMPDEP1},
		///*255 (0xFF)*/    Instruction{"impdep2", IMPDEP2},
	}

	for i, instruction := range instructions {
		this[i] = instruction
	}
}

type Instruction struct {
	mnemonic  string
	interpret func(*Thread, *Frame, *Class, *Method)
}

func assertIntCompatible(value Value) Value {
	switch value.(type) {
	case Boolean, Byte, Short, Char, Int:
	default:
		Fatal("%s is not int compatible", value.Type().Name())
	}
	return value
}

func intCompatible(value Value) Int {
	assertIntCompatible(value)

	switch value.(type) {
	case Boolean:
		return Int(value.(Boolean))
	case Byte:
		return Int(value.(Byte))
	case Short:
		return Int(value.(Short))
	case Char:
		return Int(value.(Char))
	case Int:
		return value.(Int)
	default:
		Fatal("%s is not int compatible", value.Type().Name())
	}

	// never be here
	return 0
}
