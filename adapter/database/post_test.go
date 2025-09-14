package database_test

import (
	"testing"
	"time"

	"github.com/yuita-yoshihiko/go-sample-api/adapter/database"
	"github.com/yuita-yoshihiko/go-sample-api/infrastructure/db"
	"github.com/yuita-yoshihiko/go-sample-api/models"
	"github.com/yuita-yoshihiko/go-sample-api/testutils"
)

func Test_User_FetchByUserID(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/posts/fetch_by_user_id")
	dbUtils := db.NewDBUtil(data)
	r := database.NewPostRepository(dbUtils)

	type args struct {
		userID int64
	}

	tests := []struct {
		name string
		args args
		want []*models.Post
	}{
		{
			name: "user_idで抽出した複数の投稿のデータが取得できる",
			args: args{
				userID: 1,
			},
			want: []*models.Post{
				{
					ID:        1,
					UserID:    1,
					Title:     "テスト投稿1",
					Content:   "テスト投稿1の内容",
					CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        2,
					UserID:    1,
					Title:     "テスト投稿2",
					Content:   "テスト投稿2の内容",
					CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FetchByUserID(t.Context(), tt.args.userID)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			testutils.AssertResponse(t, got, tt.want)
		})
	}
}

func Test_User_FetchByUserIDWithComments(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/posts/fetch_by_user_id_with_comments")
	dbUtils := db.NewDBUtil(data)
	r := database.NewPostRepository(dbUtils)

	type args struct {
		userID int64
	}

	tests := []struct {
		name string
		args args
		want []*models.PostWithComments
	}{
		{
			name: "user_idで抽出した複数の投稿とそのコメントのデータが取得できる",
			args: args{
				userID: 1,
			},
			want: []*models.PostWithComments{
				{
					Post: models.Post{
						ID:        1,
						UserID:    1,
						Title:     "テスト投稿1",
						Content:   "テスト投稿1の内容",
						CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					Comments: []models.Comment{
						{
							ID:        1,
							PostID:    1,
							Content:   "テストコメント1の内容",
							CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							PostID:    1,
							Content:   "テストコメント2の内容",
							CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				{
					Post: models.Post{
						ID:        2,
						UserID:    1,
						Title:     "テスト投稿2",
						Content:   "テスト投稿2の内容",
						CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					Comments: []models.Comment{
						{
							ID:        3,
							PostID:    2,
							Content:   "テストコメント3の内容",
							CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FetchByUserIDWithComments(t.Context(), tt.args.userID)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			testutils.AssertResponse(t, got, tt.want)
		})
	}
}
