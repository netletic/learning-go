package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

func Buy(b Book) (Book, error) {
	if b.Copies <= 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (b Book) NetPriceCents() int {
	discount := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - discount
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("negative price %d", price)
	}
	b.PriceCents = price
	return nil
}

func (b Book) Category() Category {
	return b.category
}

func (b *Book) SetCategory(category Category) error {
	if _, ok := validCategory[category]; !ok {
		return fmt.Errorf("unknown category %q", category)
	}
	b.category = category
	return nil
}

type Catalog map[int]Book

func (c Catalog) GetAllBooks() []Book {
	books := []Book{}
	for _, b := range c {
		books = append(books, b)
	}
	return books
}

func (c Catalog) GetBook(ID int) (Book, error) {
	b, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}
