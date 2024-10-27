package logic

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_GetToken(t *testing.T) {
	var sut, err = initializeToken("")
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("get token", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		getToken, _, err := sut.GetToken(ctx, 1)
		if err != nil {
			t.Error(err)
			return
		}
		assert.True(t, len(getToken) != 0 && getToken != "")
	})
}
