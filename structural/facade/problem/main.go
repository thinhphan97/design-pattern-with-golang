package main

import (
	"errors"
	"log"
)

type Product struct {
	Name  string
	Price float32
}

type Inventory struct {
	products []Product
}

func (iv Inventory) Lookup(name string) (*Product, error) {
	for _, product := range iv.products {
		if product.Name == name {
			return &product, nil
		}
	}
	return nil, errors.New("Produst not found")
}

type Account struct {
	Name    string
	Balance float32
}

func (acc *Account) Deposit(money float32)  { acc.Balance += money }
func (acc *Account) Withdraw(money float32) { acc.Balance -= money }
func (acc *Account) GetBalance() float32    { return acc.Balance }

type AccountStorage struct {
	accounts []Account
}

func (as AccountStorage) Lookup(name string) (*Account, error) {
	for _, account := range as.accounts {
		if account.Name == name {
			return &account, nil
		}
	}
	return nil, errors.New("Produst not found")
}

var inventory = Inventory{
	products: []Product{
		{Name: "Apple", Price: 2.5},
		{Name: "Orange", Price: 3.0},
	},
}

var accountStorage = AccountStorage{
	accounts: []Account{
		{Name: "VIP", Balance: 1000},
		{Name: "Economic", Balance: 300},
	},
}

func main() {
	// Case 1: Buy a product with an account
	productName := "Apple"
	accountName := "VIP"

	product, err := inventory.Lookup(productName)

	if err != nil {
		log.Fatal(err)
	}

	account, err := accountStorage.Lookup(accountName)

	if err != nil {
		log.Fatal(err)
	}

	if account.GetBalance() < product.Price {
		log.Fatal("not enough balance")
	}

	account.Withdraw(product.Price)
	// And more step to finish buying process....

	// Problem: I have to do it myself, too many steps to take. I'm
	// not sure if I do correctly
}
