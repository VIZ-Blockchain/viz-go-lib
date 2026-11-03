package operations

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnknownOperationRoundTrip(t *testing.T) {
	input := `[["pm_batch_settle",{"market_id":42,"future_field":[1,"x"]}],["future_operation",{"flag":true}]]`
	var ops Operations
	if err := json.Unmarshal([]byte(input), &ops); err != nil {
		t.Fatal(err)
	}
	for _, op := range ops {
		if _, ok := op.(*UnknownOperation); !ok {
			t.Fatalf("expected unknown operation, got %T", op)
		}
	}
	got, err := json.Marshal(ops)
	if err != nil {
		t.Fatal(err)
	}
	var wantValue, gotValue interface{}
	if err := json.Unmarshal([]byte(input), &wantValue); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(got, &gotValue); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(wantValue, gotValue) {
		t.Fatalf("round trip got %s, want %s", got, input)
	}
}

func TestInvalidOperationDoesNotPanic(t *testing.T) {
	for _, input := range []string{`[null]`, `[[null,{}]]`, `[["transfer",null]]`, `[["transfer"]]`} {
		t.Run(input, func(t *testing.T) {
			var ops Operations
			if err := json.Unmarshal([]byte(input), &ops); err == nil {
				t.Fatalf("accepted %s", input)
			}
		})
	}
	if _, err := json.Marshal(Operations{nil}); err == nil {
		t.Fatal("accepted nil operation")
	}
}
