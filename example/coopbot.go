// launch 10 of me at once!
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

func (bot *Bot) GoToWall() {
  status := bot.Status()
  switch status.Rotation {
    case UP: bot.RotRight()
    case RIGHT: bot.MoveBackward()
    case DOWN: bot.RotLeft()
    case LEFT: bot.MoveForward()
  }
}

func (bot *Bot) TakePosition() {
  status := bot.Status()
  for status.Y < status.Id % 11 {
    bot.MoveDown()
    status = bot.Status()
  }

  for status.Y > status.Id % 11 {
    bot.MoveUp()
    status = bot.Status()
  }

  bot.Spin()
  bot.Spin()
}

func (bot *Bot) MoveUp() {
  switch bot.Status().Rotation {
    case UP: bot.MoveForward()
    case RIGHT:
      bot.RotLeft()
      bot.MoveForward()
    case DOWN: bot.MoveBackward()
    case LEFT:
      bot.RotRight()
      bot.MoveForward()
  }
}

func (bot *Bot) MoveDown() {
  switch bot.Status().Rotation {
    case UP: bot.MoveBackward()
    case RIGHT:
      bot.RotRight()
      bot.MoveForward()
    case DOWN: bot.MoveForward()
    case LEFT:
      bot.RotLeft()
      bot.MoveForward()
  }
}

func (bot *Bot) TakeOver() {
  // face right
  switch bot.Status().Rotation {
    case UP: bot.RotRight()
    case DOWN: bot.RotLeft()
    case LEFT:
      bot.RotRight()
      bot.RotRight()
  }

  for {
    bot.ScanAndShoot()
  }
}

func main() {
  bot := &Bot{client.NewBotClient("192.168.1.111:3333", "coopBot")}
  i := 0
  for {
    i++
    if i % 2 == 0 {
      bot.Shield()
    } else {
      if !bot.OnWall() {
        bot.GoToWall()
      } else {
        bot.TakePosition()
        bot.TakeOver()
      }
    }
  }
}
// func NewBotClient(host, botname string) (*BotClient)
// type Status
//   X            int
//   Y            int
//   Rotation     int
//   Health       int
// type BotClient
//   ArenaHeight int
//   ArenaWidth  int
//   func Register(name string) (arena_width, arena_height int)
//   func MoveForward() (x, y int)
//   func MoveBackward() (x, y int)
//   func RotLeft() (rotation int)
//   func RotRight() (rotation int)
//   func Scan() ([]*Status)
//   func Status() (*Status)
//   func Shield() bool
