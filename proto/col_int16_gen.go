// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColInt16 represents Int16 column.
type ColInt16 []int16

// Compile-time assertions for ColInt16.
var (
	_ ColInput  = ColInt16{}
	_ ColResult = (*ColInt16)(nil)
	_ Column    = (*ColInt16)(nil)
)

// Type returns ColumnType of Int16.
func (ColInt16) Type() ColumnType {
	return ColumnTypeInt16
}

// Rows returns count of rows in column.
func (c ColInt16) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColInt16) Row(i int) int16 {
	return c[i]
}

// Append int16 to column.
func (c *ColInt16) Append(v int16) {
	*c = append(*c, v)
}

// AppendArr appends slice of int16 to column.
func (c *ColInt16) AppendArr(v []int16) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt16) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Int16 .
func (c *ColInt16) LowCardinality() *ColLowCardinalityOf[int16] {
	return &ColLowCardinalityOf[int16]{
		index: c,
	}
}

// Array is helper that creates Array of int16.
func (c *ColInt16) Array() *ColArrOf[int16] {
	return &ColArrOf[int16]{
		Data: c,
	}
}

// NewArrInt16 returns new Array(Int16).
func NewArrInt16() *ColArr {
	return &ColArr{
		Data: new(ColInt16),
	}
}

// AppendInt16 appends slice of int16 to Array(Int16).
func (c *ColArr) AppendInt16(data []int16) {
	d := c.Data.(*ColInt16)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
