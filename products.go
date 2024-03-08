package main

import (
  "time"
  "slices"
  "math/rand"
)
type Product struct {
  price     int
  name      string
  rel_date  time.Time
  exp_date  time.Time
}

func newProduct(price int, name string, rel_date, exp_date time.Time) Product {
  return Product{price, name, rel_date, exp_date}
}

func (prod Product) isExpired() bool {
  return (time.Now().UnixMilli() - prod.exp_date.UnixMilli()) < 0
}

type Cart struct {
  products []Product
}

func (cart *Cart) totalPrice() int {
  sum := 0;
  for _, val := range cart.products {
    sum += val.price;
  }
  return sum;
}

func (cart *Cart) filterExpiredProducts() []Product {
  return slices.DeleteFunc(cart.products, func(product Product) bool { return product.isExpired() })
}

func (cart *Cart) closestExpData() *Product {
  if len(cart.products) == 0 { return nil }
  prod := cart.products[0]
  for _, val := range cart.products {
    if val.isExpired() { continue }
    if prod.exp_date.Compare(val.exp_date) > 0 {
      prod = val
    }
  }
  return &prod
}

func randNum(max int) int {
  return rand.Intn(max) * ((rand.Intn(2) * 2) - 1)
}

func randTime() time.Time {
  return time.Now().AddDate(randNum(2), randNum(13), randNum(29))
}
