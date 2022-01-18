package repository

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	usr "github.com/mradulrathore/user-management/service/user"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewRepo()
	err := repo.Load()
	require.Nil(t, err)
}

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepositoryI(ctrl)

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	require.Nil(t, err)

	repo.EXPECT().Add(user).Return(nil).Times(1)

	err = repo.Add(user)

	require.Nil(t, err)
}
