package repositories

import (
	"context"
	"fmt"
	"learn-grpc/internal/entities"
	"learn-grpc/internal/utils"
	"time"

	"github.com/google/uuid"
)

func generateAvatar(imageUrl string) entities.Asset {
	asset, err := utils.GenerateOptimizeAsset(imageUrl)

	if err == nil {
		return *asset
	}

	// INFO: return default image for avatar
	asset, _ = utils.GenerateOptimizeAsset("https://google.com/sample-avatar.png")
	return *asset
}

type UserRepo interface {
	entities.GenericRepo[entities.UserRequest, entities.User, string]
	DoActivity(context.Context, string)
}

type _UserRepoImpl struct {
	user []entities.User
}

func NewUserRepo() UserRepo {
	return &_UserRepoImpl{
		user: make([]entities.User, 0),
	}
}

func (u *_UserRepoImpl) Create(_ context.Context, req entities.UserRequest) entities.User {
	newUser := entities.User{}
	newUser.ID = uuid.New().String()
	newUser.Avatar = generateAvatar(req.Avatar)
	newUser.Name = req.Name
	newUser.RegisteredDate = time.Now()
	newUser.LastActivityDate = time.Now()

	u.user = append(u.user, newUser)

	return newUser
}

func (u *_UserRepoImpl) FindAll(_ context.Context) []entities.User {
	return u.user
}

func (u *_UserRepoImpl) FindByID(_ context.Context, id string) (*entities.User, error) {
	var (
		user        entities.User
		isAvailable = false
	)

	for _, item := range u.user {
		if item.ID == id {
			user = item
			isAvailable = true
		}

		if isAvailable {
			break
		}
	}

	if isAvailable {
		return &user, nil
	}

	return nil, utils.NewNotFoundError(fmt.Sprintf("user %s is not found", id))
}

func (u *_UserRepoImpl) Delete(_ context.Context, id string) bool {
	isAvailable := false

	for index, item := range u.user {
		if item.ID == id {
			u.user = append(u.user[:index], u.user[index+1:]...)
			isAvailable = true
		}
	}

	if !isAvailable {
		utils.PanicIfNotFoundError(fmt.Errorf("user %s is not found", id))
		return false
	}

	return true
}

func (u *_UserRepoImpl) Update(_ context.Context, id string, req entities.UserRequest) *entities.User {
	var (
		selectedUser entities.User
		isAvailable  = false
	)

	for index, item := range u.user {
		if item.ID == id {
			selectedUser = item
			selectedUser.ID = id
			selectedUser.Name = req.Name
			selectedUser.Avatar = generateAvatar(req.Avatar)

			u.user = append(u.user[:index], u.user[index+1:]...)
			u.user = append(u.user, selectedUser)
			isAvailable = true
		}

		if isAvailable {
			break
		}
	}

	if !isAvailable {
		utils.PanicIfNotFoundError(fmt.Errorf("user %s is not found", id))
	}

	return &selectedUser
}

func (u *_UserRepoImpl) DoActivity(_ context.Context, id string) {
	for index, item := range u.user {
		if item.ID == id {
			selectedUser := item
			selectedUser.LastActivityDate = time.Now()

			u.user = append(u.user[:index], u.user[index+1:]...)
			u.user = append(u.user, selectedUser)
		}
	}
}
