package payment

import "time"

type Transaction struct {
	ID        int
	Amount    int
	CreatedAt time.Time
}
