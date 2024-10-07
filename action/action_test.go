package action

import (
	"fmt"
	"testing"
	"log/slog"
)

type MockAction struct {}

func GetMockAction(t *testing.T) *MockAction {
	t.Helper()
	return new(MockAction)
}

func TestAction(t *testing.T) {
	var mockAction *MockAction = GetMockAction(t)
	var testMockAction Action = mockAction
	slog.Debug(fmt.Sprintf("%v", testMockAction))
}
