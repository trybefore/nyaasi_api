package nyaasi

var filters = []Filter{
	{
		ID:   "0",
		Name: "No filter",
	},
	{
		ID:   "1",
		Name: "No remakes",
	},
	{
		ID:   "2",
		Name: "Trusted only",
	},
}

// Filter defines a filter on Nyaa.si
type Filter struct {
	ID   string
	Name string
}

// GetFilterByName Find filter by name
func GetFilterByName(name string) *Filter {
	for _, f := range filters {
		if f.Name == name {
			return &f
		}
	}

	return nil
}

// FilterByID Find filter by id
func FilterByID(id string) *Filter {
	for _, f := range filters {
		if f.Name == id {
			return &f
		}
	}

	return nil
}
