package utils

import (
	"encoding/json"
	"os"
)

// Save users to file
func (store *UserStore) SaveToFile() error {
	data, err := json.MarshalIndent(store.Users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(store.File, data, 0644)
}

// Load users from file
func (store *UserStore) LoadFromFile() error {
	data, err := os.ReadFile(store.File)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &store.Users)
}
