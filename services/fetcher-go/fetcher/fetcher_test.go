package fetcher

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fetcher(t *testing.T) {
	s, err := Fetcher(context.Background(), "https://qiita.com/gold-kou/items/a1cc2be6045723e242eb")
	assert.NoError(t, err)
	assert.Equal(t, "いまさらだけどgRPCに入門したので分かりやすくまとめてみた - Qiita", s)
}
