package main

import (
  "testing"
)

func testBlend(t *testing.T) {
        f := fuse{&logger{}}
        s := f.get(test)
	if s!="test" {
		t.Errorf("expect test but got %s",s)
	}
}
