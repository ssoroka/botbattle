package main

import (
	"../client"
  "math/rand"
)

const (
  UP = 90
  LEFT = 0
  RIGHT = 180
  DOWN = 270
)

type Bot struct {
  *client.BotClient
}

func (self *Bot) ScanAndShoot() {
  for len(self.Scan()) > 0 {
    self.Shield()
    self.FireGun()
    self.FireGun()
    self.FireCannon()
  }
}

func (self *Bot) Spin() {
  for i:=0; i<4; i++ {
    self.RotLeft()
    self.ScanAndShoot()
  }
}

func (bot *Bot) RandomStep() {
  new_direction := rand.Intn(4) * 90
    // orient ourselves to the right.
  for bot.Status().Rotation != new_direction {
    bot.RotLeft()
    bot.ScanAndShoot()
  }

  bot.MoveForward()
}

func main() {
  bot := &Bot{client.NewBotClient("192.168.1.111:3333", "wanderer")}
  bot.Spin()
  for {
    bot.RandomStep()
    bot.Spin()
  }
}
