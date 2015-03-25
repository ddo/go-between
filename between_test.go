package between

import "testing"

func TestBasic(t *testing.T) {
	substr := Between("hello world", "hel", "ld")

	if substr != "lo wor" {
		t.Fail()
	}
}

func TestBasic2(t *testing.T) {
	substr := Between("i'm so cool", "m", "oo")

	if substr != " so c" {
		t.Fail()
	}
}

func TestNotFoundStart(t *testing.T) {
	substr := Between("hello world", "abc", "ld")

	if substr != "" {
		t.Fail()
	}
}

func TestNotFoundEnd(t *testing.T) {
	substr := Between("hello world", "hel", "abc")

	if substr != "" {
		t.Fail()
	}
}

func TestEmptyStr(t *testing.T) {
	substr := Between("", "", "")

	if substr != "" {
		t.Fail()
	}
}

func TestEmptyStart(t *testing.T) {
	substr := Between("hello world", "", "el")

	if substr != "h" {
		t.Fail()
	}
}

func TestEmptyEnd(t *testing.T) {
	substr := Between("hello world", "hel", "")

	if substr != "" {
		t.Fail()
	}
}

func TestDuplicatedEnd(t *testing.T) {
	substr := Between("ld hello world", "hel", "ld")

	if substr != "lo wor" {
		t.Fail()
	}
}
