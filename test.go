package mod_test

import "fmt"

type HBLTest struct {
	Name string
	Age  int64
}

func (h *HBLTest) GName() {
	fmt.Println("version: v2.0.0, name: %s", h.Name)
}
