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

//TODO
func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepositoryI(ctrl)
	repo.EXPECT().Load().Return(nil).Times(1)

	user := usr.User{
		Name: "Mradul",
		Age: 21,
		Address: "Indore",
		RollNo: 43,
		Courses: ,
	}

}
