package gvm

import (
	"fmt"
	"math"
)

const MAIN_METHOD = "main([Ljava/lang/String;)V"

var opcodes = []func(uint8, *StackFrame, *Thread, *JavaClass, *JavaMethod) (int, bool) {
	/* ----- CONSTANTS -----------*/
	/*00 (0X00)*/    NOP,
	/*01 (0X01)*/    ACONST_NULL,
	/*02 (0X02)*/    ICONST_M1,
	/*03 (0X03)*/    ICONST_0,
	/*04 (0X04)*/    ICONST_1,
	/*05 (0X05)*/    ICONST_2,
	/*06 (0X06)*/    ICONST_3,
	/*07 (0X07)*/    ICONST_4,
	/*08 (0X08)*/    ICONST_5,
	/*09 (0X09)*/    LCONST_0,
	/*10 (0X0A)*/    LCONST_1,
	/*11 (0X0B)*/    FCONST_0,
	/*12 (0X0C)*/    FCONST_1,
	/*13 (0X0D)*/    FCONST_2,
	/*14 (0X0E)*/    DCONST_0,
	/*15 (0X0F)*/    DCONST_1,
	/*16 (0X10)*/    BIPUSH,
	/*17 (0X11)*/    SIPUSH,
	/*18 (0X12)*/    LDC,
	/*19 (0X13)*/    LDC_W,
	/*20 (0X14)*/    LDC2_W,
	/*--------LOADS ----------------*/
	/*21 (0X15)*/    ILOAD,
	/*22 (0X16)*/    LLOAD,
	/*23 (0X17)*/    FLOAD,
	/*24 (0X18)*/    DLOAD,
	/*25 (0X19)*/    ALOAD,
	/*26 (0X1A)*/    ILOAD_0,
	/*27 (0X1B)*/    ILOAD_1,
	/*28 (0X1C)*/    ILOAD_2,
	/*29 (0X1D)*/    ILOAD_3,
	/*30 (0X1E)*/    LLOAD_0,
	/*31 (0X1F)*/    LLOAD_1,
	/*32 (0X20)*/    LLOAD_2,
	/*33 (0X21)*/    LLOAD_3,
	/*34 (0X22)*/    FLOAD_0,
	/*35 (0X23)*/    FLOAD_1,
	/*36 (0X24)*/    FLOAD_2,
	/*37 (0X25)*/    FLOAD_3,
	/*38 (0X26)*/    DLOAD_0,
	/*39 (0X27)*/    DLOAD_1,
	/*40 (0X28)*/    DLOAD_2,
	/*41 (0X29)*/    DLOAD_3,
	/*42 (0X2A)*/    ALOAD_0,
	/*43 (0X2B)*/    ALOAD_1,
	/*44 (0X2C)*/    ALOAD_2,
	/*45 (0X2D)*/    ALOAD_3,
	/*46 (0X2E)*/    IALOAD,
	/*47 (0X2F)*/    LALOAD,
	/*48 (0X30)*/    FALOAD,
	/*49 (0X31)*/    DALOAD,
	/*50 (0X32)*/    AALOAD,
	/*51 (0X33)*/    BALOAD,
	/*52 (0X34)*/    CALOAD,
	/*53 (0X35)*/    SALOAD,
	/*--------STORES ----------------*/
	/*54 (0X36)*/    ISTORE,
	/*55 (0X37)*/    LSTORE,
	/*56 (0X38)*/    FSTORE,
	/*57 (0X39)*/    DSTORE,
	/*58 (0X3A)*/    ASTORE,
	/*59 (0X3B)*/    ISTORE_0,
	/*60 (0X3C)*/    ISTORE_1,
	/*61 (0X3D)*/    ISTORE_2,
	/*62 (0X3E)*/    ISTORE_3,
	/*63 (0X3F)*/    LSTORE_0,
	/*64 (0X40)*/    LSTORE_1,
	/*65 (0X41)*/    LSTORE_2,
	/*66 (0X42)*/    LSTORE_3,
	/*67 (0X43)*/    FSTORE_0,
	/*68 (0X44)*/    FSTORE_1,
	/*69 (0X45)*/    FSTORE_2,
	/*70 (0X46)*/    FSTORE_3,
	/*71 (0X47)*/    DSTORE_0,
	/*72 (0X48)*/    DSTORE_1,
	/*73 (0X49)*/    DSTORE_2,
	/*74 (0X4A)*/    DSTORE_3,
	/*75 (0X4B)*/    ASTORE_0,
	/*76 (0X4C)*/    ASTORE_1,
	/*77 (0X4D)*/    ASTORE_2,
	/*78 (0X4E)*/    ASTORE_3,
	/*79 (0X4F)*/    IASTORE,
	/*80 (0X50)*/    LASTORE,
	/*81 (0X51)*/    FASTORE,
	/*82 (0X52)*/    DASTORE,
	/*83 (0X53)*/    AASTORE,
	/*84 (0X54)*/    BASTORE,
	/*85 (0X55)*/    CASTORE,
	/*86 (0X56)*/    SASTORE,
	/*--------STACK---------------*/
	/*87 (0X57)*/    POP,
	/*88 (0X58)*/    POP2,
	/*89 (0X59)*/    DUP,
	/*90 (0X5A)*/    DUP_X1,
	/*91 (0X5B)*/    DUP_X2,
	/*92 (0X5C)*/    DUP2,
	/*93 (0X5D)*/    DUP2_X1,
	/*94 (0X5E)*/    DUP2_X2,
	/*95 (0X5F)*/    SWAP,
	/*---------MATH -------------*/
	/*96 (0X60)*/    IADD,
	/*97 (0X61)*/    LADD,
	/*98 (0X62)*/    FADD,
	/*99 (0X63)*/    DADD,
	/*100 (0X64)*/    ISUB,
	/*101 (0X65)*/    LSUB,
	/*102 (0X66)*/    FSUB,
	/*103 (0X67)*/    DSUB,
	/*104 (0X68)*/    IMUL,
	/*105 (0X69)*/    LMUL,
	/*106 (0X6A)*/    FMUL,
	/*107 (0X6B)*/    DMUL,
	/*108 (0X6C)*/    IDIV,
	/*109 (0X6D)*/    LDIV,
	/*110 (0X6E)*/    FDIV,
	/*111 (0X6F)*/    DDIV,
	/*112 (0X70)*/    IREM,
	/*113 (0X71)*/    LREM,
	/*114 (0X72)*/    FREM,
	/*115 (0X73)*/    DREM,
	/*116 (0X74)*/    INEG,
	/*117 (0X75)*/    LNEG,
	/*118 (0X76)*/    FNEG,
	/*119 (0X77)*/    DNEG,
	/*120 (0X78)*/    ISHL,
	/*121 (0X79)*/    LSHL,
	/*122 (0X7A)*/    ISHR,
	/*123 (0X7B)*/    LSHR,
	/*124 (0X7C)*/    IUSHR,
	/*125 (0X7D)*/    LUSHR,
	/*126 (0X7E)*/    IAND,
	/*127 (0X7F)*/    LAND,
	/*128 (0X80)*/    IOR,
	/*129 (0X81)*/    LOR,
	/*130 (0X82)*/    IXOR,
	/*131 (0X83)*/    LXOR,
	/*132 (0X84)*/    IINC,
	/*--------CONVERSIONS-----------*/
	/*133 (0X85)*/    I2L,
	/*134 (0X86)*/    I2F,
	/*135 (0X87)*/    I2D,
	/*136 (0X88)*/    L2I,
	/*137 (0X89)*/    L2F,
	/*138 (0X8A)*/    L2D,
	/*139 (0X8B)*/    F2I,
	/*140 (0X8C)*/    F2L,
	/*141 (0X8D)*/    F2D,
	/*142 (0X8E)*/    D2I,
	/*143 (0X8F)*/    D2L,
	/*144 (0X90)*/    D2F,
	/*145 (0X91)*/    I2B,
	/*146 (0X92)*/    I2C,
	/*147 (0X93)*/    I2S,
	/*-----------COMPARASON -----------*/
	/*148 (0X94)*/    LCMP,
	/*149 (0X95)*/    FCMPL,
	/*150 (0X96)*/    FCMPG,
	/*151 (0X97)*/    DCMPL,
	/*152 (0X98)*/    DCMPG,
	/*153 (0X99)*/    IFEQ,
	/*154 (0X9A)*/    IFNE,
	/*155 (0X9B)*/    IFLT,
	/*156 (0X9C)*/    IFGE,
	/*157 (0X9D)*/    IFGT,
	/*158 (0X9E)*/    IFLE,
	/*159 (0X9F)*/    IF_ICMPEQ,
	/*160 (0XA0)*/    IF_ICMPNE,
	/*161 (0XA1)*/    IF_ICMPLT,
	/*162 (0XA2)*/    IF_ICMPGE,
	/*163 (0XA3)*/    IF_ICMPGT,
	/*164 (0XA4)*/    IF_ICMPLE,
	/*165 (0XA5)*/    IF_ACMPEQ,
	/*166 (0XA6)*/    IF_ACMPNE,
	/*---------REFERENCES -------------*/
	/*167 (0XA7)*/    GOTO,
	/*168 (0XA8)*/    JSR,
	/*169 (0XA9)*/    RET,
	/*170 (0XAA)*/    TABLESWITCH,
	/*171 (0XAB)*/    LOOKUPSWITCH,
	/*172 (0XAC)*/    IRETURN,
	/*173 (0XAD)*/    LRETURN,
	/*174 (0XAE)*/    FRETURN,
	/*175 (0XAF)*/    DRETURN,
	/*176 (0XB0)*/    ARETURN,
	/*177 (0XB1)*/    RETURN,
	/*-------CONTROLS------------------*/
	/*178 (0XB2)*/    GETSTATIC,
	/*179 (0XB3)*/    PUTSTATIC,
	/*180 (0XB4)*/    GETFIELD,
	/*181 (0XB5)*/    PUTFIELD,
	/*182 (0XB6)*/    INVOKEVIRTUAL,
	/*183 (0XB7)*/    INVOKESPECIAL,
	/*184 (0XB8)*/    INVOKESTATIC,
	/*185 (0XB9)*/    INVOKEINTERFACE,
	/*186 (0XBA)*/    INVOKEDYNAMIC,
	/*187 (0XBB)*/    NEW,
	/*188 (0XBC)*/    NEWARRAY,
	/*189 (0XBD)*/    ANEWARRAY,
	/*190 (0XBE)*/    ARRAYLENGTH,
	/*191 (0XBF)*/    ATHROW,
	/*192 (0XC0)*/    CHECKCAST,
	/*193 (0XC1)*/    INSTANCEOF,
	/*194 (0XC2)*/    MONITORENTER,
	/*195 (0XC3)*/    MONITOREXIT,
	/*--------EXTENDED-----------------*/
	/*196 (0XC4)*/    WIDE,
	/*197 (0XC5)*/    MULTIANEWARRAY,
	/*198 (0XC6)*/    IFNULL,
	/*199 (0XC7)*/    IFNONNULL,
	/*200 (0XC8)*/    GOTO_W,
	/*201 (0XC9)*/    JSR_W,
	/*----------RESERVED ---------------*/
	/*202 (0XCA)*/    BREAKPOINT,
	///*254 (0XFE)*/    IMPDEP1,
	///*255 (0XFF)*/    IMPDEP2,
}

