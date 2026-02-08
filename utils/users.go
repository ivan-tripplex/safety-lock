package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash []byte `json:"password_hash"`
	OpenAfter    string `json:"open_after"`
	BlockAfter   string `json:"block_after"`
}

type UserStore struct {
	Users map[string]User
	File  string
}

// Initialize a new user store
func NewUserStore(file string) *UserStore {
	return &UserStore{
		Users: make(map[string]User),
		File:  file,
	}
}

// Working with hashing passwords
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

// Create a new user
func (store *UserStore) CreateUser(name, password, openTime, blockTime string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := User{
		ID:           uuid.NewString(),
		Username:     name,
		PasswordHash: hash,
		OpenAfter:    openTime,
		BlockAfter:   blockTime,
	}

	// Save the user to json file
	store.Users[user.ID] = user
	return store.SaveToFile()
}

func (store *UserStore) UserByUsername(username string) (User, error) {
	for _, user := range store.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

// Add time to open and block
func (store *UserStore) AddToTime(open, close string) (time.Time, time.Time) {
	now := time.Now()

	timeToOpen := strings.Split(open, ":")
	timeToClose := strings.Split(close, ":")

	hourToOpen, _ := strconv.Atoi(timeToOpen[0])
	minuteToOpen, _ := strconv.Atoi(timeToOpen[1])

	hourToClose, _ := strconv.Atoi(timeToClose[0])
	minuteToClose, _ := strconv.Atoi(timeToClose[1])

	openTime := time.Date(now.Year(), now.Month(), now.Day(), hourToOpen, minuteToOpen, 0, 0, now.Location())
	closeTime := time.Date(now.Year(), now.Month(), now.Day(), hourToClose, minuteToClose, 0, 0, now.Location())

	return openTime, closeTime
}
