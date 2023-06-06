package services

import (
	"errors"
	"strings"
)

func CheckContentLength(content string) error {
	if len(content) < 2 || len(content) > 99 {
		return errors.New("todoの内容は3文字以上、100文字以下にしましょう")
	}

	if strings.Contains(content, "\n") {
		return errors.New("ひとつのtodoに改行を含んではいけません")
	}

	return nil
}