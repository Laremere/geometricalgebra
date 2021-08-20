// e₀²=-1
// e₁²=0
// e₂²=0
// e₃²=1
package r121

import (
	"fmt"
)

type V1 struct {
	E0 float64
	E1 float64
	E2 float64
	E3 float64
}

func (v V1) String() string {
	return fmt.Sprintf("(%ve₀ + %ve₁ + %ve₂ + %ve₃)", v.E0, v.E1, v.E2, v.E3)
}
func (v V1) MulS(s float64) V1 {
	return V1{
		E0: v.E0 * s,
		E1: v.E1 * s,
		E2: v.E2 * s,
		E3: v.E3 * s,
	}
}
func (v V1) MulV1(w V1) (float64, V1) {
	r0 := -1*v.E0*v.E0 + 0*v.E1*v.E1 + 0*v.E2*v.E2 + 1*v.E3*v.E3
	r1 := V2{
		E01: 1*v.E0*v.E1 + -1*v.E1*v.E0,
		E02: 1*v.E0*v.E2 + -1*v.E2*v.E0,
		E12: 1*v.E1*v.E2 + -1*v.E2*v.E1,
		E03: 1*v.E0*v.E3 + -1*v.E3*v.E0,
		E13: 1*v.E1*v.E3 + -1*v.E3*v.E1,
		E23: 1*v.E2*v.E3 + -1*v.E3*v.E2,
	}
	return r0, r1
}
func (v V1) MulV2(w V2) (V0, V2) {
	r0 := V1{
		E0: 0*v.E1*v.E01 + 0*v.E2*v.E02 + -1*v.E3*v.E03,
		E1: -1*v.E0*v.E01 + 0*v.E2*v.E12 + -1*v.E3*v.E13,
		E2: -1*v.E0*v.E02 + 0*v.E1*v.E12 + -1*v.E3*v.E23,
		E3: -1*v.E0*v.E03 + 0*v.E1*v.E13 + 0*v.E2*v.E23,
	}
	r1 := V3{
		E012: 1*v.E0*v.E12 + -1*v.E1*v.E02 + 1*v.E2*v.E01,
		E013: 1*v.E0*v.E13 + -1*v.E1*v.E03 + 1*v.E3*v.E01,
		E023: 1*v.E0*v.E23 + -1*v.E2*v.E03 + 1*v.E3*v.E02,
		E123: 1*v.E1*v.E23 + -1*v.E2*v.E13 + 1*v.E3*v.E12,
	}
	return r0, r1
}
func (v V1) MulV3(w V3) (V1, V3) {
	r0 := V2{
		E01: 0*v.E2*v.E012 + 1*v.E3*v.E013,
		E02: 0*v.E1*v.E012 + 1*v.E3*v.E023,
		E12: -1*v.E0*v.E012 + 1*v.E3*v.E123,
		E03: 0*v.E1*v.E013 + 0*v.E2*v.E023,
		E13: -1*v.E0*v.E013 + 0*v.E2*v.E123,
		E23: -1*v.E0*v.E023 + 0*v.E1*v.E123,
	}
	r1 := V4{
		E0123: 1*v.E0*v.E123 + -1*v.E1*v.E023 + 1*v.E2*v.E013 + -1*v.E3*v.E012,
	}
	return r0, r1
}
func (v V1) MulV4(w V4) V2 {
	r0 := V3{
		E012: -1 * v.E3 * v.E0123,
		E013: 0 * v.E2 * v.E0123,
		E023: 0 * v.E1 * v.E0123,
		E123: -1 * v.E0 * v.E0123,
	}
	return r0
}

type V2 struct {
	E01 float64
	E02 float64
	E12 float64
	E03 float64
	E13 float64
	E23 float64
}

