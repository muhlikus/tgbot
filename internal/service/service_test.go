package service

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhlikus/tgbot/internal/service/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GetMessages(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(m *mock.Mockrepository)
		expectedError assert.ErrorAssertionFunc
	}{
		{
			name: "OK",
			mockSetup: func(m *mock.Mockrepository) {
				m.EXPECT().
					GetMessages()
			},
			expectedError: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			rep := mock.NewMockrepository(ctrl)
			tt.mockSetup(rep)

			service, err := NewService(rep)
			require.NoError(t, err)

			messages, err := service.GetMessages()
			tt.expectedError(t, err)
			_ = messages
		})
	}
}

func TestService_SendFile(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(m *mock.Mockrepository)
		expectedError assert.ErrorAssertionFunc
	}{
		{
			name: "OK",
			mockSetup: func(m *mock.Mockrepository) {
				m.EXPECT().
					SendFile(gomock.Any(), gomock.Any(), gomock.Any())
			},
			expectedError: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			repo := mock.NewMockrepository(ctrl)
			tt.mockSetup(repo)

			service, err := NewService(repo)
			require.NoError(t, err)

			err = service.SendFile(1, "file.name", bytes.NewBufferString("file content"))
			tt.expectedError(t, err)
		})
	}
}

func TestService_SendMessage(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(m *mock.Mockrepository)
		expectedError assert.ErrorAssertionFunc
	}{
		{
			name: "OK",
			mockSetup: func(m *mock.Mockrepository) {
				m.EXPECT().
					SendMessage(gomock.Any(), gomock.Any())
			},
			expectedError: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			repo := mock.NewMockrepository(ctrl)
			tt.mockSetup(repo)

			rep, err := NewService(repo)
			require.NoError(t, err)

			err = rep.SendMessage(1, "Any text message.")
			tt.expectedError(t, err)
		})
	}
}
