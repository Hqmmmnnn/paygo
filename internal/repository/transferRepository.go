package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Hqqm/paygo/internal/entities"
	"github.com/Hqqm/paygo/internal/interfaces"
	"github.com/jmoiron/sqlx"
)

type transferRepository struct {
	db *sqlx.DB
}

func NewTransferRepository(db *sqlx.DB) interfaces.TransferRepository {
	return &transferRepository{db: db}
}

func (transferRepository *transferRepository) GetDbConnection() *sqlx.DB {
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

func (transferRepository *transferRepository) GetTransfers(ctx context.Context, login string) (*[]entities.Transfer, error) {
	transfers := []entities.Transfer{}
	query := "SELECT * FROM (SELECT * FROM transfers WHERE sender_login=$1 OR recipient_login=$1) allUserTransfers order by date DESC"

	err := transferRepository.db.SelectContext(ctx, &transfers, query, login)

	for i := 0; i < len(transfers); i++ {
		d := strings.Split(transfers[i].Date, "T")
		t := strings.Split(d[1], ".")
		transfers[i].Date = fmt.Sprintf("%s %s", d[0], t[0])
	}

	if err != nil {
		return nil, err
	}

	return &transfers, nil
}
