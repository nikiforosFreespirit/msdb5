package card

import "errors"

// Card type
type Card struct {
	number uint8
	seed   Seed
}

// ByID func
func ByID(id uint8) (*Card, error) {
	if id < 1 {
		return nil, errors.New("Index cannot be less than 1")
	} else if id > 40 {
		return nil, errors.New("Index cannot be more than 40")
	} else {
		seedIndex := (id - 1) / 10
		number := uint8(id - (10 * seedIndex))
		seed := Seed(seedIndex)
		return &Card{number: number, seed: seed}, nil
	}
}

func (card *Card) points() uint8 {
	switch card.number {
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