func run(mainClass *JavaClass)  {
	mainMethod := mainClass.findMethod(MAIN_METHOD)
	thread := newThread("main")
	thread.vmStack.push(NewStackFrame(mainMethod))
	for thread.vmStack.size != 0 { // per stack frame
		f := thread.vmStack.peek()
		bytecode := f.method.code
		for f.pc < uint32(len(f.method.code)) {
				opcode := bytecode[f.pc]
				progress, next := opcodes[opcode](opcode, f, thread, f.method.class, f.method)
				f.pc += uint32(progress)
				if !next {
					break
				}
			}
		}
}

/*00 (0X00)*/
func NOP(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	return 1, true
}

/*01 (0X01)*/
func ACONST_NULL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(nil)
	return 1, true
}

/*02 (0X02)*/
func ICONST_M1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_int(-1))
	return 1, true
}

/*03 (0X03)*/
func ICONST_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_int(0))
	return 1, true
}

/*04 (0X04)*/
func ICONST_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_int(1))
	return 1, true
}

/*05 (0X05)*/
func ICONST_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_int(2))
	return 1, true
}

/*06 (0X06)*/
func ICONST_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_int(3))
	return 1, true
}

/*07 (0X07)*/
func ICONST_4(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_int(4))
	return 1, true
}

/*08 (0X08)*/
func ICONST_5(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_int(5))
	return 1, true
}

