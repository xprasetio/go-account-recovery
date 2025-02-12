package memberships

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"xprasetio/go-account-recovery.git/internal/models/memberships"
)

func Test_repository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	ctx := context.Background()

	type args struct {
		ctx   context.Context
		model memberships.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				model: memberships.User{
					Email:     "xprasetio@gmail.com",
					Username:  "xprasetio",
					Password:  "admin789",
					CreatedBy: "xprasetio@gmail.com",
					UpdatedBy: "xprasetio@gmail.com",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.Email,
						args.model.Username,
						args.model.Password,
						args.model.RecoverCode,
						args.model.CreatedBy,
						args.model.UpdatedBy,
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
				model: memberships.User{
					Email:     "xprasetio@gmail.com",
					Username:  "xprasetio",
					Password:  "admin789",
					CreatedBy: "xprasetio@gmail.com",
					UpdatedBy: "xprasetio@gmail.com",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.Email,
						args.model.Username,
						args.model.Password,
						args.model.RecoverCode,
						args.model.CreatedBy,
						args.model.UpdatedBy,
					).
					WillReturnError(assert.AnError)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			if err := r.CreateUser(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			//memastikan semua mock terpanggil dengan benar
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	ctx := context.Background()

	now := time.Now()
	type args struct {
		ctx         context.Context
		email       string
		username    string
		id          uint
		recoverCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *memberships.User
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx:         ctx,
				email:       "xprasetio@gmail.com",
				username:    "xprasetio",
				id:          1,
				recoverCode: "coderecover",
			},
			want: &memberships.User{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: gorm.DeletedAt{},
				},
				Email:       "xprasetio@gmail.com",
				Username:    "xprasetio",
				Password:    "admin789",
				RecoverCode: "coderecover",
				CreatedBy:   "xprasetio@gmail.com",
				UpdatedBy:   "xprasetio@gmail.com",
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" WHERE \(email = \$1 OR username = \$2 OR recover_code = \$3 OR id = \$4\) AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$5`).
					WithArgs(args.email, args.username, args.recoverCode, args.id, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "email", "username", "password", "recover_code", "created_by", "updated_by"}).
						AddRow(1, now, now, nil, "xprasetio@gmail.com", "xprasetio", "admin789", "coderecover", "xprasetio@gmail.com", "xprasetio@gmail.com"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			got, err := r.GetUser(tt.args.ctx, tt.args.email, tt.args.username, tt.args.id, tt.args.recoverCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetUser() = %v, want %v", got, tt.want)
			}
			// memastikan semua mock terpanggil dengan benar
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
