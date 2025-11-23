// repo/dedication_repo.go
package repo

import (
	"context"
	"database/sql"
	"dedicationWall/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DedicationRepo interface {
	Create(ctx context.Context, d domain.Dedication) (*domain.Dedication, error)
	List(ctx context.Context, page, limit int64) ([]*domain.Dedication, int64, error)
	GetByID(ctx context.Context, id uint64) (*domain.Dedication, error)
	Delete(ctx context.Context, id uint64) error
}

type dedicationRepo struct {
	db *sqlx.DB
}

// Constructor — same as your product
func NewDedicationRepo(db *sqlx.DB) DedicationRepo {
	return &dedicationRepo{
		db: db,
	}
}

// Create — same style as your product Create()
func (r *dedicationRepo) Create(ctx context.Context, d domain.Dedication) (*domain.Dedication, error) {
	query := `
	INSERT INTO dedications (
		name,
		message,
		song_url
	) VALUES (
		$1, $2, $3
	)
	RETURNING id, created_at
	`

	var createdAt string
	err := r.db.QueryRowContext(ctx, query, d.Name, d.Message, d.SongURL).
		Scan(&d.ID, &createdAt)
	if err != nil {
		return nil, err
	}

	d.CreatedAt = createdAt
	return &d, nil
}

// List — same as your product List() + Count()
func (r *dedicationRepo) List(ctx context.Context, page, limit int64) ([]*domain.Dedication, int64, error) {
	offset := (page - 1) * limit

	var dedications []*domain.Dedication
	query := `
	SELECT 
		id, name, message, song_url, created_at
	FROM dedications
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2
	`

	err := r.db.SelectContext(ctx, &dedications, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Count total
	var total int64
	countQuery := `SELECT COUNT(*) FROM dedications`
	err = r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return dedications, total, nil
}

// GetByID — same as your Get()
func (r *dedicationRepo) GetByID(ctx context.Context, id uint64) (*domain.Dedication, error) {
	var d domain.Dedication

	query := `
	SELECT id, name, message, song_url, created_at
	FROM dedications
	WHERE id = $1
	`

	err := r.db.GetContext(ctx, &d, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		return nil, err
	}

	return &d, nil
}

// Delete — same as your Delete()
func (r *dedicationRepo) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM dedications WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("dedication with id %d not found", id)
	}

	return nil
}
