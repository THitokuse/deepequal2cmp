package a

import (
	"fmt"
	"reflect"
)

func f() {
	m1 := 1
	m2 := 1
	eq := reflect.DeepEqual(m1, m2) // want "DeepEqual is used"
	fmt.Println(eq)
}
