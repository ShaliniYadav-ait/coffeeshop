package coffee

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type CoffeeDetails struct {
	Name  string  `mapstructure:"name"`
	Price float32 `mapstructure:"price"`
}

type CoffeeList struct {
	List []CoffeeDetails `mapstructure:"coffeelist"`
}

var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	//get current filepath along with the filename in the path
	_, b, _, _ := runtime.Caller(0)
	// trim the filename and current directory to find the root.
	basepath := filepath.Dir(filepath.Dir(b))

	viper.AddConfigPath(basepath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: %w", err)
		return nil, err
	}

	err = viper.Unmarshal(&Coffees)
	if err != nil {
		return nil, err
	}

	return &Coffees, nil
}

func IsCoffeeAvailable(coffeetype string) bool {
	for _, coffee := range Coffees.List {
		if coffee.Name == coffeetype {
			result := fmt.Sprintf("%s for $%v", coffee.Name, coffee.Price)
			fmt.Println(result)
			return true
		}

	}
	return false
}
