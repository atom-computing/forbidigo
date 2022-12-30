package testdata

import (
	"fmt"
	alias "fmt"

	anotherpkg "example.com/another/pkg"
	somepkg "example.com/some/pkg"
)

func myCustom() somepkg.CustomType {
	return somepkg.CustomType{}
}

var myCustomFunc = myCustom

type myCustomStruct struct {
	somepkg.CustomType
}

type myCustomInterface interface {
	AlsoForbidden()
}

func Foo() {
	fmt.Println("here I am") // want "forbidden by pattern"
	fmt.Printf("this is ok") //permit:fmt.Printf // this is ok
	print("not ok")          // want "forbidden by pattern"
	println("also not ok")   // want "forbidden by pattern"
	alias.Println("hello")   // not matched by default pattern fmt.Println
	somepkg.Forbidden()      // want "somepkg.Forbidden.*forbidden by pattern .*example.com/some/pkg.*Forbidden"

	c := somepkg.CustomType{}
	c.AlsoForbidden() // want "c.AlsoForbidden.*forbidden by pattern.*example.com/some/pkg.CustomType.*AlsoForbidden"

	// Selector expression with result of function call in package.
	somepkg.NewCustom().AlsoForbidden() // want "somepkg.NewCustom...AlsoForbidden.*forbidden by pattern.*example.com/some/pkg.CustomType.*AlsoForbidden"

	// Selector expression with result of normal function call.
	myCustom().AlsoForbidden() // want "myCustom...AlsoForbidden.*forbidden by pattern.*example.com/some/pkg.CustomType.*AlsoForbidden"

	// Selector expression with result of normal function call.
	myCustomFunc().AlsoForbidden() // want "myCustomFunc...AlsoForbidden.*forbidden by pattern.*example.com/some/pkg.CustomType.*AlsoForbidden"

	// Type alias and pointer.
	c2 := &anotherpkg.CustomType{}
	c2.AlsoForbidden() // want "c2.AlsoForbidden.*forbidden by pattern.*example.com/some/pkg.CustomType.*AlsoForbidden"

	// Interface.
	var ci somepkg.CustomInterface
	ci.StillForbidden() // want "ci.StillForbidden.*forbidden by pattern.*example.com/some/pkg.CustomInterface.*StillForbidden"

	// Struct embedded inside another - not handled!
	myCustomStruct{}.AlsoForbidden()

	// Forbidden method called via interface - not handled!
	var ci2 myCustomInterface = somepkg.CustomType{}
	ci2.AlsoForbidden()
}

func Bar() string {
	fmt := struct {
		Println string
	}{}
	return fmt.Println // want "forbidden by pattern"
}
