package jwt

import (
	"testing"
	"time"

	"github.com/google/uuid"
	common_utils "github.com/kholiqdev/go-common/utils"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	config, err := common_utils.LoadBaseConfig("../", "test")
	if err != nil {
		t.Error(err)
	}
	jwt, err := NewJWTMaker(config)

	if err != nil {
		t.Error(err)
	}

	userId := uuid.New()

	payload := PayloadParams{
		UserId: userId,
		Email:  "test@mail.com",
		Role:   "admin",
		Status: "active",
	}

	t.Run("Create token", func(t *testing.T) {
		token, _, err := jwt.CreateToken("access", payload, 60*time.Second)

		if err != nil {
			t.Error(err)
		}

		assert.NotEmpty(t, token)
	})

	t.Run("Verify token", func(t *testing.T) {
		token, _, err := jwt.CreateToken("access", payload, 60*time.Second)

		if err != nil {
			t.Error(err)
		}

		payload, err := jwt.VerifyToken(token)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, userId, payload.UserId)
	})

}
