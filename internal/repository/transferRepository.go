package repository

import (
	"context"

	"github.com/Hqqm/paygo/internal/interfaces"
	"github.com/jmoiron/sqlx"
)

type transferRepository struct {
	db *sqlx.DB
}

func NewTransferRepository(db *sqlx.DB) interfaces.TransferRepository {
	return &transferRepository{db: db}
}

func (transferRepository *transferRepository) getDbConnection() *sqlx.DB {
	return transferRepository.db
}

func (transferRepository *transferRepository) InsertMoneyTransferData(ctx context.Context, tx *sqlx.Tx, moneyTransferID, senderLogin, recipientLogin, comment string, amount float64) error {
	insertIntoTransfersQuery := `	
		INSERT INTO transfers(id, sender_login, recipient_login, comment, amount)
		VALUES (:id, :sender_login, :recipient_login, :comment, :amount)
	`

	_, err := tx.NamedExecContext(ctx, insertIntoTransfersQuery, map[string]interface{}{
		"id":              moneyTransferID,
		"sender_login":    senderLogin,
		"recipient_login": recipientLogin,
		"comment":         comment,
		"amount":          amount,
	})

	if err != nil {
		return err
	}

	return nil
}

func (transferRepository *transferRepository) Transaction(txFunc func(*sqlx.Tx) error) (err error) {
	db := transferRepository.getDbConnection()
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)

	return err
}
