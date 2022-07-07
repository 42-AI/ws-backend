package db

import (
	"context"
	"github.com/google/uuid"
	"time"
)

func Bootstrap(ctx context.Context, dbal Dbal) error {
	err := dbal.CreateIndexes(ctx)
	if err != nil {
		return err
	}
	userSimple := User{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Email:     "simple-user@email.com",
		Password:  "97c440035331438ac703d2c6e82290c890f8e2ffaaf0e21d2641aa27919e2866",
		Admin:     false,
	}
	err = dbal.CreateUser(ctx, userSimple)
	if err != nil {
		return err
	}
	userAdmin := User{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Email:     "admin-user@email.com",
		Password:  "5d4bd9bf825f3220ffea4a9e442ef4442586c1d7356161fbe85fc6ce1d2ae683",
		Admin:     true,
	}
	err = dbal.CreateUser(ctx, userAdmin)
	if err != nil {
		return err
	}
	return err
}
