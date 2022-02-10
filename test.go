package mod_test

import "fmt"

type HBLTest struct {
	Name string
	Age  int64
}

func (h *HBLTest) GetName() {
	fmt.Println("[GetName]V1...")
}
