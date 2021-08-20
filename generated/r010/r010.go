// e₀²=0
package r010

import (
	"fmt"
)

type V1 struct {
	E0 float64
}

func (v V1) String() string {
	return fmt.Sprintf("(%ve₀)", v.E0)
}
func (v V1) MulS(s float64) V1 {
	return V1{
		E0: v.E0 * s,
	}
}
func (v V1) MulV1(w V1) float64 {
	r0 := 0 * v.E0 * v.E0
	return r0
}
