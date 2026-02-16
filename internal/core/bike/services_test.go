package bike

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type MockBikeRepo struct{ mock.Mock }

func (m *MockBikeRepo) ListAvailableBikes() ([]pedalea.Bike, error) {
	args := m.Called()
	return args.Get(0).([]pedalea.Bike), args.Error(1)
}

// NOTE: We shouldn't need to check rentals table if we make
// sure to update is_available on start / end of a rental.
func TestNoAvailableBikes(t *testing.T) {
	repo := new(MockBikeRepo)

	repo.On("ListAvailableBikes").Return([]pedalea.Bike{}, nil)

	s := NewService(repo)

	bikes, err := s.ListAvailableBikes()

	require.Nil(t, err)
	require.Empty(t, bikes)
}