/*09 (0X09)*/
func LCONST_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_long(0))
	return 1, true
}

/*10 (0X0A)*/
func LCONST_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_long(1))
	return 1, true
}

/*11 (0X0B)*/
func FCONST_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_float(0.0))
	return 1, true
}

/*12 (0X0C)*/
func FCONST_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_float(1.0))
	return 1, true
}

/*13 (0X0D)*/
func FCONST_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_float(2.0))
	return 1, true
}

/*14 (0X0E)*/
func DCONST_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_double(0.0))
	return 1, true
}

/*15 (0X0F)*/
func DCONST_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(t_double(1.0))
	return 1, true
}

/*16 (0X10)*/
func BIPUSH(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	b := t_byte(m.code[f.pc+1])
	f.push(t_int(b))
	return 2, true
}

/*17 (0X11)*/
func SIPUSH(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	s := t_short((m.code[f.pc+1] << 8) | m.code[f.pc+2])
	f.push(t_int(s))
	return 2, true
}

/*18 (0X12)*/
func LDC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	cpInfo := c.constantPool[index].resolve()
	switch cpInfo.(type) {
	case *RuntimeConstantIntegerInfo:
		f.push(cpInfo.(*RuntimeConstantIntegerInfo).value)
	case *RuntimeConstantFloatInfo:
		f.push(cpInfo.(*RuntimeConstantFloatInfo).value)
	case *RuntimeConstantStringInfo:
		f.push(cpInfo.(*RuntimeConstantStringInfo).value)
	default:
		panic("Not int, float or string in constant pool")
	}
	return 2, true
}

