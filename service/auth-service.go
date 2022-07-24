package service

import (
	"example/web-service-gin/entity"
	"example/web-service-gin/repository"
	"example/web-service-gin/request"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user request.UserRequest) request.UserRequest
	FindByEmail(email string) request.UserRequest
	IsDupplicate(email string) bool
}

type authService struct {
	userRepo repository.UserRepository
}

// CreateUser implements AuthService
func (service *authService) CreateUser(userRequest request.UserRequest) request.UserRequest {
	user := entity.User{}
	err := smapping.FillStruct(&userRequest, smapping.MapFields(&user))
	if err != nil {
		log.Fatal("Failed to mapping %v", err)
	}
	service.userRepo.CreateUser(user)
	return userRequest
}

// FindByEmail implements AuthService
func (service *authService) FindByEmail(email string) request.UserRequest {
	user := service.userRepo.FindByEmail(email)
	response := request.UserRequest{}
	smapping.FillStruct(&user, smapping.MapFields(&response))
	return response
}

// IsDupplicate implements AuthService
func (service *authService) IsDupplicate(email string) bool {
	isDuplicate := service.userRepo.IsDuplicateEmail(email)
	return !(isDuplicate.Error == nil)
}

// VerifyCredential implements AuthService
func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepo.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparePassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparePassword {
			return res
		}
		return false
	}
	return false
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func comparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
