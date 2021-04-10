package chkr_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/pdk/chkr"
)

func TestIsNil(t *testing.T) {

	chk := chkr.New(t)

	chk.Nil(nil)
	chk.Nil(5)
	chk.Nil(4, "should not be 4")
	chk.Nil(0, "and not %v?", 1)
}

func TestNotNil(t *testing.T) {
	chk := chkr.New(t)

	chk.NotNil(nil)
	chk.NotNil(nil, "gah!")
	chk.NotNil("ack", "pass, actually")
}

func TestEqual(t *testing.T) {
	chk := chkr.New(t)

	chk.Equal(1, 1, "should pass")
	chk.Equal(0, 1)
	chk.Equal(int64(1), 1)
}

func TestNotEqual(t *testing.T) {
	chk := chkr.New(t)

	chk.NotEqual(nil, nil)
	chk.NotEqual(nil, "foo")
	chk.NotEqual(int64(1), 1)
	chk.NotEqual(0, 0, "zeroes are equal to %d", 0)
}

func TestTrue(t *testing.T) {
	chk := chkr.New(t)

	chk.True(false, "arg")
	chk.True(false)
	one := 1
	chk.True(one == 1)

	a := 1
	b := 2
	chk.True(a == b, "a and b are equal, %v == %v", a, b)
}

func TestFalse(t *testing.T) {
	chk := chkr.New(t)

	chk.False("foo" == "bar", "foo and bar are not equal")
	chk.False("foo" == "foo", "foo and foo are equal")

	chk.Errorf("ack")
}

func TestErrorIs(t *testing.T) {
	chk := chkr.New(t)

	err := sql.ErrNoRows
	w := fmt.Errorf("there aint no rows: %w", err)

	chk.ErrIs(err, sql.ErrNoRows)
	chk.ErrIs(w, sql.ErrNoRows)
	chk.ErrIs(nil, sql.ErrNoRows, "let's see that message")
}

// minimal custom error
type myError struct {
	msg string
}

func (e myError) Error() string {
	return e.msg
}

func TestErrorAs(t *testing.T) {
	chk := chkr.New(t)

	err := myError{"bork!"}

	matched := myError{}
	chk.ErrAs(err, &matched, "should pass")
	chk.ErrAs(nil, &matched, "fail!!")

	w := fmt.Errorf("wrap that myError: %w", err)
	chk.ErrAs(w, &matched, "should pass")
	chk.ErrAs(sql.ErrNoRows, &matched, "fail!!")
}

func TestAllTheThings(t *testing.T) {
	chk := chkr.New(t)

	err := fmt.Errorf("FAIL!")
	foo := "foo"

	chk.True(1 == 0)
	chk.False(foo == "foo", "check all the things")
	chk.Nil(err, "should not have gotten error from %s", "that code")
	chk.NotNil(1, "one is not nil. duh")
	chk.Equal(foo, "bar")
	chk.NotEqual(foo, "foo", "just do not understand %d, %d, %d", 1, 2, 3)
	chk.ErrIs(err, sql.ErrNoRows, "wrong error")
}
