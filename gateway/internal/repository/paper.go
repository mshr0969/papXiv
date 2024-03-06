package repository

import (
	"context"
	"database/sql"
	"gateway/domain"

	"github.com/jmoiron/sqlx"
)

type PaperRepository struct {
	db *sqlx.DB
}

func NewPaperRepository(db *sqlx.DB) *PaperRepository {
	return &PaperRepository{
		db: db,
	}
}

func (pr *PaperRepository) CreatePaper(ctx context.Context, do domain.Paper) error {
	tx, err := pr.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var subjectId int64
	err = tx.GetContext(ctx, &subjectId, "SELECT id FROM subjects WHERE name = ?", do.Subject)
	if err == sql.ErrNoRows {
		result, err := tx.ExecContext(ctx, "INSERT INTO subjects (name) VALUES (?)", do.Subject)
		if err != nil {
			return err
		}

		subjectId, err = result.LastInsertId()
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO papers(id, title, published, url) VALUES (?, ?, ?, ?)",
		do.Id, do.Title, do.Published, do.Url)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO paper_subjects (paper_id, subject_id) VALUES (?, ?)",
		do.Id, subjectId)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PaperRepository) ListPapers(ctx context.Context) (domain.Papers, error) {
	var papers domain.Papers

	err := pr.db.SelectContext(ctx, &papers, "SELECT id, title FROM papers")
	if err != nil {
		return nil, err
	}

	return papers, nil
}
func (pr *PaperRepository) SelectPaper(ctx context.Context, paperID string) (*domain.Paper, error) {
	var paper domain.Paper
	err := pr.db.GetContext(ctx, &paper, "SELECT id, published, title, url, created_at, updated_at FROM papers WHERE id=?", paperID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNonExistentPaper
		}
		return nil, err
	}

	var subjectId int64
	err = pr.db.GetContext(ctx, &subjectId, "SELECT subject_id FROM paper_subjects WHERE paper_id=?", paperID)
	if err != nil {
		return nil, err
	}

	var subjectName string
	err = pr.db.GetContext(ctx, &subjectName, "SELECT name FROM subjects WHERE id=?", subjectId)
	if err != nil {
		return nil, err
	}

	paper.Subject = subjectName

	return &paper, nil
}
func (pr *PaperRepository) UpdatePaper(ctx context.Context, paperID string) error {
	return nil
}
func (pr *PaperRepository) DeletePaper(ctx context.Context, paperID string) error {
	return nil
}
