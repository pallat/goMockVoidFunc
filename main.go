package main

import "fmt"

type logger struct {
 msg string
}

func test() string {
 return "test"
}

type fuse struct{
 *logger
}

type printer func() string

func (f *fuse) get(print printer) string {
 f.msg = print()
 return f.msg
}

func main() {
	f := fuse{&logger{}}
	s := f.get(test)
	fmt.Println(f.msg)
}
