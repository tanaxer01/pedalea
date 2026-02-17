package admin

import "github.com/tanaxer01/pedalea/pkg/pedalea"

func (s *Service) UpdateRental(ID int, data pedalea.RentalData) error {
	return s.rentalRepo.UpdateRental(ID, data)
}

func (s *Service) GetRentalData(ID int) (*pedalea.Rental, error) {
	return s.rentalRepo.GetRentalByID(ID)
}

func (s *Service) ListRentals() ([]pedalea.Rental, error) {
	return s.rentalRepo.ListRentals()
}
