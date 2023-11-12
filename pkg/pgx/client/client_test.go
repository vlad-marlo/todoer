package client

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"os"
	"testing"
)

type testConfig struct {
	Uri string `env:"TEST_DB_URI"`
}

func (cfg *testConfig) URI() string {
	return cfg.Uri
}

func testCfg(t testing.TB) Config {
	var cfg testConfig
	cfg.Uri = os.Getenv("TEST_DB_URI")
	if cfg.Uri == "" {
		t.Skip("db uri was not provided")
	}
	return &cfg
}

func TestClient_L(t *testing.T) {
	t.Run("nil client", func(t *testing.T) {
		l := (*Client)(nil).L()
		assert.Equal(t, zap.L(), l)
	})
	t.Run("non nil client", func(t *testing.T) {
		log, err := zap.NewProduction()
		require.NoError(t, err)
		cli := &Client{
			log: log,
		}
		assert.Equal(t, log, cli.L())
	})
}

func TestClient_P(t *testing.T) {
	t.Run("nil client", func(t *testing.T) {
		l := (*Client)(nil).P()
		assert.Nil(t, l)
	})
	t.Run("non nil client", func(t *testing.T) {
		cli := &Client{
			pool: &pgxpool.Pool{},
		}
		assert.Empty(t, cli.P())
	})
}

func TestNew_Positive(t *testing.T) {
	cfg := testCfg(t)
	lc := fxtest.NewLifecycle(t)
	cli, err := New(lc, cfg, zap.L())
	assert.NoError(t, err)
	assert.NotNil(t, cli)
}

type badCfg struct{}

func (badCfg) URI() string { return "bad uri" }

func TestNew_Negative_BadConfig(t *testing.T) {
	lc := fxtest.NewLifecycle(t)
	cli, err := New(lc, badCfg{}, zap.L())
	assert.Nil(t, cli)
	assert.Error(t, err)
}

func TestNewTest_DefaultClient(t *testing.T) {
	cli, td := NewTest(t)
	defer td()
	assert.NoError(t, cli.pool.Ping(context.Background()))
}

func TestNewTest_BadClient(t *testing.T) {
	require.NoError(t, os.Setenv("TEST_DB_URI", "postgres://postgres:password@localhost:5432/unknonnnonnononnnononno"))
	_, _ = NewTest(t)
	t.Log("test is unexpectedly was not skipped")
	t.Fail()
}

func TestTD(t *testing.T) {
	td := teardown(&pgxpool.Pool{}, "")
	assert.Panics(t, td)
}

func TestBadCli(t *testing.T) {
	cli := BadCli(t)
	assert.Error(t, cli.P().Ping(context.Background()))
}
