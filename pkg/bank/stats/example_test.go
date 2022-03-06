 package stats

import (
	"github.com/AzamatZarif/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg_positive() {
	payments := []types.Payment {
		{Amount: 100,},
		{Amount: 200},
		{Amount: 300},
	}
	result := Avg(payments)
	fmt.Println(result)
	//Output:
	//200
}

func ExampleTotalInCategory_positive() {
	payments := []types.Payment {
		{Amount: 100, Category: "avto",},
		{Amount: 200, Category: "avto"},
		{Amount: 300, Category: "aptec"},
		{Amount: 500, Category: "aptec"},
		{Amount: 900, Category: "avto"},
	}
	result := TotalInCategory(payments, "avto")
	fmt.Println(result)
	//Output:
	//1200
}