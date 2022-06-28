package config_test

import (
	"testing"

	"github.com/acetimesolutions/goledger-network/config"
)

func TestInit(t *testing.T) {
	var config config.Config
	config.LoadConfigs()

}
