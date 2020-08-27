package validations

import (
	"errors"
	"regexp"
)

// StationSearchValidate 駅名検索機能で入力されたキーワードのフォーマット確認をするためのバリデーション処理
func StationSearchValidate(word string) (string, error) {
	// 全文字平仮名か片仮名の場合
	match, _ := regexp.MatchString("^[ぁ-んァ-ン]+$", word)
	if match {
		return "HiraganaOrKatakana", nil
	}
	// 漢字が含まれている場合
	match, _ = regexp.MatchString("^[ぁ-んァ-ン一-龥]+$", word)
	if match {
		return "AndKanji", nil
	}
	// それ以外（ローマ字、空欄等の場合）
	return "unknown", errors.New("不正な形式です")
}
