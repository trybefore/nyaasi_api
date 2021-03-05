package nyaasi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryByID(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string
		wanted interface{}
	}{
		{input: "1", wanted: "Anime"},
		{input: "2", wanted: "Audio"},
		{input: "3", wanted: "Literature"},
		{input: "0", wanted: "All categories"},
		{input: "4", wanted: "Live Action"},
		{input: "7", wanted: nil},
	}

	for _, test := range tests {
		if test.wanted == nil {
			assert.Nil(GetCategoryByID(test.input))
		} else {
			assert.Equal(test.wanted, GetCategoryByID(test.input).Name)
		}
	}
}

func TestCategoryByName(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string
		wanted interface{}
	}{
		{wanted: "1", input: "Anime"},
		{wanted: "2", input: "Audio"},
		{wanted: "3", input: "Literature"},
		{wanted: "0", input: "All categories"},
		{wanted: "4", input: "Live Action"},
		{wanted: nil, input: "Asdfasdfasdf"}}
	for _, test := range tests {
		if test.wanted == nil {
			assert.Nil(GetCategoryByName(test.input))
		} else {
			assert.Equal(test.wanted, GetCategoryByName(test.input).ID)
		}
	}
}

func TestCategoryFormat(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		categoryID    string
		subcategoryID string
		wanted        string // url-formatted string
	}{
		{
			categoryID:    "1",
			subcategoryID: "1",
			wanted:        "1_1",
		},
		{
			categoryID:    "1",
			subcategoryID: "-1",
			wanted:        "1_0",
		},
		{
			categoryID:    "2",
			subcategoryID: "2",
			wanted:        "2_2",
		},
		{
			categoryID:    "4",
			subcategoryID: "2",
			wanted:        "4_2",
		},
	}

	for _, test := range tests {
		cat := GetCategoryByID(test.categoryID)

		formatted := cat.FormatWithSubCategoryID(test.subcategoryID)

		assert.Equal(test.wanted, formatted)
	}
}

func TestParseCategory(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string
		wanted *Category
	}{
		{
			input: "Anime - English-translated",
			wanted: &Category{
				ID:   "1",
				Name: "Anime",
				SubCategories: []Category{
					{
						ID:   "2",
						Name: "English-translated",
					},
				},
			},
		},
		{
			input: "Anime - Non-English-translated",
			wanted: &Category{
				ID:   "1",
				Name: "Anime",
				SubCategories: []Category{
					{
						ID:   "3",
						Name: "Non-English-translated",
					},
				},
			},
		},
		{
			input: "Audio - Lossless",
			wanted: &Category{
				ID:   "2",
				Name: "Audio",
				SubCategories: []Category{
					{
						ID:   "1",
						Name: "Lossless",
					},
				},
			},
		},
	}

	for _, test := range tests {
		actual := ParseCategory(test.input)
		assert.Equal(test.wanted, actual)
	}
}
