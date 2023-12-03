package db

// TransactDatastore is an interface that outlines methods for interacting with the transactional database.
// It abstracts the database implementation, allowing for flexible data storage strategies.
type TransactDatastore interface {
	Init() error

	GetUserById(id uint) (*User, error) // Retrieves a user by their ID.
	GetUser(username string) (*User, error)
	CreateUser(user *User) (*User, error) // Creates a new user record.
	UpdateUser(user *User) error          // Updates an existing user record.
	DeleteUser(id uint) error

	//GetVault(id uint) (*models.Vault, error)        // Fetches a vault by its ID.
	//GetVaults(userId uint) ([]*models.Vault, error) // Retrieves all vaults for a given user.
	//CreateVault(vault *models.Vault) error          // Adds a new vault record.
	//UpdateVault(vault *models.Vault) error
	//DeleteVault(id uint) error

	//GetTag(id uint) error
	//CreateTag(tag *models.Tag) error
	//UpdateTag(tag *models.Tag) error
	//DeleteTag(id uint) error
}

type AnalyticDatastore interface{}
