package persistence

// func TestComment_FindAll(t *testing.T) {
// 	type args struct {
// 		pid uint32
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *[]models.Comment
// 		wantErr bool
// 	}{
// 		{
// 			name: "指定した投稿IDに紐付くコメント情報を取得出来ること",
// 			args: args{
// 				pid: 1,
// 			},
// 			want: &[]models.Comment{
// 				{
// 					ID:        1,
// 					Text:      "mikuのコメント001",
// 					UserID:    1,
// 					PostID:    1,
// 					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 				},
// 				{
// 					ID:        2,
// 					Text:      "takaのコメント001",
// 					UserID:    2,
// 					PostID:    1,
// 					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "指定した投稿IDに紐付くコメントがない場合、空の値を返す",
// 			args: args{
// 				pid: 2,
// 			},
// 			want:    &[]models.Comment{},
// 			wantErr: false,
// 		},
// 	}
// 	tx := db.Begin()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			cp := NewCommentPersistence(tx)
// 			got, err := cp.FindAll(tt.args.pid)
// 			// 予期しないエラーの場合
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("commentPersistence.FindAll() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			// 返り値が期待しない値の場合
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("commentPersistence.FindAll() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// 	tx.Rollback()
// }

// func TestComment_Save(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		args    *models.Comment
// 		want    *models.Comment
// 		wantErr bool
// 	}{
// 		{
// 			name: "255文字以内のコメントを保存出来ること",
// 			args: &models.Comment{
// 				Text:      strings.Repeat("a", 255),
// 				UserID:    1,
// 				PostID:    1,
// 				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 			},
// 			want: &models.Comment{
// 				ID:        3,
// 				Text:      strings.Repeat("a", 255),
// 				UserID:    1,
// 				PostID:    1,
// 				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "コメントが空でも保存出来ること",
// 			args: &models.Comment{
// 				Text:      "",
// 				UserID:    1,
// 				PostID:    1,
// 				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 			},
// 			want: &models.Comment{
// 				ID:        4,
// 				Text:      "",
// 				UserID:    1,
// 				PostID:    1,
// 				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "コメントが256字以上は保存出来ないこと",
// 			args: &models.Comment{
// 				Text:      strings.Repeat("a", 256),
// 				UserID:    1,
// 				PostID:    1,
// 				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
// 			},
// 			want:    &models.Comment{},
// 			wantErr: true,
// 		},
// 	}
// 	tx := db.Begin()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			cp := NewCommentPersistence(tx)
// 			err := cp.Save(tt.args)
// 			// 予期しないエラーの場合
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("commentPersistence.Save() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// 	tx.Rollback()
// }

// func TestComment_Update(t *testing.T) {
// 	type args struct {
// 		cid     uint32
// 		comment *models.Comment
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    int64
// 		wantErr bool
// 	}{
// 		{
// 			name: "255文字以内のコメントを更新出来ること",
// 			args: args{
// 				cid: 1,
// 				comment: &models.Comment{
// 					Text: strings.Repeat("a", 255),
// 				},
// 			},
// 			want:    1,
// 			wantErr: false,
// 		},
// 		{
// 			name: "256文字以上のコメントは更新出来ないこと",
// 			args: args{
// 				cid: 1,
// 				comment: &models.Comment{
// 					Text: strings.Repeat("a", 256),
// 				},
// 			},
// 			want:    0,
// 			wantErr: true,
// 		},
// 		{
// 			name: "指定のコメントIDが存在しなければ、エラーが返ること",
// 			args: args{
// 				cid: 10,
// 				comment: &models.Comment{
// 					Text: strings.Repeat("a", 255),
// 				},
// 			},
// 			want:    0,
// 			wantErr: true,
// 		},
// 	}
// 	tx := db.Begin()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			cp := NewCommentPersistence(tx)
// 			got, err := cp.Update(tt.args.cid, tt.args.comment)
// 			// 予期しないエラーの場合
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("commentPersistence.Update() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			// 返り値が期待しない値の場合
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("commentPersistence.Update() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// 	tx.Rollback()
// }

// func TestComment_Delete(t *testing.T) {
// 	type args struct {
// 		cid uint32
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    int64
// 		wantErr bool
// 	}{
// 		{
// 			name: "コメント情報を削除出来ること",
// 			args: args{
// 				cid: 1,
// 			},
// 			want:    1,
// 			wantErr: false,
// 		},
// 	}
// 	tx := db.Begin()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			cp := NewCommentPersistence(tx)
// 			err := cp.Delete(tt.args.cid)
// 			// 予期しないエラーの場合
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("commentPersistence.Delete() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// 	tx.Rollback()
// }
