package vproductbus

import (
	"time"

	"github.com/mydomain/see-other/business/types/money"
	"github.com/mydomain/see-other/business/types/name"
	"github.com/mydomain/see-other/business/types/quantity"
	"github.com/google/uuid"
)

// Product represents an individual product with extended information.
type Product struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        name.Name
	Cost        money.Money
	Quantity    quantity.Quantity
	DateCreated time.Time
	DateUpdated time.Time
	UserName    name.Name
}