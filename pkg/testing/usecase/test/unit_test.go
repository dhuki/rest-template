package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/dhuki/rest-template/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAllData(t *testing.T) {
	mockRepo := new(MockTestTableRepo)
	mockEmail := new(MockEmail)

	utils := utils.NewUtils().WireWithEmail(mockEmail)

	var data entity.TestTable
	err := faker.FakeData(&data)
	if err != nil {
		t.Errorf("%s", assert.AnError)
		return
	}

	mockRepo.On("GetAll", context.TODO()).Return([]entity.TestTable{data}, nil)

	usecase := usecase.UsecaseImpl{
		TestTableRepo: mockRepo,
		Utils:         utils,
	}
	actual := usecase.GetAllData(context.TODO())

	expected := common.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    []entity.TestTable{data},
		Error:   nil,
	}

	fmt.Println(data)
	assert.Equal(t, actual, expected, "it's supposed to same")
}
