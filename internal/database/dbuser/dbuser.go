package dbuser

import (
	"context"
	"encoding/json"
	"fmt"
)

type PhoneToCreate struct {
	Phone       string `json:"phone"`
	PhoneTypeID string `json:"phone_type_id"`
}
type AccountToCreate struct {
	Username    string          `json:"username"`
	LastName    string          `json:"last_name"`
	FirstName   string          `json:"first_name"`
	DateOfBirth string          `json:"dob"`
	Phones      []PhoneToCreate `json:"phones"`
}

type PhoneCreated struct {
	PhoneID     int    `json:"phone_id"`
	PhoneNumber string `json:"phone_number"`
	PhoneTypeID int    `json:"phone_type_id"`
	PhoneType   string `json:"phone_type"`
}
type AccountCreated struct {
	AccountID   int            `json:"account_id"`
	Username    string         `json:"username"`
	LastName    string         `json:"last_name"`
	FirstName   string         `json:"first_name"`
	DateOfBirth string         `json:"dob"`
	Phones      []PhoneCreated `json:"phones"`
}

func (d *Database) CreateAccount(ctx context.Context, acc *AccountToCreate) (AccountCreated, error) {
	query := `select * from norm.account_create($$%s$$::json);`

	data, err := json.Marshal(acc)
	if err != nil {
		return AccountCreated{}, fmt.Errorf("cannot marshal account to create: %w", err)
	}

	var rawRes *string
	err = d.db.GetContext(ctx, &rawRes, fmt.Sprintf(query, string(data)))
	if err != nil {
		return AccountCreated{}, fmt.Errorf("cannot insert user: %w", err)
	}

	if rawRes == nil {
		return AccountCreated{}, nil
	}

	res := AccountCreated{}
	err = json.Unmarshal([]byte(*rawRes), &res)
	if err != nil {
		return AccountCreated{}, fmt.Errorf("cannot unmarshal result: %w", err)
	}

	return res, nil
}
