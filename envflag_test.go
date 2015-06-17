package envflag

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var TESTKEY = "__ENVFLAG_TEST__"
var NOARGS = []string{}

func TestTrueBool(t *testing.T) {
	if err := boolTest("true", true, false); err != nil {
		t.Fatal(err)
	}

}
func TestFalseBool(t *testing.T) {
	if err := boolTest("false", false, false); err != nil {
		t.Fatal(err)
	}
}
func TestErrBool(t *testing.T) {
	if err := boolTest("monkey", true, true); err != nil {
		t.Fatal(err)
	}
}

func boolTest(val string, expected bool, expectErr bool) error {
	fs := setup(val)
	b := fs.Bool(TESTKEY, !expected, "Test")
	err := chkErr(val, expectErr, ParseFlagSet(fs, NOARGS))

	if err != nil || expectErr {
		return err
	}
	if *b != expected {
		return fmt.Errorf("Expected: %s to result in %v", val, expected)
	}
	return nil
}

func chkErr(val string, expectErr bool, err error) error {
	if expectErr {
		if err == nil {
			return fmt.Errorf("Expected error for value: %s", val)
		}
		return nil
	} else {
		if err != nil {
			return fmt.Errorf("Expected no error for value: %s, got: %s", val, err)
		}
	}
	return nil
}
func setup(val string) *flag.FlagSet {
	os.Setenv(TESTKEY, val)
	return flag.NewFlagSet("test", flag.ContinueOnError)
}
