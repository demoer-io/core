package repository

// GitHub struct manager
type GitHub struct{}

func (github *GitHub) getAllRepositories() ([]Repository, error) {
	repositories := []Repository{
		{ID: 1, Name: "Repo 1"},
		{ID: 2, Name: "Repo 2"},
	}

	return repositories, nil
}
