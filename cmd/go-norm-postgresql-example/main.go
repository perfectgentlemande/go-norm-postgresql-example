package main

import (
	"context"

	"github.com/perfectgentlemande/go-norm-postgresql-example/internal/database/dbuser"
	"github.com/perfectgentlemande/go-norm-postgresql-example/internal/logger"
	dbUserMigrations "github.com/perfectgentlemande/go-norm-postgresql-example/migrations/dbuser"
	"github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func strToPtr(v string) *string {
	return &v
}

func main() {
	ctx := context.Background()
	log := logger.DefaultLogger()

	dbUser, err := dbuser.NewDatabase(ctx, &dbuser.Config{
		ConnStr: "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable",
	}, dbUserMigrations.MigrationAssets)
	if err != nil {
		log.WithError(err).Fatal("cannot create db")
	}
	defer dbUser.Close(ctx)

	err = dbUser.Ping(ctx)
	if err != nil {
		log.WithError(err).Error("cannot ping db")
		return
	}

	newAcc, err := dbUser.CreateAccount(ctx, &dbuser.AccountToCreate{
		Username:    "johnsmithaccount",
		LastName:    "john",
		FirstName:   "smith",
		DateOfBirth: "1991-04-01",
		Phones: []dbuser.PhoneToCreate{
			{
				Phone:       "3123334556",
				PhoneTypeID: "1",
			},
		},
	})
	if err != nil {
		log.WithError(err).Error("cannot create account")
		return
	}
	log.WithFields(logrus.Fields{
		"account_id": newAcc.AccountID,
		"username":   newAcc.Username,
		"last_name":  newAcc.LastName,
		"first_name": newAcc.FirstName,
		"phones":     newAcc.Phones,
	}).Info("created account")

	acc, err := dbUser.GetAccountByID(ctx, 1)
	if err != nil {
		log.WithError(err).Error("cannot get account by id")
		return
	}
	log.WithFields(logrus.Fields{
		"account_id":      acc.AccountID,
		"username":        acc.Username,
		"last_name":       acc.LastName,
		"first_name":      acc.FirstName,
		"phones":          acc.Phones,
		"email_addresses": acc.EmailAddresses,
	}).Info("got account")

	accs, err := dbUser.ListAcccounts(ctx, &dbuser.ListAccountsParams{
		LastName: strToPtr("john"),
		Username: strToPtr("ali"),
	})
	if err != nil {
		log.WithError(err).Error("cannot list accounts")
		return
	}
	for i := range accs {
		log.WithFields(logrus.Fields{
			"account_id":      accs[i].AccountID,
			"username":        accs[i].Username,
			"last_name":       accs[i].LastName,
			"first_name":      accs[i].FirstName,
			"phones":          accs[i].Phones,
			"email_addresses": accs[i].EmailAddresses,
		}).Info("listed account")
	}
	log.Info("list account done")

	accs, err = dbUser.ListAcccounts(ctx, &dbuser.ListAccountsParams{
		LastName: strToPtr("john"),
	})
	if err != nil {
		log.WithError(err).Error("cannot list accounts")
		return
	}
	for i := range accs {
		log.WithFields(logrus.Fields{
			"account_id":      accs[i].AccountID,
			"username":        accs[i].Username,
			"last_name":       accs[i].LastName,
			"first_name":      accs[i].FirstName,
			"phones":          accs[i].Phones,
			"email_addresses": accs[i].EmailAddresses,
		}).Info("listed account")
	}
	log.Info("list account done")

	accs, err = dbUser.ListAcccounts(ctx, &dbuser.ListAccountsParams{
		LastName: strToPtr("john"),
		Username: strToPtr("john"),
	})
	if err != nil {
		log.WithError(err).Error("cannot list accounts")
		return
	}
	for i := range accs {
		log.WithFields(logrus.Fields{
			"account_id":      accs[i].AccountID,
			"username":        accs[i].Username,
			"last_name":       accs[i].LastName,
			"first_name":      accs[i].FirstName,
			"phones":          accs[i].Phones,
			"email_addresses": accs[i].EmailAddresses,
		}).Info("listed account")
	}
	log.Info("list account done")

	accs, err = dbUser.ListAcccounts(ctx, &dbuser.ListAccountsParams{
		FirstName: strToPtr("john"),
		Username:  strToPtr("john"),
	})
	if err != nil {
		log.WithError(err).Error("cannot list accounts")
		return
	}
	for i := range accs {
		log.WithFields(logrus.Fields{
			"account_id":      accs[i].AccountID,
			"username":        accs[i].Username,
			"last_name":       accs[i].LastName,
			"first_name":      accs[i].FirstName,
			"phones":          accs[i].Phones,
			"email_addresses": accs[i].EmailAddresses,
		}).Info("listed account")
	}
	log.Info("list account done")

	accs, err = dbUser.ListAcccounts(ctx, &dbuser.ListAccountsParams{
		Username: strToPtr("alic"),
		Phone:    strToPtr("202"),
	})
	if err != nil {
		log.WithError(err).Error("cannot list accounts")
		return
	}
	for i := range accs {
		log.WithFields(logrus.Fields{
			"account_id":      accs[i].AccountID,
			"username":        accs[i].Username,
			"last_name":       accs[i].LastName,
			"first_name":      accs[i].FirstName,
			"phones":          accs[i].Phones,
			"email_addresses": accs[i].EmailAddresses,
		}).Info("listed account")
	}
	log.Info("list account done")

	accs, err = dbUser.ListAcccounts(ctx, &dbuser.ListAccountsParams{
		Phone: strToPtr("202"),
	})
	if err != nil {
		log.WithError(err).Error("cannot list accounts")
		return
	}
	for i := range accs {
		log.WithFields(logrus.Fields{
			"account_id":      accs[i].AccountID,
			"username":        accs[i].Username,
			"last_name":       accs[i].LastName,
			"first_name":      accs[i].FirstName,
			"phones":          accs[i].Phones,
			"email_addresses": accs[i].EmailAddresses,
		}).Info("listed account")
	}
	log.Info("list account done")

	accs, err = dbUser.ListAcccounts(ctx, &dbuser.ListAccountsParams{
		LastName: strToPtr("john"),
		Email:    strToPtr("john"),
	})
	if err != nil {
		log.WithError(err).Error("cannot list accounts")
		return
	}
	for i := range accs {
		log.WithFields(logrus.Fields{
			"account_id":      accs[i].AccountID,
			"username":        accs[i].Username,
			"last_name":       accs[i].LastName,
			"first_name":      accs[i].FirstName,
			"phones":          accs[i].Phones,
			"email_addresses": accs[i].EmailAddresses,
		}).Info("listed account")
	}
	log.Info("list account done")
}
