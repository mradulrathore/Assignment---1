package service

// cmd go test -coverprofile=coverage.out

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mradulrathore/item-inventory/service/enum"
)

//TODO
func TestInitialize(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := NewMockDB(ctrl)

	raw, _ := enum.ItemTypeString("raw")
	listItem := &ListItems{
		Items: []Item{{
			Name:     "Keybord",
			Price:    100.00,
			Quantity: 1,
			Type:     raw,
		}}}

	db.EXPECT().GetItems().Return(&listItem, nil).Times(1)

	// repo := NewRepo(db)
	// result, err := getItemsFromDB(repo)
	// require.Nil(t, err)
	// require.NotNil(t, result)
}
