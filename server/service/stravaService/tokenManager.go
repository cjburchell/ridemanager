package stravaService

import (
	"time"

	"github.com/cjburchell/ridemanager/service/data"
	"github.com/cjburchell/ridemanager/service/data/models"
)

type TokenManager interface {
	Get() (*models.Token, error)
}

type tokenManager struct {
	authenticator Authenticator
	userId        models.AthleteId
	dataService   data.IService
	token         *models.Token
}

func GetTokenManager(authenticator Authenticator, userId models.AthleteId, dataService data.IService, token *models.Token) TokenManager {
	return &tokenManager{
		authenticator,
		userId,
		dataService,
		token,
	}
}

func (m *tokenManager) Get() (*models.Token, error) {
	if m.token == nil {
		user, err := m.dataService.GetUser(m.userId)
		if err != nil {
			return nil, err
		}

		m.token = &user.StravaToken
	}

	if m.token.ExpiresAt < time.Now().Unix() {
		newToken, err := m.authenticator.Renew(m.token.RefreshToken)
		if err != nil {
			return nil, err
		}

		user, err := m.dataService.GetUser(m.userId)
		if err != nil {
			return nil, err
		}

		user.StravaToken = *newToken

		err = m.dataService.UpdateUser(*user)
		if err != nil {
			return nil, err
		}
	}

	return m.token, nil
}
