// e₀²=-1
// e₁²=0
// e₂²=1
package r111

import (
	"fmt"
)

type V1 struct {
	E0 float64
	E1 float64
	E2 float64
}

func (v V1) String() string {
	return fmt.Sprintf("(%ve₀ + %ve₁ + %ve₂)", v.E0, v.E1, v.E2)
}
func (v V1) MulS(s float64) V1 {
	return V1{
		E0: v.E0 * s,
		E1: v.E1 * s,
		E2: v.E2 * s,
	}
}
func (v V1) MulV1(w V1) (float64, V1) {
	r0 := -1*v.E0*v.E0 + 0*v.E1*v.E1 + 1*v.E2*v.E2
	r1 := V2{
		E01: 1*v.E0*v.E1 + -1*v.E1*v.E0,
		E02: 1*v.E0*v.E2 + -1*v.E2*v.E0,
		E12: 1*v.E1*v.E2 + -1*v.E2*v.E1,
	}
	return r0, r1
}
func (v V1) MulV2(w V2) (V0, V2) {
	r0 := V1{
		E0: 0*v.E1*v.E01 + -1*v.E2*v.E02,
		E1: -1*v.E0*v.E01 + -1*v.E2*v.E12,
		E2: -1*v.E0*v.E02 + 0*v.E1*v.E12,
	}
	r1 := V3{
		E012: 1*v.E0*v.E12 + -1*v.E1*v.E02 + 1*v.E2*v.E01,
	}
	return r0, r1
}
func (v V1) MulV3(w V3) V1 {
	r0 := V2{
		E01: 1 * v.E2 * v.E012,
		E02: 0 * v.E1 * v.E012,
		E12: -1 * v.E0 * v.E012,
	}
	return r0
}

type V2 struct {
	E01 float64
	E02 float64
	E12 float64
}

func (v V2) String() string {
	return fmt.Sprintf("(%ve₀₁ + %ve₀₂ + %ve₁₂)", v.E01, v.E02, v.E12)
}
func (v V2) MulS(s float64) V2 {
	return V2{
		E01: v.E01 * s,
		E02: v.E02 * s,
		E12: v.E12 * s,
	}
}
func (v V2) MulV1(w V1) (V0, V2) {
	r0 := V1{
		E0: 0*v.E01*v.E1 + 1*v.E02*v.E2,
		E1: 1*v.E01*v.E0 + 1*v.E12*v.E2,
		E2: 1*v.E02*v.E0 + 0*v.E12*v.E1,
	}
	r1 := V3{
		E012: 1*v.E01*v.E2 + -1*v.E02*v.E1 + 1*v.E12*v.E0,
	}
	return r0, r1
}
func (v V2) MulV2(w V2) (float64, V1) {
	r0 := 0*v.E01*v.E01 + 1*v.E02*v.E02 + 0*v.E12*v.E12
	r1 := V2{
		E01: -1*v.E02*v.E12 + 1*v.E12*v.E02,
		E02: 0*v.E01*v.E12 + 0*v.E12*v.E01,
		E12: 1*v.E01*v.E02 + -1*v.E02*v.E01,
	}
	return r0, r1
}
func (v V2) MulV3(w V3) V0 {
	r0 := V1{
		E0: 0 * v.E12 * v.E012,
		E1: -1 * v.E02 * v.E012,
		E2: 0 * v.E01 * v.E012,
	}
	return r0
}

type V3 struct {
	E012 float64
}

func (v V3) String() string {
	return fmt.Sprintf("(%ve₀₁₂)", v.E012)
}
func (v V3) MulS(s float64) V3 {
	return V3{
		E012: v.E012 * s,
	}
}
func (v V3) MulV1(w V1) V1 {
	r0 := V2{
		E01: 1 * v.E012 * v.E2,
		E02: 0 * v.E012 * v.E1,
		E12: -1 * v.E012 * v.E0,
	}
	return r0
}
func (v V3) MulV2(w V2) V0 {
	r0 := V1{
		E0: 0 * v.E012 * v.E12,
		E1: -1 * v.E012 * v.E02,
		E2: 0 * v.E012 * v.E01,
	}
	return r0
}
func (v V3) MulV3(w V3) float64 {
	r0 := 0 * v.E012 * v.E012
	return r0
}
