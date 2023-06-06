package services_test

import (
	"errors"
	"go-todo-sample/app/services"
	"testing"
)

func TestCheckContentLength(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected error
	}{
		{
			name:     "Valid content",
			content:  "Valid todo content",
			expected: nil,
		},
		{
			name:     "Content too short",
			content:  "A",
			expected: errors.New("todoの内容は3文字以上、100文字以下にしましょう"),
		},
		{
			name:     "Content too long",
			content:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eget leo et ligula ultricies rutrum ut non ligula.",
			expected: errors.New("todoの内容は3文字以上、100文字以下にしましょう"),
		},
		{
			name:     "Content with newline",
			content:  "Todo\nwith\nnewline",
			expected: errors.New("ひとつのtodoに改行を含んではいけません"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := services.CheckContentLength(test.content)
			if err != nil && err.Error() != test.expected.Error() {
				t.Errorf("Expected error: %v, but got: %v", test.expected, err)
			} else if err == nil && test.expected != nil {
				t.Errorf("Expected error: %v, but got nil", test.expected)
			}
		})
	}
}