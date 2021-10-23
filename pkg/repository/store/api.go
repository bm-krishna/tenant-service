package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bm-krishna/tenant-service/pkg/models"
	"github.com/bm-krishna/tenant-service/pkg/repository/crud"
)

type sqlRepo struct {
	Conn *sql.DB
}

// NewSQLPostRepo retunrs implement of post repository interface
func NewSQLPostRepo(Conn *sql.DB) crud.TenantRepo {
	return &sqlRepo{
		Conn: Conn,
	}
}

// implement registry here

func (m *sqlRepo) Fetch(ctx context.Context, tenantCount int64) ([]*models.Tenant, error) {
	query := "Select id, title, content From tenant limit ?"
	connection := &SQLConnetion{
		Conn: m.Conn,
	}
	return connection.QueryParser(ctx, query, tenantCount)
}

func (m *sqlRepo) GetByID(ctx context.Context, id int64) (*models.Tenant, error) {
	query := "Select id, title, content From tenant where id=?"
	connection := &SQLConnetion{
		Conn: m.Conn,
	}
	rows, err := connection.QueryParser(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Tenant{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, errors.New("resquesed item is not found")
	}

	return payload, nil
}

func (m *sqlRepo) Create(ctx context.Context, p *models.Tenant) (int64, error) {
	query := "Insert tenant SET name=?, ssn=?, loanAmount=?, loanCount=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Name, p.SSN, p.LoanAmount, p.LoanCount)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *sqlRepo) Update(ctx context.Context, p *models.Tenant) (*models.Tenant, error) {
	query := "Update tenant set title=?, content=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.LoanAmount,
		p.LoanCount,
		p.SSN,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *sqlRepo) Delete(ctx context.Context, tenantID int64) (bool, error) {
	query := "Delete From tenant Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, tenantID)
	if err != nil {
		return false, err
	}
	return true, nil
}
