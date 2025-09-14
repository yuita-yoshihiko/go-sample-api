package database_test

import (
	"errors"
	"testing"
	"time"

	"github.com/yuita-yoshihiko/go-sample-api/adapter/database"
	"github.com/yuita-yoshihiko/go-sample-api/infrastructure/db"
	"github.com/yuita-yoshihiko/go-sample-api/models"
	"github.com/yuita-yoshihiko/go-sample-api/testutils"
)

func Test_User_Fetch(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/users/fetch")
	dbUtils := db.NewDBUtil(data)
	r := database.NewUserRepository(dbUtils)

	type args struct {
		id int64
	}

	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "idで抽出した単一のユーザーのデータが取得できる",
			args: args{
				id: 1,
			},
			want: &models.User{
				ID:        1,
				Name:      "テストユーザー1",
				Email:     "user1@example.com",
				CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Fetch(t.Context(), tt.args.id)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			testutils.AssertResponse(t, got, tt.want)
		})
	}
}

func Test_User_FetchWithPosts(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/users/fetch_with_posts")
	dbUtils := db.NewDBUtil(data)
	r := database.NewUserRepository(dbUtils)

	type args struct {
		id int64
	}

	tests := []struct {
		name string
		args args
		want *models.UserWithPosts
	}{
		{
			name: "idで抽出した単一のユーザーのデータと、そのユーザーが投稿した記事のデータが取得できる",
			args: args{
				id: 1,
			},
			want: &models.UserWithPosts{
				User: models.User{
					ID:        1,
					Name:      "テストユーザー1",
					Email:     "user1@example.com",
					CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Posts: []models.Post{
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FetchWithPosts(t.Context(), tt.args.id)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			testutils.AssertResponse(t, got, tt.want)
		})
	}
}

func Test_User_Create(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/users/create")
	dbUtils := db.NewDBUtil(data)
	r := database.NewUserRepository(dbUtils)

	type args struct {
		user *models.User
	}

	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "ユーザーのデータが登録できる",
			args: args{
				user: &models.User{
					Name:  "テストユーザー1",
					Email: "user1@example.com",
				},
			},
			want: &models.User{
				Name:  "テストユーザー1",
				Email: "user1@example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := r.Create(t.Context(), tt.args.user)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			got, err := r.Fetch(t.Context(), id)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			opt := testutils.DefaultIgnoreFieldsOptions(models.User{})
			testutils.AssertResponseWithOption(t, got, tt.want, opt)
		})
	}
}

func Test_User_Update(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/users/update")
	dbUtils := db.NewDBUtil(data)
	r := database.NewUserRepository(dbUtils)

	type args struct {
		user *models.User
	}

	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "ユーザーのデータが更新できる",
			args: args{
				user: &models.User{
					ID:    1,
					Name:  "テストユーザー1-1",
					Email: "user1-1@example.com",
				},
			},
			want: &models.User{
				Name:  "テストユーザー1-1",
				Email: "user1-1@example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.Update(t.Context(), tt.args.user); err != nil {
				t.Errorf("error = %v", err)
			}
			got, err := r.Fetch(t.Context(), tt.args.user.ID)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			opt := testutils.DefaultIgnoreFieldsOptions(models.User{})
			testutils.AssertResponseWithOption(t, got, tt.want, opt)
		})
	}
}

func Test_User_Delete(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/users/delete")
	dbUtils := db.NewDBUtil(data)
	r := database.NewUserRepository(dbUtils)

	type args struct {
		id int64
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "ユーザーのデータが削除できる",
			args: args{
				id: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.Delete(t.Context(), tt.args.id); err != nil {
				t.Errorf("error = %v", err)
			}
			got, err := r.Fetch(t.Context(), tt.args.id)
			if err != nil && errors.Is(err, db.ErrNotFound) {
				t.Errorf("error = %v", err)
			}
			if got != nil {
				t.Errorf("Fetch() got = %v, want nil", got)
			}
		})
	}
}
