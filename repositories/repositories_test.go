package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/tavo/prueba/coupon/models"
)

func Test_GetItemByID(t *testing.T) {
	type args struct {
		IDs      string
		bodyItem []BodyItem
	}

	tests := []struct {
		name    string
		args    args
		mocker  func(args, *testing.T) *httptest.Server
		want    []models.Item
		wantErr bool
	}{
		{
			"get_item_by_id_test_success",
			args{
				IDs: "MLA905913105",
				bodyItem: []BodyItem{
					{
						Code: 200,
						Body: models.Item{
							ID:    "MLA905913105",
							Price: 500,
						},
					},
				},
			},
			func(a args, t *testing.T) *httptest.Server {
				hs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Println(r.URL.Path)
					if r.URL.Path != "/items" {
						t.Error("Invalid path")
					}

					w.WriteHeader(http.StatusOK)

					body, err := json.Marshal(a.bodyItem)
					if err != nil {
						t.Error("Error while marshal body")
					}

					w.Write(body)
				}))
				return hs
			},
			[]models.Item{
				{
					ID:    "MLA905913105",
					Price: 500,
				},
			},
			false,
		},
		{
			"get_item_by_id_test_not_found",
			args{
				IDs: "MLA905913105",
				bodyItem: []BodyItem{
					{
						Code: 200,
						Body: models.Item{
							ID:    "MLA905913105",
							Price: 500,
						},
					},
				},
			},
			func(a args, t *testing.T) *httptest.Server {
				hs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
					body, err := json.Marshal(map[string]interface{}{
						"code": "404",
					})
					if err != nil {
						t.Error("Error while marshal body")
					}
					w.Write(body)
				}))
				return hs
			},
			[]models.Item{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := tt.mocker(tt.args, t)
			defer hs.Close()
			repository := NewRepository(
				hs.Client(),
				hs.URL,
			)

			got, err := repository.GetItemByID(tt.args.IDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItemByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
