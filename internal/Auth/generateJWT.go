package Auth


import (
	"errors"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"Amora/internal/user"
	"Amora/internal/Logging"
)

func GenerateJWT(user *user.Users) (string, error) {
	// Load the secret key from environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY") 

	// if the thing breaks with the secret again, just uncomment this line 
	// it will work so yea don't know how that one works but just go with it no touchy please
	//fmt.Println("JWT Secret Key:", secretKey) 

	if secretKey == "" {
		Logging.Logs.Error("JWT secret key is not set in environment variables")
		return "", errors.New("JWT secret key is not set in environment variables") 
	}

	// createing the claim for what the JWT token will contain and return within the login response 
	claims := jwt.MapClaims{
		"id": user.Id,
		"email": user.Email,
		"FirstName": user.Firstname,
		"LastName": user.Lastname,
		"phone_number": user.PhoneNumber,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func GenerateResetToken(user *user.Users) (string, error) {
	// Load the secret key from environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY") 

	// if the thing breaks with the secret again, just uncomment this line 
	// it will work so yea don't know how that one works but just go with it no touchy please
	//fmt.Println("JWT Secret Key:", secretKey) 

	if secretKey == "" {
		Logging.Logs.Error("JWT secret key is not set in environment variables")
		return "", errors.New("JWT secret key is not set in environment variables") 
	}

	// createing the claim for what the JWT token will contain and return within the login response 
	claims := jwt.MapClaims{
		"id": user.Id,
		"email": user.Email,
		"phone_number": user.PhoneNumber,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"reset": true, // you can use this to distinguish reset tokens
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
