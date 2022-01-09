//go:build go1.18
// +build go1.18

package xtuple

// IfThen returns value if condition is true, otherwise returns the default value of type T.
func IfThen[T any](condition bool, value T) (v T) {
	if condition {
		v = value
	}
	return v
}

// IfThenElse returns value1 if condition is true, otherwise returns value2.
func IfThenElse[T any](condition bool, value1, value2 T) T {
	if condition {
		return value1
	}
	return value2
}

// TupleItem1 returns given item1 in given tuple.
func TupleItem1[T1, T2 any](item1 T1, _ T2) T1 {
	return item1
}

// TupleItem2 returns given item2 in given tuple.
func TupleItem2[T1, T2 any](_ T1, item2 T2) T2 {
	return item2
}

// TripleItem1 returns given item1 in given triple.
func TripleItem1[T1, T2, T3 any](item1 T1, _ T2, _ T3) T1 {
	return item1
}

// TripleItem2 returns given item2 in given triple.
func TripleItem2[T1, T2, T3 any](_ T1, item2 T2, _ T3) T2 {
	return item2
}

// TripleItem3 returns given item3 in given triple.
func TripleItem3[T1, T2, T3 any](_ T1, _ T2, item3 T3) T3 {
	return item3
}

// QuadrupleItem1 returns given item1 in given quadruple.
func QuadrupleItem1[T1, T2, T3, T4 any](item1 T1, _ T2, _ T3, _ T4) T1 {
	return item1
}

// QuadrupleItem2 returns given item2 in given quadruple.
func QuadrupleItem2[T1, T2, T3, T4 any](_ T1, item2 T2, _ T3, _ T4) T2 {
	return item2
}

// QuadrupleItem3 returns given item3 in given quadruple.
func QuadrupleItem3[T1, T2, T3, T4 any](_ T1, _ T2, item3 T3, _ T4) T3 {
	return item3
}

// QuadrupleItem4 returns given item4 in given quadruple.
func QuadrupleItem4[T1, T2, T3, T4 any](_ T1, _ T2, _ T3, item4 T4) T4 {
	return item4
}

// QuintupleItem1 returns given item1 in given quintuple.
func QuintupleItem1[T1, T2, T3, T4, T5 any](item1 T1, _ T2, _ T3, _ T4, _ T5) T1 {
	return item1
}

// QuintupleItem2 returns given item2 in given quintuple.
func QuintupleItem2[T1, T2, T3, T4, T5 any](_ T1, item2 T2, _ T3, _ T4, _ T5) T2 {
	return item2
}

// QuintupleItem3 returns given item3 in given quintuple.
func QuintupleItem3[T1, T2, T3, T4, T5 any](_ T1, _ T2, item3 T3, _ T4, _ T5) T3 {
	return item3
}

// QuintupleItem4 returns given item4 in given quintuple.
func QuintupleItem4[T1, T2, T3, T4, T5 any](_ T1, _ T2, _ T3, item4 T4, _ T5) T4 {
	return item4
}

// QuintupleItem5 returns given item5 in given quintuple.
func QuintupleItem5[T1, T2, T3, T4, T5 any](_ T1, _ T2, _ T3, _ T4, item5 T5) T5 {
	return item5
}

// SextupleItem1 returns given item1 in given sextuple.
func SextupleItem1[T1, T2, T3, T4, T5, T6 any](item1 T1, _ T2, _ T3, _ T4, _ T5, _ T6) T1 {
	return item1
}

// SextupleItem2 returns given item2 in given sextuple.
func SextupleItem2[T1, T2, T3, T4, T5, T6 any](_ T1, item2 T2, _ T3, _ T4, _ T5, _ T6) T2 {
	return item2
}

// SextupleItem3 returns given item3 in given sextuple.
func SextupleItem3[T1, T2, T3, T4, T5, T6 any](_ T1, _ T2, item3 T3, _ T4, _ T5, _ T6) T3 {
	return item3
}

// SextupleItem4 returns given item4 in given sextuple.
func SextupleItem4[T1, T2, T3, T4, T5, T6 any](_ T1, _ T2, _ T3, item4 T4, _ T5, _ T6) T4 {
	return item4
}

// SextupleItem5 returns given item5 in given sextuple.
func SextupleItem5[T1, T2, T3, T4, T5, T6 any](_ T1, _ T2, _ T3, _ T4, item5 T5, _ T6) T5 {
	return item5
}

// SextupleItem6 returns given item6 in given sextuple.
func SextupleItem6[T1, T2, T3, T4, T5, T6 any](_ T1, _ T2, _ T3, _ T4, _ T5, item6 T6) T6 {
	return item6
}

// SeptupleItem1 returns given item1 in given septuple.
func SeptupleItem1[T1, T2, T3, T4, T5, T6, T7 any](item1 T1, _ T2, _ T3, _ T4, _ T5, _ T6, _ T7) T1 {
	return item1
}

// SeptupleItem2 returns given item2 in given septuple.
func SeptupleItem2[T1, T2, T3, T4, T5, T6, T7 any](_ T1, item2 T2, _ T3, _ T4, _ T5, _ T6, _ T7) T2 {
	return item2
}

// SeptupleItem3 returns given item3 in given septuple.
func SeptupleItem3[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, item3 T3, _ T4, _ T5, _ T6, _ T7) T3 {
	return item3
}

// SeptupleItem4 returns given item4 in given septuple.
func SeptupleItem4[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, _ T3, item4 T4, _ T5, _ T6, _ T7) T4 {
	return item4
}

// SeptupleItem5 returns given item5 in given septuple.
func SeptupleItem5[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, _ T3, _ T4, item5 T5, _ T6, _ T7) T5 {
	return item5
}

// SeptupleItem6 returns given item6 in given septuple.
func SeptupleItem6[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, _ T3, _ T4, _ T5, item6 T6, _ T7) T6 {
	return item6
}

// SeptupleItem7 returns given item7 in given septuple.
func SeptupleItem7[T1, T2, T3, T4, T5, T6, T7 any](_ T1, _ T2, _ T3, _ T4, _ T5, _ T6, item7 T7) T7 {
	return item7
}
