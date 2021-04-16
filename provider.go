package mysqlusers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gotuna/gotuna"
	"golang.org/x/crypto/bcrypt"
)

// DBUser is a sample gotuna.User implementation used with mysql
type DBUser struct {
	ID           string
	Email        string
	Name         string
	Phone        string
	PasswordHash string
}

// GetID should return auto-incremented, unique "id" field from mysql database for this user
func (u DBUser) GetID() string {
	return u.ID
}

// NewRepository returns a new mysql implementation of gotuna.UserRepository
func NewRepository(db *sql.DB) gotuna.UserRepository {
	return mysqlUserRepository{db}
}

type mysqlUserRepository struct {
	db *sql.DB
}

func (u mysqlUserRepository) Authenticate(w http.ResponseWriter, r *http.Request) (gotuna.User, error) {

	email := strings.ToLower(strings.TrimSpace(r.FormValue("email")))
	Password := r.FormValue("password")

	if email == "" {
		return DBUser{}, errors.New("this field is required")
	}
	if Password == "" {
		return DBUser{}, errors.New("this field is required")
	}

	found, err := u.getUserByField("email", email)
	if err != nil {
		return DBUser{}, fmt.Errorf("cannot find user with this email: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(found.PasswordHash), []byte(Password)); err != nil {
		return DBUser{}, fmt.Errorf("passwords don't match %v", err)
	}

	return found, nil
}

func (u mysqlUserRepository) GetUserByID(id string) (gotuna.User, error) {
	return u.getUserByField("id", id)
}

func (u mysqlUserRepository) getUserByField(field, value string) (DBUser, error) {
	user := DBUser{}

	err := u.db.QueryRow(fmt.Sprintf(`
	SELECT id, email, name, phone, password_hash
	FROM users
	WHERE %s = ?
	`, field), value).
		Scan(&user.ID, &user.Email, &user.Name, &user.Phone, &user.PasswordHash)

	if err == sql.ErrNoRows {
		return DBUser{}, errors.New("user not found")
	}
	if err != nil {
		return DBUser{}, fmt.Errorf("user cannot be checked: %v", err)
	}

	return user, nil
}
