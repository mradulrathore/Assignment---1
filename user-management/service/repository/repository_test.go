package repository

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestLoad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepositoryI(ctrl)

	repo.EXPECT().Load().Return(nil).Times(1)

}
