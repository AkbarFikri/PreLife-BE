package authRepository

import (
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type AuthRepository interface {
	Save(ctx context.Context, user domain.User) error
	CountEmail(ctx context.Context, email string) (int, error)
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
}

type authRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Save(ctx context.Context, user domain.User) error {
	arg := map[string]interface{}{
		"id":            user.ID,
		"email":         user.Email,
		"full_name":     user.FullName,
		"date_of_birth": user.DateOfBirth,
	}

	_, err := r.db.NamedExecContext(ctx, CreateUser, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r *authRepository) CountEmail(ctx context.Context, email string) (int, error) {
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

func (r *authRepository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	arg := map[string]interface{}{
		"email": email,
	}

	query, args, err := sqlx.Named(GetUserByEmail, arg)
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
