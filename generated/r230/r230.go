// e₀²=-1
// e₁²=-1
// e₂²=0
// e₃²=0
// e₄²=0
package r230

import (
	"fmt"
)

type V1 struct {
	E0 float64
	E1 float64
	E2 float64
	E3 float64
	E4 float64
}

func (v V1) String() string {
	return fmt.Sprintf("(%ve₀ + %ve₁ + %ve₂ + %ve₃ + %ve₄)", v.E0, v.E1, v.E2, v.E3, v.E4)
}
func (v V1) MulS(s float64) V1 {
	return V1{
		E0: v.E0 * s,
		E1: v.E1 * s,
		E2: v.E2 * s,
		E3: v.E3 * s,
		E4: v.E4 * s,
	}
}
func (v V1) MulV1(w V1) (float64, V1) {
	r0 := -1*v.E0*v.E0 + -1*v.E1*v.E1 + 0*v.E2*v.E2 + 0*v.E3*v.E3 + 0*v.E4*v.E4
	r1 := V2{
		E01: 1*v.E0*v.E1 + -1*v.E1*v.E0,
		E02: 1*v.E0*v.E2 + -1*v.E2*v.E0,
		E12: 1*v.E1*v.E2 + -1*v.E2*v.E1,
		E03: 1*v.E0*v.E3 + -1*v.E3*v.E0,
		E13: 1*v.E1*v.E3 + -1*v.E3*v.E1,
		E23: 1*v.E2*v.E3 + -1*v.E3*v.E2,
		E04: 1*v.E0*v.E4 + -1*v.E4*v.E0,
		E14: 1*v.E1*v.E4 + -1*v.E4*v.E1,
		E24: 1*v.E2*v.E4 + -1*v.E4*v.E2,
		E34: 1*v.E3*v.E4 + -1*v.E4*v.E3,
	}
	return r0, r1
}
func (v V1) MulV2(w V2) (V0, V2) {
	r0 := V1{
		E0: 1*v.E1*v.E01 + 0*v.E2*v.E02 + 0*v.E3*v.E03 + 0*v.E4*v.E04,
		E1: -1*v.E0*v.E01 + 0*v.E2*v.E12 + 0*v.E3*v.E13 + 0*v.E4*v.E14,
		E2: -1*v.E0*v.E02 + -1*v.E1*v.E12 + 0*v.E3*v.E23 + 0*v.E4*v.E24,
		E3: -1*v.E0*v.E03 + -1*v.E1*v.E13 + 0*v.E2*v.E23 + 0*v.E4*v.E34,
		E4: -1*v.E0*v.E04 + -1*v.E1*v.E14 + 0*v.E2*v.E24 + 0*v.E3*v.E34,
	}
	r1 := V3{
		E012: 1*v.E0*v.E12 + -1*v.E1*v.E02 + 1*v.E2*v.E01,
		E013: 1*v.E0*v.E13 + -1*v.E1*v.E03 + 1*v.E3*v.E01,
		E023: 1*v.E0*v.E23 + -1*v.E2*v.E03 + 1*v.E3*v.E02,
		E123: 1*v.E1*v.E23 + -1*v.E2*v.E13 + 1*v.E3*v.E12,
		E014: 1*v.E0*v.E14 + -1*v.E1*v.E04 + 1*v.E4*v.E01,
		E024: 1*v.E0*v.E24 + -1*v.E2*v.E04 + 1*v.E4*v.E02,
		E124: 1*v.E1*v.E24 + -1*v.E2*v.E14 + 1*v.E4*v.E12,
		E034: 1*v.E0*v.E34 + -1*v.E3*v.E04 + 1*v.E4*v.E03,
		E134: 1*v.E1*v.E34 + -1*v.E3*v.E14 + 1*v.E4*v.E13,
		E234: 1*v.E2*v.E34 + -1*v.E3*v.E24 + 1*v.E4*v.E23,
	}
	return r0, r1
}
func (v V1) MulV3(w V3) (V1, V3) {
	r0 := V2{
		E01: 0*v.E2*v.E012 + 0*v.E3*v.E013 + 0*v.E4*v.E014,
		E02: 1*v.E1*v.E012 + 0*v.E3*v.E023 + 0*v.E4*v.E024,
		E12: -1*v.E0*v.E012 + 0*v.E3*v.E123 + 0*v.E4*v.E124,
		E03: 1*v.E1*v.E013 + 0*v.E2*v.E023 + 0*v.E4*v.E034,
		E13: -1*v.E0*v.E013 + 0*v.E2*v.E123 + 0*v.E4*v.E134,
		E23: -1*v.E0*v.E023 + -1*v.E1*v.E123 + 0*v.E4*v.E234,
		E04: 1*v.E1*v.E014 + 0*v.E2*v.E024 + 0*v.E3*v.E034,
		E14: -1*v.E0*v.E014 + 0*v.E2*v.E124 + 0*v.E3*v.E134,
		E24: -1*v.E0*v.E024 + -1*v.E1*v.E124 + 0*v.E3*v.E234,
		E34: -1*v.E0*v.E034 + -1*v.E1*v.E134 + 0*v.E2*v.E234,
	}
	r1 := V4{
		E0123: 1*v.E0*v.E123 + -1*v.E1*v.E023 + 1*v.E2*v.E013 + -1*v.E3*v.E012,
		E0124: 1*v.E0*v.E124 + -1*v.E1*v.E024 + 1*v.E2*v.E014 + -1*v.E4*v.E012,
		E0134: 1*v.E0*v.E134 + -1*v.E1*v.E034 + 1*v.E3*v.E014 + -1*v.E4*v.E013,
		E0234: 1*v.E0*v.E234 + -1*v.E2*v.E034 + 1*v.E3*v.E024 + -1*v.E4*v.E023,
		E1234: 1*v.E1*v.E234 + -1*v.E2*v.E134 + 1*v.E3*v.E124 + -1*v.E4*v.E123,
	}
	return r0, r1
}
func (v V1) MulV4(w V4) (V2, V4) {
	r0 := V3{
		E012: 0*v.E3*v.E0123 + 0*v.E4*v.E0124,
		E013: 0*v.E2*v.E0123 + 0*v.E4*v.E0134,
		E023: 1*v.E1*v.E0123 + 0*v.E4*v.E0234,
		E123: -1*v.E0*v.E0123 + 0*v.E4*v.E1234,
		E014: 0*v.E2*v.E0124 + 0*v.E3*v.E0134,
		E024: 1*v.E1*v.E0124 + 0*v.E3*v.E0234,
		E124: -1*v.E0*v.E0124 + 0*v.E3*v.E1234,
		E034: 1*v.E1*v.E0134 + 0*v.E2*v.E0234,
		E134: -1*v.E0*v.E0134 + 0*v.E2*v.E1234,
		E234: -1*v.E0*v.E0234 + -1*v.E1*v.E1234,
	}
	r1 := V5{
		E01234: 1*v.E0*v.E1234 + -1*v.E1*v.E0234 + 1*v.E2*v.E0134 + -1*v.E3*v.E0124 + 1*v.E4*v.E0123,
	}
	return r0, r1
}
func (v V1) MulV5(w V5) V3 {
	r0 := V4{
		E0123: 0 * v.E4 * v.E01234,
		E0124: 0 * v.E3 * v.E01234,
		E0134: 0 * v.E2 * v.E01234,
		E0234: 1 * v.E1 * v.E01234,
		E1234: -1 * v.E0 * v.E01234,
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
	E04 float64
	E14 float64
	E24 float64
	E34 float64
}

func (v V2) String() string {
	return fmt.Sprintf("(%ve₀₁ + %ve₀₂ + %ve₁₂ + %ve₀₃ + %ve₁₃ + %ve₂₃ + %ve₀₄ + %ve₁₄ + %ve₂₄ + %ve₃₄)", v.E01, v.E02, v.E12, v.E03, v.E13, v.E23, v.E04, v.E14, v.E24, v.E34)
}
func (v V2) MulS(s float64) V2 {
	return V2{
		E01: v.E01 * s,
		E02: v.E02 * s,
		E12: v.E12 * s,
		E03: v.E03 * s,
		E13: v.E13 * s,
		E23: v.E23 * s,
		E04: v.E04 * s,
		E14: v.E14 * s,
		E24: v.E24 * s,
		E34: v.E34 * s,
	}
}
func (v V2) MulV1(w V1) (V0, V2) {
	r0 := V1{
		E0: -1*v.E01*v.E1 + 0*v.E02*v.E2 + 0*v.E03*v.E3 + 0*v.E04*v.E4,
		E1: 1*v.E01*v.E0 + 0*v.E12*v.E2 + 0*v.E13*v.E3 + 0*v.E14*v.E4,
		E2: 1*v.E02*v.E0 + 1*v.E12*v.E1 + 0*v.E23*v.E3 + 0*v.E24*v.E4,
		E3: 1*v.E03*v.E0 + 1*v.E13*v.E1 + 0*v.E23*v.E2 + 0*v.E34*v.E4,
		E4: 1*v.E04*v.E0 + 1*v.E14*v.E1 + 0*v.E24*v.E2 + 0*v.E34*v.E3,
	}
	r1 := V3{
		E012: 1*v.E01*v.E2 + -1*v.E02*v.E1 + 1*v.E12*v.E0,
		E013: 1*v.E01*v.E3 + -1*v.E03*v.E1 + 1*v.E13*v.E0,
		E023: 1*v.E02*v.E3 + -1*v.E03*v.E2 + 1*v.E23*v.E0,
		E123: 1*v.E12*v.E3 + -1*v.E13*v.E2 + 1*v.E23*v.E1,
		E014: 1*v.E01*v.E4 + -1*v.E04*v.E1 + 1*v.E14*v.E0,
		E024: 1*v.E02*v.E4 + -1*v.E04*v.E2 + 1*v.E24*v.E0,
		E124: 1*v.E12*v.E4 + -1*v.E14*v.E2 + 1*v.E24*v.E1,
		E034: 1*v.E03*v.E4 + -1*v.E04*v.E3 + 1*v.E34*v.E0,
		E134: 1*v.E13*v.E4 + -1*v.E14*v.E3 + 1*v.E34*v.E1,
		E234: 1*v.E23*v.E4 + -1*v.E24*v.E3 + 1*v.E34*v.E2,
	}
	return r0, r1
}
func (v V2) MulV2(w V2) (float64, V1, V3) {
	r0 := -1*v.E01*v.E01 + 0*v.E02*v.E02 + 0*v.E12*v.E12 + 0*v.E03*v.E03 + 0*v.E13*v.E13 + 0*v.E23*v.E23 + 0*v.E04*v.E04 + 0*v.E14*v.E14 + 0*v.E24*v.E24 + 0*v.E34*v.E34
	r1 := V2{
		E01: 0*v.E02*v.E12 + 0*v.E12*v.E02 + 0*v.E03*v.E13 + 0*v.E13*v.E03 + 0*v.E04*v.E14 + 0*v.E14*v.E04,
		E02: -1*v.E01*v.E12 + 1*v.E12*v.E01 + 0*v.E03*v.E23 + 0*v.E23*v.E03 + 0*v.E04*v.E24 + 0*v.E24*v.E04,
		E12: 1*v.E01*v.E02 + -1*v.E02*v.E01 + 0*v.E13*v.E23 + 0*v.E23*v.E13 + 0*v.E14*v.E24 + 0*v.E24*v.E14,
		E03: -1*v.E01*v.E13 + 0*v.E02*v.E23 + 1*v.E13*v.E01 + 0*v.E23*v.E02 + 0*v.E04*v.E34 + 0*v.E34*v.E04,
		E13: 1*v.E01*v.E03 + 0*v.E12*v.E23 + -1*v.E03*v.E01 + 0*v.E23*v.E12 + 0*v.E14*v.E34 + 0*v.E34*v.E14,
		E23: 1*v.E02*v.E03 + 1*v.E12*v.E13 + -1*v.E03*v.E02 + -1*v.E13*v.E12 + 0*v.E24*v.E34 + 0*v.E34*v.E24,
		E04: -1*v.E01*v.E14 + 0*v.E02*v.E24 + 0*v.E03*v.E34 + 1*v.E14*v.E01 + 0*v.E24*v.E02 + 0*v.E34*v.E03,
		E14: 1*v.E01*v.E04 + 0*v.E12*v.E24 + 0*v.E13*v.E34 + -1*v.E04*v.E01 + 0*v.E24*v.E12 + 0*v.E34*v.E13,
		E24: 1*v.E02*v.E04 + 1*v.E12*v.E14 + 0*v.E23*v.E34 + -1*v.E04*v.E02 + -1*v.E14*v.E12 + 0*v.E34*v.E23,
		E34: 1*v.E03*v.E04 + 1*v.E13*v.E14 + 0*v.E23*v.E24 + -1*v.E04*v.E03 + -1*v.E14*v.E13 + 0*v.E24*v.E23,
	}
	r2 := V4{
		E0123: 1*v.E01*v.E23 + -1*v.E02*v.E13 + 1*v.E12*v.E03 + 1*v.E03*v.E12 + -1*v.E13*v.E02 + 1*v.E23*v.E01,
		E0124: 1*v.E01*v.E24 + -1*v.E02*v.E14 + 1*v.E12*v.E04 + 1*v.E04*v.E12 + -1*v.E14*v.E02 + 1*v.E24*v.E01,
		E0134: 1*v.E01*v.E34 + -1*v.E03*v.E14 + 1*v.E13*v.E04 + 1*v.E04*v.E13 + -1*v.E14*v.E03 + 1*v.E34*v.E01,
		E0234: 1*v.E02*v.E34 + -1*v.E03*v.E24 + 1*v.E23*v.E04 + 1*v.E04*v.E23 + -1*v.E24*v.E03 + 1*v.E34*v.E02,
		E1234: 1*v.E12*v.E34 + -1*v.E13*v.E24 + 1*v.E23*v.E14 + 1*v.E14*v.E23 + -1*v.E24*v.E13 + 1*v.E34*v.E12,
	}
	return r0, r1, r2
}
func (v V2) MulV3(w V3) (V0, V2, V4) {
	r0 := V1{
		E0: 0*v.E12*v.E012 + 0*v.E13*v.E013 + 0*v.E23*v.E023 + 0*v.E14*v.E014 + 0*v.E24*v.E024 + 0*v.E34*v.E034,
		E1: 0*v.E02*v.E012 + 0*v.E03*v.E013 + 0*v.E23*v.E123 + 0*v.E04*v.E014 + 0*v.E24*v.E124 + 0*v.E34*v.E134,
		E2: -1*v.E01*v.E012 + 0*v.E03*v.E023 + 0*v.E13*v.E123 + 0*v.E04*v.E024 + 0*v.E14*v.E124 + 0*v.E34*v.E234,
		E3: -1*v.E01*v.E013 + 0*v.E02*v.E023 + 0*v.E12*v.E123 + 0*v.E04*v.E034 + 0*v.E14*v.E134 + 0*v.E24*v.E234,
		E4: -1*v.E01*v.E014 + 0*v.E02*v.E024 + 0*v.E12*v.E124 + 0*v.E03*v.E034 + 0*v.E13*v.E134 + 0*v.E23*v.E234,
	}
	r1 := V3{
		E012: 0*v.E03*v.E123 + 0*v.E13*v.E023 + 0*v.E23*v.E013 + 0*v.E04*v.E124 + 0*v.E14*v.E024 + 0*v.E24*v.E014,
		E013: 0*v.E02*v.E123 + 0*v.E12*v.E023 + 0*v.E23*v.E012 + 0*v.E04*v.E134 + 0*v.E14*v.E034 + 0*v.E34*v.E014,
		E023: -1*v.E01*v.E123 + 1*v.E12*v.E013 + -1*v.E13*v.E012 + 0*v.E04*v.E234 + 0*v.E24*v.E034 + 0*v.E34*v.E024,
		E123: 1*v.E01*v.E023 + -1*v.E02*v.E013 + 1*v.E03*v.E012 + 0*v.E14*v.E234 + 0*v.E24*v.E134 + 0*v.E34*v.E124,
		E014: 0*v.E02*v.E124 + 0*v.E12*v.E024 + 0*v.E03*v.E134 + 0*v.E13*v.E034 + 0*v.E24*v.E012 + 0*v.E34*v.E013,
		E024: -1*v.E01*v.E124 + 1*v.E12*v.E014 + 0*v.E03*v.E234 + 0*v.E23*v.E034 + -1*v.E14*v.E012 + 0*v.E34*v.E023,
		E124: 1*v.E01*v.E024 + -1*v.E02*v.E014 + 0*v.E13*v.E234 + 0*v.E23*v.E134 + 1*v.E04*v.E012 + 0*v.E34*v.E123,
		E034: -1*v.E01*v.E134 + 0*v.E02*v.E234 + 1*v.E13*v.E014 + 0*v.E23*v.E024 + -1*v.E14*v.E013 + 0*v.E24*v.E023,
		E134: 1*v.E01*v.E034 + 0*v.E12*v.E234 + -1*v.E03*v.E014 + 0*v.E23*v.E124 + 1*v.E04*v.E013 + 0*v.E24*v.E123,
		E234: 1*v.E02*v.E034 + 1*v.E12*v.E134 + -1*v.E03*v.E024 + -1*v.E13*v.E124 + 1*v.E04*v.E023 + 1*v.E14*v.E123,
	}
	r2 := V5{
		E01234: 1*v.E01*v.E234 + -1*v.E02*v.E134 + 1*v.E12*v.E034 + 1*v.E03*v.E124 + -1*v.E13*v.E024 + 1*v.E23*v.E014 + -1*v.E04*v.E123 + 1*v.E14*v.E023 + -1*v.E24*v.E013 + 1*v.E34*v.E012,
	}
	return r0, r1, r2
}
func (v V2) MulV4(w V4) (V1, V3) {
	r0 := V2{
		E01: 0*v.E23*v.E0123 + 0*v.E24*v.E0124 + 0*v.E34*v.E0134,
		E02: 0*v.E13*v.E0123 + 0*v.E14*v.E0124 + 0*v.E34*v.E0234,
		E12: 0*v.E03*v.E0123 + 0*v.E04*v.E0124 + 0*v.E34*v.E1234,
		E03: 0*v.E12*v.E0123 + 0*v.E14*v.E0134 + 0*v.E24*v.E0234,
		E13: 0*v.E02*v.E0123 + 0*v.E04*v.E0134 + 0*v.E24*v.E1234,
		E23: -1*v.E01*v.E0123 + 0*v.E04*v.E0234 + 0*v.E14*v.E1234,
		E04: 0*v.E12*v.E0124 + 0*v.E13*v.E0134 + 0*v.E23*v.E0234,
		E14: 0*v.E02*v.E0124 + 0*v.E03*v.E0134 + 0*v.E23*v.E1234,
		E24: -1*v.E01*v.E0124 + 0*v.E03*v.E0234 + 0*v.E13*v.E1234,
		E34: -1*v.E01*v.E0134 + 0*v.E02*v.E0234 + 0*v.E12*v.E1234,
	}
	r1 := V4{
		E0123: 0*v.E04*v.E1234 + 0*v.E14*v.E0234 + 0*v.E24*v.E0134 + 0*v.E34*v.E0124,
		E0124: 0*v.E03*v.E1234 + 0*v.E13*v.E0234 + 0*v.E23*v.E0134 + 0*v.E34*v.E0123,
		E0134: 0*v.E02*v.E1234 + 0*v.E12*v.E0234 + 0*v.E23*v.E0124 + 0*v.E24*v.E0123,
		E0234: -1*v.E01*v.E1234 + 1*v.E12*v.E0134 + -1*v.E13*v.E0124 + 1*v.E14*v.E0123,
		E1234: 1*v.E01*v.E0234 + -1*v.E02*v.E0134 + 1*v.E03*v.E0124 + -1*v.E04*v.E0123,
	}
	return r0, r1
}
func (v V2) MulV5(w V5) V2 {
	r0 := V3{
		E012: 0 * v.E34 * v.E01234,
		E013: 0 * v.E24 * v.E01234,
		E023: 0 * v.E14 * v.E01234,
		E123: 0 * v.E04 * v.E01234,
		E014: 0 * v.E23 * v.E01234,
		E024: 0 * v.E13 * v.E01234,
		E124: 0 * v.E03 * v.E01234,
		E034: 0 * v.E12 * v.E01234,
		E134: 0 * v.E02 * v.E01234,
		E234: -1 * v.E01 * v.E01234,
	}
	return r0
}

type V3 struct {
	E012 float64
	E013 float64
	E023 float64
	E123 float64
	E014 float64
	E024 float64
	E124 float64
	E034 float64
	E134 float64
	E234 float64
}

func (v V3) String() string {
	return fmt.Sprintf("(%ve₀₁₂ + %ve₀₁₃ + %ve₀₂₃ + %ve₁₂₃ + %ve₀₁₄ + %ve₀₂₄ + %ve₁₂₄ + %ve₀₃₄ + %ve₁₃₄ + %ve₂₃₄)", v.E012, v.E013, v.E023, v.E123, v.E014, v.E024, v.E124, v.E034, v.E134, v.E234)
}
func (v V3) MulS(s float64) V3 {
	return V3{
		E012: v.E012 * s,
		E013: v.E013 * s,
		E023: v.E023 * s,
		E123: v.E123 * s,
		E014: v.E014 * s,
		E024: v.E024 * s,
		E124: v.E124 * s,
		E034: v.E034 * s,
		E134: v.E134 * s,
		E234: v.E234 * s,
	}
}
func (v V3) MulV1(w V1) (V1, V3) {
	r0 := V2{
		E01: 0*v.E012*v.E2 + 0*v.E013*v.E3 + 0*v.E014*v.E4,
		E02: 1*v.E012*v.E1 + 0*v.E023*v.E3 + 0*v.E024*v.E4,
		E12: -1*v.E012*v.E0 + 0*v.E123*v.E3 + 0*v.E124*v.E4,
		E03: 1*v.E013*v.E1 + 0*v.E023*v.E2 + 0*v.E034*v.E4,
		E13: -1*v.E013*v.E0 + 0*v.E123*v.E2 + 0*v.E134*v.E4,
		E23: -1*v.E023*v.E0 + -1*v.E123*v.E1 + 0*v.E234*v.E4,
		E04: 1*v.E014*v.E1 + 0*v.E024*v.E2 + 0*v.E034*v.E3,
		E14: -1*v.E014*v.E0 + 0*v.E124*v.E2 + 0*v.E134*v.E3,
		E24: -1*v.E024*v.E0 + -1*v.E124*v.E1 + 0*v.E234*v.E3,
		E34: -1*v.E034*v.E0 + -1*v.E134*v.E1 + 0*v.E234*v.E2,
	}
	r1 := V4{
		E0123: 1*v.E012*v.E3 + -1*v.E013*v.E2 + 1*v.E023*v.E1 + -1*v.E123*v.E0,
		E0124: 1*v.E012*v.E4 + -1*v.E014*v.E2 + 1*v.E024*v.E1 + -1*v.E124*v.E0,
		E0134: 1*v.E013*v.E4 + -1*v.E014*v.E3 + 1*v.E034*v.E1 + -1*v.E134*v.E0,
		E0234: 1*v.E023*v.E4 + -1*v.E024*v.E3 + 1*v.E034*v.E2 + -1*v.E234*v.E0,
		E1234: 1*v.E123*v.E4 + -1*v.E124*v.E3 + 1*v.E134*v.E2 + -1*v.E234*v.E1,
	}
	return r0, r1
}
func (v V3) MulV2(w V2) (V0, V2, V4) {
	r0 := V1{
		E0: 0*v.E012*v.E12 + 0*v.E013*v.E13 + 0*v.E023*v.E23 + 0*v.E014*v.E14 + 0*v.E024*v.E24 + 0*v.E034*v.E34,
		E1: 0*v.E012*v.E02 + 0*v.E013*v.E03 + 0*v.E123*v.E23 + 0*v.E014*v.E04 + 0*v.E124*v.E24 + 0*v.E134*v.E34,
		E2: -1*v.E012*v.E01 + 0*v.E023*v.E03 + 0*v.E123*v.E13 + 0*v.E024*v.E04 + 0*v.E124*v.E14 + 0*v.E234*v.E34,
		E3: -1*v.E013*v.E01 + 0*v.E023*v.E02 + 0*v.E123*v.E12 + 0*v.E034*v.E04 + 0*v.E134*v.E14 + 0*v.E234*v.E24,
		E4: -1*v.E014*v.E01 + 0*v.E024*v.E02 + 0*v.E124*v.E12 + 0*v.E034*v.E03 + 0*v.E134*v.E13 + 0*v.E234*v.E23,
	}
	r1 := V3{
		E012: 0*v.E013*v.E23 + 0*v.E023*v.E13 + 0*v.E123*v.E03 + 0*v.E014*v.E24 + 0*v.E024*v.E14 + 0*v.E124*v.E04,
		E013: 0*v.E012*v.E23 + 0*v.E023*v.E12 + 0*v.E123*v.E02 + 0*v.E014*v.E34 + 0*v.E034*v.E14 + 0*v.E134*v.E04,
		E023: 1*v.E012*v.E13 + -1*v.E013*v.E12 + 1*v.E123*v.E01 + 0*v.E024*v.E34 + 0*v.E034*v.E24 + 0*v.E234*v.E04,
		E123: -1*v.E012*v.E03 + 1*v.E013*v.E02 + -1*v.E023*v.E01 + 0*v.E124*v.E34 + 0*v.E134*v.E24 + 0*v.E234*v.E14,
		E014: 0*v.E012*v.E24 + 0*v.E013*v.E34 + 0*v.E024*v.E12 + 0*v.E124*v.E02 + 0*v.E034*v.E13 + 0*v.E134*v.E03,
		E024: 1*v.E012*v.E14 + 0*v.E023*v.E34 + -1*v.E014*v.E12 + 1*v.E124*v.E01 + 0*v.E034*v.E23 + 0*v.E234*v.E03,
		E124: -1*v.E012*v.E04 + 0*v.E123*v.E34 + 1*v.E014*v.E02 + -1*v.E024*v.E01 + 0*v.E134*v.E23 + 0*v.E234*v.E13,
		E034: 1*v.E013*v.E14 + 0*v.E023*v.E24 + -1*v.E014*v.E13 + 0*v.E024*v.E23 + 1*v.E134*v.E01 + 0*v.E234*v.E02,
		E134: -1*v.E013*v.E04 + 0*v.E123*v.E24 + 1*v.E014*v.E03 + 0*v.E124*v.E23 + -1*v.E034*v.E01 + 0*v.E234*v.E12,
		E234: -1*v.E023*v.E04 + -1*v.E123*v.E14 + 1*v.E024*v.E03 + 1*v.E124*v.E13 + -1*v.E034*v.E02 + -1*v.E134*v.E12,
	}
	r2 := V5{
		E01234: 1*v.E012*v.E34 + -1*v.E013*v.E24 + 1*v.E023*v.E14 + -1*v.E123*v.E04 + 1*v.E014*v.E23 + -1*v.E024*v.E13 + 1*v.E124*v.E03 + 1*v.E034*v.E12 + -1*v.E134*v.E02 + 1*v.E234*v.E01,
	}
	return r0, r1, r2
}
func (v V3) MulV3(w V3) (float64, V1, V3) {
	r0 := 0*v.E012*v.E012 + 0*v.E013*v.E013 + 0*v.E023*v.E023 + 0*v.E123*v.E123 + 0*v.E014*v.E014 + 0*v.E024*v.E024 + 0*v.E124*v.E124 + 0*v.E034*v.E034 + 0*v.E134*v.E134 + 0*v.E234*v.E234
	r1 := V2{
		E01: 0*v.E023*v.E123 + 0*v.E123*v.E023 + 0*v.E024*v.E124 + 0*v.E124*v.E024 + 0*v.E034*v.E134 + 0*v.E134*v.E034,
		E02: 0*v.E013*v.E123 + 0*v.E123*v.E013 + 0*v.E014*v.E124 + 0*v.E124*v.E014 + 0*v.E034*v.E234 + 0*v.E234*v.E034,
		E12: 0*v.E013*v.E023 + 0*v.E023*v.E013 + 0*v.E014*v.E024 + 0*v.E024*v.E014 + 0*v.E134*v.E234 + 0*v.E234*v.E134,
		E03: 0*v.E012*v.E123 + 0*v.E123*v.E012 + 0*v.E014*v.E134 + 0*v.E024*v.E234 + 0*v.E134*v.E014 + 0*v.E234*v.E024,
		E13: 0*v.E012*v.E023 + 0*v.E023*v.E012 + 0*v.E014*v.E034 + 0*v.E124*v.E234 + 0*v.E034*v.E014 + 0*v.E234*v.E124,
		E23: -1*v.E012*v.E013 + 1*v.E013*v.E012 + 0*v.E024*v.E034 + 0*v.E124*v.E134 + 0*v.E034*v.E024 + 0*v.E134*v.E124,
		E04: 0*v.E012*v.E124 + 0*v.E013*v.E134 + 0*v.E023*v.E234 + 0*v.E124*v.E012 + 0*v.E134*v.E013 + 0*v.E234*v.E023,
		E14: 0*v.E012*v.E024 + 0*v.E013*v.E034 + 0*v.E123*v.E234 + 0*v.E024*v.E012 + 0*v.E034*v.E013 + 0*v.E234*v.E123,
		E24: -1*v.E012*v.E014 + 0*v.E023*v.E034 + 0*v.E123*v.E134 + 1*v.E014*v.E012 + 0*v.E034*v.E023 + 0*v.E134*v.E123,
		E34: -1*v.E013*v.E014 + 0*v.E023*v.E024 + 0*v.E123*v.E124 + 1*v.E014*v.E013 + 0*v.E024*v.E023 + 0*v.E124*v.E123,
	}
	r2 := V4{
		E0123: 0*v.E014*v.E234 + 0*v.E024*v.E134 + 0*v.E124*v.E034 + 0*v.E034*v.E124 + 0*v.E134*v.E024 + 0*v.E234*v.E014,
		E0124: 0*v.E013*v.E234 + 0*v.E023*v.E134 + 0*v.E123*v.E034 + 0*v.E034*v.E123 + 0*v.E134*v.E023 + 0*v.E234*v.E013,
		E0134: 0*v.E012*v.E234 + 0*v.E023*v.E124 + 0*v.E123*v.E024 + 0*v.E024*v.E123 + 0*v.E124*v.E023 + 0*v.E234*v.E012,
		E0234: 1*v.E012*v.E134 + -1*v.E013*v.E124 + 1*v.E123*v.E014 + 1*v.E014*v.E123 + -1*v.E124*v.E013 + 1*v.E134*v.E012,
		E1234: -1*v.E012*v.E034 + 1*v.E013*v.E024 + -1*v.E023*v.E014 + -1*v.E014*v.E023 + 1*v.E024*v.E013 + -1*v.E034*v.E012,
	}
	return r0, r1, r2
}
func (v V3) MulV4(w V4) (V0, V2) {
	r0 := V1{
		E0: 0*v.E123*v.E0123 + 0*v.E124*v.E0124 + 0*v.E134*v.E0134 + 0*v.E234*v.E0234,
		E1: 0*v.E023*v.E0123 + 0*v.E024*v.E0124 + 0*v.E034*v.E0134 + 0*v.E234*v.E1234,
		E2: 0*v.E013*v.E0123 + 0*v.E014*v.E0124 + 0*v.E034*v.E0234 + 0*v.E134*v.E1234,
		E3: 0*v.E012*v.E0123 + 0*v.E014*v.E0134 + 0*v.E024*v.E0234 + 0*v.E124*v.E1234,
		E4: 0*v.E012*v.E0124 + 0*v.E013*v.E0134 + 0*v.E023*v.E0234 + 0*v.E123*v.E1234,
	}
	r1 := V3{
		E012: 0*v.E034*v.E1234 + 0*v.E134*v.E0234 + 0*v.E234*v.E0134,
		E013: 0*v.E024*v.E1234 + 0*v.E124*v.E0234 + 0*v.E234*v.E0124,
		E023: 0*v.E014*v.E1234 + 0*v.E124*v.E0134 + 0*v.E134*v.E0124,
		E123: 0*v.E014*v.E0234 + 0*v.E024*v.E0134 + 0*v.E034*v.E0124,
		E014: 0*v.E023*v.E1234 + 0*v.E123*v.E0234 + 0*v.E234*v.E0123,
		E024: 0*v.E013*v.E1234 + 0*v.E123*v.E0134 + 0*v.E134*v.E0123,
		E124: 0*v.E013*v.E0234 + 0*v.E023*v.E0134 + 0*v.E034*v.E0123,
		E034: 0*v.E012*v.E1234 + 0*v.E123*v.E0124 + 0*v.E124*v.E0123,
		E134: 0*v.E012*v.E0234 + 0*v.E023*v.E0124 + 0*v.E024*v.E0123,
		E234: -1*v.E012*v.E0134 + 1*v.E013*v.E0124 + -1*v.E014*v.E0123,
	}
	return r0, r1
}
func (v V3) MulV5(w V5) V1 {
	r0 := V2{
		E01: 0 * v.E234 * v.E01234,
		E02: 0 * v.E134 * v.E01234,
		E12: 0 * v.E034 * v.E01234,
		E03: 0 * v.E124 * v.E01234,
		E13: 0 * v.E024 * v.E01234,
		E23: 0 * v.E014 * v.E01234,
		E04: 0 * v.E123 * v.E01234,
		E14: 0 * v.E023 * v.E01234,
		E24: 0 * v.E013 * v.E01234,
		E34: 0 * v.E012 * v.E01234,
	}
	return r0
}

type V4 struct {
	E0123 float64
	E0124 float64
	E0134 float64
	E0234 float64
	E1234 float64
}

func (v V4) String() string {
	return fmt.Sprintf("(%ve₀₁₂₃ + %ve₀₁₂₄ + %ve₀₁₃₄ + %ve₀₂₃₄ + %ve₁₂₃₄)", v.E0123, v.E0124, v.E0134, v.E0234, v.E1234)
}
func (v V4) MulS(s float64) V4 {
	return V4{
		E0123: v.E0123 * s,
		E0124: v.E0124 * s,
		E0134: v.E0134 * s,
		E0234: v.E0234 * s,
		E1234: v.E1234 * s,
	}
}
func (v V4) MulV1(w V1) (V2, V4) {
	r0 := V3{
		E012: 0*v.E0123*v.E3 + 0*v.E0124*v.E4,
		E013: 0*v.E0123*v.E2 + 0*v.E0134*v.E4,
		E023: -1*v.E0123*v.E1 + 0*v.E0234*v.E4,
		E123: 1*v.E0123*v.E0 + 0*v.E1234*v.E4,
		E014: 0*v.E0124*v.E2 + 0*v.E0134*v.E3,
		E024: -1*v.E0124*v.E1 + 0*v.E0234*v.E3,
		E124: 1*v.E0124*v.E0 + 0*v.E1234*v.E3,
		E034: -1*v.E0134*v.E1 + 0*v.E0234*v.E2,
		E134: 1*v.E0134*v.E0 + 0*v.E1234*v.E2,
		E234: 1*v.E0234*v.E0 + 1*v.E1234*v.E1,
	}
	r1 := V5{
		E01234: 1*v.E0123*v.E4 + -1*v.E0124*v.E3 + 1*v.E0134*v.E2 + -1*v.E0234*v.E1 + 1*v.E1234*v.E0,
	}
	return r0, r1
}
func (v V4) MulV2(w V2) (V1, V3) {
	r0 := V2{
		E01: 0*v.E0123*v.E23 + 0*v.E0124*v.E24 + 0*v.E0134*v.E34,
		E02: 0*v.E0123*v.E13 + 0*v.E0124*v.E14 + 0*v.E0234*v.E34,
		E12: 0*v.E0123*v.E03 + 0*v.E0124*v.E04 + 0*v.E1234*v.E34,
		E03: 0*v.E0123*v.E12 + 0*v.E0134*v.E14 + 0*v.E0234*v.E24,
		E13: 0*v.E0123*v.E02 + 0*v.E0134*v.E04 + 0*v.E1234*v.E24,
		E23: -1*v.E0123*v.E01 + 0*v.E0234*v.E04 + 0*v.E1234*v.E14,
		E04: 0*v.E0124*v.E12 + 0*v.E0134*v.E13 + 0*v.E0234*v.E23,
		E14: 0*v.E0124*v.E02 + 0*v.E0134*v.E03 + 0*v.E1234*v.E23,
		E24: -1*v.E0124*v.E01 + 0*v.E0234*v.E03 + 0*v.E1234*v.E13,
		E34: -1*v.E0134*v.E01 + 0*v.E0234*v.E02 + 0*v.E1234*v.E12,
	}
	r1 := V4{
		E0123: 0*v.E0124*v.E34 + 0*v.E0134*v.E24 + 0*v.E0234*v.E14 + 0*v.E1234*v.E04,
		E0124: 0*v.E0123*v.E34 + 0*v.E0134*v.E23 + 0*v.E0234*v.E13 + 0*v.E1234*v.E03,
		E0134: 0*v.E0123*v.E24 + 0*v.E0124*v.E23 + 0*v.E0234*v.E12 + 0*v.E1234*v.E02,
		E0234: -1*v.E0123*v.E14 + 1*v.E0124*v.E13 + -1*v.E0134*v.E12 + 1*v.E1234*v.E01,
		E1234: 1*v.E0123*v.E04 + -1*v.E0124*v.E03 + 1*v.E0134*v.E02 + -1*v.E0234*v.E01,
	}
	return r0, r1
}
func (v V4) MulV3(w V3) (V0, V2) {
	r0 := V1{
		E0: 0*v.E0123*v.E123 + 0*v.E0124*v.E124 + 0*v.E0134*v.E134 + 0*v.E0234*v.E234,
		E1: 0*v.E0123*v.E023 + 0*v.E0124*v.E024 + 0*v.E0134*v.E034 + 0*v.E1234*v.E234,
		E2: 0*v.E0123*v.E013 + 0*v.E0124*v.E014 + 0*v.E0234*v.E034 + 0*v.E1234*v.E134,
		E3: 0*v.E0123*v.E012 + 0*v.E0134*v.E014 + 0*v.E0234*v.E024 + 0*v.E1234*v.E124,
		E4: 0*v.E0124*v.E012 + 0*v.E0134*v.E013 + 0*v.E0234*v.E023 + 0*v.E1234*v.E123,
	}
	r1 := V3{
		E012: 0*v.E0134*v.E234 + 0*v.E0234*v.E134 + 0*v.E1234*v.E034,
		E013: 0*v.E0124*v.E234 + 0*v.E0234*v.E124 + 0*v.E1234*v.E024,
		E023: 0*v.E0124*v.E134 + 0*v.E0134*v.E124 + 0*v.E1234*v.E014,
		E123: 0*v.E0124*v.E034 + 0*v.E0134*v.E024 + 0*v.E0234*v.E014,
		E014: 0*v.E0123*v.E234 + 0*v.E0234*v.E123 + 0*v.E1234*v.E023,
		E024: 0*v.E0123*v.E134 + 0*v.E0134*v.E123 + 0*v.E1234*v.E013,
		E124: 0*v.E0123*v.E034 + 0*v.E0134*v.E023 + 0*v.E0234*v.E013,
		E034: 0*v.E0123*v.E124 + 0*v.E0124*v.E123 + 0*v.E1234*v.E012,
		E134: 0*v.E0123*v.E024 + 0*v.E0124*v.E023 + 0*v.E0234*v.E012,
		E234: -1*v.E0123*v.E014 + 1*v.E0124*v.E013 + -1*v.E0134*v.E012,
	}
	return r0, r1
}
func (v V4) MulV4(w V4) (float64, V1) {
	r0 := 0*v.E0123*v.E0123 + 0*v.E0124*v.E0124 + 0*v.E0134*v.E0134 + 0*v.E0234*v.E0234 + 0*v.E1234*v.E1234
	r1 := V2{
		E01: 0*v.E0234*v.E1234 + 0*v.E1234*v.E0234,
		E02: 0*v.E0134*v.E1234 + 0*v.E1234*v.E0134,
		E12: 0*v.E0134*v.E0234 + 0*v.E0234*v.E0134,
		E03: 0*v.E0124*v.E1234 + 0*v.E1234*v.E0124,
		E13: 0*v.E0124*v.E0234 + 0*v.E0234*v.E0124,
		E23: 0*v.E0124*v.E0134 + 0*v.E0134*v.E0124,
		E04: 0*v.E0123*v.E1234 + 0*v.E1234*v.E0123,
		E14: 0*v.E0123*v.E0234 + 0*v.E0234*v.E0123,
		E24: 0*v.E0123*v.E0134 + 0*v.E0134*v.E0123,
		E34: 0*v.E0123*v.E0124 + 0*v.E0124*v.E0123,
	}
	return r0, r1
}
func (v V4) MulV5(w V5) V0 {
	r0 := V1{
		E0: 0 * v.E1234 * v.E01234,
		E1: 0 * v.E0234 * v.E01234,
		E2: 0 * v.E0134 * v.E01234,
		E3: 0 * v.E0124 * v.E01234,
		E4: 0 * v.E0123 * v.E01234,
	}
	return r0
}

type V5 struct {
	E01234 float64
}

func (v V5) String() string {
	return fmt.Sprintf("(%ve₀₁₂₃₄)", v.E01234)
}
func (v V5) MulS(s float64) V5 {
	return V5{
		E01234: v.E01234 * s,
	}
}
func (v V5) MulV1(w V1) V3 {
	r0 := V4{
		E0123: 0 * v.E01234 * v.E4,
		E0124: 0 * v.E01234 * v.E3,
		E0134: 0 * v.E01234 * v.E2,
		E0234: 1 * v.E01234 * v.E1,
		E1234: -1 * v.E01234 * v.E0,
	}
	return r0
}
func (v V5) MulV2(w V2) V2 {
	r0 := V3{
		E012: 0 * v.E01234 * v.E34,
		E013: 0 * v.E01234 * v.E24,
		E023: 0 * v.E01234 * v.E14,
		E123: 0 * v.E01234 * v.E04,
		E014: 0 * v.E01234 * v.E23,
		E024: 0 * v.E01234 * v.E13,
		E124: 0 * v.E01234 * v.E03,
		E034: 0 * v.E01234 * v.E12,
		E134: 0 * v.E01234 * v.E02,
		E234: -1 * v.E01234 * v.E01,
	}
	return r0
}
func (v V5) MulV3(w V3) V1 {
	r0 := V2{
		E01: 0 * v.E01234 * v.E234,
		E02: 0 * v.E01234 * v.E134,
		E12: 0 * v.E01234 * v.E034,
		E03: 0 * v.E01234 * v.E124,
		E13: 0 * v.E01234 * v.E024,
		E23: 0 * v.E01234 * v.E014,
		E04: 0 * v.E01234 * v.E123,
		E14: 0 * v.E01234 * v.E023,
		E24: 0 * v.E01234 * v.E013,
		E34: 0 * v.E01234 * v.E012,
	}
	return r0
}
func (v V5) MulV4(w V4) V0 {
	r0 := V1{
		E0: 0 * v.E01234 * v.E1234,
		E1: 0 * v.E01234 * v.E0234,
		E2: 0 * v.E01234 * v.E0134,
		E3: 0 * v.E01234 * v.E0124,
		E4: 0 * v.E01234 * v.E0123,
	}
	return r0
}
func (v V5) MulV5(w V5) float64 {
	r0 := 0 * v.E01234 * v.E01234
	return r0
}
