// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColDecimal256 represents Decimal256 column.
type ColDecimal256 []Decimal256

// Compile-time assertions for ColDecimal256.
var (
	_ ColInput  = ColDecimal256{}
	_ ColResult = (*ColDecimal256)(nil)
	_ Column    = (*ColDecimal256)(nil)
)

// Type returns ColumnType of Decimal256.
func (ColDecimal256) Type() ColumnType {
	return ColumnTypeDecimal256
}

// Rows returns count of rows in column.
func (c ColDecimal256) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColDecimal256) Row(i int) Decimal256 {
	return c[i]
}

// Append Decimal256 to column.
func (c *ColDecimal256) Append(v Decimal256) {
	*c = append(*c, v)
}

// AppendArr appends slice of Decimal256 to column.
func (c *ColDecimal256) AppendArr(v []Decimal256) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDecimal256) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Decimal256 .
func (c *ColDecimal256) LowCardinality() *ColLowCardinalityOf[Decimal256] {
	return &ColLowCardinalityOf[Decimal256]{
		index: c,
	}
}

// Array is helper that creates Array of Decimal256.
func (c *ColDecimal256) Array() *ColArrOf[Decimal256] {
	return &ColArrOf[Decimal256]{
		Data: c,
	}
}

// NewArrDecimal256 returns new Array(Decimal256).
func NewArrDecimal256() *ColArr {
	return &ColArr{
		Data: new(ColDecimal256),
	}
}

// AppendDecimal256 appends slice of Decimal256 to Array(Decimal256).
func (c *ColArr) AppendDecimal256(data []Decimal256) {
	d := c.Data.(*ColDecimal256)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
