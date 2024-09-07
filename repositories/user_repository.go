package repositories

import (
	"task-list/config"
	"task-list/models"
)

// UserRepositoryInterface is an interface for the UserRepository
// It has the methods CreateUser, GetUserByID, and GetAllUsers
type UserRepositoryInterface interface {
	CreateUser(display_name string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserIdByDisplayName(display_name string) (int, error)
}

// UserRepositoryImpl is a struct that implements the UserRepositoryInterface
type UserRepositoryImpl struct{}

// UserRepository is an instance of the UserRepositoryInterface
var UserRepository UserRepositoryInterface = &UserRepositoryImpl{}

// CreateUser creates a user
// It takes a display_name string as a parameter
// It returns a user and an error
func (r *UserRepositoryImpl) CreateUser(display_name string) (models.User, error) {
	var query string = `
	INSERT INTO
	Users (
	DisplayName
	)
	VALUES (
	?
	)
	`

	result, err := config.Database_Connection.Exec(query, display_name)

	if err != nil {
		return models.User{}, err
	}

	id, err := result.LastInsertId()

	if id < 1 {
		return models.User{}, err
	}

	int_id := int(id)

	var user models.User = models.User{
		Id:          int_id,
		DisplayName: display_name,
	}

	return user, nil
}

// GetUserByID gets a user by id
// It takes an id int as a parameter
// It returns a user and an error
func (r *UserRepositoryImpl) GetUserByID(id int) (models.User, error) {
	var query string = `
	SELECT
	U.Id,
	U.DisplayName
	FROM
	Users U
	WHERE
	U.Id = ?
	`

	result, err := config.Database_Connection.Query(query, id)

	if err != nil {
		return models.User{}, err
	}

	var user models.User

	for result.Next() {
		err := result.Scan(&user.Id, &user.DisplayName)

		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// GetAllUsers gets all users
// It returns a slice of users and an error
// It returns an empty slice of users if there are no users
func (r *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var query string = `
	SELECT
	U.Id,
	U.DisplayName
	FROM
	Users U
	`

	result, err := config.Database_Connection.Query(query)

	if err != nil {
		return []models.User{}, err
	}

	users := []models.User{}

	for result.Next() {
		var user models.User
		err := result.Scan(&user.Id, &user.DisplayName)

		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetUserIdByDisplayName(display_name string) (int, error) {
	var query string = `
	SELECT
	U.Id
	FROM
	Users U
	WHERE
	U.DisplayName = ?
	`

	result, err := config.Database_Connection.Query(query, display_name)

	if err != nil {
		return 0, err
	}

	var id int

	for result.Next() {
		err := result.Scan(&id)

		if err != nil {
			return 0, err
		}
	}

	return id, nil
}
