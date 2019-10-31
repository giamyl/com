package bjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Float1 struct {
	Value JSONFloat64One `json:"value"`
}

type Float2 struct {
	Value JSONFloat64Two `json:"value"`
}

type Float6 struct {
	Value JSONFloat64Six `json:"value"`
}

func TestFloat(t *testing.T) {
	const f = 273.2873894309280430

	f1 := Float1{f}
	f2 := Float2{f}
	f6 := Float6{f}

	s1, _ := json.Marshal(f1)
	s2, _ := json.Marshal(f2)
	s6, _ := json.Marshal(f6)

	fmt.Println(string(s1))
	fmt.Println(string(s2))
	fmt.Println(string(s6))

}
