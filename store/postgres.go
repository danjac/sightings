package store

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/danjac/sightings/models"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
	sq sq.StatementBuilderType
}

const pageSize = 30

func newDB(db *sqlx.DB) *DB {
	builder := sq.StatementBuilder.
		RunWith(sq.NewStmtCacher(db.DB)).
		PlaceholderFormat(sq.Dollar)
	return &DB{db, builder}
}

type DBReader struct {
	*DB
}

func (r *DBReader) Get(id int64) (*models.Sighting, error) {

	sql, args, err := r.sq.
		Select("*").
		From("sightings").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	s := &models.Sighting{}

	if err := r.DB.Get(s, sql, args...); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *DBReader) Find(pageNumber int64) (*models.Page, error) {
	countQuery := r.sq.
		Select("COUNT(id)").
		From("sightings")

	selectQuery := r.sq.
		Select("*").
		From("sightings").
		OrderBy("occurred_at DESC")

	return r.paginate(countQuery, selectQuery, pageNumber)
}

func (r *DBReader) Search(search string, pageNumber int64) (*models.Page, error) {

	where := sq.Expr("tsv @@ plainto_tsquery(?)", search)

	countQuery := r.sq.
		Select("COUNT(id)").
		From("sightings").
		Where(where)

	selectQuery := r.sq.
		Select("*").
		From("sightings").
		Where(where).
		OrderBy("occurred_at DESC")

	return r.paginate(countQuery, selectQuery, pageNumber)
}

func (r *DBReader) paginate(countQuery sq.SelectBuilder,
	selectQuery sq.SelectBuilder,
	pageNumber int64) (*models.Page, error) {

	page := &models.Page{Number: pageNumber, PageSize: pageSize}

	if err := countQuery.
		QueryRow().
		Scan(&page.TotalRows); err != nil {
		return nil, err
	}

	page.TotalPages = page.TotalRows / pageSize

	offset := uint64((page.Number - 1) * pageSize)

	sql, args, err := selectQuery.
		Limit(pageSize).
		Offset(offset).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := r.Select(&page.Sightings, sql, args...); err != nil {
		return nil, err
	}

	return page, nil
}

type DBWriter struct {
	*DB
}

func (w *DBWriter) Insert(s *models.Sighting) error {
	q := w.sq.
		Insert("sightings").
		Columns(
			"occurred_at",
			"reported_at",
			"description",
			"shape",
			"duration",
			"location",
			"latitude",
			"longitude").
		Values(
			s.OccurredAt,
			s.ReportedAt,
			s.Description,
			s.Shape,
			s.Duration,
			s.Location,
			s.Latitude,
			s.Longitude).
		Suffix("RETURNING \"id\"")

	return q.QueryRow().Scan(&s.ID)

}