func (v V2) String() string {
	return fmt.Sprintf("(%ve₀₁ + %ve₀₂ + %ve₁₂ + %ve₀₃ + %ve₁₃ + %ve₂₃)", v.E01, v.E02, v.E12, v.E03, v.E13, v.E23)
}
func (v V2) MulS(s float64) V2 {
	return V2{
		E01: v.E01 * s,
		E02: v.E02 * s,
		E12: v.E12 * s,
		E03: v.E03 * s,
		E13: v.E13 * s,
		E23: v.E23 * s,
	}
}
func (v V2) MulV1(w V1) (V0, V2) {
	r0 := V1{
		E0: 0*v.E01*v.E1 + 0*v.E02*v.E2 + 1*v.E03*v.E3,
		E1: 1*v.E01*v.E0 + 0*v.E12*v.E2 + 1*v.E13*v.E3,
		E2: 1*v.E02*v.E0 + 0*v.E12*v.E1 + 1*v.E23*v.E3,
		E3: 1*v.E03*v.E0 + 0*v.E13*v.E1 + 0*v.E23*v.E2,
	}
	r1 := V3{
		E012: 1*v.E01*v.E2 + -1*v.E02*v.E1 + 1*v.E12*v.E0,
		E013: 1*v.E01*v.E3 + -1*v.E03*v.E1 + 1*v.E13*v.E0,
		E023: 1*v.E02*v.E3 + -1*v.E03*v.E2 + 1*v.E23*v.E0,
		E123: 1*v.E12*v.E3 + -1*v.E13*v.E2 + 1*v.E23*v.E1,
	}
	return r0, r1
}
func (v V2) MulV2(w V2) (float64, V1, V3) {
	r0 := 0*v.E01*v.E01 + 0*v.E02*v.E02 + 0*v.E12*v.E12 + 1*v.E03*v.E03 + 0*v.E13*v.E13 + 0*v.E23*v.E23
	r1 := V2{
		E01: 0*v.E02*v.E12 + 0*v.E12*v.E02 + -1*v.E03*v.E13 + 1*v.E13*v.E03,
		E02: 0*v.E01*v.E12 + 0*v.E12*v.E01 + -1*v.E03*v.E23 + 1*v.E23*v.E03,
		E12: 1*v.E01*v.E02 + -1*v.E02*v.E01 + -1*v.E13*v.E23 + 1*v.E23*v.E13,
		E03: 0*v.E01*v.E13 + 0*v.E02*v.E23 + 0*v.E13*v.E01 + 0*v.E23*v.E02,
		E13: 1*v.E01*v.E03 + 0*v.E12*v.E23 + -1*v.E03*v.E01 + 0*v.E23*v.E12,
		E23: 1*v.E02*v.E03 + 0*v.E12*v.E13 + -1*v.E03*v.E02 + 0*v.E13*v.E12,
	}
	r2 := V4{
		E0123: 1*v.E01*v.E23 + -1*v.E02*v.E13 + 1*v.E12*v.E03 + 1*v.E03*v.E12 + -1*v.E13*v.E02 + 1*v.E23*v.E01,
	}
	return r0, r1, r2
}
func (v V2) MulV3(w V3) (V0, V2) {
	r0 := V1{
		E0: 0*v.E12*v.E012 + 0*v.E13*v.E013 + 0*v.E23*v.E023,
		E1: 0*v.E02*v.E012 + -1*v.E03*v.E013 + 0*v.E23*v.E123,
		E2: 0*v.E01*v.E012 + -1*v.E03*v.E023 + 0*v.E13*v.E123,
		E3: 0*v.E01*v.E013 + 0*v.E02*v.E023 + 0*v.E12*v.E123,
	}
	r1 := V3{
		E012: 1*v.E03*v.E123 + -1*v.E13*v.E023 + 1*v.E23*v.E013,
		E013: 0*v.E02*v.E123 + 0*v.E12*v.E023 + 0*v.E23*v.E012,
		E023: 0*v.E01*v.E123 + 0*v.E12*v.E013 + 0*v.E13*v.E012,
		E123: 1*v.E01*v.E023 + -1*v.E02*v.E013 + 1*v.E03*v.E012,
	}
	return r0, r1
}
func (v V2) MulV4(w V4) V1 {
	r0 := V2{
		E01: 0 * v.E23 * v.E0123,
		E02: 0 * v.E13 * v.E0123,
		E12: 1 * v.E03 * v.E0123,
		E03: 0 * v.E12 * v.E0123,
		E13: 0 * v.E02 * v.E0123,
		E23: 0 * v.E01 * v.E0123,
	}
	return r0
}

type V3 struct {
	E012 float64
	E013 float64
	E023 float64
	E123 float64
}

