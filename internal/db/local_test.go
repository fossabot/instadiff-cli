package db

import (
	"context"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"

	"github.com/oleg-balunenko/instadiff-cli/internal/models"
)

func Test_localDB_GetLastUsersBatchByType(t *testing.T) {
	type args struct {
		batchType models.UsersBatchType
	}

	tests := []struct {
		name    string
		args    args
		want    models.UsersBatch
		wantErr bool
	}{
		{
			name: "get followers",
			args: args{
				batchType: models.UsersBatchTypeFollowers,
			},
			want: models.UsersBatch{
				Users: followersFixture,
				Type:  models.UsersBatchTypeFollowers,
			},
			wantErr: false,
		},
		{
			name: "get unknown type",
			args: args{
				batchType: models.UsersBatchTypeUnknown,
			},
			want:    models.EmptyUsersBatch,
			wantErr: true,
		},
		{
			name: "get invalid type",
			args: args{
				batchType: 5,
			},
			want:    models.EmptyUsersBatch,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			l := setUpDbWithFixtures(t)

			got, err := l.GetLastUsersBatchByType(context.TODO(), tt.args.batchType)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_localDB_InsertUsersBatch(t *testing.T) {
	ldb := newLocalDB()
	bType := models.UsersBatchTypeFollowers

	gotBatch, err := ldb.GetLastUsersBatchByType(context.TODO(), bType)
	require.NoError(t, err)
	assert.Equal(t, models.EmptyUsersBatch, gotBatch)

	goldenBatch := models.UsersBatch{
		Users: followersFixture,
		Type:  models.UsersBatchTypeFollowers,
	}

	err = ldb.InsertUsersBatch(context.TODO(), goldenBatch)
	require.NoError(t, err)

	gotBatch, err = ldb.GetLastUsersBatchByType(context.TODO(), bType)
	require.NoError(t, err)

	assert.Equal(t, goldenBatch, gotBatch)
}

var (
	followersFixture = []models.User{
		{
			ID:       1,
			UserName: "user1",
			FullName: "test user 1",
		},
		{
			ID:       2,
			UserName: "user2",
			FullName: "test user 2",
		},
	}
	followingsFixture = []models.User{
		{
			ID:       3,
			UserName: "user3",
			FullName: "test user 3",
		},
		{
			ID:       4,
			UserName: "user4",
			FullName: "test user 4",
		},
	}
)

func setUpDbWithFixtures(t testing.TB) DB {
	t.Helper()

	fixtures := map[models.UsersBatchType]models.UsersBatch{
		models.UsersBatchTypeFollowers: {
			Users: followersFixture,
			Type:  models.UsersBatchTypeFollowers,
		},
		models.UsersBatchTypeFollowings: {
			Users: followingsFixture,
			Type:  models.UsersBatchTypeFollowings,
		},
	}

	db := newLocalDB()
	db.users = fixtures

	return db
}
