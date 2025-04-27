package repository

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhlikus/tgbot/internal/repository/external/telegram/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_SendMessage(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(m *mock.MocktelegramClient)
		expectedError assert.ErrorAssertionFunc
	}{
		{
			name: "OK",
			mockSetup: func(m *mock.MocktelegramClient) {
				m.EXPECT().
					SendMessage(gomock.Any(), gomock.Any())
			},
			expectedError: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			tgClientMock := mock.NewMocktelegramClient(ctrl)
			tt.mockSetup(tgClientMock)

			rep, err := NewRepository(tgClientMock)
			require.NoError(t, err)

			err = rep.SendMessage(1, "Any text message.")
			tt.expectedError(t, err)
		})
	}
}
