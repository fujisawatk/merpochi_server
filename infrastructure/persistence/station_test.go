package persistence

import (
	"testing"
)

func TestStation_SearchKanaWord(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    int
		wantErr bool
	}{
		{
			name:    "平仮名のキーワードで駅名を検索出来ること",
			args:    "つるみ",
			want:    6,
			wantErr: false,
		},
		{
			name:    "カタカナのキーワードで駅名を検索出来ること",
			args:    "ツルミ",
			want:    6,
			wantErr: false,
		},
		{
			name:    "半角カタカナのキーワードで駅名を検索出来ること",
			args:    "ﾂﾙﾐ",
			want:    6,
			wantErr: false,
		},
		{
			name:    "該当する駅名は10件まで検索出来ること",
			args:    "つる",
			want:    10,
			wantErr: false,
		},
		{
			name:    "該当する駅名が存在しない場合、空の配列を返す",
			args:    "わわ",
			want:    0,
			wantErr: false,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewStationPersistence(tx)
			got, err := sp.SearchKanaWord(tt.args)
			if (err != nil) != tt.wantErr {
				// 予期しないエラーの場合
				t.Errorf("stationPersistence.SearchKanaWord() error = %v, wantErr %v", err, tt.wantErr)
			}
			// テストケースが冗長になるため、取得するレコード数で比較
			if len(got) != tt.want {
				t.Errorf("stationPersistence.SearchKanaWord() = %v, want %v", len(got), tt.want)
			}
		})
	}
	tx.Rollback()
}

func TestStation_SearchKanjiWord(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    int
		wantErr bool
	}{
		{
			name:    "漢字のキーワードで駅名を検索出来ること",
			args:    "鶴見",
			want:    6,
			wantErr: false,
		},
		{
			name:    "該当する駅名は10件まで検索出来ること",
			args:    "鶴",
			want:    10,
			wantErr: false,
		},
		{
			name:    "該当する駅名が存在しない場合、空の配列を返す",
			args:    "亜",
			want:    0,
			wantErr: false,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewStationPersistence(tx)
			got, err := sp.SearchKanjiWord(tt.args)
			if (err != nil) != tt.wantErr {
				// 予期しないエラーの場合
				t.Errorf("stationPersistence.SearchKanjiWord() error = %v, wantErr %v", err, tt.wantErr)
			}
			// テストケースが冗長になるため、取得するレコード数で比較
			if len(got) != tt.want {
				t.Errorf("stationPersistence.SearchKanjiWord() = %v, want %v", len(got), tt.want)
			}
		})
	}
	tx.Rollback()
}
