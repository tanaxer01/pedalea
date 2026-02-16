package bike

import "github.com/tanaxer01/pedalea/pkg/pedalea"

type Service struct {
	bikeRepo BikeRepo
}

type BikeRepo interface {
	ListAvailableBikes() ([]pedalea.Bike, error)
}

func NewService(repo BikeRepo) *Service {
	return &Service{bikeRepo: repo}
}

func (s *Service) ListAvailableBikes() ([]pedalea.Bike, error) {
	return s.bikeRepo.ListAvailableBikes()
}
