package repo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/danjac/sightings/models"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
	sq sq.StatementBuilderType
}

const (
	pageSize       = 30
	sightingsTable = "sightings"
	defaultOrderBy = "occurred_at DESC"
)

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

	sql, args, err := r.doSelect().Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	s := &models.Sighting{}

	if err := r.DB.Get(s, sql, args...); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *DBReader) Find(pageNumber int64) (*models.SightingsPage, error) {
	return r.paginate(r.doCount(), r.doSelect(), pageNumber)
}

func (r *DBReader) Search(search string, pageNumber int64) (*models.SightingsPage, error) {

	where := sq.Expr("tsv @@ plainto_tsquery(?)", search)

	countQuery := r.doCount().Where(where)
	selectQuery := r.doSelectAll().Where(where)

	return r.paginate(countQuery, selectQuery, pageNumber)
}

func (r *DBReader) doCount() sq.SelectBuilder {
	return r.sq.Select("COUNT(id)").From(sightingsTable)
}

func (r *DBReader) doSelect() sq.SelectBuilder {
	return r.sq.Select("*").From(sightingsTable)
}

func (r *DBReader) doSelectAll() sq.SelectBuilder {
	return r.doSelect().OrderBy(defaultOrderBy)
}

func (r *DBReader) paginate(countQuery sq.SelectBuilder,
	selectQuery sq.SelectBuilder,
	pageNumber int64) (*models.SightingsPage, error) {

	page := &models.SightingsPage{
		Page: &models.Page{
			Number:   pageNumber,
			PageSize: pageSize,
		},
	}

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
		Insert(sightingsTable).
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
