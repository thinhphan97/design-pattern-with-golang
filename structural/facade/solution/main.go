package main

import (
	"errors"
	"fmt"
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
	accounts []*Account
}

func (as AccountStorage) Lookup(name string) (*Account, error) {
	for _, account := range as.accounts {
		if account.Name == name {
			return account, nil
		}
	}
	return nil, errors.New("Produst not found")
}

type FacadeService struct {
	inventory      Inventory
	accountStorage AccountStorage
}

func (s *FacadeService) BuyProduct(name, accountName string) error {
	Product, err := s.inventory.Lookup(name)

	if err != nil {
		return err
	}
	account, err := s.accountStorage.Lookup(accountName)

	if err != nil {
		return err
	}

	if account.GetBalance() < Product.Price {
		return errors.New("not enough balance in account")
	}
	account.Withdraw(Product.Price)

	return nil
}

func (s *FacadeService) Deposit(accountName string, money float32) error {
	account, err := s.accountStorage.Lookup(accountName)

	if err != nil {
		return err
	}
	account.Deposit(money)

	return nil
}

func (s *FacadeService) FetchBalance(accountName string) float32 {
	account, err := s.accountStorage.Lookup(accountName)

	if err != nil {
		return 0
	}

	return account.GetBalance()
}

func NewFacadeService() FacadeService {
	return FacadeService{
		inventory: Inventory{
			products: []Product{
				{Name: "Apple", Price: 2.5},
				{Name: "Orange", Price: 3.0},
			},
		},
		accountStorage: AccountStorage{
			accounts: []*Account{
				{Name: "VIP", Balance: 1000},
				{Name: "Economic", Balance: 300},
			},
		},
	}
}

func main() {
	service := NewFacadeService()

	productName := "Apple"
	accountName := "VIP"

	if err := service.BuyProduct(productName, accountName); err != nil {
		log.Fatal(err)
	}

	// Check my balance
	fmt.Println("Account Balance:", service.FetchBalance(accountName))

	// Case 2: Deposit 100 into VIP Account

	if err := service.Deposit(accountName, 100); err != nil {
		log.Fatal(err)
	}

	// Check my balance again
	fmt.Println("Account Balance:", service.FetchBalance(accountName))
}
