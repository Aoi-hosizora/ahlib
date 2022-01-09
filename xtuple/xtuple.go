//go:build go1.18
// +build go1.18

package xtuple

import (
	"fmt"
)

// Tuple represents a tuple, or 2-pair data structure.
type Tuple[T1, T2 any] struct {
	Item1 T1
	Item2 T2
}

// NewTuple creates a Tuple structure from given parameters.
func NewTuple[T1, T2 any](item1 T1, item2 T2) Tuple[T1, T2] {
	return Tuple[T1, T2]{item1, item2}
}

// NewTuplePtr creates a Tuple pointer from given parameters.
func NewTuplePtr[T1, T2 any](item1 T1, item2 T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{item1, item2}
}

// String implements fmt.Stringer for Tuple.
func (t Tuple[T1, T2]) String() string {
	return fmt.Sprintf("[%v, %v]", t.Item1, t.Item2)
}

// Triple represents a triple, or 3-pair data structure.
type Triple[T1, T2, T3 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
}

// NewTriple creates a Triple structure from given parameters.
func NewTriple[T1, T2, T3 any](item1 T1, item2 T2, item3 T3) Triple[T1, T2, T3] {
	return Triple[T1, T2, T3]{item1, item2, item3}
}

// NewTriplePtr creates a Triple pointer from given parameters.
func NewTriplePtr[T1, T2, T3 any](item1 T1, item2 T2, item3 T3) *Triple[T1, T2, T3] {
	return &Triple[T1, T2, T3]{item1, item2, item3}
}

// String implements fmt.Stringer for Triple.
func (t Triple[T1, T2, T3]) String() string {
	return fmt.Sprintf("[%v, %v, %v]", t.Item1, t.Item2, t.Item3)
}

// Quadruple represents a quadruple, or 4-pair data structure.
type Quadruple[T1, T2, T3, T4 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
}

// NewQuadruple creates a Quadruple structure from given parameters.
func NewQuadruple[T1, T2, T3, T4 any](item1 T1, item2 T2, item3 T3, item4 T4) Quadruple[T1, T2, T3, T4] {
	return Quadruple[T1, T2, T3, T4]{item1, item2, item3, item4}
}

// NewQuadruplePtr creates a Quadruple pointer from given parameters.
func NewQuadruplePtr[T1, T2, T3, T4 any](item1 T1, item2 T2, item3 T3, item4 T4) *Quadruple[T1, T2, T3, T4] {
	return &Quadruple[T1, T2, T3, T4]{item1, item2, item3, item4}
}

// String implements fmt.Stringer for Quadruple.
func (t Quadruple[T1, T2, T3, T4]) String() string {
	return fmt.Sprintf("[%v, %v, %v, %v]", t.Item1, t.Item2, t.Item3, t.Item4)
}

// Quintuple represents a quintuple, or 5-pair data structure.
type Quintuple[T1, T2, T3, T4, T5 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
	Item5 T5
}

// NewQuintuple creates a Quintuple structure from given parameters.
func NewQuintuple[T1, T2, T3, T4, T5 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5) Quintuple[T1, T2, T3, T4, T5] {
	return Quintuple[T1, T2, T3, T4, T5]{item1, item2, item3, item4, item5}
}

// NewQuintuplePtr creates a Quintuple pointer from given parameters.
func NewQuintuplePtr[T1, T2, T3, T4, T5 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5) *Quintuple[T1, T2, T3, T4, T5] {
	return &Quintuple[T1, T2, T3, T4, T5]{item1, item2, item3, item4, item5}
}

// String implements fmt.Stringer for Quintuple.
func (t Quintuple[T1, T2, T3, T4, T5]) String() string {
	return fmt.Sprintf("[%v, %v, %v, %v, %v]", t.Item1, t.Item2, t.Item3, t.Item4, t.Item5)
}

// Sextuple represents a sextuple, or 6-pair data structure.
type Sextuple[T1, T2, T3, T4, T5, T6 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
	Item5 T5
	Item6 T6
}

// NewSextuple creates a Sextuple structure from given parameters.
func NewSextuple[T1, T2, T3, T4, T5, T6 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5, item6 T6) Sextuple[T1, T2, T3, T4, T5, T6] {
	return Sextuple[T1, T2, T3, T4, T5, T6]{item1, item2, item3, item4, item5, item6}
}

// NewSextuplePtr creates a Sextuple pointer from given parameters.
func NewSextuplePtr[T1, T2, T3, T4, T5, T6 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5, item6 T6) *Sextuple[T1, T2, T3, T4, T5, T6] {
	return &Sextuple[T1, T2, T3, T4, T5, T6]{item1, item2, item3, item4, item5, item6}
}

// String implements fmt.Stringer for Sextuple.
func (t Sextuple[T1, T2, T3, T4, T5, T6]) String() string {
	return fmt.Sprintf("[%v, %v, %v, %v, %v, %v]", t.Item1, t.Item2, t.Item3, t.Item4, t.Item5, t.Item6)
}

// Septuple represents a septuple, or 7-pair data structure.
type Septuple[T1, T2, T3, T4, T5, T6, T7 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
	Item5 T5
	Item6 T6
	Item7 T7
}

// NewSeptuple creates a Septuple structure from given parameters.
func NewSeptuple[T1, T2, T3, T4, T5, T6, T7 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5, item6 T6, item7 T7) Septuple[T1, T2, T3, T4, T5, T6, T7] {
	return Septuple[T1, T2, T3, T4, T5, T6, T7]{item1, item2, item3, item4, item5, item6, item7}
}

// NewSeptuplePtr creates a Septuple pointer from given parameters.
func NewSeptuplePtr[T1, T2, T3, T4, T5, T6, T7 any](item1 T1, item2 T2, item3 T3, item4 T4, item5 T5, item6 T6, item7 T7) *Septuple[T1, T2, T3, T4, T5, T6, T7] {
	return &Septuple[T1, T2, T3, T4, T5, T6, T7]{item1, item2, item3, item4, item5, item6, item7}
}

// String implements fmt.Stringer for Septuple.
func (t Septuple[T1, T2, T3, T4, T5, T6, T7]) String() string {
	return fmt.Sprintf("[%v, %v, %v, %v, %v, %v, %v]", t.Item1, t.Item2, t.Item3, t.Item4, t.Item5, t.Item6, t.Item7)
}
