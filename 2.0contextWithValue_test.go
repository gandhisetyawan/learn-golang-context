package learn_golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f")) //ambil value F dari key : f ke context F : dapat value
	fmt.Println(contextF.Value("c")) //ambil value C dari key : c ke context F : dapat value milik perent
	fmt.Println(contextF.Value("b")) //ambil value B dari key : b ke context F : tidak dapat value , karena beda perent
	fmt.Println(contextA.Value("b")) //ambil value B dari key : b ke context A : tidak dapat value , tidak dapat ambil data child
}
