// Magic file header to confirm directory for geometric algebgra

package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	{
		const expectedHeader = "// Magic file header to confirm directory for geometric algebgra"
		buildmain, err := os.Open("build.go")
		if err != nil {
			fmt.Println("Did not find expected build.go header, are you running from the right directory?")
			fmt.Println(err)
			os.Exit(1)
		}
		foundHeader, err := bufio.NewReader(buildmain).ReadString('\n')
		if err != nil {
			fmt.Println("Did not find expected build.go header, are you running from the right directory?")
			fmt.Println(err)
			os.Exit(1)
		}

		if (foundHeader != (expectedHeader + "\n")) && (foundHeader != (expectedHeader + "\r\n")) {
			fmt.Println("Did not find expected build.go header, are you running from the right directory?")
			os.Exit(1)
		}
	}

	const max = 6
	for i := 0; i < max; i++ {
		for j := 0; j < max-i; j++ {
			for k := 0; k < max-i-j; k++ {
				if i+j+k > 0 {
					fmt.Printf("Generating r%d%d%d\n", i, j, k)
					err := genPackage(i, j, k)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}

var names []string
var prettyNames []string

func init() {
	digits := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	subscripts := []string{
		"₀", "₁", "₂", "₃", "₄", "₅", "₆", "₇", "₈", "₉",
	}

	for i := uint16(0); i < (1 << 9); i++ {
		name := "E"
		pretty := "e"
		ii := i
		for j := 0; ii > 0; j++ {
			if ii&1 == 1 {
				name += digits[j]
				pretty += subscripts[j]
			}
			ii >>= 1
		}
		names = append(names, name)
		prettyNames = append(prettyNames, pretty)
	}
	names[0] = "1"
	prettyNames[0] = "1"
}

func genPackage(neg, zed, pos int) error {
	squares := make([]int, 0)
	for i := 0; i < neg; i++ {
		squares = append(squares, -1)
	}
	for i := 0; i < zed; i++ {
		squares = append(squares, 0)
	}
	for i := 0; i < pos; i++ {
		squares = append(squares, 1)
	}

	spaceName := fmt.Sprintf("r%d%d%d", neg, zed, pos)
	var w *writer

	{
		err := os.MkdirAll("generated/"+spaceName, os.ModeDir)
		if err != nil {
			return err
		}
		f, err := os.Create(fmt.Sprintf("generated/%s/%s.go", spaceName, spaceName))
		if err != nil {
			return err
		}
		w = &writer{file: f}
	}

	mvectors := make([][]uint16, len(squares)+1)
	for i := uint16(0); i < (1 << len(squares)); i++ {
		oc := bits.OnesCount16(i)
		mvectors[oc] = append(mvectors[oc], i)
	}

	for i, v := range squares {
		w.ln("// %s²=%d", prettyNames[1<<i], v)
	}
	w.ln("package %s", spaceName)

	w.ln("import (")
	w.ln("\"fmt\"")
	w.ln(")")

	for i := 1; i <= len(squares); i++ {
		w.ln("type V%d struct {", i)
		for _, v := range mvectors[i] {
			w.ln("%s float64", names[v])
		}
		w.ln("}")

		///////////////////////////
		// String
		w.ln("func (v V%d) String() string {", i)
		w.f("return fmt.Sprintf(\"(")
		for j, v := range mvectors[i] {
			if j != 0 {
				w.f(" + ")
			}
			w.f("%%v%s", prettyNames[v])
		}
		w.f(")\"")
		for _, v := range mvectors[i] {
			w.f(", v.%s", names[v])
		}
		w.ln(")")
		w.ln("}")

		///////////////////////////
		// MulS
		w.ln("func (v V%d) MulS(s float64) V%d {", i, i)
		w.ln("return V%d{", i)
		for _, v := range mvectors[i] {
			w.ln("%s: v.%s * s,", names[v], names[v])
		}
		w.ln("}")
		w.ln("}")

		///////////////////////////
		// MulV
		for j := 1; j <= len(squares); j++ {
			// fmt.Println("=====", i, j)
			prods := map[uint16][]string{}
			for _, ie := range mvectors[i] {
				for _, je := range mvectors[j] {
					mul, e := vMuls(squares, ie, je)
					// This includes when the multiplier is 0, to make seeing
					// that the code is considering all of the cases easier.
					prods[e] = append(prods[e], fmt.Sprintf("%d * v.%s * v.%s", mul, names[ie], names[je]))
					// fmt.Println(prettyNames[ie], "*", prettyNames[je], "=", mul, "*", prettyNames[e])
				}
			}

			outs := make([]bool, len(squares)+1)
			for k := range prods {
				outs[bits.OnesCount16(k)] = true
			}

			w.f("func (v V%d) MulV%d(w V%d) (", i, j, j)
			if outs[0] {
				w.f("float64,")
			}
			for k, v := range outs[1:] {
				if v {
					w.f("V%d,", k)
				}
			}
			w.ln(") {")
			var r int
			if outs[0] {

			}
			for out, v := range outs {
				if v {
					if out == 0 {
						w.ln("r0 := %s", strings.Join(prods[0], " + "))
						r++
					} else {
						w.ln("r%d := V%d{", r, out)
						for _, k := range mvectors[out] {
							w.ln("%s: %s,", names[k], strings.Join(prods[k], " + "))
						}
						w.ln("}")
						r++
					}
				}
			}
			w.f("return ")
			for k := 0; k < r; k++ {
				w.f("r%d", k)
				if k+1 < r {
					w.f(", ")
				}
			}
			w.ln("")
			w.ln("}")
		}
	}

	w.close()
	return w.err
}

func vMuls(squares []int, a, b uint16) (int, uint16) {
	axis := []int{}

	for i := 0; a > 0; i++ {
		if a&1 == 1 {
			axis = append(axis, i)
		}
		a >>= 1
	}
	for i := 0; b > 0; i++ {
		if b&1 == 1 {
			axis = append(axis, i)
		}
		b >>= 1
	}

	mul := 1
	// Yes, this is a bad sort in O(n) terms, but it's at most 18 elements, so
	// it will always be fast.  Makes it easier to merge elements.
outer:
	for {
		for i := 0; i < len(axis)-1; i++ {
			if axis[i] > axis[i+1] {
				axis[i], axis[i+1] = axis[i+1], axis[i]
				mul *= -1
				continue outer
			}
			if axis[i] == axis[i+1] {
				mul *= squares[axis[i]]
				axis = append(axis[:i], axis[i+2:]...)
				continue outer
			}
		}

		break
	}

	var final uint16
	for _, v := range axis {
		final |= 1 << v
	}
	return mul, final
}

type writer struct {
	file *os.File
	err  error
}

func (w *writer) ln(s string, a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = w.file.WriteString(fmt.Sprintf(s, a...))
	w.f("\n")
}

func (w *writer) f(s string, a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = w.file.WriteString(fmt.Sprintf(s, a...))
}

func (w *writer) close() {
	if w.err == nil {
		w.err = w.file.Close()
	} else {
		w.file.Close()
	}
}
