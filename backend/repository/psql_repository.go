package repository

import (
	"context"

	"github.com/BlackMocca/sqlx"
	"github.com/Blackmocca/wedding-ui/models"
)

type PsqlClient struct {
	client *sqlx.DB
}

func NewPsqlClient(db *sqlx.DB) *PsqlClient {
	return &PsqlClient{client: db}
}

func (p PsqlClient) Create(ctx context.Context, ptr *models.Celebrate) error {
	sql := `INSERT INTO celebrate (celebrate_text, celebrate_from, created_at, updated_at)
	VALUES (?,?,?,?)
	`
	sql = sqlx.Rebind(sqlx.DOLLAR, sql)
	stmt, err := p.client.PreparexContext(ctx, sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.ExecContext(ctx, ptr.CelebrateText, ptr.CelebrateFrom, ptr.CreatedAt, ptr.UpdatedAt); err != nil {
		return err
	}
	return nil
}
