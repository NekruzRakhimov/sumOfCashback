package main

import "fmt"

type Bill struct {
	purchases []int64
	percentCashback int64
	SumOfCashback int64
}

func (client *Bill) CountSumOfCashback() {
	sumOfPurchases := int64(0)
	for _, purchase := range client.purchases {
		sumOfPurchases += purchase
	}
	client.SumOfCashback = sumOfPurchases * client.percentCashback / 100
}

func main(){
	var client Bill
	client.purchases = []int64{500, 1000, 1500}
	client.percentCashback = 15
	client.CountSumOfCashback()
	fmt.Println(client.SumOfCashback)
}