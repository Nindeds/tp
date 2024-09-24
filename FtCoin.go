package poulpe

func Ft_coin(coins []int, amount int) int {
	x := 0
	for i := len(coins) - 1; 0 <= i; i-- {
		if amount > coins[i] && coins[i]+coins[i] > amount {
			return -1
		}
		for amount >= coins[i] {
			amount -= coins[i]
			x++
		}
	}
	return x
}