func (v V3) String() string {
	return fmt.Sprintf("(%ve₀₁₂ + %ve₀₁₃ + %ve₀₂₃ + %ve₁₂₃)", v.E012, v.E013, v.E023, v.E123)
}
func (v V3) MulS(s float64) V3 {
	return V3{
		E012: v.E012 * s,
		E013: v.E013 * s,
		E023: v.E023 * s,
		E123: v.E123 * s,
	}
}
func (v V3) MulV1(w V1) (V1, V3) {
	r0 := V2{
		E01: 0*v.E012*v.E2 + 1*v.E013*v.E3,
		E02: 0*v.E012*v.E1 + 1*v.E023*v.E3,
		E12: -1*v.E012*v.E0 + 1*v.E123*v.E3,
		E03: 0*v.E013*v.E1 + 0*v.E023*v.E2,
		E13: -1*v.E013*v.E0 + 0*v.E123*v.E2,
		E23: -1*v.E023*v.E0 + 0*v.E123*v.E1,
	}
	r1 := V4{
		E0123: 1*v.E012*v.E3 + -1*v.E013*v.E2 + 1*v.E023*v.E1 + -1*v.E123*v.E0,
	}
	return r0, r1
}
func (v V3) MulV2(w V2) (V0, V2) {
	r0 := V1{
		E0: 0*v.E012*v.E12 + 0*v.E013*v.E13 + 0*v.E023*v.E23,
		E1: 0*v.E012*v.E02 + -1*v.E013*v.E03 + 0*v.E123*v.E23,
		E2: 0*v.E012*v.E01 + -1*v.E023*v.E03 + 0*v.E123*v.E13,
		E3: 0*v.E013*v.E01 + 0*v.E023*v.E02 + 0*v.E123*v.E12,
	}
	r1 := V3{
		E012: -1*v.E013*v.E23 + 1*v.E023*v.E13 + -1*v.E123*v.E03,
		E013: 0*v.E012*v.E23 + 0*v.E023*v.E12 + 0*v.E123*v.E02,
		E023: 0*v.E012*v.E13 + 0*v.E013*v.E12 + 0*v.E123*v.E01,
		E123: -1*v.E012*v.E03 + 1*v.E013*v.E02 + -1*v.E023*v.E01,
	}
	return r0, r1
}
func (v V3) MulV3(w V3) (float64, V1) {
	r0 := 0*v.E012*v.E012 + 0*v.E013*v.E013 + 0*v.E023*v.E023 + 0*v.E123*v.E123
	r1 := V2{
		E01: 0*v.E023*v.E123 + 0*v.E123*v.E023,
		E02: 0*v.E013*v.E123 + 0*v.E123*v.E013,
		E12: 1*v.E013*v.E023 + -1*v.E023*v.E013,
		E03: 0*v.E012*v.E123 + 0*v.E123*v.E012,
		E13: 0*v.E012*v.E023 + 0*v.E023*v.E012,
		E23: 0*v.E012*v.E013 + 0*v.E013*v.E012,
	}
	return r0, r1
}
func (v V3) MulV4(w V4) V0 {
	r0 := V1{
		E0: 0 * v.E123 * v.E0123,
		E1: 0 * v.E023 * v.E0123,
		E2: 0 * v.E013 * v.E0123,
		E3: 0 * v.E012 * v.E0123,
	}
	return r0
}

type V4 struct {
	E0123 float64
}

func (v V4) String() string {
	return fmt.Sprintf("(%ve₀₁₂₃)", v.E0123)
}
func (v V4) MulS(s float64) V4 {
	return V4{
		E0123: v.E0123 * s,
	}
}
func (v V4) MulV1(w V1) V2 {
	r0 := V3{
		E012: 1 * v.E0123 * v.E3,
		E013: 0 * v.E0123 * v.E2,
		E023: 0 * v.E0123 * v.E1,
		E123: 1 * v.E0123 * v.E0,
	}
	return r0
}
func (v V4) MulV2(w V2) V1 {
	r0 := V2{
		E01: 0 * v.E0123 * v.E23,
		E02: 0 * v.E0123 * v.E13,
		E12: 1 * v.E0123 * v.E03,
		E03: 0 * v.E0123 * v.E12,
		E13: 0 * v.E0123 * v.E02,
		E23: 0 * v.E0123 * v.E01,
	}
	return r0
}
func (v V4) MulV3(w V3) V0 {
	r0 := V1{
		E0: 0 * v.E0123 * v.E123,
		E1: 0 * v.E0123 * v.E023,
		E2: 0 * v.E0123 * v.E013,
		E3: 0 * v.E0123 * v.E012,
	}
	return r0
}
func (v V4) MulV4(w V4) float64 {
	r0 := 0 * v.E0123 * v.E0123
	return r0
}
