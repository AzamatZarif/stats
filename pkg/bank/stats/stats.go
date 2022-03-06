 package stats

 import "github.com/AzamatZarif/bank/pkg/bank/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	length := 0
	for _, payment := range payments {
		sum += payment.Amount
		length++
	}
	result := sum/types.Money(length)
	return result
}

// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	length := 0
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
			length++
		}
	}
	return sum
}