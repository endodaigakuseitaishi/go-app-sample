package services_test

import (
	"errors"
	"go-todo-sample/app/models"
	"go-todo-sample/app/services"
	"regexp"
	"testing"
)

// DBのモックを容易して以下のケーそもテストしないと。。
// if passwordCount > 0 { return errors.New("既に存在するメールアドレスです") }

func TestCheckUser(t *testing.T) {
	tests := []struct {
		name     string
		user     models.User
		expected error
	}{
		{
			name: "Valid user",
			user: models.User{
				Name:     "John Doe",
				Email:    "john@example.com",
				PassWord: "password123",
			},
			expected: nil,
		},
		{
			name: "Short name",
			user: models.User{
				Name:     "Jo",
				Email:    "john@example.com",
				PassWord: "password123",
			},
			expected: errors.New("ユーザー名は3文字以上で入力してください"),
		},
		{
			name: "Long name",
			user: models.User{
				Name:     "ttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
				Email:    "john@example.com",
				PassWord: "password123",
			},
			expected: errors.New("ユーザー名は100文字以内で入力してください"),
		},
		{
			name: "Invalid email",
			user: models.User{
				Name:     "John Doe",
				Email:    "invalid_email",
				PassWord: "password123",
			},
			expected: errors.New("正しいメールアドレスの形式で入力してください"),
		},
		{
			name: "Short password",
			user: models.User{
				Name:     "John Doe",
				Email:    "john@example.com",
				PassWord: "pass",
			},
			expected: errors.New("パスワードは6文字以上で入力してください"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := services.CheckUser(test.user, models.Db)
			if (err != nil && test.expected == nil) || (err == nil && test.expected != nil) || (err != nil && test.expected != nil && err.Error() != test.expected.Error()) {
				t.Errorf("Expected error: %v, but got: %v", test.expected, err)
			}
		})
	}
}

func TestEmailRegex(t *testing.T) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	validEmails := []string{
		"test@example.com",
		"john.doe@example.com",
		"jane_doe123@example.co.uk",
	}

	invalidEmails := []string{
		"invalid_email",
		"test@example",
		"test@.com",
		"test@example.",
	}

	for _, email := range validEmails {
		if !emailRegex.MatchString(email) {
			t.Errorf("Expected email %s to match the regex pattern, but it did not", email)
		}
	}

	for _, email := range invalidEmails {
		if emailRegex.MatchString(email) {
			t.Errorf("Expected email %s to not match the regex pattern, but it did", email)
		}
	}
}