package services

import (
	"database/sql"
	"errors"
	"fmt"
	"go-todo-sample/app/models"
	"regexp"
)

func CheckUser(user models.User, db *sql.DB) error {
	// Name
	if len(user.Name) < 3 {
		return errors.New("ユーザー名は3文字以上で入力してください")
	}

	if len(user.Name) > 100 {
		return errors.New("ユーザー名は100文字以内で入力してください")
	}

	// Email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return errors.New("正しいメールアドレスの形式で入力してください")
	}

	cmdExtractEmail := `SELECT COUNT(*) FROM users WHERE email = ?`
	fmt.Printf(cmdExtractEmail)
	var emailCount int
	emailExtractiongErr := db.QueryRow(cmdExtractEmail, user.Email).Scan(&emailCount)
	if emailExtractiongErr != nil { return emailExtractiongErr }
	if emailCount > 0 { return errors.New("既に存在するメールアドレスです") }

	// PassWord
	if len(user.PassWord) < 6 { return errors.New("パスワードは6文字以上で入力してください") }

	cmdExtractPassWord := `SELECT COUNT(*) FROM users WHERE email = ?`
	var passwordCount int
	passwordExtractingErr := db.QueryRow(cmdExtractPassWord, user.PassWord).Scan(&passwordCount)
	if passwordExtractingErr != nil { return passwordExtractingErr }
	if passwordCount > 0 { return errors.New("既に存在するメールアドレスです") }

	return nil
}