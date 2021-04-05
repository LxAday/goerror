package goerror

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	t.Log(New("test"))
	t.Logf("%#v", New("test1"))
}

func TestWrap(t *testing.T) {
	test1 := Wrap(errors.New("test"), "test1")
	t.Logf("%#v", test1)

	test2 := Wrap(test1, "test2")
	t.Logf("%#v", test2)

	test3 := Wrap(test2, "test3")
	t.Logf("%#v", test3)

	test4 := Wrap(nil, "test4")
	t.Logf("%#v", test4)

	test5 := Wrap(errors.New("test5"), "test55")
	t.Logf("%#v", test5)

	test6 := Wrap(errors.New("test6"), "")
	t.Logf("%#v", test6)
}

func TestError_Error(t *testing.T) {
	t.Log(New("test").Error())
}

func TestError_Where(t *testing.T) {
	t.Log(New("test").Where())
}
