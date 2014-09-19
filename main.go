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

func (f *fuse) blend(print printer) string {
 f.msg = print()
 return f.msg
}

func main() {
	f := fuse{&logger{}}
	s := f.blend(test)
	fmt.Println(f.msg)
}
