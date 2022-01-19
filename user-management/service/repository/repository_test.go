package repository

// import (
// 	"errors"
// 	"testing"

// 	gomock "github.com/golang/mock/gomock"
// 	usr "github.com/mradulrathore/user-management/service/user"

// 	"github.com/stretchr/testify/require"
// )

// func TestLoad(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	repo := NewRepo()
// 	err := repo.Load()
// 	require.Nil(t, err)
// }

// func TestAdd(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	repo := NewMockRepositoryI(ctrl)

// 	user, err := usr.New("Mradul", 21, "", 43, []string{"A", "B", "C", "D"})
// 	require.NotNil(t, err)

// 	repo.EXPECT().Add(user).Return(errors.New("dummy error")).Times(1)

// 	err = repo.Add(user)

// 	require.NotNil(t, err)
// }

// func TestCheckDataExistence(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	repo := NewMockRepositoryI(ctrl)

// 	user, err := usr.New("Mradul", 1, "Indore", 43, []string{"A", "B", "C", "D"})
// 	require.NotNil(t, err)
// 	repo.EXPECT().Add(user).Return(nil).Times(1)

// 	repo.EXPECT().CheckDataExistence(1).Return(false).Times(1)

// 	exist := repo.CheckDataExistence(1)

// 	require.Equal(t, exist, false)
// }
