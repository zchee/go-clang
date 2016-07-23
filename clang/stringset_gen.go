package clang

// #include "./clang-c/CXString.h"
// #include "go-clang.h"
import "C"

type StringSet struct {
	c C.CXStringSet
}

func (ss StringSet) Strings() string {
	o := cxstring{*ss.c.Strings}
	defer o.Dispose()

	return o.String()
}

func (ss StringSet) Count() uint32 {
	return uint32(ss.c.Count)
}
