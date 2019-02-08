package main

import (
	"fmt"

	gomodlib "github.com/travisjeffery/gomodlib"
	"github.com/travisjeffery/gomodlib/subpkga"
)

func main() {
	srv := gomodlib.Service{A: &subpkga.Service{}}
	srv.A.Name = "test we can refer to the name"
	fmt.Println(srv.A.Name)
}
