// e₀²=1
// e₁²=1
package r002

import (
	"fmt"
)

type V1 struct {
	E0 float64
	E1 float64
}

func (v V1) String() string {
	return fmt.Sprintf("(%ve₀ + %ve₁)", v.E0, v.E1)
}
func (v V1) MulS(s float64) V1 {
	return V1{
		E0: v.E0 * s,
		E1: v.E1 * s,
	}
}
func (v V1) MulV1(w V1) (float64, V1) {
	r0 := 1*v.E0*v.E0 + 1*v.E1*v.E1
	r1 := V2{
		E01: 1*v.E0*v.E1 + -1*v.E1*v.E0,
	}
	return r0, r1
}
func (v V1) MulV2(w V2) V0 {
	r0 := V1{
		E0: -1 * v.E1 * v.E01,
		E1: 1 * v.E0 * v.E01,
	}
	return r0
}

type V2 struct {
	E01 float64
}

func (v V2) String() string {
	return fmt.Sprintf("(%ve₀₁)", v.E01)
}
func (v V2) MulS(s float64) V2 {
	return V2{
		E01: v.E01 * s,
	}
}
func (v V2) MulV1(w V1) V0 {
	r0 := V1{
		E0: 1 * v.E01 * v.E1,
		E1: -1 * v.E01 * v.E0,
	}
	return r0
}
func (v V2) MulV2(w V2) float64 {
	r0 := -1 * v.E01 * v.E01
	return r0
}
