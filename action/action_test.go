package action

import "testing"

func TestAction(t *testing.T) {
	var testAction Action
	if testAction != nil {
		t.Error("test action is not nil")
	}
}
