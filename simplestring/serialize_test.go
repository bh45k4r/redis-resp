package simplestring

import "testing"

func TestSerializeNonReplaced(t *testing.T) {
	e := "+OK\r\n"
	if v := Serialize("OK"); v != e {
		t.Errorf("expected: %s, but got: %s", e, v)
	}
}

func TestSerializeReplaced(t *testing.T) {
	e := "+Hello__World\r\n"
	if v := Serialize("Hello\r\nWorld"); v != e {
		t.Errorf("expected: %s, but got: %s", e, v)
	}
}
