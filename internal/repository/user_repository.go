package repository

import (
	"database/sql"

	"github.com/user/quantum-server/internal/models"
)

type UserRepository interface {
	FindByTgID(tgID string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByTgID(tgID string) (*models.User, error) {
	query := `SELECT id, tg_id, username, first_name, last_name, lang_code, invited_by, is_active, created_at, updated_at FROM users WHERE tg_id = $1`
	user := &models.User{}
	err := r.db.QueryRow(query, tgID).Scan(
		&user.ID, &user.TgID, &user.Username, &user.FirstName, &user.LastName, &user.LangCode, &user.InvitedBy, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(user *models.User) error {
	query := `INSERT INTO users (tg_id, username, first_name, last_name, lang_code, invited_by, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, user.TgID, user.Username, user.FirstName, user.LastName, user.LangCode, user.InvitedBy, user.IsActive).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}