/*19 (0X13)*/
func LDC_W(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])
	cpInfo := c.constantPool[index].resolve()
	switch cpInfo.(type) {
	case *RuntimeConstantIntegerInfo:
		f.push(cpInfo.(*RuntimeConstantIntegerInfo).value)
	case *RuntimeConstantFloatInfo:
		f.push(cpInfo.(*RuntimeConstantFloatInfo).value)
	case *RuntimeConstantStringInfo:
		f.push(cpInfo.(*RuntimeConstantStringInfo).value)
	default:
		panic("Not int, float or string in constant pool")
	}
	return 3, true
}

/*20 (0X14)*/
func LDC2_W(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])
	cpInfo := c.constantPool[index].resolve()
	switch cpInfo.(type) {
	case *RuntimeConstantLongInfo:
		f.push(cpInfo.(*RuntimeConstantLongInfo).value)
	case *RuntimeConstantDoubleInfo:
		f.push(cpInfo.(*RuntimeConstantDoubleInfo).value)
	default:
		panic("Not long or double in constant pool")
	}
	return 3, true
}


/*21 (0X15)*/
func ILOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_int))
	return 2, true
}

/*22 (0X16)*/
func LLOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_long))
	return 2, true
}

/*23 (0X17)*/
func FLOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_float))
	return 2, true
}

/*24 (0X18)*/
func DLOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_double))
	return 2, true
}

/*25 (0X19)*/
func ALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_object))
	return 2, true
}

/*26 (0X1A)*/
func ILOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(0).(t_int))
	return 1, true
}

/*27 (0X1B)*/
func ILOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(1).(t_int))
	return 1, true
}

/*28 (0X1C)*/
func ILOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(2).(t_int))
	return 1, true
}

/*29 (0X1D)*/
func ILOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(3).(t_int))
	return 1, true
}

/*30 (0X1E)*/
func LLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(0)).(t_long))
	return 1, true
}

/*31 (0X1F)*/
func LLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(1)).(t_long))
	return 1, true
}

/*32 (0X20)*/
func LLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(2)).(t_long))
	return 1, true
}

/*33 (0X21)*/
func LLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(3)).(t_long))
	return 1, true
}

/*34 (0X22)*/
func FLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(0)).(t_float))
	return 1, true
}

/*35 (0X23)*/
func FLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(1)).(t_float))
	return 1, true
}

/*36 (0X24)*/
func FLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(2)).(t_float))
	return 1, true
}

/*37 (0X25)*/
func FLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(3)).(t_float))
	return 1, true
}

/*38 (0X26)*/
func DLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(0)).(t_double))
	return 1, true
}

/*39 (0X27)*/
func DLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(1)).(t_double))
	return 1, true
}

/*40 (0X28)*/
func DLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(2)).(t_double))
	return 1, true
}

/*41 (0X29)*/
func DLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(uint(3)).(t_double))
	return 1, true
}

/*42 (0X2A)*/
func ALOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(0).(t_object))
	return 1, true
}

/*43 (0X2B)*/
func ALOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(1).(t_object))
	return 1, true
}

/*44 (0X2C)*/
func ALOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(2).(t_object))
	return 1, true
}

/*45 (0X2D)*/
func ALOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.loadVar(3).(t_object))
	return 1, true
}

/*46 (0X2E)
Run-time Exceptions

If arrayref is null, iaload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the iaload instruction throws an ArrayIndexOutOfBoundsException.
*/
func IALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_INT {
		panic("Not an int array")
	}
	f.push(arrayref.elements[index])
	return 1, true
}

/*47 (0X2F)*/
func LALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_LONG {
		panic("Not an long array")
	}
	f.push(arrayref.elements[index])
	return 1, true
}

/*48 (0X30)*/
func FALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_FLOAT {
		panic("Not an long array")
	}
	f.push(arrayref.elements[index])
	return 1, true
}

