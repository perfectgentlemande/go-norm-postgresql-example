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

type Phone struct {
	PhoneID     int    `json:"phone_id"`
	PhoneNumber string `json:"phone_number"`
	PhoneTypeID int    `json:"phone_type_id"`
	PhoneType   string `json:"phone_type"`
}
type EmailAddress struct {
	EmailID         int    `json:"email_id"`
	Email           string `json:"email"`
	EmailPriorityID int    `json:"email_priority_id"`
	EmailPriority   string `json:"email_priority"`
}

type Account struct {
	AccountID      int            `json:"account_id"`
	Username       string         `json:"username"`
	LastName       string         `json:"last_name"`
	FirstName      string         `json:"first_name"`
	DateOfBirth    string         `json:"dob"`
	Phones         []Phone        `json:"phones"`
	EmailAddresses []EmailAddress `json:"email_addresses"`
}

func (d *Database) CreateAccount(ctx context.Context, acc *AccountToCreate) (Account, error) {
	query := `select * from norm.account_create($$%s$$::json);`

	data, err := json.Marshal(acc)
	if err != nil {
		return Account{}, fmt.Errorf("cannot marshal account to create: %w", err)
	}

	var rawRes *string
	err = d.db.GetContext(ctx, &rawRes, fmt.Sprintf(query, string(data)))
	if err != nil {
		return Account{}, fmt.Errorf("cannot insert account: %w", err)
	}

	if rawRes == nil {
		return Account{}, nil
	}

	res := Account{}
	err = json.Unmarshal([]byte(*rawRes), &res)
	if err != nil {
		return Account{}, fmt.Errorf("cannot unmarshal result: %w", err)
	}

	return res, nil
}

func (d *Database) GetAccountByID(ctx context.Context, id int) (Account, error) {
	query := `select * from norm.account_select_by_id(%d)`

	var rawRes *string
	err := d.db.GetContext(ctx, &rawRes, fmt.Sprintf(query, id))
	if err != nil {
		return Account{}, fmt.Errorf("cannot get account by id: %w", err)
	}

	if rawRes == nil {
		return Account{}, nil
	}

	res := Account{}
	err = json.Unmarshal([]byte(*rawRes), &res)
	if err != nil {
		return Account{}, fmt.Errorf("cannot unmarshal result: %w", err)
	}

	return res, nil
}
