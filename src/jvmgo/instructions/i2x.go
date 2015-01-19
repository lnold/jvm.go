package instructions

import "jvmgo/rtda"

// Convert int to byte
type i2b struct {}
func (self *i2b) fetchOperands(bcr *BytecodeReader) {}
func (self *i2b) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    b := int32(int8(i))
    stack.PushInt(b)
}

// Convert int to char
type i2c struct {}
func (self *i2c) fetchOperands(bcr *BytecodeReader) {}
func (self *i2c) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    c := int32(uint16(i))
    stack.PushInt(c)
}

// Convert int to double
type i2d struct {}
func (self *i2d) fetchOperands(bcr *BytecodeReader) {}
func (self *i2d) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    d := float64(i)
    stack.PushDouble(d)
}