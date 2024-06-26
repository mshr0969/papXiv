package repository

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/test/db"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePaper(t *testing.T) {
	t.Parallel()

	type args struct {
		createPaper *domain.Paper
	}

	type want struct {
		err bool
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create case",
			args: args{
				createPaper: &domain.Paper{
					Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
					Subject:   "physics",
					Title:     "test title",
					Url:       "https://arxiv.org/hogehoge",
					Published: "2024/01/01",
				},
			},
			want: want{
				err: false,
			},
		},
		{
			name: "invalid id case",
			args: args{
				createPaper: &domain.Paper{
					Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e-hogehoge",
					Subject:   "physics",
					Url:       "https://arxiv.org/hogehoge",
					Published: "2024/01/01",
				},
			},
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resource, pool := db.CreateContainer()
			defer db.CloseContainer(resource, pool)
			d := db.ConnectDB(resource, pool)

			if err := db.MigrateDB(d); err != nil {
				t.Error("could not migrate")
			}

			pr := NewPaperRepository(d)

			ctx := context.Background()
			err := pr.CreatePaper(ctx, *tt.args.createPaper)

			if diff := cmp.Diff(tt.want.err, err != nil); diff != "" {
				t.Errorf("err mismatch (-want +got):\n%s", diff)
			}

			if !tt.want.err {
				var paper domain.Paper
				err = d.Get(&paper, "SELECT id, published, title, url FROM papers WHERE id=?", tt.args.createPaper.Id)
				if err != nil {
					t.Errorf("could not select paper: %s", err)
				}

				var subjectId int64
				err = pr.db.GetContext(ctx, &subjectId, "SELECT subject_id FROM paper_subjects WHERE paper_id=?", tt.args.createPaper.Id)
				if err != nil {
					t.Errorf("could not select subject_id: %s", err)
				}

				var subjectName string
				err = pr.db.GetContext(ctx, &subjectName, "SELECT name FROM subjects WHERE id=?", subjectId)
				if err != nil {
					t.Errorf("could not select subject name: %s", err)
				}

				paper.Subject = subjectName

				if diff := cmp.Diff(tt.args.createPaper, &paper); diff != "" {
					t.Errorf("paper mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestListPapers(t *testing.T) {
	t.Parallel()

	type want struct {
		err bool
		res domain.Papers
	}

	type args struct {
		existPapers domain.Papers
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful list case",
			args: args{
				existPapers: domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
			},
			want: want{
				err: false,
				res: domain.Papers{
					{
						Id:    "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Title: "test title",
					},
					{
						Id:    "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Title: "test title2",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resource, pool := db.CreateContainer()
			defer db.CloseContainer(resource, pool)
			d := db.ConnectDB(resource, pool)

			if err := db.MigrateDB(d); err != nil {
				t.Error("could not migrate")
			}

			pr := NewPaperRepository(d)

			// データの準備
			for _, paper := range tt.args.existPapers {
				_, err := d.Exec("INSERT INTO papers(id, title, published, url) VALUES (?, ?, ?, ?)",
					paper.Id, paper.Title, paper.Published, paper.Url)
				if err != nil {
					t.Errorf("could not insert paper: %s", err)
				}

				// ListPapersにおいては不要な処理
				_, err = d.Exec("INSERT INTO subjects(name) VALUES (?)", paper.Subject)
				if err != nil {
					t.Errorf("could not insert subject: %s", err)
				}

				var subjectId int64
				err = d.Get(&subjectId, "SELECT id FROM subjects WHERE name=?", paper.Subject)
				if err != nil {
					t.Errorf("could not select subject_id: %s", err)
				}

				_, err = d.Exec("INSERT INTO paper_subjects(paper_id, subject_id) VALUES (?, ?)", paper.Id, subjectId)
				if err != nil {
					t.Errorf("could not insert paper_subjects: %s", err)
				}
			}

			ctx := context.Background()
			papers, err := pr.ListPapers(ctx)

			if diff := cmp.Diff(tt.want.err, err != nil); diff != "" {
				t.Errorf("err mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.want.res, papers); diff != "" {
				t.Errorf("papers mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSelectPaper(t *testing.T) {
	t.Parallel()

	type want struct {
		err bool
		res *domain.Paper
	}

	type args struct {
		paperID     string
		existPapers *domain.Papers
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful select case",
			args: args{
				paperID: "2d1423b3-15ff-498e-b978-241f2b87de9e",
				existPapers: &domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
			},
			want: want{
				err: false,
				res: &domain.Paper{
					Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
					Title:     "test title",
					Subject:   "physics",
					Url:       "https://arxiv.org/hogehoge",
					Published: "2024/01/01",
				},
			},
		},
		{
			name: "not found case",
			args: args{
				paperID: "FA20022B-76D2-43F0-865F-BA45D6370D22",
				existPapers: &domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
			},
			want: want{
				err: true,
				res: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resource, pool := db.CreateContainer()
			defer db.CloseContainer(resource, pool)
			d := db.ConnectDB(resource, pool)

			if err := db.MigrateDB(d); err != nil {
				t.Error("could not migrate")
			}

			pr := NewPaperRepository(d)

			// データの準備
			for _, paper := range *tt.args.existPapers {
				_, err := d.Exec("INSERT INTO papers(id, title, published, url) VALUES (?, ?, ?, ?)",
					paper.Id, paper.Title, paper.Published, paper.Url)
				if err != nil {
					t.Errorf("could not insert paper: %s", err)
				}

				_, err = d.Exec("INSERT INTO subjects(name) VALUES (?)", paper.Subject)
				if err != nil {
					t.Errorf("could not insert subject: %s", err)
				}

				var subjectId int64
				err = d.Get(&subjectId, "SELECT id FROM subjects WHERE name=?", paper.Subject)
				if err != nil {
					t.Errorf("could not select subject_id: %s", err)
				}

				_, err = d.Exec("INSERT INTO paper_subjects(paper_id, subject_id) VALUES (?, ?)", paper.Id, subjectId)
				if err != nil {
					t.Errorf("could not insert paper_subjects: %s", err)
				}
			}

			ctx := context.Background()
			paper, err := pr.SelectPaper(ctx, tt.args.paperID)

			if diff := cmp.Diff(tt.want.err, err != nil); diff != "" {
				t.Errorf("err mismatch (-want +got):\n%s", diff)
			}

			if tt.want.res != nil {
				if paper.CreatedAt == "" {
					t.Errorf("CreatedAt is zero")
				}

				if paper.UpdatedAt == "" {
					t.Errorf("UpdatedAt is zero")
				}
				paper.CreatedAt = ""
				paper.UpdatedAt = ""

				if diff := cmp.Diff(tt.want.res, paper); diff != "" {
					t.Errorf("paper mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestUpdatePaper(t *testing.T) {
	t.Parallel()

	type want struct {
		err bool
		paper *domain.Paper
	}

	type args struct {
		existPapers *domain.Papers
		updatePaper *domain.Paper
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful update case",
			args: args{
				existPapers: &domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
				updatePaper: &domain.Paper{
					Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
					Subject:   "test",
					Title:     "test",
					Url:       "https://arxiv.org/hogehoge3",
					Published: "2024/01/03",
				},
			},
			want: want{
				err: false,
				paper: &domain.Paper{
					Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
					Subject:   "test",
					Title:     "test",
					Url:       "https://arxiv.org/hogehoge3",
					Published: "2024/01/03",
				},
			},
		},
		{
			name: "not found case",
			args: args{
				existPapers: &domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
				updatePaper: &domain.Paper{
					Id:        "FA20022B-76D2-43F0-865F-BA45D6370D22",
					Subject:   "test",
					Title:     "test",
					Url:       "https://arxiv.org/hogehoge3",
					Published: "2024/01/03",
				},
			},
			want: want{
				err: true,
			},
		},
		{
			name: "invalid id case",
			args: args{
				existPapers: &domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
				updatePaper: &domain.Paper{
					Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e-hogehoge",
					Subject:   "physics",
					Title:     "test title",
					Url:       "https://arxiv.org/hogehoge",
					Published: "2024/01/01",
				},
			},
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resource, pool := db.CreateContainer()
			defer db.CloseContainer(resource, pool)
			d := db.ConnectDB(resource, pool)

			if err := db.MigrateDB(d); err != nil {
				t.Error("could not migrate")
			}

			pr := NewPaperRepository(d)

			// データの準備
			for _, paper := range *tt.args.existPapers {
				_, err := d.Exec("INSERT INTO papers(id, title, published, url) VALUES (?, ?, ?, ?)",
					paper.Id, paper.Title, paper.Published, paper.Url)
				if err != nil {
					t.Errorf("could not insert paper: %s", err)
				}

				_, err = d.Exec("INSERT INTO subjects(name) VALUES (?)", paper.Subject)
				if err != nil {
					t.Errorf("could not insert subject: %s", err)
				}

				var subjectId int64
				err = d.Get(&subjectId, "SELECT id FROM subjects WHERE name=?", paper.Subject)
				if err != nil {
					t.Errorf("could not select subject_id: %s", err)
				}

				_, err = d.Exec("INSERT INTO paper_subjects(paper_id, subject_id) VALUES (?, ?)", paper.Id, subjectId)
				if err != nil {
					t.Errorf("could not insert paper_subjects: %s", err)
				}
			}

			ctx := context.Background()
			err := pr.UpdatePaper(ctx, *tt.args.updatePaper)

			if diff := cmp.Diff(tt.want.err, err != nil); diff != "" {
				t.Errorf("err mismatch (-want +got):\n%s", diff)
			}

			if !tt.want.err {
				var paper domain.Paper
				err = d.Get(&paper, "SELECT id, published, title, url FROM papers WHERE id=?", tt.want.paper.Id)
				if err != nil {
					t.Errorf("could not select paper: %s", err)
				}

				var subjectId int64
				err = pr.db.GetContext(ctx, &subjectId, "SELECT subject_id FROM paper_subjects WHERE paper_id=?", tt.want.paper.Id)
				if err != nil {
					t.Errorf("could not select subject_id: %s", err)
				}

				var subjectName string
				err = pr.db.GetContext(ctx, &subjectName, "SELECT name FROM subjects WHERE id=?", subjectId)
				if err != nil {
					t.Errorf("could not select subject name: %s", err)
				}

				paper.Subject = subjectName

				if diff := cmp.Diff(tt.want.paper, &paper); diff != "" {
					t.Errorf("paper mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestDeletePaper(t *testing.T) {
	t.Parallel()

	type want struct {
		err bool
	}

	type args struct {
		existPapers *domain.Papers
		paperId    string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful delete case",
			args: args{
				existPapers: &domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
				paperId: "2d1423b3-15ff-498e-b978-241f2b87de9e",
			},
			want: want{
				err: false,
			},
		},
		{
			name: "not found case",
			args: args{
				existPapers: &domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
				paperId: "FA20022B-76D2-43F0-865F-BA45D6370D22",
			},
			want: want{
				err: true,
			},
		},
		{
			name: "invalid id case",
			args: args{
				existPapers: &domain.Papers{
					{
						Id:        "2d1423b3-15ff-498e-b978-241f2b87de9e",
						Subject:   "physics",
						Title:     "test title",
						Url:       "https://arxiv.org/hogehoge",
						Published: "2024/01/01",
					},
					{
						Id:        "74459C44-3815-4DAE-BC5E-DC1F92407A1B",
						Subject:   "mathematics",
						Title:     "test title2",
						Url:       "https://arxiv.org/hogehoge2",
						Published: "2024/01/02",
					},
				},
				paperId: "2d1423b3-15ff-498e-b978-241f2b87de9e-hogehoge",
			},
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resource, pool := db.CreateContainer()
			defer db.CloseContainer(resource, pool)
			d := db.ConnectDB(resource, pool)

			if err := db.MigrateDB(d); err != nil {
				t.Error("could not migrate")
			}

			pr := NewPaperRepository(d)

			// データの準備
			for _, paper := range *tt.args.existPapers {
				_, err := d.Exec("INSERT INTO papers(id, title, published, url) VALUES (?, ?, ?, ?)",
					paper.Id, paper.Title, paper.Published, paper.Url)
				if err != nil {
					t.Errorf("could not insert paper: %s", err)
				}

				_, err = d.Exec("INSERT INTO subjects(name) VALUES (?)", paper.Subject)
				if err != nil {
					t.Errorf("could not insert subject: %s", err)
				}

				var subjectId int64
				err = d.Get(&subjectId, "SELECT id FROM subjects WHERE name=?", paper.Subject)
				if err != nil {
					t.Errorf("could not select subject_id: %s", err)
				}

				_, err = d.Exec("INSERT INTO paper_subjects(paper_id, subject_id) VALUES (?, ?)", paper.Id, subjectId)
				if err != nil {
					t.Errorf("could not insert paper_subjects: %s", err)
				}
			}

			ctx := context.Background()
			err := pr.DeletePaper(ctx, tt.args.paperId)

			if diff := cmp.Diff(tt.want.err, err != nil); diff != "" {
				t.Errorf("err mismatch (-want +got):\n%s", diff)
			}

			if !tt.want.err {
				var paper domain.Paper
				err = d.Get(&paper, "SELECT id, published, title, url FROM papers WHERE id=?", tt.args.paperId)
				if err == nil {
					t.Errorf("could select paper: %s", err)
				}
			}
		})
	}

}
