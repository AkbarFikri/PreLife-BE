package userRepository

import (
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindById(ctx context.Context, id string) (domain.User, error)
	CountEmail(ctx context.Context, email string) (int, error)
}

type userRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

func (r userRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	arg := map[string]interface{}{"id": id}

	query, args, err := sqlx.Named(GetById, arg)
	if err != nil {
		return domain.User{}, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return domain.User{}, err
	}
	query = r.db.Rebind(query)

	var user domain.User
	if err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r userRepository) Save(ctx context.Context, user domain.User) error {
	if user.FullName == "" {
		user.FullName = "default-name"
	}

	arg := map[string]interface{}{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
	}

	_, err := r.db.NamedExecContext(ctx, CreateUser, arg)
	if err != nil {
		return ErrorExecContext
	}
	return nil
}

func (r userRepository) CountEmail(ctx context.Context, email string) (int, error) {
	arg := map[string]interface{}{
		"email": email,
	}

	query, args, err := sqlx.Named(CountEmail, arg)
	if err != nil {
		return -1, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return -1, err
	}
	query = r.db.Rebind(query)

	var count int
	if err := r.db.QueryRowxContext(ctx, query, args...).Scan(&count); err != nil {
		return -1, err
	}
	return count, nil
}
