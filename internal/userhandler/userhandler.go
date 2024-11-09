package userhandler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/iamhi/leo/db/postgres"
	"github.com/iamhi/leo/db/postgres/models"
	"github.com/iamhi/leo/internal/errors"
)

func CreateUser(username string, password string, email string) (UserDetails, errors.UserHandlerError) {
	var existingUser models.User

	postgres.Db.Model(&models.User{}).Where("username =?", username).Where("email =?", email).First(&existingUser)

	if existingUser.ID != 0 {
		existingUserError := errors.UserExistsError{}

		existingUserError.Email = email
		existingUserError.Username = username

		return UserDetails{}, existingUserError
	}

	hashedPassword, err := generateHashPassword(password)

	if err != nil {
		fmt.Printf("Error while hashing user password: %s", err)

		return UserDetails{}, nil
	}

	user := models.User{}

	user.Uuid = uuid.New().String()
	user.Username = username
	user.Password = hashedPassword
	user.Email = email

	postgres.Db.Create(&user)

	return LoginUser(username, password)
}

func LoginUser(username string, password string) (UserDetails, errors.UserHandlerError) {
	var user models.User

	postgres.Db.Model(&models.User{}).Where("username =?", username).First(&user)

	if user.ID == 0 {
		userNotFoundError := errors.UserNotFoundError{}

		userNotFoundError.Username = username

		return UserDetails{}, userNotFoundError
	}

	passwordEquals := compareHashPassword(password, user.Password)

	if !passwordEquals {
		badCredentials := errors.UserBadCredentialsError{}

		badCredentials.Username = username

		return UserDetails{}, badCredentials
	}

	userDetails := UserDetails{
		Uuid:     user.Uuid,
		Username: user.Username,
		Email:    user.Email,
		Token:    generateToken(username),
	}

	return userDetails, nil
}

func GetUserDetails(token string) (UserDetails, errors.UserHandlerError) {
	username, err := getUserName(token)

	if err != nil {
		return UserDetails{}, err
	}

	var user models.User

	postgres.Db.Model(&models.User{}).Where("username =?", username).First(&user)

	if user.ID == 0 {
		userNotFoundError := errors.UserNotFoundError{}

		userNotFoundError.Username = username

		return UserDetails{}, userNotFoundError
	}

	return UserDetails{
		Token:    token,
		Username: user.Username,
		Email:    user.Email,
		Uuid:     user.Uuid,
	}, nil
}

func RefreshToken(user_details UserDetails) (UserDetails, errors.UserHandlerError) {
	new_token, err := replaceToken(user_details.Token)

	if err != nil {
		return UserDetails{}, err
	}

	return UserDetails{
		Token:    new_token,
		Username: user_details.Username,
		Email:    user_details.Email,
	}, nil
}

func LogoutUser(user_details UserDetails) {
	invalidateToken(user_details.Token)
}
