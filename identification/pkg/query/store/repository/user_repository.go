package repository

import (
	"database/sql"

	models "github.com/L4B0MB4/PRYVT/identification/pkg/models/query"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	if db == nil {
		return nil
	}
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUserById(userId uuid.UUID) (*models.UserInfo, error) {
	var user models.UserInfo
	stmt, err := repo.db.Prepare("SELECT display_name, name, email, change_date FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(userId).Scan(&user.DisplayName, &user.Name, &user.Email, &user.ChangeDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetAllUsers(limit, offset int) ([]models.UserInfo, error) {
	if limit > 100 {
		limit = 100
	}
	stmt, err := repo.db.Prepare("SELECT id, display_name FROM users LIMIT ? OFFSET ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserInfo
	for rows.Next() {
		var user models.UserInfo
		if err := rows.Scan(&user.ID, &user.DisplayName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) AddOrReplaceUser(user *models.UserInfo) error {
	stmt, err := repo.db.Prepare(`
		INSERT INTO users (id, display_name, name, email, change_date)
		VALUES (?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			display_name = excluded.display_name,
			name = excluded.name,
			email = excluded.email,
			change_date = excluded.change_date
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.DisplayName, user.Name, user.Email, user.ChangeDate)
	if err != nil {
		return err
	}
	return nil
}
