package chkr

import (
	"errors"
	"reflect"
	"testing"
)

type Checker struct {
	*testing.T
}

// New returns a Checker that provides some helper methods on a *testing.T
func New(t *testing.T) Checker {
	return Checker{
		t,
	}
}

// True tests if the value is true.
func (c Checker) True(v bool, args ...interface{}) {
	if v {
		return
	}
	c.Helper()

	mesg := "expected true, but got false"
	if len(args) > 0 {
		mesg += ": " + args[0].(string)
		args = args[1:]
	}

	c.Errorf(mesg, args...)
}

// False tests if the value is false.
func (c Checker) False(v bool, args ...interface{}) {
	if !v {
		return
	}
	c.Helper()

	mesg := "expected false, but got true"
	if len(args) > 0 {
		mesg += ": " + args[0].(string)
		args = args[1:]
	}

	c.Errorf(mesg, args...)
}

// Nil tests if the value is nil.
func (c Checker) Nil(v interface{}, args ...interface{}) {
	if v == nil {
		return
	}
	c.Helper()

	mesg := "expected nil, but got %#v"
	if len(args) > 0 {
		mesg += ": " + args[0].(string)
		args[0] = v
	} else {
		args = []interface{}{v}
	}

	c.Errorf(mesg, args...)
}

// NotNil tests if the value is not nil.
func (c Checker) NotNil(v interface{}, args ...interface{}) {
	if v != nil {
		return
	}
	c.Helper()

	mesg := "expected not nil, but got nil"
	if len(args) > 0 {
		mesg += ": " + args[0].(string)
		args = args[1:]
	}

	c.Errorf(mesg, args...)
}

// Equal tests if the pair of values are equal.
func (c Checker) Equal(v1, v2 interface{}, args ...interface{}) {

	if reflect.DeepEqual(v1, v2) {
		return
	}
	c.Helper()

	mesg := "expected equal, but %#v != %#v"
	if reflect.TypeOf(v1) != reflect.TypeOf(v2) {
		mesg += " (types)"
	}
	if len(args) > 0 {
		mesg += ": " + args[0].(string)
		args = append([]interface{}{v1, v2}, args[1:]...)
	} else {
		args = []interface{}{v1, v2}
	}

	c.Errorf(mesg, args...)
}

// NotEqual tests if the pair of values are not equal.
func (c Checker) NotEqual(v1, v2 interface{}, args ...interface{}) {

	if !reflect.DeepEqual(v1, v2) {
		return
	}
	c.Helper()

	mesg := "expected not equal, but both are %#v"
	if len(args) > 0 {
		mesg += ": " + args[0].(string)
		args[0] = v1
	} else {
		args = []interface{}{v1}
	}

	c.Errorf(mesg, args...)
}

// ErrIs tests if errors.Is() returns true.
func (c Checker) ErrIs(v1, v2 error, args ...interface{}) {

	if errors.Is(v1, v2) {
		return
	}
	c.Helper()

	mesg := "expected errors.Is(%#v, %#v) to return true"
	if len(args) > 0 {
		mesg += ": " + args[0].(string)
		args = append([]interface{}{v1, v2}, args[1:]...)
	} else {
		args = []interface{}{v1, v2}
	}

	c.Errorf(mesg, args...)
}

// ErrAs tests if errors.As() returns true.
func (c Checker) ErrAs(v1 error, v2 interface{}, args ...interface{}) {

	if errors.As(v1, v2) {
		return
	}
	c.Helper()

	mesg := "expected errors.As(%#v, %#v) to return true"
	if len(args) > 0 {
		mesg += ": " + args[0].(string)
		args = append([]interface{}{v1, v2}, args[1:]...)
	} else {
		args = []interface{}{v1, v2}
	}

	c.Errorf(mesg, args...)
}
