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

func (bot *Bot) ScanAndShoot() {
  for len(bot.Scan()) > 0 {
    bot.Shield()
    bot.FireGun()
    bot.FireGun()
    bot.FireCannon()
  }
}

func (bot *Bot) Spin() {
  for i:=0; i<4; i++ {
    bot.RotLeft()
    bot.ScanAndShoot()
  }
}

// slightly optimized spin. Don't scan left or stop facing left.
func (bot *Bot) WallSpin() {
  for i:=0; i<4; i++ {
    if !(i == 4 && bot.Status().Rotation == UP) {
      bot.RotLeft()
    }

    if bot.Status().Rotation != LEFT {
      bot.ScanAndShoot()
    }
  }
}

func (bot *Bot) OnWall() bool {
  status := bot.Status()
  return status.X == 0;
}

func (bot *Bot) AtTop() bool {
  status := bot.Status()
  return status.Y == 0;
}

func (bot *Bot) AtBottom() bool {
  status := bot.Status()
  return status.Y == 10; // bottom row bug.
}

func (bot *Bot) DefendCorner() {
  for i := 0; i < 30; i++ {
    bot.RotRight()
    bot.ScanAndShoot()
    bot.RotLeft()
    bot.ScanAndShoot()
  }
}

func (bot *Bot) GoToWall() {
  for !bot.OnWall() {
    status := bot.Status()
    switch status.Rotation {
      case UP: bot.RotRight()
      case RIGHT: bot.MoveBackward()
      case DOWN: bot.RotLeft()
      case LEFT: bot.MoveForward()
    }
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
      bot.DefendCorner()
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
  bot := &Bot{client.NewBotClient("192.168.1.111:3333", "cornerbot")}
  bot.Spin()
  bot.Shield()
  bot.GoToWall()
  bot.WallSpin()
  bot.Sweep()
}