/*49 (0X31)*/
func DALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_DOUBLE {
		panic("Not an long array")
	}
	f.push(arrayref.elements[index])
	return 1, true
}

/*50 (0X32)*/
func AALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != 0 || arrayref.aclass == nil {
		panic("Not an object array")
	}
	f.push(arrayref.elements[index])
	return 1, true
}

/*51 (0X33)*/
func BALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_BOOLEAN {
		panic("Not an boolean array")
	}
	f.push(arrayref.elements[index])
	return 1, true
}

/*52 (0X34)*/
func CALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_CHAR {
		panic("Not an char array")
	}
	f.push(arrayref.elements[index])
	return 1, true
}

/*53 (0X35)*/
func SALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_SHORT {
		panic("Not an short array")
	}
	f.push(arrayref.elements[index])
	return 1, true
}


/*54 (0X36)*/
func ISTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.storeVar(uint(index), f.pop().(t_int))
	return 2, true
}

/*55 (0X37)*/
func LSTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.storeVar(uint(index), f.pop().(t_long))
	return 2, true
}

/*56 (0X38)*/
func FSTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.storeVar(uint(index), f.pop().(t_float))
	return 2, true
}

/*57 (0X39)*/
func DSTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.storeVar(uint(index), f.pop().(t_double))
	return 2, true
}

/*58 (0X3A)*/
func ASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	f.storeVar(uint(index), f.pop().(t_object))
	return 2, true
}

/*59 (0X3B)*/
func ISTORE_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(0, f.pop().(t_int))
	return 1, true
}

/*60 (0X3C)*/
func ISTORE_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(1, f.pop().(t_int))
	return 1, true
}

/*61 (0X3D)*/
func ISTORE_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(2, f.pop().(t_int))
	return 1, true
}

/*62 (0X3E)*/
func ISTORE_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(3, f.pop().(t_int))
	return 1, true
}

/*63 (0X3F)*/
func LSTORE_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(0, f.pop().(t_long))
	return 1, true
}

/*64 (0X40)*/
func LSTORE_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(1, f.pop().(t_long))
	return 1, true
}

/*65 (0X41)*/
func LSTORE_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(2, f.pop().(t_long))
	return 1, true
}

/*66 (0X42)*/
func LSTORE_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(3, f.pop().(t_long))
	return 1, true
}

/*67 (0X43)*/
func FSTORE_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(0, f.pop().(t_float))
	return 1, true
}

/*68 (0X44)*/
func FSTORE_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(1, f.pop().(t_float))
	return 1, true
}

/*69 (0X45)*/
func FSTORE_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(2, f.pop().(t_float))
	return 1, true
}

/*70 (0X46)*/
func FSTORE_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(3, f.pop().(t_float))
	return 1, true
}

/*71 (0X47)*/
func DSTORE_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(0, f.pop().(t_double))
	return 1, true
}

/*72 (0X48)*/
func DSTORE_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(1, f.pop().(t_double))
	return 1, true
}

/*73 (0X49)*/
func DSTORE_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(2, f.pop().(t_double))
	return 1, true
}

/*74 (0X4A)*/
func DSTORE_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(3, f.pop().(t_double))
	return 1, true
}

/*75 (0X4B)*/
func ASTORE_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(0, f.pop().(t_object))
	return 1, true
}

/*76 (0X4C)*/
func ASTORE_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(1, f.pop().(t_object))
	return 1, true
}

/*77 (0X4D)*/
func ASTORE_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(2, f.pop().(t_object))
	return 1, true
}

/*78 (0X4E)*/
func ASTORE_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.storeVar(3, f.pop().(t_object))
	return 1, true
}

/*79 (0X4F)*/
func IASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_int)
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	arrayref.elements[index] = value
	//TODO check component type and boundary
	return 1, true
}

/*80 (0X50)*/
func LASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_long)
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	arrayref.elements[index] = value
	//TODO check component type and boundary
	return 1, true
}

/*81 (0X51)*/
func FASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_float)
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	arrayref.elements[index] = value
	//TODO check component type and boundary
	return 1, true
}

/*82 (0X52)*/
func DASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_double)
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	arrayref.elements[index] = value
	//TODO check component type and boundary
	return 1, true
}

/*83 (0X53)*/
func AASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_object)
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	arrayref.elements[index] = value
	//TODO check component type, boundary and subtypes
	return 1, true
}

/*84 (0X54)*/
func BASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_int)
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	arrayref.elements[index] = t_byte(value)
	//TODO check component type and boundary
	return 1, true
}

/*85 (0X55)*/
func CASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_int)
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	arrayref.elements[index] = t_char(value)
	//TODO check component type and boundary
	return 1, true
}

