package card

import "errors"

// Data type
type Data struct {
	number uint8
	seed   Seed
}

// StrData is a card represented by string
type StrData struct {
	number, seed string
}

// ID is the id of a card from 1 to 40
type ID uint8

// Creator interface to represent what's needed to create a card
type Creator interface {
	ToNumber() uint8
	ToSeed() Seed
}

// By func
func By(sCard ID) (ID, error) {
	var c Data
	var err error
	if c.number, err = sCard.toNumber(); err == nil {
		c.seed, err = sCard.toSeed()
	}
	return c.ID(), err
}

// Card func
func Card(index uint8) (id ID, err error) {
	if index < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if index > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		id = ID(index)
	}
	return
}

// ByName func
func ByName(number, seed string) (ID, error) {
	sCard := StrData{number, seed}
	var c Data
	var err error
	if c.number, err = sCard.toNumber(); err == nil {
		c.seed, err = sCard.toSeed()
	}
	return c.ID(), err
}

// ID func
func (card *Data) ID() ID {
	return ID(card.number + (uint8)(card.seed)*10)
}

// Number func
func (card *ID) Number() uint8 {
	return card.ToNumber()
}

// Seed func
func (card *ID) Seed() Seed {
	return card.ToSeed()
}

// Points func
func (card *ID) Points() uint8 {
	switch card.ToNumber() {
	case 1:
		return 11
	case 3:
		return 10
	case 8:
		return 2
	case 9:
		return 3
	case 10:
		return 4
	default:
		return 0
	}
}
