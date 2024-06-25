package userRepository

import (
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	CountEmail(ctx context.Context, email string) (int, error)
	FindUserById(ctx context.Context, id string) (domain.User, error)
	FindAllUserProfiles(ctx context.Context, userId string) ([]domain.UserProfile, error)
	FindPregnantProfileById(ctx context.Context, id string) (domain.PregnantProfile, error)
	FindNonPregnantProfileById(ctx context.Context, id string) (domain.NotPregnantProfile, error)
	SaveUserPregnantProfile(ctx context.Context, profile domain.PregnantProfile) error
	SaveUserNotPregnantProfile(ctx context.Context, profile domain.NotPregnantProfile) error
}

type userRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

func (r userRepository) Save(ctx context.Context, user domain.User) error {
	arg := map[string]interface{}{
		"id":            user.ID,
		"email":         user.Email,
		"full_name":     user.FullName,
		"date_of_birth": user.DateOfBirth,
	}

	_, err := r.db.NamedExecContext(ctx, CreateUser, arg)
	if err != nil {
		return ErrorExecContext
	}
	return nil
}

func (r userRepository) SaveUserPregnantProfile(ctx context.Context, profile domain.PregnantProfile) error {
	arg := map[string]interface{}{
		"id":            profile.ID,
		"user_id":       profile.UserId,
		"is_pregnant":   profile.IsPregnant,
		"profile_name":  profile.ProfileName,
		"pregnant_date": profile.PregnantDate,
	}

	_, err := r.db.NamedExecContext(ctx, CreateProfilPregnant, arg)
	if err != nil {
		return ErrorExecContext
	}
	return nil
}

func (r userRepository) SaveUserNotPregnantProfile(ctx context.Context, profile domain.NotPregnantProfile) error {
	arg := map[string]interface{}{
		"id":           profile.ID,
		"user_id":      profile.UserId,
		"is_pregnant":  profile.IsPregnant,
		"profile_name": profile.ProfileName,
		"birth_date":   profile.BirthDate,
		"height":       profile.Height,
		"weight":       profile.Weight,
	}

	_, err := r.db.NamedExecContext(ctx, CreateProfilNotPregnant, arg)
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

func (r userRepository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
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

func (r userRepository) FindUserById(ctx context.Context, id string) (domain.User, error) {
	arg := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(GetUserById, arg)
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

func (r userRepository) FindPregnantProfileById(ctx context.Context, id string) (domain.PregnantProfile, error) {
	arg := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(GetPregnantProfileById, arg)
	if err != nil {
		return domain.PregnantProfile{}, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return domain.PregnantProfile{}, err
	}
	query = r.db.Rebind(query)

	var profile domain.PregnantProfile
	if err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&profile); err != nil {
		return domain.PregnantProfile{}, err
	}

	return profile, nil
}

func (r userRepository) FindNonPregnantProfileById(ctx context.Context, id string) (domain.NotPregnantProfile, error) {
	arg := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(GetNonPregnantProfileById, arg)
	if err != nil {
		return domain.NotPregnantProfile{}, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return domain.NotPregnantProfile{}, err
	}
	query = r.db.Rebind(query)

	var profile domain.NotPregnantProfile
	if err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&profile); err != nil {
		return domain.NotPregnantProfile{}, err
	}

	return profile, nil
}

func (r userRepository) FindAllUserProfiles(ctx context.Context, userId string) ([]domain.UserProfile, error) {
	arg := map[string]interface{}{
		"user_id": userId,
	}

	var userProfiles []domain.UserProfile
	query, args, err := sqlx.Named(GetAllProfilesByUserId, arg)
	if err != nil {
		return userProfiles, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return userProfiles, err
	}
	query = r.db.Rebind(query)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return userProfiles, err
	}

	for rows.Next() {
		var userProfile domain.UserProfile
		if err := rows.StructScan(&userProfile); err != nil {
			return userProfiles, err
		}
		userProfiles = append(userProfiles, userProfile)
	}

	return userProfiles, nil
}
