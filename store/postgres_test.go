package store

import (
	"fmt"
	"github.com/danjac/sightings/models"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"testing"
	"time"
)

var db *DB

func TestMain(m *testing.M) {

	viper.SetDefault("test_db_name", "sightings_test")
	viper.SetDefault("test_db_user", "postgres")
	viper.SetDefault("test_db_password", "postgres")
	viper.SetDefault("test_db_host", "127.0.0.1")

	viper.SetConfigName("test.yml")
	viper.AddConfigPath("..")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Unable to find config file, using defaults")
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations/postgres",
	}

	connString := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s sslmode=disable",
		viper.GetString("test_db_name"),
		viper.GetString("test_db_user"),
		viper.GetString("test_db_password"),
		viper.GetString("test_db_host"),
	)
	var err error

	if db, err = Connect(connString); err != nil {
		panic(err)
	}

	defer db.Close()

	if _, err := migrate.Exec(db.DB.DB, "postgres", migrations, migrate.Up); err != nil {
		panic(err)
	}

	result := m.Run()

	if _, err := migrate.Exec(db.DB.DB, "postgres", migrations, migrate.Down); err != nil {
		panic(err)
	}

	os.Exit(result)

}

func truncateDB() {
	if _, err := db.Exec("TRUNCATE sightings"); err != nil {
		panic(err)
	}
}

func createSighting() *models.Sighting {
	s := &models.Sighting{
		OccurredAt:  time.Now(),
		ReportedAt:  time.Now(),
		Location:    "Iowa City, IA",
		Description: "Strange glowing object",
		Shape:       "Saucer",
		Latitude:    41.6611277,
		Longitude:   -91.46778,
		Duration:    "10 minutes",
	}

	w := &DBWriter{db}

	if err := w.Insert(s); err != nil {
		panic(err)
	}

	return s
}

func TestGetOneIfNone(t *testing.T) {
	truncateDB()

	r := &DBReader{db}

	s, err := r.GetOne("1234")
	if err == nil || s != nil {
		t.Fatal("Should fail here")
	}

}

func TestGetOne(t *testing.T) {
	truncateDB()

	fixture := createSighting()

	r := &DBReader{db}

	s, err := r.GetOne(strconv.Itoa(int(fixture.ID)))
	if err != nil {
		t.Fatal(err)
	}

	if s.ID != fixture.ID {
		t.Fatalf("Invalid ID, expected %d and got %d", fixture.ID, s.ID)
	}

}

func TestFind(t *testing.T) {
	truncateDB()

	numRows := 3

	for i := 0; i < numRows; i++ {
		createSighting()
	}

	r := &DBReader{db}

	page, err := r.Find(1)
	if err != nil {
		t.Fatal(err)
	}

	if int(page.TotalRows) != numRows {
		t.Fatalf("Incorrect number of rows, expected %d got %d", numRows, page.TotalRows)
	}
}

func TestSearch(t *testing.T) {
	truncateDB()

	numRows := 3

	for i := 0; i < numRows; i++ {
		createSighting()
	}

	r := &DBReader{db}

	page, err := r.Search("iowa", 1)
	if err != nil {
		t.Fatal(err)
	}

	if int(page.TotalRows) != numRows {
		t.Fatalf("Incorrect number of rows, expected %d got %d", numRows, page.TotalRows)
	}
}

func TestInsert(t *testing.T) {

	truncateDB()

	s := &models.Sighting{
		OccurredAt:  time.Now(),
		ReportedAt:  time.Now(),
		Location:    "Iowa City, IA",
		Description: "Strange glowing object",
		Shape:       "Saucer",
		Latitude:    41.6611277,
		Longitude:   -91.46778,
		Duration:    "10 minutes",
	}

	w := &DBWriter{db}

	if err := w.Insert(s); err != nil {
		t.Fatal(err)
	}

	if s.ID == 0 {
		t.Fatal("New ID not assigned")
	}

}
