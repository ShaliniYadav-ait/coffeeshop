package coffee

import (
	"reflect"
	"testing"
)

func TestGetCoffees(t *testing.T) {

	tests := []struct {
		name    string
		want    *CoffeeList
		wantErr bool
	}{
		{
			name: "get coffees",
			want: &CoffeeList{
				List: []CoffeeDetails{
					{
						Name:  "Latte",
						Price: 3,
					},
					{
						Name:  "Cappuccino",
						Price: 2.75,
					},
					{
						Name:  "Flat White",
						Price: 2.25,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCoffees()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoffees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoffees() = %v, want %v", got, tt.want)
			}
		})
	}
}
