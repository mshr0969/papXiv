package repository

import (
	"context"
	"gateway/internal/db"
	"gateway/internal/domain"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePaper(t *testing.T)  {
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
					Id: "2d1423b3-15ff-498e-b978-241f2b87de9e",
					Subject: "physics",
					Title: "test title",
					Url: "https://arxiv.org/hogehoge",
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
					Id: "2d1423b3-15ff-498e-b978-241f2b87de9e-hogehoge",
					Subject: "physics",
					Url: "https://arxiv.org/hogehoge",
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
			db := db.ConnectDB(resource, pool)

			pr := NewPaperRepository(db)

			ctx := context.Background()
			err := pr.CreatePaper(ctx, *tt.args.createPaper)

			if diff := cmp.Diff(tt.want.err, err != nil); diff != ""{
				t.Errorf("err mismatch (-want +got):\n%s", diff)
			}

			if !tt.want.err {
				var paper domain.Paper
				err = db.Get(&paper, "SELECT id, published, title, url FROM papers WHERE id=?", tt.args.createPaper.Id)
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

				if diff := cmp.Diff(tt.args.createPaper, &paper); diff != ""  {
					t.Errorf("paper mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
