package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	hash := "e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746"
	s := struct{ Test string }{Test: "test"}

	t.Run("Hash is always same", func(t *testing.T) {
		x := Hash(s)
		if x != hash {
			t.Errorf("\nExpected\t%s\nGot\t\t%s", hash, x)
		}
	})

	t.Run("Hash is hex encoded", func(t *testing.T) {
		x := Hash(s)
		_, err := hex.DecodeString(x)
		if err != nil {
			t.Error("Hash should be hex encoded")
		}
	})
}

func ExampleHash() {
	s := struct{ Test string }{Test: "test"}
	x := Hash(s)
	fmt.Println(x)
	// Output: e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746
}

func TestTobytes(t *testing.T) {
	s := "test"
	b := ToBytes(s)
	k := reflect.TypeOf(b).Kind()
	if k != reflect.Slice {
		t.Errorf("Tobytes should return a slice of bytes got %s", k)
	}
}

func TestSplitter(t *testing.T) {
	type test struct {
		input  string
		sep    string
		index  int
		output string
	}

	tests := []test{
		{input: "0:1:2", sep: ":", index: 1, output: "1"},
		{input: "0:1:2", sep: ":", index: 10, output: ""},
		{input: "0:1:2", sep: "/", index: 0, output: "0:1:2"},
		{input: "0:1:2", sep: "/", index: 1, output: ""},
	}

	for _, tc := range tests {
		got := Splitter(tc.input, tc.sep, tc.index)
		if got != tc.output {
			t.Errorf("\nExpected\t%s\nGot\t\t%s", tc.output, got)
		}
	}
}

func TestHandleErr(t *testing.T) {
	oldLogFn := logFn
	defer func() {
		logFn = oldLogFn
	}()
	called := false
	logFn = func(v ...interface{}) {
		called = true
	}
	err := errors.New("test")
	HandleErr(err)
	if !called {
		t.Error("HandleError should call logFn")
	}
}

func TestDecodeStringOrErr(t *testing.T) {
	oldLogFn := logFn
	defer func() {
		logFn = oldLogFn
	}()
	called := false
	logFn = func(v ...interface{}) {
		called = true
	}

	tests := []string{
		"!",
		"ABCDEFG",
		"41g",
		"000BAA@",
	}

	for _, tc := range tests {
		DecodeStringOrErr(tc)
		if !called {
			t.Error("DecodeStringOrErr should call logFn if the input is improper")
		}
		called = false
	}
}

func TestFromBytes(t *testing.T) {
	type testStruct struct {
		Test string
	}
	var restored testStruct

	ts := testStruct{"test"}
	b := ToBytes(ts)

	FromeBytes(&restored, b)
	if !reflect.DeepEqual(ts, restored) {
		t.Error("FromBytes() should restore struct.")
	}
}

func TestToJson(t *testing.T) {
	type testStruct struct {
		Test string
	}
	var restored testStruct

	ts := testStruct{"test"}
	b := ToJSON(ts)

	t.Run("JSON Marsharling", func(t *testing.T) {
		k := reflect.TypeOf(b).Kind()
		if k != reflect.Slice {
			t.Errorf("Expected %v and got %v", reflect.Slice, k)
		}
	})

	t.Run("JSON Unmarsharling", func(t *testing.T) {
		json.Unmarshal(b, &restored)
		if !reflect.DeepEqual(ts, restored) {
			t.Error("ToJSON() should encode to JSON corretly.")
		}
	})
}
