package usecase

import (
	"context"
	"effective-task/internal/users"
	"effective-task/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type usersUc struct {
	usersRepo users.Repository
}

func NewUsersUc(usersRepo users.Repository) users.Usecase {
	return &usersUc{usersRepo: usersRepo}
}

func (s *usersUc) DeleteById(ctx context.Context, id int) error {
	return s.usersRepo.DeleteById(ctx, id)
}
func (c *usersUc) Add(ctx context.Context, user *users.UserDto) error {

	enrichData, err := enrichUser(user.Name)
	if err != nil {
		return fmt.Errorf("users.usecase.add: %v", err)
	}

	userRet := users.UserDto{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         enrichData.Age,
		Gender:      enrichData.Gender,
		Nationalize: enrichData.Country,
	}

	return c.usersRepo.Add(ctx, &userRet)
}

type resp struct {
	Age     int
	Gender  string
	Country []struct {
		Country_id  string  `json:"country_id"`
		Probability float64 `json:"probability"`
	}
}
type respRet struct {
	Age     int
	Gender  string
	Country string
}

func enrichUser(username string) (respRet, error) {

	urls := []string{
		fmt.Sprintf("https://api.agify.io/?name=%v", username),
		fmt.Sprintf("https://api.genderize.io/?name=%v", username),
		fmt.Sprintf("https://api.nationalize.io/?name=%v", username),
	}

	wg := sync.WaitGroup{}
	respArr := make([]resp, 3)
	for i, url := range urls {
		wg.Add(1)
		i := i
		url := url
		go func() {
			defer wg.Done()
			ageReq, err := http.Get(url)
			if err != nil {
				log.Printf("http get: %v", err.Error())
			}
			ageData, err := utils.JsonUnmarshal[resp](ageReq)
			if err != nil {
				log.Printf("response parse: %v", err.Error())
			}
			respArr[i] = *ageData
		}()
	}

	wg.Wait()
	var ret respRet
	for _, resp := range respArr {
		if resp.Age != 0 {
			ret.Age = resp.Age
		}
		if len(resp.Country) > 0 {
			ret.Country = resp.Country[0].Country_id
		}
		if resp.Gender != "" {
			ret.Gender = resp.Gender
		}
	}

	return ret, nil
}

var GetAllParamKeys = []string{"name", "surname", "patronymic", "age", "nationalize"}

func (s *usersUc) GetAllFiltered(ctx context.Context, filter *users.UserFilter) ([]users.User, error) {
	return s.usersRepo.GetAllFiltered(ctx, filter)
}

func (s *usersUc) UpdateById(ctx context.Context, user *users.User) (*users.User, error) {
	return s.usersRepo.UpdateById(ctx, user)
}
