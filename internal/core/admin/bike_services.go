package admin

import "github.com/tanaxer01/pedalea/pkg/pedalea"

func (s *Service) InsertBike(data pedalea.BikeData) error {
	return s.bikeRepo.InsertBike(data)
}

func (s *Service) UpdateBike(ID int, data pedalea.BikeData) error {
	return s.bikeRepo.UpdateBike(ID, data)
}

func (s *Service) ListBikes() ([]pedalea.Bike, error) {
	return s.bikeRepo.ListBikes()
}
