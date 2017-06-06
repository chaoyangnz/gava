package gvm

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	ReadInfo(reader *ClassReader)
}

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	attributeNameIndex   u2
	attributeLength      u4
	maxStack             u2
	maxLocals            u2
	codeLength           u4
	code                 []u1               //u4 code_length
	exceptionTableLength u2
	exceptionTable       []ExceptionTableEntry //u2 exception_table_length
	attributesCount      u2
	attributes           []AttributeInfo  //u2 attributes_count
}

func (this *CodeAttribute) ReadInfo(reader *ClassReader) {
	this.maxStack = reader.readU2()
	this.maxLocals = reader.readU2()
	codeLength := reader.readU4()
	this.codeLength = codeLength
	this.code = reader.readU1s(uint32(codeLength))
	exceptionTableLength := reader.readU2()
	this.exceptionTableLength = exceptionTableLength
	this.exceptionTable = make([]ExceptionTableEntry, exceptionTableLength)
	for i := u2(0); i < exceptionTableLength; i++ {
		exceptionTable := ExceptionTableEntry{}
		exceptionTable.ReadInfo(reader)
		this.exceptionTable[i] = exceptionTable
	}
	attributesCount := reader.readU2()
	this.attributesCount = attributesCount
	this.attributes = make([]AttributeInfo, attributesCount)
	for i := u2(0); i < attributesCount; i++ {
		reader.readAttribute(&this.attributes[i])
	}
}

type ExceptionTableEntry struct {
	startPc   u2
	endPc     u2
	handlerPc u2
	catchType u2
}

func (this *ExceptionTableEntry) ReadInfo(reader *ClassReader) {
	this.startPc = reader.readU2()
	this.endPc = reader.readU2()
	this.handlerPc = reader.readU2()
	this.catchType = reader.readU2()
}

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	attributeNameIndex    u2
	attributeLength       u4
	lineNumberTableLength u2
	lineNumberTable       []LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    u2
	lineNumber u2
}

func (this *LineNumberTableAttribute) ReadInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readU2()
	this.lineNumberTableLength = lineNumberTableLength
	this.lineNumberTable = make([]LineNumberTableEntry, lineNumberTableLength)
	for i := u2(0); i < lineNumberTableLength; i++ {
		this.lineNumberTable[i] = LineNumberTableEntry{reader.readU2(), reader.readU2()}
	}
}

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
 */

type SourceFileAttribue struct {
	attributeNameIndex u2
	attributeLength    u4
	sourceFileIndex u2
}

func (this *SourceFileAttribue) ReadInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.readU2()
}

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
 */
type LocalVariableTableAttribute struct {
	attributeNameIndex       u2
	attributeLength          u4
	localVariableTableLength u2
	localVariableTable       []LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc             u2
	length              u2
	nameIndex           u2
	descriptorIndex     u2
	index               u2
}

func (this *LocalVariableTableAttribute) ReadInfo(reader *ClassReader) {
	localVariableTableLength := reader.readU2()
	this.localVariableTableLength = localVariableTableLength
	this.localVariableTable = make([]LocalVariableTableEntry, localVariableTableLength)
	for i := u2(0); i < localVariableTableLength; i++ {
		this.localVariableTable[i] = LocalVariableTableEntry{
			reader.readU2(),
			reader.readU2(),
			reader.readU2(),
			reader.readU2(),
			reader.readU2()}
	}
}

