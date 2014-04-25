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

type Position struct {
  X, Y int
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
  positions := make(map[int]Position, 20)
  positions[0] = Position{X: 0, Y: 0}
  positions[1] = Position{X: 0, Y: 1}
  positions[2] = Position{X: 0, Y: 2}
  positions[3] = Position{X: 0, Y: 3}
  positions[4] = Position{X: 0, Y: 4}
  positions[5] = Position{X: 0, Y: 5}
  positions[6] = Position{X: 0, Y: 6}
  positions[7] = Position{X: 0, Y: 7}
  positions[8] = Position{X: 0, Y: 8}
  positions[9] = Position{X: 0, Y: 9}

  positions[10] = Position{X: 1, Y: 10}
  positions[11] = Position{X: 2, Y: 10}
  positions[12] = Position{X: 3, Y: 10}
  positions[13] = Position{X: 4, Y: 10}
  positions[14] = Position{X: 5, Y: 10}
  positions[15] = Position{X: 6, Y: 10}
  positions[16] = Position{X: 7, Y: 10}
  positions[17] = Position{X: 8, Y: 10}
  positions[18] = Position{X: 9, Y: 10}
  positions[19] = Position{X: 10, Y: 10}

  positions[20] = Position{X: 11, Y: 10}
  positions[21] = Position{X: 12, Y: 10}
  positions[22] = Position{X: 13, Y: 10}
  positions[23] = Position{X: 14, Y: 10}
  positions[24] = Position{X: 15, Y: 10}
  positions[25] = Position{X: 16, Y: 10}
  positions[26] = Position{X: 17, Y: 10}
  positions[27] = Position{X: 18, Y: 10}
  positions[28] = Position{X: 19, Y: 10}
  positions[29] = Position{X: 20, Y: 10}

  status := bot.Status()
  bot.GoTo(positions[status.Id % 30])

  bot.Spin()
  bot.Spin()
}

func (bot *Bot) GoTo(pos Position) {
  status := bot.Status()
  for status.Y < pos.Y {
    bot.MoveDown()
    status = bot.Status()
  }

  for status.Y > pos.Y {
    bot.MoveUp()
    status = bot.Status()
  }

  for status.X < pos.X {
    bot.MoveRight()
    status = bot.Status()
  }

  for status.X > pos.X {
    bot.MoveLeft()
    status = bot.Status()
  }

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

func (bot *Bot) MoveLeft() {
  switch bot.Status().Rotation {
    case UP:
      bot.RotRight()
      bot.MoveBackward()
    case RIGHT:
      bot.MoveBackward()
    case DOWN:
      bot.RotLeft()
      bot.MoveBackward()
    case LEFT:
      bot.MoveForward()
  }
}

func (bot *Bot) MoveRight() {
  switch bot.Status().Rotation {
    case UP:
      bot.RotRight()
      bot.MoveForward()
    case RIGHT:
      bot.MoveForward()
    case DOWN:
      bot.RotLeft()
      bot.MoveForward()
    case LEFT:
      bot.MoveBackward()
  }
}

func (bot *Bot) TakeOver() {
  // face right
  status := bot.Status()
  if status.X == 0 {
    switch status.Rotation {
      case UP: bot.RotRight()
      case DOWN: bot.RotLeft()
      case LEFT:
        bot.RotRight()
        bot.RotRight()
    }

    for {
      bot.ScanAndShoot()
    }
  } else {
    switch status.Rotation {
      case RIGHT: bot.RotLeft()
      case DOWN:
        bot.RotLeft()
        bot.RotLeft()
      case LEFT:
        bot.RotRight()
    }

    for {
      bot.ScanAndShoot()
    }
  }
}

func main() {
  bot := &Bot{client.NewBotClient("192.168.1.111:3333", "coopBot2")}
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
