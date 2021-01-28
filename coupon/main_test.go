package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"github.com/tavo/prueba/coupon/models"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) GetItemByID(ids string) ([]models.Item, error) {
	args := r.Called(ids)
	return args.Get(0).([]models.Item), args.Error(1)
}

func TestHandler_Handler(t *testing.T) {
	type mocks struct {
		repository *RepositoryMock
	}
	type args struct {
		ctx      context.Context
		req      events.APIGatewayProxyRequest
		ids      string
		itemsGet []models.Item
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    Response
		wantErr bool
		mocker  func(args, mocks)
	}{
		{
			name:    "success_response",
			wantErr: false,
			args: args{
				ctx: context.Background(),
				req: events.APIGatewayProxyRequest{
					Headers: map[string]string{
						"Content-Type":                     "application/json",
						"Access-Control-Allow-Origin":      "*",
						"Access-Control-Allow-Credentials": "true",
					},
					Body: `{
						"item_ids": ["MLA1"],
						"amount": 350
					}`,
				},
				ids: "MLA1",
				itemsGet: []models.Item{
					{
						ID:    "MLA1",
						Price: 100,
					},
				},
			},
			want: Response{
				StatusCode: http.StatusOK,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"item_ids":["MLA1"],"total":100}`,
			},
			mocks: mocks{
				repository: &RepositoryMock{},
			},
			mocker: func(a args, m mocks) {
				m.repository.On("GetItemByID", a.ids).Return(a.itemsGet, nil).Once()
			},
		},
		{
			name:    "error_amount",
			wantErr: true,
			args: args{
				ctx: context.Background(),
				req: events.APIGatewayProxyRequest{
					Headers: map[string]string{
						"Content-Type":                     "application/json",
						"Access-Control-Allow-Origin":      "*",
						"Access-Control-Allow-Credentials": "true",
					},
					Body: `{}`,
				},
			},
			want: Response{
				StatusCode: http.StatusNotFound,
				Body:       "ammount insufficiency : ValidatePriceMin",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			mocks: mocks{
				repository: &RepositoryMock{},
			},
			mocker: func(a args, m mocks) {
				m.repository.On("GetItemByID", a.ids).Return(a.itemsGet, nil).Once()
			},
		},
		{
			name:    "error_unmarshal",
			wantErr: true,
			args: args{
				ctx: context.Background(),
				req: events.APIGatewayProxyRequest{
					Headers: map[string]string{
						"Content-Type":                     "application/json",
						"Access-Control-Allow-Origin":      "*",
						"Access-Control-Allow-Credentials": "true",
					},
				},
			},
			want: Response{
				StatusCode: http.StatusNotFound,
			},
			mocks: mocks{
				repository: &RepositoryMock{},
			},
			mocker: func(a args, m mocks) {},
		},
		{
			name:    "error_getItem()",
			wantErr: true,
			args: args{
				ctx: context.Background(),
				req: events.APIGatewayProxyRequest{
					Headers: map[string]string{
						"Content-Type":                     "application/json",
						"Access-Control-Allow-Origin":      "*",
						"Access-Control-Allow-Credentials": "true",
					},
					Body: `{
						"item_ids": ["MLA1"],
						"amount": 350
					}`,
				},
				itemsGet: []models.Item{
					{
						ID:    "MLA1",
						Price: 100,
					},
				},
				ids: "MLA1",
			},
			want: Response{
				StatusCode: http.StatusNotFound,
			},
			mocks: mocks{
				repository: &RepositoryMock{},
			},
			mocker: func(a args, m mocks) {
				m.repository.On("GetItemByID", a.ids).Return([]models.Item{}, errors.New("error")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.mocks)
			Handler := Adapter(tt.mocks.repository)
			got, _ := Handler(tt.args.ctx, tt.args.req)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Error in handler, (-want,+got)\n%s", diff)
			}
			tt.mocks.repository.AssertExpectations(t)
		})
	}
}
