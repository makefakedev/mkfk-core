package action

import (
	"fmt"
	"log/slog"
	"testing"
)

type MockAction struct{}

func (MockAction) Act() Result {
	return "result"
}

func (MockAction) Options() map[string]any {
	return map[string]any{"opt": "option"}
}

func GetMockAction(t *testing.T) *MockAction {
	t.Helper()
	return new(MockAction)
}

func TestAction(t *testing.T) {
	t.Run("action is an interface", func(t *testing.T) {
		var (
			mockAction     *MockAction = GetMockAction(t)
			testMockAction Action      = mockAction
		)
		slog.Debug(fmt.Sprintf("%v", testMockAction))
	})
	t.Run("action has method act yielding result", func(t *testing.T) {
		var (
			mockAction       *MockAction = GetMockAction(t)
			testMockAction   Action      = mockAction
			testActionResult Result      = testMockAction.Act()
		)
		slog.Debug(fmt.Sprintf("%v", testActionResult))
	})
	t.Run("action has method options yielding str to any", func(t *testing.T) {
		var(
			mockAction *MockAction = GetMockAction(t)
			testMockAction Action = mockAction
			testActionOptions map[string]any = testMockAction.Options()
		)
		slog.Debug(fmt.Sprintf("%v", testActionOptions))
	})
}
