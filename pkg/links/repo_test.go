package links

import (
	"reflect"
	"testing"
	"testingTask/internal/shorter"
)

type TestCase struct {
	longURL  string
	shortULR string
}

func TestAdd(t *testing.T) {
	TestCases := []TestCase{
		{
			longURL:  "https://www.ozon.ru/",
			shortULR: shorter.Shorter("https://www.ozon.ru/"),
		},
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

	LinksMemoryRepo := NewLinksMemoryRepo()

	for index, testCase := range TestCases {
		result, _ := LinksMemoryRepo.Add(testCase.longURL)
		if !reflect.DeepEqual(testCase.shortULR, result) {
			t.Errorf("error at [%d] test case:\n\tExpected: %v\n\tGot: %v", index, testCase.shortULR, result)
		}
	}
}

func TestGet(t *testing.T) {

	TestCases := []TestCase{
		{
			longURL:  "https://www.ozon.ru/",
			shortULR: shorter.Shorter("https://www.ozon.ru/"),
		},
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

	LinksMemoryRepo := NewLinksMemoryRepo()

	for _, value := range TestCases {
		_, err := LinksMemoryRepo.Add(value.longURL)
		if err != nil {
			t.Errorf("error while itial: %s", err.Error())
		}
	}

	for index, testCase := range TestCases {
		result, _ := LinksMemoryRepo.Add(testCase.longURL)
		if !reflect.DeepEqual(testCase.shortULR, result) {
			t.Errorf("error at [%d] test case:\n\tExpected: %v\n\tGot: %v", index, testCase.shortULR, result)
		}
	}
}

func TestError(t *testing.T) {

	TestCases := []TestCase{
		{
			longURL:  "https://www.ozon.ru/",
			shortULR: "",
		},
		{
			longURL:  "https://www.ozon.ru/",
			shortULR: "",
		},
		{
			longURL:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			shortULR: "",
		},
	}

	LinksMemoryRepo := NewLinksMemoryRepo()

	for index, testCase := range TestCases {
		result, _ := LinksMemoryRepo.Add(testCase.longURL)
		if !reflect.DeepEqual(testCase.shortULR, result) {
			t.Errorf("error at [%d] test case:\n\tExpected: %v\n\tGot: %v", index, testCase.shortULR, result)
		}
	}
}
