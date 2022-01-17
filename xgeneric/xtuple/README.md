# xtuple

## Dependencies

+ None

## Documents

### Types

+ `type Tuple[T1, T2 any] struct`
+ `type Triple[T1, T2, T3 any] struct`
+ `type Quadruple[T1, T2, T3, T4 any] struct`
+ `type Quintuple[T1, T2, T3, T4, T5 any] struct`
+ `type Sextuple[T1, T2, T3, T4, T5, T6 any] struct`
+ `type Septuple[T1, T2, T3, T4, T5, T6, T7 any] struct`

### Variables

+ None

### Constants

+ None

### Functions

+ `func NewTuple[T1, T2 any](item1 T1, item2 T2) Tuple[T1, T2]`
+ `func NewTuplePtr[T1, T2 any](item1 T1, item2 T2) *Tuple[T1, T2]`
+ `func NewTriple[T1, T2, T3 any](item1 T1, item2 T2, item3 T3) Triple[T1, T2, T3]`
+ `func NewTriplePtr[T1, T2, T3 any](item1 T1, item2 T2, item3 T3) *Triple[T1, T2, T3]`
+ `func NewQuadruple[T1, T2, T3, T4 any](item1 T1, item2 T2, item3 T3, item4 T4) Quadruple[T1, T2, T3, T4]`
+ `func NewQuadruplePtr[T1, T2, T3, T4 any](item1 T1, item2 T2, item3 T3, item4 T4) *Quadruple[T1, T2, T3, T4]`
+ `func NewQuintuple[T1, T2, T3, T4, T5 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5) Quintuple[T1, T2, T3, T4, T5]`
+ `func NewQuintuplePtr[T1, T2, T3, T4, T5 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5) *Quintuple[T1, T2, T3, T4, T5]`
+ `func NewSextuple[T1, T2, T3, T4, T5, T6 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5, item6 T6) Sextuple[T1, T2, T3, T4, T5, T6]`
+ `func NewSextuplePtr[T1, T2, T3, T4, T5, T6 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5, item6 T6) *Sextuple[T1, T2, T3, T4, T5, T6]`
+ `func NewSeptuple[T1, T2, T3, T4, T5, T6, T7 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5, item6 T6, item7 T7) Septuple[T1, T2, T3, T4, T5, T6, T7]`
+ `func NewSeptuplePtr[T1, T2, T3, T4, T5, T6, T7 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5, item6 T6, item7 T7) *Septuple[T1, T2, T3, T4, T5, T6, T7]`
+ `func IfThen[T any](condition bool, value T) T`
+ `func IfThenElse[T any](condition bool, value1, value2 T) T`
+ `func ValPtr[T any](t T) *T`
+ `func PtrVal[T any](t *T, o T) T`
+ `func TupleItem1[T1, T2 any](item1 T1, _ T2) T1`
+ `func TupleItem2[T1, T2 any](_ T1, item2 T2) T2`
+ `func TripleItem1[T1, T2, T3 any](item1 T1, _ T2, _ T3) T1`
+ `func TripleItem2[T1, T2, T3 any](_ T1, item2 T2, _ T3) T2`
+ `func TripleItem3[T1, T2, T3 any](_ T1, _ T2, item3 T3) T3`
+ `func QuadrupleItem1[T1, T2, T3, T4 any](item1 T1, _ T2, _ T3, _ T4) T1`
+ `func QuadrupleItem2[T1, T2, T3, T4 any](_ T1, item2 T2, _ T3, _ T4) T2`
+ `func QuadrupleItem3[T1, T2, T3, T4 any](_ T1, _ T2, item3 T3, _ T4) T3`
+ `func QuadrupleItem4[T1, T2, T3, T4 any](_ T1, _ T2, _ T3, item4 T4) T4`
+ `func QuintupleItem1[T1, T2, T3, T4, T5 any](item1 T1, _ T2, _ T3, _ T4, _ T5) T1`
+ `func QuintupleItem2[T1, T2, T3, T4, T5 any](_ T1, item2 T2, _ T3, _ T4, _ T5) T2`
+ `func QuintupleItem3[T1, T2, T3, T4, T5 any](_ T1, _ T2, item3 T3, _ T4, _ T5) T3`
+ `func QuintupleItem4[T1, T2, T3, T4, T5 any](_ T1, _ T2, _ T3, item4 T4, _ T5) T4`
+ `func QuintupleItem5[T1, T2, T3, T4, T5 any](_ T1, _ T2, _ T3, _ T4, item5 T5) T5`
+ `func SextupleItem1[T1, T2, T3, T4, T5, T6 any](item1 T1, _ T2, _ T3, _ T4, _ T5, _ T6) T1`
+ `func SextupleItem2[T1, T2, T3, T4, T5, T6 any](_ T1, item2 T2, _ T3, _ T4, _ T5, _ T6) T2`
+ `func SextupleItem3[T1, T2, T3, T4, T5, T6 any](_ T1, _ T2, item3 T3, _ T4, _ T5, _ T6) T3`
+ `func SextupleItem4[T1, T2, T3, T4, T5, T6 any](_ T1, _ T2, _ T3, item4 T4, _ T5, _ T6) T4`
+ `func SextupleItem5[T1, T2, T3, T4, T5, T6 any](_ T1, _ T2, _ T3, _ T4, item5 T5, _ T6) T5`
+ `func SextupleItem6[T1, T2, T3, T4, T5, T6 any](_ T1, _ T2, _ T3, _ T4, _ T5, item6 T6) T6`
+ `func SeptupleItem1[T1, T2, T3, T4, T5, T6, T7 any](item1 T1, _ T2, _ T3, _ T4, _ T5, _ T6, _ T7) T1`
+ `func SeptupleItem2[T1, T2, T3, T4, T5, T6, T7 any](_ T1, item2 T2, _ T3, _ T4, _ T5, _ T6, _ T7) T2`
+ `func SeptupleItem3[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, item3 T3, _ T4, _ T5, _ T6, _ T7) T3`
+ `func SeptupleItem4[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, _ T3, item4 T4, _ T5, _ T6, _ T7) T4`
+ `func SeptupleItem5[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, _ T3, _ T4, item5 T5, _ T6, _ T7) T5`
+ `func SeptupleItem6[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, _ T3, _ T4, _ T5, item6 T6, _ T7) T6`
+ `func SeptupleItem7[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, _ T3, _ T4, _ T5, _ T6, item7 T7) T7`

### Methods

+ `func (t Tuple[T1, T2]) String() string`
+ `func (t Triple[T1, T2, T3]) String() string`
+ `func (t Quadruple[T1, T2, T3, T4]) String() string`
+ `func (t Quintuple[T1, T2, T3, T4, T5]) String() string`
+ `func (t Sextuple[T1, T2, T3, T4, T5, T6]) String() string`
+ `func (t Septuple[T1, T2, T3, T4, T5]) String() string`
