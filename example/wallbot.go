// this is an original draft of the wallbot, kept around just so we can see what it's not good at.
// see wallbot2 for a better implementation.
package main

import (
	"../client"
)

type Bot struct {
  *client.BotClient
}
func (self *Bot) OnWall() bool {
  status := self.Status()
  return status.X == 0;
}

func (self *Bot) AtTop() bool {
  status := self.Status()
  return status.Y == 0;
}

func (self *Bot) AtBottom() bool {
  status := self.Status()
  return status.Y == 10;
}

func (self *Bot) GoToWall() {
  status := self.Status()
  switch status.Rotation {
    case 90: self.RotRight()
    case 180: self.MoveBackward()
    case 270: self.RotLeft()
    case 0: self.MoveForward()
  }
}

func (bot *Bot) Sweep() {
  dir := -1
  for {
    // status := bot.Status()
    bot.Shield()
    for bot.Status().Rotation != 180 {
      if dir == -1 {
        bot.RotRight()
      } else {
        bot.RotLeft()
      }
    }
    bot.Shield()
    bot.FireGun()
    bot.FireGun()
    bot.FireCannon()
    bot.Shield()
    if dir == -1 && bot.AtTop() {
      dir = 1
    }
    if dir == 1 && bot.AtBottom() {
      dir = -1
    }
    if dir == -1 {
      bot.RotLeft()
    } else {
      bot.RotRight()
    }
    bot.FireGun()
    bot.MoveForward()
  }
}

func main() {
  bot := &Bot{client.NewBotClient("192.168.1.111:3333", "wallbot")}
  i := 0
  for {
    i++
    if i % 2 == 0 {
      bot.Shield()
    } else {
      if !bot.OnWall() {
        bot.GoToWall()
      } else {
        bot.Sweep()
      }
    }
  }
}
