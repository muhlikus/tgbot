package repository

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhlikus/tgbot/internal/repository/external/telegram/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_SendFile(t *testing.T) {
	// type args struct {
	// 	chatID      int
	// 	fileName    string
	// 	fileContent string
	// }

	tests := []struct {
		name string
		//args           args
		mockSetup     func(m *mock.MocktelegramClient)
		expectedError assert.ErrorAssertionFunc
	}{
		{
			name: "OK",
			// args: args{
			// 	chatID: 1,

			// },
			mockSetup: func(m *mock.MocktelegramClient) {
				m.EXPECT().
					SendDocument(gomock.Any(), gomock.Any(), gomock.Any())
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

			err = rep.SendFile(1, "file.name", bytes.NewBufferString("file content"))
			tt.expectedError(t, err)
		})
	}
}
