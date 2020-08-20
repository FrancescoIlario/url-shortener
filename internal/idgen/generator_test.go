package idgen_test

import (
	"testing"

	"github.com/FrancescoIlario/url-shortener/internal/idgen"
)

const idlength = 7

func Test_LengthNewID(t *testing.T) {
	s, err := idgen.NewID()
	if err != nil {
		t.Fatal(err)
	}

	if obt, exp := len(s), idlength; exp != obt {
		t.Errorf("Invalid length of the produced id, expected %d, obtained %d", exp, obt)
	}
}
