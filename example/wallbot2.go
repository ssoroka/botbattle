package main

import (
	"../client"
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

// slightly optimized spin. Don't scan left or stop facing left.
func (self *Bot) WallSpin() {
  for i:=0; i<4; i++ {
    if !(i == 4 && self.Status().Rotation == UP) {
      self.RotLeft()
    }

    if self.Status().Rotation != LEFT {
      self.ScanAndShoot()
    }
  }
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
  return status.Y == 10; // bottom row bug.
}

func (self *Bot) GoToWall() {
  status := self.Status()
  switch status.Rotation {
    case UP: self.RotRight()
    case RIGHT: self.MoveBackward()
    case DOWN: self.RotLeft()
    case LEFT: self.MoveForward()
  }
}

func (bot *Bot) Sweep() {
  dir := -1
  for {
    // orient ourselves to the right.
    for bot.Status().Rotation != RIGHT {
      if dir == -1 {
        bot.RotRight()
      } else {
        bot.RotLeft()
      }
    }
    bot.ScanAndShoot()
    // turn around when you hit boundaries
    if dir == -1 && bot.AtTop() {
      dir = 1
    }
    if dir == 1 && bot.AtBottom() {
      dir = -1
    }

    // cover your rear
    if dir == -1 {
      bot.RotRight()
      bot.ScanAndShoot()
      bot.RotLeft()
    } else {
      if !bot.AtTop() {
        bot.RotLeft()
        bot.ScanAndShoot()
        bot.RotRight()
      }
    }

    // look ahead
    if dir == -1 {
      bot.RotLeft()
    } else {
      bot.RotRight()
    }
    bot.ScanAndShoot()
    bot.MoveForward()
  }
}

func main() {
  bot := &Bot{client.NewBotClient("192.168.1.111:3333", "wallbot2")}
  bot.Spin()
  i := 0
  for {
    i++
    if i % 2 == 0 {
      bot.Shield()
    } else {
      if !bot.OnWall() {
        bot.GoToWall()
      } else {
        bot.WallSpin()
        bot.Sweep()
      }
    }
  }
}
