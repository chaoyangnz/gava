package classfile

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
	// attributeNameIndex uint16
	// attributeLength    uint32
	maxStack   uint16
	maxLocals  uint16
	code       []uint8          //u4 code_length
	exceptions []ExceptionTable //u2 exception_table_length
	attributes []AttributeInfo  //u2 attributes_count
}

func (this *CodeAttribute) ReadInfo(reader *ClassReader) {
	this.maxStack = reader.ReadUint16()
	this.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	this.code = reader.ReadBytes(int(codeLength))
	exceptionTableLength := reader.ReadUint16()
	this.exceptions = make([]ExceptionTable, exceptionTableLength)
	for i := uint16(0); i < exceptionTableLength; i++ {
		exceptionTable := ExceptionTable{}
		exceptionTable.ReadInfo(reader)
		this.exceptions[i] = exceptionTable
	}
	attributesCount := reader.ReadUint16()
	this.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		this.attributes[i] = reader.ReadAsAttribute()
	}
}

type ExceptionTable struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (this *ExceptionTable) ReadInfo(reader *ClassReader) {
	this.startPc = reader.ReadUint16()
	this.endPc = reader.ReadUint16()
	this.handlerPc = reader.ReadUint16()
	this.catchType = reader.ReadUint16()
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
	lineNumberTables []LineNumberTable
}

type LineNumberTable struct {
	startPc    uint16
	lineNumber uint16
}

func (this *LineNumberTableAttribute) ReadInfo(reader *ClassReader) {
	lineNumberTableLength := reader.ReadUint16()
	this.lineNumberTables = make([]LineNumberTable, lineNumberTableLength)
	for i := uint16(0); i < lineNumberTableLength; i++ {
		this.lineNumberTables[i] = LineNumberTable{reader.ReadUint16(), reader.ReadUint16()}
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
	sourceFileIndex uint16
}

func (this *SourceFileAttribue) ReadInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.ReadUint16()
}