package repository

import (
	"context"
	"fmt"

	"github.com/BlackMocca/sqlx"
	"github.com/Blackmocca/wedding-ui/models"
	pg "github.com/lib/pq"
)

type PsqlClient struct {
	client *sqlx.DB
}

func getPsqlClient(uri string) *sqlx.DB {
	addr, err := pg.ParseURL(uri)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		panic(err)
	}

	return db
}

func NewPsqlClient(uri string) *PsqlClient {
	return &PsqlClient{client: getPsqlClient(uri)}
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
