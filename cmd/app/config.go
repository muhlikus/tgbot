package main

import (
	"github.com/muhlikus/telegramclient"
)

type config struct {
	debug    bool
	TgClient telegramclient.Config
}
