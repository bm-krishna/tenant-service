package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bm-krishna/tenant-service/pkg/models"
)

// SQLConnetion is a struct
type SQLConnetion struct {
	Conn *sql.DB
}

//QueryParser will take sql query and will return tenant model
func (conn *SQLConnetion) QueryParser(ctx context.Context, query string, args ...interface{}) ([]*models.Tenant, error) {
	rows, err := conn.Conn.QueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.New("Faild to parse Query " + err.Error())
	}
	defer rows.Close()
	result := make([]*models.Tenant, 0)
	for rows.Next() {
		data := new(models.Tenant)

		err := rows.Scan(
			&data.Name,
			&data.LoanAmount,
			&data.SSN,
			&data.LoanCount,
		)
		if err != nil {
			return nil, errors.New("Failed to map the result to tenant" + err.Error())
		}
		result = append(result, data)
	}
	return result, nil
}
