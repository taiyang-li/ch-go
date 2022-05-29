// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColDateTime64 represents DateTime64 column.
type ColDateTime64 []DateTime64

// Compile-time assertions for ColDateTime64.
var (
	_ ColInput  = ColDateTime64{}
	_ ColResult = (*ColDateTime64)(nil)
	_ Column    = (*ColDateTime64)(nil)
)

// Type returns ColumnType of DateTime64.
func (ColDateTime64) Type() ColumnType {
	return ColumnTypeDateTime64
}

// Rows returns count of rows in column.
func (c ColDateTime64) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColDateTime64) Row(i int) DateTime64 {
	return c[i]
}

// Append DateTime64 to column.
func (c *ColDateTime64) Append(v DateTime64) {
	*c = append(*c, v)
}

// AppendArr appends slice of DateTime64 to column.
func (c *ColDateTime64) AppendArr(v []DateTime64) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDateTime64) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for DateTime64 .
func (c *ColDateTime64) LowCardinality() *ColLowCardinalityOf[DateTime64] {
	return &ColLowCardinalityOf[DateTime64]{
		index: c,
	}
}

// Array is helper that creates Array of DateTime64.
func (c *ColDateTime64) Array() *ColArrOf[DateTime64] {
	return &ColArrOf[DateTime64]{
		Data: c,
	}
}

// NewArrDateTime64 returns new Array(DateTime64).
func NewArrDateTime64() *ColArr {
	return &ColArr{
		Data: new(ColDateTime64),
	}
}

// AppendDateTime64 appends slice of DateTime64 to Array(DateTime64).
func (c *ColArr) AppendDateTime64(data []DateTime64) {
	d := c.Data.(*ColDateTime64)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
