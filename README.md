# chkr

minimal test helper methods

A very small package to provide some testing 1-liners to make go testing more concise.

8 helper methods are provided:

1. `Nil(v, args...)`
2. `NotNil(v, args...)`
3. `True(v, args...)`
4. `False(v, args...)`
5. `Equal(v1, v2, args...)`
6. `NotEqual(v1, v2, args...)`
7. `ErrIs(err, err, args...)`
8. `ErrAs(err, target, args...)`

Each takes a value (1-4) or a pair of values (5-8), and optional arguments. The
first argument after the value(s) is a format/message to be appended to the default failure
message, and subsequent arguments are values to provide to the format.

Example:

    import "github.com/pdk/chkr"

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

Failures for those look like:

    --- FAIL: TestAllTheThings (0.00s)
        chkr_test.go:108: expected true, but got false
        chkr_test.go:109: expected false, but got true: check all the things
        chkr_test.go:110: expected nil, but got &errors.errorString{s:"FAIL!"}: should not have gotten error from that code
        chkr_test.go:112: expected equal, but "foo" != "bar"
        chkr_test.go:113: expected not equal, but both are "foo": just do not understand 1, 2, 3
        chkr_test.go:114: expected errors.Is(&errors.errorString{s:"FAIL!"}, &errors.errorString{s:"sql: no rows in result set"}) to return true: wrong error