/*86 (0X56)*/
func SASTORE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_int)
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	arrayref.elements[index] = t_short(value)
	//TODO check component type and boundary
	return 1, true
}


/*87 (0X57)*/
func POP(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.pop()
	return 1, false
}

/*88 (0X58)*/
func POP2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*89 (0X59)*/
func DUP(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop()
	f.push(value)
	f.push(value)
	return 1, true
}

/*90 (0X5A)*/
func DUP_X1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value1 := f.pop()
	value2 := f.pop()
	f.push(value1)
	f.push(value2)
	f.push(value1)
	return 1, true
}

/*91 (0X5B)*/
func DUP_X2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*92 (0X5C)*/
func DUP2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*93 (0X5D)*/
func DUP2_X1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*94 (0X5E)*/
func DUP2_X2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*95 (0X5F)*/
func SWAP(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value1 := f.pop()
	value2 := f.pop()
	f.push(value1)
	f.push(value2)
	return 1, true
}


/*96 (0X60)*/
func IADD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(value1 + value2)
	return 1, true
}

/*97 (0X61)*/
func LADD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_long)
	value1 := f.pop().(t_long)
	f.push(value1 + value2)
	return 1, true
}

/*98 (0X62)*/
func FADD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_float)
	value1 := f.pop().(t_float)
	f.push(value1 + value2)
	return 1, true
}

/*99 (0X63)*/
func DADD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_double)
	value1 := f.pop().(t_double)
	f.push(value1 + value2)
	return 1, true
}

/*100 (0X64)*/
func ISUB(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(value1 - value2)
	return 1, true
}

/*101 (0X65)*/
func LSUB(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_long)
	value1 := f.pop().(t_long)
	f.push(t_long(value1 - value2))
	return 1, true
}

/*102 (0X66)*/
func FSUB(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_float)
	value1 := f.pop().(t_float)
	f.push(t_float(value1 - value2))
	return 1, true
}

/*103 (0X67)*/
func DSUB(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_double)
	value1 := f.pop().(t_double)
	f.push(t_double(value1 - value2))
	return 1, true
}

/*104 (0X68)*/
func IMUL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(value1 * value2))
	return 1, true
}

/*105 (0X69)*/
func LMUL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_long)
	value1 := f.pop().(t_long)
	f.push(t_long(value1 * value2))
	return 1, true
}

/*106 (0X6A)*/
func FMUL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_float)
	value1 := f.pop().(t_float)
	f.push(t_float(value1 * value2))
	return 1, true
}

/*107 (0X6B)*/
func DMUL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_double)
	value1 := f.pop().(t_double)
	f.push(t_double(value1 * value2))
	return 1, true
}

/*108 (0X6C)*/
func IDIV(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(value1 / value2))
	return 1, true
}

/*109 (0X6D)*/
func LDIV(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_long)
	value1 := f.pop().(t_long)
	f.push(t_long(value1 / value2))
	return 1, true
}

/*110 (0X6E)*/
func FDIV(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_float)
	value1 := f.pop().(t_float)
	f.push(t_float(value1 / value2))
	return 1, true
}

/*111 (0X6F)*/
func DDIV(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_double)
	value1 := f.pop().(t_double)
	f.push(t_double(value1 / value2))
	return 1, true
}

/*112 (0X70)*/
func IREM(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(value1 % value2))
	return 1, true
}

/*113 (0X71)*/
func LREM(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_long)
	value1 := f.pop().(t_long)
	f.push(t_long(value1 % value2))
	return 1, true
}

/*114 (0X72)*/
func FREM(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_float)
	value1 := f.pop().(t_float)
	f.push(math.Remainder(float64(value1), float64(value2)))
	return 1, true
}

/*115 (0X73)*/
func DREM(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_double)
	value1 := f.pop().(t_double)
	f.push(t_float(math.Remainder(float64(value1), float64(value2))))
	return 1, true
}

/*116 (0X74)*/
func INEG(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_int)
	f.push(t_int(-value))
	return 1, true
}

/*117 (0X75)*/
func LNEG(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_long)
	f.push(t_long(-value))
	return 1, true
}

/*118 (0X76)*/
func FNEG(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_float)
	f.push(t_float(-value))
	return 1, true
}

/*119 (0X77)*/
func DNEG(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_double)
	f.push(t_double(-value))
	return 1, true
}

/*120 (0X78)*/
func ISHL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(value1 << uint(value2 & 0x1F))) // low 5 bits
	return 1, true
}

/*121 (0X79)*/
func LSHL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)

	value1 := f.pop().(t_long)
	f.push(t_long(value1 << uint(value2 & 0x3F))) // low 6 bits
	return 1, true
}

