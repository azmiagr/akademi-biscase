package service

import (
	"akademi-business-case/entity"
	"akademi-business-case/internal/repository"
	"akademi-business-case/model"
	"akademi-business-case/pkg/bcrypt"
	"akademi-business-case/pkg/database/mariadb"
	"akademi-business-case/pkg/jwt"
	"akademi-business-case/pkg/mail"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(param *model.UserRegister) (*model.RegisterResponse, error)
	VerifyUser(param model.VerifyUser) error
	Login(param model.UserLogin) (*model.LoginResponse, error)
	GetUserProfile(userID uuid.UUID) (*model.GetUserProfileResponse, error)
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserService struct {
	db             *gorm.DB
	UserRepository repository.IUserRepository
	CartRepository repository.ICartRepository
	OtpRepository  repository.IOtpRepository
	Bcrypt         bcrypt.Interface
	JwtAuth        jwt.Interface
}

func NewUserService(userRepository repository.IUserRepository, cartRepository repository.ICartRepository, otpRepository repository.IOtpRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) IUserService {
	return &UserService{
		db:             mariadb.Connection,
		UserRepository: userRepository,
		CartRepository: cartRepository,
		OtpRepository:  otpRepository,
		Bcrypt:         bcrypt,
		JwtAuth:        jwtAuth,
	}
}

func (s *UserService) Register(param *model.UserRegister) (*model.RegisterResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	var result *model.RegisterResponse

	_, err := s.UserRepository.GetUser(tx, model.UserParam{
		Email: param.Email,
	})

	if err == nil {
		return nil, errors.New("email already registered")
	}

	if param.Password != param.ConfirmPassword {
		return nil, errors.New("password not match")
	}

	hash, err := s.Bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		UserID:    id,
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Email:     param.Email,
		Password:  hash,
		RoleID:    2,
	}

	_, err = s.UserRepository.CreateUser(tx, user)
	if err != nil {
		return nil, err
	}

	token, err := s.JwtAuth.CreateJWTToken(user.UserID, false)
	if err != nil {
		return nil, err
	}

	cartId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	cart := &entity.Cart{
		CartID: cartId,
		UserID: id,
	}

	_, err = s.CartRepository.CreateCart(tx, cart)
	if err != nil {
		return nil, err
	}

	code := mail.GenerateCode()
	otp := &entity.OtpCode{
		OtpID:  uuid.New(),
		UserID: user.UserID,
		Code:   code,
	}

	err = s.OtpRepository.CreateOtp(tx, otp)
	if err != nil {
		return nil, err
	}

	err = mail.SendEmail(user.Email, "OTP Verification", "Your OTP verification code is "+code+"")
	if err != nil {
		return nil, err
	}

	result = &model.RegisterResponse{
		Token: token,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) VerifyUser(param model.VerifyUser) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	otp, err := s.OtpRepository.GetOtp(tx, model.GetOtp{
		UserID: param.UserID,
	})
	if err != nil {
		return err
	}

	if otp.Code != param.OtpCode {
		return errors.New("invalid otp code")
	}

	expiredTime, err := strconv.Atoi(os.Getenv("EXPIRED_OTP"))
	if err != nil {
		return err
	}

	expiredThreshold := time.Now().UTC().Add(-time.Duration(expiredTime) * time.Minute)
	if otp.UpdatedAt.Before(expiredThreshold) {
		return errors.New("otp expired")
	}

	user, err := s.UserRepository.GetUser(tx, model.UserParam{
		UserID: param.UserID,
	})
	if err != nil {
		return err
	}

	user.Status = "active"
	_, err = s.UserRepository.UpdateUser(tx, user)
	if err != nil {
		return err
	}

	err = s.OtpRepository.DeleteOtp(tx, otp)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(param model.UserLogin) (*model.LoginResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	var (
		result  *model.LoginResponse
		isAdmin bool
	)

	user, err := s.UserRepository.GetUser(tx, model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return nil, errors.New("email or password incorrect")
	}

	if user.RoleID == 1 {
		isAdmin = true
	} else {
		isAdmin = false
	}

	err = s.Bcrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return nil, errors.New("email or password incorrect")
	}

	token, err := s.JwtAuth.CreateJWTToken(user.UserID, isAdmin)
	if err != nil {
		return nil, err
	}

	result = &model.LoginResponse{
		Token: token,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *UserService) GetUserProfile(userID uuid.UUID) (*model.GetUserProfileResponse, error) {
	var result *model.GetUserProfileResponse

	tx := s.db.Begin()
	defer tx.Rollback()

	user, err := s.UserRepository.GetUser(tx, model.UserParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	result = &model.GetUserProfileResponse{
		Username: user.FirstName + " " + user.LastName,
		Email:    user.Email,
	}

	return result, nil
}

func (u *UserService) GetUser(param model.UserParam) (*entity.User, error) {
	return u.UserRepository.GetUser(u.db, param)
}
