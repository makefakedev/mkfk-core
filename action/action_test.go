package action

import (
	"fmt"
	"testing"
	"log/slog"
)

type MockAction struct {}

func (MockAction) Act() Result {
	return "result"
}

func GetMockAction(t *testing.T) *MockAction {
	t.Helper()
	return new(MockAction)
}

func TestAction(t *testing.T) {
	var mockAction *MockAction = GetMockAction(t)
	var testMockAction Action = mockAction
	var testActionResult Result = testMockAction.Act()
	slog.Debug(fmt.Sprintf("%v", testActionResult))
}