/*122 (0X7A)*/
func ISHR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(value1 >> uint(value2 & 0x1F))) // low 5 bits
	return 1, true
}

/*123 (0X7B)*/
func LSHR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)

	value1 := f.pop().(t_long)
	f.push(t_long(value1 >> uint(value2 & 0x3F))) // low 6 bits
	return 1, true
}

/*124 (0X7C)*/
func IUSHR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(uint32(value1) >> uint(value2 & 0x1F))) // low 5 bits
	return 1, true
}

/*125 (0X7D)*/
func LUSHR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)

	value1 := f.pop().(t_long)
	f.push(t_long(uint64(value1) >> uint(value2 & 0x3F))) // low 6 bits
	return 1, true
}

/*126 (0X7E)*/
func IAND(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(value1 & value2))
	return 1, true
}

/*127 (0X7F)*/
func LAND(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_long)
	value1 := f.pop().(t_long)
	f.push(t_long(value1 & value2))
	return 1, true
}

/*128 (0X80)*/
func IOR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(value1 | value2))
	return 1, true
}

/*129 (0X81)*/
func LOR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_long)
	value1 := f.pop().(t_long)
	f.push(t_long(value1 | value2))
	return 1, true
}

/*130 (0X82)*/
func IXOR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	f.push(t_int(value1 ^ value2))
	return 1, true
}

/*131 (0X83)*/
func LXOR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value2 := f.pop().(t_long)
	value1 := f.pop().(t_long)
	f.push(t_long(value1 ^ value2))
	return 1, true
}

/*132 (0X84)*/
func IINC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := m.code[f.pc+1]
	const_value := m.code[f.pc+1]
	value := f.loadVar(uint(index)).(t_int)
	f.storeVar(uint(index), value + t_int(const_value))
	return 3, true
}


/*133 (0X85)*/
func I2L(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	value := f.pop().(t_int)
	f.push(t_long(value))
	return 1, true
}

