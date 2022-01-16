package usecase

import (
	"context"
	"fmt"
	"github.com/Eretic431/fibonacci/internal/fibonacci/mock"
	"github.com/Eretic431/fibonacci/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestFibonacciUseCase_GetSliceArgumentsValidation(t *testing.T) {
	input := []struct {
		x int
		y int
	}{
		{
			0, 0,
		},
		{
			0, 3,
		},
		{
			0, -3,
		},
		{
			-1, 0,
		},
		{
			-1, 2,
		},
		{
			-1, -2,
		},
		{
			1, 0,
		},
		{
			1, -1,
		},
		{
			3, 1,
		},
	}

	for i, v := range input {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock.NewMockFibonacciRepository(ctrl)
			uc := NewFibonacciUseCase(repo)

			ctx := context.Background()
			output, err := uc.GetSlice(ctx, v.x, v.y)
			require.Nil(t, output)
			require.EqualError(t, err, models.ErrInvalidArguments.Error())
		})
	}
}

func TestFibonacciUseCase_GetSliceCalculation(t *testing.T) {
	type mockBehaviour func(s *mock.MockFibonacciRepository)
	testCases := []struct {
		name          string
		x             int
		y             int
		mockBehaviour mockBehaviour
	}{
		{
			"#1 case",
			5,
			10,
			func(s *mock.MockFibonacciRepository) {
				s.EXPECT().GetLastTwoNumbers(gomock.Any()).Return([]int64{21, 9, 34, 10}, nil)
				s.EXPECT().GetInterval(gomock.Any(), 5, 10).Return([]int64{3, 5, 8, 13, 21, 34}, nil)
			},
		},
		{
			"#2 case",
			5,
			10,
			func(s *mock.MockFibonacciRepository) {
				s.EXPECT().GetLastTwoNumbers(gomock.Any()).Return([]int64{2, 4, 3, 5}, nil)
				s.EXPECT().Set(gomock.Any(), 6, int64(5)).Return(nil)
				s.EXPECT().Set(gomock.Any(), 7, int64(8)).Return(nil)
				s.EXPECT().Set(gomock.Any(), 8, int64(13)).Return(nil)
				s.EXPECT().Set(gomock.Any(), 9, int64(21)).Return(nil)
				s.EXPECT().Set(gomock.Any(), 10, int64(34)).Return(nil)
			},
		},
		{
			"#3 case",
			5,
			10,
			func(s *mock.MockFibonacciRepository) {
				s.EXPECT().GetLastTwoNumbers(gomock.Any()).Return([]int64{8, 7, 13, 8}, nil)
				s.EXPECT().GetInterval(gomock.Any(), 5, 8).Return([]int64{3, 5, 8, 13}, nil)
				s.EXPECT().Set(gomock.Any(), 9, int64(21)).Return(nil)
				s.EXPECT().Set(gomock.Any(), 10, int64(34)).Return(nil)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock.NewMockFibonacciRepository(ctrl)
			testCase.mockBehaviour(repo)
			uc := NewFibonacciUseCase(repo)

			ctx := context.Background()
			output, err := uc.GetSlice(ctx, testCase.x, testCase.y)
			require.NoError(t, err)
			require.True(t, reflect.DeepEqual([]int64{3, 5, 8, 13, 21, 34}, output))
		})
	}
}
