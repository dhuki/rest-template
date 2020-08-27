package test

import (
	"context"
	"testing"

	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/dhuki/rest-template/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTestTableRepo struct {
	mock.Mock
}

func (m mockTestTableRepo) GetAll(ctx context.Context) ([]entity.TestTable, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.TestTable), args.Error(1)
}

func (m mockTestTableRepo) Get(ctx context.Context) (entity.TestTable, error) {
	return entity.TestTable{}, nil
}

func (m mockTestTableRepo) GetByName(ctx context.Context, name string) (entity.TestTable, error) {
	return entity.TestTable{}, nil
}

func (m mockTestTableRepo) Create(ctx context.Context, data entity.TestTable) error {
	return nil
}

type mockEmail struct {
	mock.Mock
}

func (m mockEmail) SendEmail(ctx context.Context, data entity.TestTable) error {
	return nil
}

func TestGetAllData(t *testing.T) {
	mockRepo := new(mockTestTableRepo)
	mockEmail := new(mockEmail)

	utils := utils.NewUtils().WireWithEmail(mockEmail)

	mockRepo.On("GetAll", context.TODO()).Return([]entity.TestTable{
		{
			Name:        "testing",
			Description: "testing",
		},
		{
			Name:        "testing",
			Description: "testing",
		},
	}, nil)

	usecase := usecase.NewUsecase(mockRepo, utils)
	actual := usecase.GetAllData(context.TODO())

	expected := common.BaseResponse{
		Success: true,
		Message: "Success",
		Data: []entity.TestTable{
			{
				Name:        "testing",
				Description: "testing",
			},
			{
				Name:        "testing",
				Description: "testing",
			},
		},
		Error: nil,
	}

	assert.Equal(t, actual, expected, "it's supposed to same")
}