/*134 (0X86)*/
func I2F(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*135 (0X87)*/
func I2D(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*136 (0X88)*/
func L2I(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*137 (0X89)*/
func L2F(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*138 (0X8A)*/
func L2D(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*139 (0X8B)*/
func F2I(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*140 (0X8C)*/
func F2L(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*141 (0X8D)*/
func F2D(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*142 (0X8E)*/
func D2I(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*143 (0X8F)*/
func D2L(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*144 (0X90)*/
func D2F(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*145 (0X91)*/
func I2B(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*146 (0X92)*/
func I2C(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*147 (0X93)*/
func I2S(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}


/*148 (0X94)*/
func LCMP(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*149 (0X95)*/
func FCMPL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*150 (0X96)*/
func FCMPG(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*151 (0X97)*/
func DCMPL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*152 (0X98)*/
func DCMPG(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*153 (0X99)*/
func IFEQ(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(t_int)
	if value  == 0 {
		return int(offset), true
	}
	return 3, true
}

/*154 (0X9A)*/
func IFNE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(t_int)
	if value != 0 {
		return int(offset), true
	}
	return 3, true
}

/*155 (0X9B)*/
func IFLT(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(t_int)
	if value < 0 {
		return int(offset), true
	}
	return 3, true
}

/*156 (0X9C)*/
func IFGE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(t_int)
	if value >= 0 {
		return int(offset), true
	}
	return 3, true
}

/*157 (0X9D)*/
func IFGT(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(t_int)
	if value > 0 {
		return int(offset), true
	}
	return 3, true
}

/*158 (0X9E)*/
func IFLE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(t_int)
	if value <= 0 {
		return int(offset), true
	}
	return 3, true
}

/*159 (0X9F)*/
func IF_ICMPEQ(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	if value1 == value2 {
		return int(offset), true
	}
	return 3, true
}

/*160 (0XA0)*/
func IF_ICMPNE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	if value1 != value2 {
		return int(offset), true
	}
	return 3, true
}

/*161 (0XA1)*/
func IF_ICMPLT(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	if value1 < value2 {
		return int(offset), true
	}
	return 3, true
}

/*162 (0XA2)*/
func IF_ICMPGE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	if value1 >= value2 {
		return int(offset), true
	}
	return 3, true
}

/*163 (0XA3)*/
func IF_ICMPGT(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	if value1 > value2 {
		return int(offset), true
	}
	return 3, true
}

/*164 (0XA4)*/
func IF_ICMPLE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(t_int)
	value1 := f.pop().(t_int)
	if value1 <= value2 {
		return int(offset), true
	}
	return 3, true
}

/*165 (0XA5)*/
func IF_ACMPEQ(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(t_object)
	value1 := f.pop().(t_object)
	if value1 == value2 {
		return int(offset), true
	}
	return 3, true
}

/*166 (0XA6)*/
func IF_ACMPNE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(t_object)
	value1 := f.pop().(t_object)
	if value1 != value2 {
		return int(offset), true
	}
	return 3, true
}


/*167 (0XA7)*/
func GOTO(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*168 (0XA8)*/
func JSR(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*169 (0XA9)*/
func RET(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*170 (0XAA)*/
func TABLESWITCH(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*171 (0XAB)*/
func LOOKUPSWITCH(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*172 (0XAC)*/
func IRETURN(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	t.vmStack.pop()
	// return value
	f.passReturn(t.vmStack.peek())
	return 1, false
}

/*173 (0XAD)*/
func LRETURN(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	t.vmStack.pop()
	// return value
	f.passReturn(t.vmStack.peek())
	return 1, false
}

/*174 (0XAE)*/
func FRETURN(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	t.vmStack.pop()
	// return value
	f.passReturn(t.vmStack.peek())
	return 1, false
}

/*175 (0XAF)*/
func DRETURN(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	t.vmStack.pop()
	// return value
	f.passReturn(t.vmStack.peek())
	return 1, false
}

/*176 (0XB0)*/
func ARETURN(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	t.vmStack.pop()
	// return value
	f.passReturn(t.vmStack.peek())
	return 1, false
}

/*177 (0XB1)*/
func RETURN(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	t.vmStack.pop()
	return 1, false
}


/*178 (0XB2)*/
func GETSTATIC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*179 (0XB3)*/
func PUTSTATIC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*180 (0XB4)*/
func GETFIELD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	jreference := f.pop().(t_object)

	f.push((*jreference).getField(c, index))
	return 3, true
}

/*181 (0XB5)*/
func PUTFIELD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])

	value := f.pop()
	jreference := f.pop().(t_object)

	(*jreference).putField(c, index, value)
	return 3, true
}

/*182 (0XB6)*/
func INVOKEVIRTUAL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*183 (0XB7)*/
func INVOKESPECIAL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])

	method := f.method.class.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo).method
	frame := NewStackFrame(method)
	// pass parameters
	f.passParameters(frame)

	t.vmStack.push(frame)
	return 3, false
}

/*184 (0XB8)*/
func INVOKESTATIC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])

	methodRef := f.method.class.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo)
	method := methodRef.method
	if method.isNative() {
		GVM_print(f.pop().(java_lang_string))
		return 3, true
	}
	frame := NewStackFrame(method)
	// pass parameters
	f.passParameters(frame)

	t.vmStack.push(frame)
	return 3, false
}

/*185 (0XB9)*/
func INVOKEINTERFACE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*186 (0XBA)*/
func INVOKEDYNAMIC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*187 (0XBB)*/
func NEW(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	class := c.constantPool[index].resolve().(*RuntimeConstantClassInfo).class
	jreference := class.new()
	f.push(jreference)
	return 3, true
}

/*188 (0XBC)*/
func NEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	atype := uint8(m.code[f.pc+1])
	count := f.pop().(t_int)
	jreference := t_array(&JavaArray{atype: atype, length: count})
	jreference.elements = make([]t_any, count)
	f.push(jreference)
	return 2, true
}

/*189 (0XBD)*/
func ANEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	count := f.pop().(t_int)

	classInfo := c.constantPool[index].resolve().(*RuntimeConstantClassInfo)
	jreference := newArray(classInfo.class, count)
	f.push(jreference)
	return 3, true
}

/*190 (0XBE)*/
func ARRAYLENGTH(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	f.push(f.pop().(t_array).length)
	return 1, true
}

/*191 (0XBF)*/
func ATHROW(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*192 (0XC0)*/
func CHECKCAST(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*193 (0XC1)*/
func INSTANCEOF(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*194 (0XC2)*/
func MONITORENTER(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*195 (0XC3)*/
func MONITOREXIT(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}


/*196 (0XC4)*/
func WIDE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*197 (0XC5)*/
func MULTIANEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*198 (0XC6)*/
func IFNULL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*199 (0XC7)*/
func IFNONNULL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*200 (0XC8)*/
func GOTO_W(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*201 (0XC9)*/
func JSR_W(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}


/*202 (0XCA)*/
func BREAKPOINT(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*254 (0XFE)*/
func IMPDEP1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}

/*255 (0XFF)*/
func IMPDEP2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) (int, bool) {
	panic(fmt.Sprintf("Not implemented for opcode %s\n", opcode))
}


