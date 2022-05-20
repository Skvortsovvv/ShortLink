package links

import (
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"testingTask/internal/shorter"
)

func TestSQLAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err.Error())
	}
	defer db.Close()

	repo := NewLinksSQLRepo(db)

	TestCases := []TestCase{
		{
			longURL:  "https://www.ozon.ru/",
			shortULR: shorter.Shorter("https://www.ozon.ru/"),
		},
		{
			longURL:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			shortULR: shorter.Shorter("https://www.youtube.com/watch?v=dQw4w9WgXcQ"),
		},
		{
			longURL: `https://ru.wikipedia.org/wiki/Go#%D0%9D%D0%B0%D0%B7%D0%BD%D0%B0%D1%87%D0%B5%D0%BD%D0%B8%D0%B5,
_%D0%B8%D0%B4%D0%B5%D0%BE%D0%BB%D0%BE%D0%B3%D0%B8%D1%8F`,
			shortULR: shorter.Shorter(`https://ru.wikipedia.org/wiki/Go#%D0%9D%D0%B0%D0%B7%D0%BD%D0%B0%
D1%87%D0%B5%D0%BD%D0%B8%D0%B5,_%D0%B8%D0%B4%D0%B5%D0%BE%D0%BB%D0%BE%D0%B3%D0%B8%D1%8F`),
		},
		{
			longURL:  "https://ya.ru/",
			shortULR: shorter.Shorter("https://ya.ru/"),
		},
	}

	for index, testCase := range TestCases {
		mock.
			ExpectQuery("INSERT INTO links").
			WithArgs(testCase.shortULR, testCase.longURL).
			WillReturnRows(func() *sqlmock.Rows {
				result := sqlmock.NewRows([]string{"short_URL"})
				result.AddRow(testCase.shortULR)
				return result
			}())

		result, err := repo.Add(testCase.longURL)

		if err != nil {
			t.Errorf("unexpected err: %s", err.Error())
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if result != testCase.shortULR {
			t.Errorf("error at [%d] test case:\n\tExpected: %s\n\tGot: %s",
				index,
				testCase.shortULR,
				result)
		}

	}

}

func TestSQLGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("cant cerate mock: %s", err.Error())
	}
	defer db.Close()

	TestCases := []TestCase{
		{
			longURL:  "https://www.ozon.ru/",
			shortULR: shorter.Shorter("https://www.ozon.ru/"),
		},
		{
			longURL:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			shortULR: shorter.Shorter("https://www.youtube.com/watch?v=dQw4w9WgXcQ"),
		},
		{
			longURL: `https://ru.wikipedia.org/wiki/Go#%D0%9D%D0%B0%D0%B7%D0%BD%D0%B0%D1%87%D0%B5%D0%BD%D0%B8%D0%B5,
_%D0%B8%D0%B4%D0%B5%D0%BE%D0%BB%D0%BE%D0%B3%D0%B8%D1%8F`,
			shortULR: shorter.Shorter(`https://ru.wikipedia.org/wiki/Go#%D0%9D%D0%B0%D0%B7%D0%BD%D0%B0%
D1%87%D0%B5%D0%BD%D0%B8%D0%B5,_%D0%B8%D0%B4%D0%B5%D0%BE%D0%BB%D0%BE%D0%B3%D0%B8%D1%8F`),
		},
		{
			longURL:  "https://ya.ru/",
			shortULR: shorter.Shorter("https://ya.ru/"),
		},
	}

	repo := NewLinksSQLRepo(db)

	for index, testCase := range TestCases {
		mock.ExpectQuery("SELECT long_URL FROM links WHERE").
			WithArgs(testCase.shortULR).
			WillReturnRows(func() *sqlmock.Rows {
				result := sqlmock.NewRows([]string{"long_URL"})
				result.AddRow(testCase.longURL)
				return result
			}())

		result, err := repo.Get(testCase.shortULR)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}

		if result != testCase.longURL {
			t.Errorf("error at [%d] test case:\n\tExpected: %s\n\tGot: %s",
				index,
				testCase.shortULR,
				result)
		}
	}
}

func TestSQLError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("cant cerate mock: %s", err.Error())
	}
	defer db.Close()

	LinksSQLRepo := NewLinksSQLRepo(db)

	// Get method error
	TestCaseGet := TestCase{
		shortULR: shorter.Shorter("https://ozon.ru"),
		longURL:  "",
	}

	mock.ExpectQuery("SELECT long_URL FROM links WHERE").
		WithArgs(TestCaseGet.shortULR).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = LinksSQLRepo.Get(TestCaseGet.shortULR)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

}
