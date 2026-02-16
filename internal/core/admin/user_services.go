package admin

import "github.com/tanaxer01/pedalea/pkg/pedalea"

func (s *Service) UpdateUser(ID int, data pedalea.UserData) error {
	return s.userRepo.UpdateUser(ID, data)
}

func (s *Service) GetUserData(ID int) (*pedalea.UserData, error) {
	user, err := s.userRepo.GetUserData(ID)
	if err != nil {
		return nil, err
	}

	return &user.UserData, nil
}

func (s *Service) ListUsers() ([]pedalea.User, error) {
	return s.userRepo.ListUsers()
}
