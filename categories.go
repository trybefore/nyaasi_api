package nyaasi

import (
	"fmt"
	"strings"
)

var categories = []Category{
	{
		ID:   "0",
		Name: "All categories",
	},
	{
		ID:   "1",
		Name: "Anime",
		SubCategories: []Category{
			{
				ID:   "1",
				Name: "Anime Music Video",
			},
			{
				ID:   "2",
				Name: "English-translated",
			},
			{
				ID:   "3",
				Name: "Non-English-translated",
			},
			{
				ID:   "4",
				Name: "Raw",
			},
		},
	},
	{
		ID:   "2",
		Name: "Audio",
		SubCategories: []Category{
			{
				ID:   "1",
				Name: "Lossless",
			},
			{
				ID:   "2",
				Name: "Lossy",
			},
		},
	},
	{
		ID:   "3",
		Name: "Literature",
		SubCategories: []Category{
			{
				ID:   "1",
				Name: "English-translated",
			},
			{
				ID:   "2",
				Name: "Non-English-translated",
			},
			{
				ID:   "3",
				Name: "Raw",
			},
		},
	},
	{
		ID:   "4",
		Name: "Live Action",
		SubCategories: []Category{
			{
				ID:   "1",
				Name: "English-translated",
			},
			{
				ID:   "2",
				Name: "Idol/Promotional Video",
			},
			{
				ID:   "3",
				Name: "Non-English-translated",
			},
			{
				ID:   "4",
				Name: "Raw",
			},
		},
	},
	{
		ID:   "5",
		Name: "Pictures",
		SubCategories: []Category{
			{
				ID:   "1",
				Name: "Graphics",
			},
			{
				ID:   "2",
				Name: "Photos",
			},
		},
	},
	{
		ID:   "6",
		Name: "Software",
		SubCategories: []Category{
			{
				ID:   "1",
				Name: "Applications",
			},
			{
				ID:   "2",
				Name: "Games",
			},
		},
	},
}

// Category defines a category on Nyaa.si
type Category struct {
	ID            string
	Name          string
	SubCategories []Category
}

func (c *Category) clone() *Category {
	return &Category{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (c *Category) cloneWithSubcategory(subCategory Category) *Category {
	return &Category{
		ID:            c.ID,
		Name:          c.Name,
		SubCategories: []Category{subCategory},
	}
}

// Format format the category for use in URL without subcategory included
func (c *Category) Format() string {
	return c.FormatWithSubCategoryID("0")
}

// FormatWithSubCategoryID format the category for use in URL with subcategory included
func (c *Category) FormatWithSubCategoryID(id string) string {
	sub := c.GetSubCategoryByID(id)

	if sub == nil {
		return fmt.Sprintf("%s_%s", c.ID, "0")
	}

	return fmt.Sprintf("%s_%s", c.ID, sub.ID)
}

// FormatWithSubCategoryName format the category for use in URL with subcategory included
func (c *Category) FormatWithSubCategoryName(name string) string {
	sub := c.GetSubCategoryByName(name)

	if sub == nil {
		return fmt.Sprintf("%s_%s", c.ID, "0")
	}

	return fmt.Sprintf("%s_%s", c.ID, sub.ID)
}

// GetSubCategoryByName Find subcategory by name
func (c *Category) GetSubCategoryByName(name string) *Category {
	for _, c := range c.SubCategories {
		if c.Name == name {
			return &c
		}
	}

	return nil
}

// GetSubCategoryByID Find subcategory by id
func (c *Category) GetSubCategoryByID(id string) *Category {
	for _, c := range c.SubCategories {
		if c.ID == id {
			return &c
		}
	}

	return nil
}

// GetCategoryByID Find category by id
func GetCategoryByID(id string) *Category {
	for _, c := range categories {
		if c.ID == id {
			return &c
		}
	}

	return nil
}

// GetCategoryByName Find category(or subcategory) by name
func GetCategoryByName(name string) *Category {
	for _, c := range categories {
		if c.Name == name {
			return &c
		}
	}

	return nil
}

// ParseCategory parses the category(and subcategory, if exists) from string
func ParseCategory(s string) *Category {
	if split := strings.Split(s, " - "); len(split) > 0 {
		categoryName := split[0]
		subCategory := strings.Join(split[1:], " ")

		cat := GetCategoryByName(categoryName)

		if cat == nil {
			return nil
		}

		subCat := cat.GetSubCategoryByName(subCategory)

		if subCat != nil {
			cat = cat.cloneWithSubcategory(*subCat)
		}

		return cat
	} else if cat := GetCategoryByName(s); cat != nil {
		return cat.clone()
	}

	return nil
}
