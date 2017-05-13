package store

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/danjac/sightings"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
	sq sq.StatementBuilderType
}

const pageSize = 30

func newDB(db *sqlx.DB) *DB {
	builder := sq.StatementBuilder.
		RunWith(db.DB).
		PlaceholderFormat(sq.Dollar)
	return &DB{db, builder}
}

func (db *DB) GetOne(id string) (*sightings.Sighting, error) {

	sql, args, err := db.sq.
		Select("*").
		From("sightings").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	s := &sightings.Sighting{}

	if err := db.Get(s, sql, args...); err != nil {
		return nil, err
	}

	return s, nil
}

func (db *DB) GetAll(pageNumber int64) (*sightings.Page, error) {
	countQuery := db.sq.
		Select("COUNT(id)").
		From("sightings")

	selectQuery := db.sq.
		Select("*").
		From("sightings").
		OrderBy("occurred_at DESC")

	return db.paginate(countQuery, selectQuery, pageNumber)
}

func (db *DB) Search(search string, pageNumber int64) (*sightings.Page, error) {

	q := "%" + search + "%"
	cols := []string{"location", "shape", "description"}

	clauses := []sq.Sqlizer{}

	for _, col := range cols {
		clauses = append(clauses, sq.Expr(fmt.Sprintf("%s ILIKE ?", col), q))
	}

	where := sq.Or(clauses)

	countQuery := db.sq.
		Select("COUNT(id)").
		From("sightings").
		Where(where)

	selectQuery := db.sq.
		Select("*").
		From("sightings").
		Where(where).
		OrderBy("occurred_at DESC")

	return db.paginate(countQuery, selectQuery, pageNumber)
}

func (db *DB) Insert(s *sightings.Sighting) error {

	q := db.sq.
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

func (db *DB) paginate(countQuery sq.SelectBuilder,
	selectQuery sq.SelectBuilder,
	pageNumber int64) (*sightings.Page, error) {

	page := &sightings.Page{Number: pageNumber, PageSize: pageSize}

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

	if err := db.Select(&page.Sightings, sql, args...); err != nil {
		return nil, err
	}

	return page, nil
}
