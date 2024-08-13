package service_test

import (
	"AsiaYo-Backend-pre-test/internal/model"
	"AsiaYo-Backend-pre-test/internal/service"
	"errors"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	type args struct {
		in *model.CreateOrder
	}

	tests := []struct {
		name     string
		args     args
		newPrice string
		want     error
	}{
		{
			name: "ok",
			args: args{
				in: &model.CreateOrder{
					ID:   "A0000001",
					Name: "Melody Holiday Inn",
					Address: model.Address{
						City:     "taipei-city",
						District: "da-an-district",
						Street:   "fuxing-south-road",
					},
					Price:    "1900",
					Currency: "TWD",
				},
			},
			want: nil,
		},
		{
			name: "Name contains non-English characters",
			args: args{
				in: &model.CreateOrder{
					ID:   "A0000001",
					Name: "Melody我 Holiday Inn",
					Address: model.Address{
						City:     "taipei-city",
						District: "da-an-district",
						Street:   "fuxing-south-road",
					},
					Price:    "2050",
					Currency: "TWD",
				},
			},
			want: model.ErrNameContainNonEnglishChar,
		},
		{
			name: "Name is not capitalized",
			args: args{
				in: &model.CreateOrder{
					ID:   "A0000001",
					Name: "Melody holiday Inn",
					Address: model.Address{
						City:     "taipei-city",
						District: "da-an-district",
						Street:   "fuxing-south-road",
					},
					Price:    "2050",
					Currency: "TWD",
				},
			},
			want: model.ErrNameIsNotCapitalized,
		},
		{
			name: "Price is over 2000",
			args: args{
				in: &model.CreateOrder{
					ID:   "A0000001",
					Name: "Melody Holiday Inn",
					Address: model.Address{
						City:     "taipei-city",
						District: "da-an-district",
						Street:   "fuxing-south-road",
					},
					Price:    "2050",
					Currency: "TWD",
				},
			},
			want: model.ErrPriceIsOver2000,
		},
		{
			name: "Currency format is wrong",
			args: args{
				in: &model.CreateOrder{
					ID:   "A0000001",
					Name: "Melody Holiday Inn",
					Address: model.Address{
						City:     "taipei-city",
						District: "da-an-district",
						Street:   "fuxing-south-road",
					},
					Price:    "1900",
					Currency: "JPY",
				},
			},
			want: model.ErrCurrencyFormatWrong,
		},
		{
			name: "exchange ok",
			args: args{
				in: &model.CreateOrder{
					ID:   "A0000001",
					Name: "Melody Holiday Inn",
					Address: model.Address{
						City:     "taipei-city",
						District: "da-an-district",
						Street:   "fuxing-south-road",
					},
					Price:    "50",
					Currency: "USD",
				},
			},
			newPrice: "1550",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.CreateOrder(tt.args.in); !errors.Is(got, tt.want) {
				t.Errorf("service.CreateOrder() = %v, want %v", got, tt.want)
				if tt.newPrice != "" && tt.newPrice != tt.args.in.Price {
					t.Errorf("exchange wrong, price = %v, want %v", tt.args.in.Price, tt.newPrice)
				}
			}
		})
	}
}
