package game

import (
	"botbattle/conn"
	"time"
)

const (
	ARENA_HEIGHT = 12
	ARENA_WIDTH  = 24
)

type Scene struct {
	serv *conn.Server
	bots map[int]*Bot
}

func NewScene() *Scene {
	server := conn.NewServer()
	newScene := &Scene{
		server,
		make(map[int]*Bot),
	}
	newScene.bindActions()
	return newScene
}

func (self *Scene) bindActions() {
	self.serv.Handle("connected", self.onWebSocketConnected)
	self.serv.Handle("register", self.onRegister)
	self.serv.Handle("status", self.onStatus)
	self.serv.Handle("disconnected", self.onBotDisconnect)
	self.serv.Handle("rotate left", self.onBotRotLeft)
	self.serv.Handle("rotate right", self.onBotRotRight)
	self.serv.Handle("move forward", self.onBotMoveForward)
	self.serv.Handle("move backward", self.onBotMoveBackward)
	self.serv.Handle("fire gun", self.onFireGun)
	self.serv.Handle("fire cannon", self.onFireCannon)
	self.serv.Handle("scan", self.onScan)
	self.serv.Handle("shield", self.onShield)
}

type Status struct {
  Id        int     `json:"id"`
  X         int     `json:"x"`
  Y         int     `json:"y"`
  Rotation  int     `json:"rotation"`
  Name      string  `json:"name"`
  Health    int     `json:"health"`
}

func (self *Scene) onWebSocketConnected(client *conn.Client) {
  result := []Status{}
  for _, bot := range self.bots {
    result = append(result, Status{bot.client.Id, bot.x, bot.y, bot.rotation, bot.name, bot.health})
  }
  go func(){
	  time.Sleep(1000 * time.Millisecond)
    client.Emit("connected", result)
  }()
  return
}

func (self *Scene) onRegister(client *conn.Client, name string) (int, int) {
	newBot := NewBot(self, client, name)
	self.bots[client.Id] = newBot
	self.serv.Broadcast("register", newBot.client.Id, newBot.x, newBot.y, newBot.rotation, name)
	return ARENA_WIDTH, ARENA_HEIGHT
}

func (self *Scene) onStatus(client *conn.Client) (int, int, int, int, int) {
	if bot := self.bots[client.Id]; bot != nil {
		return bot.Status()
	}
	return 0, 0, 0, 0, 0
}

func (self *Scene) onBotDisconnect(client *conn.Client) {
	if bot := self.bots[client.Id]; bot != nil {
		self.KillBot(bot)
	}
}

func (self *Scene) KillBot(bot *Bot) {
  bot.client.Close()
	delete(self.bots, bot.client.Id)
}

func (self *Scene) onBotRotLeft(client *conn.Client) int {
	if bot := self.bots[client.Id]; bot != nil {
    rot := bot.RotLeft()
		self.serv.Broadcast("rotate", bot.client.Id, rot)
		return rot
	}
	return 0
}

func (self *Scene) onBotRotRight(client *conn.Client) int {
	if bot := self.bots[client.Id]; bot != nil {
    rot := bot.RotRight()
		self.serv.Broadcast("rotate", bot.client.Id, rot)
		return rot
	}
	return 0
}

func (self *Scene) onBotMoveForward(client *conn.Client) (int, int) {
	if bot := self.bots[client.Id]; bot != nil {
    x, y := bot.MoveForward()
		self.serv.Broadcast("move", bot.client.Id, x, y)
		return x, y
	}
	return 0, 0
}

func (self *Scene) onBotMoveBackward(client *conn.Client) (int, int) {
	if bot := self.bots[client.Id]; bot != nil {
    x, y := bot.MoveBackward()
		self.serv.Broadcast("move", bot.client.Id, x, y)
		return x, y
	}
	return 0, 0
}

func (self *Scene) onFireGun(client *conn.Client) bool {
	if bot := self.bots[client.Id]; bot != nil {
		self.serv.Broadcast("fire gun", bot.client.Id)
		return bot.FireGun()
	}
	return false
}

func (self *Scene) onFireCannon(client *conn.Client) bool {
	if bot := self.bots[client.Id]; bot != nil {
		self.serv.Broadcast("fire cannon", bot.client.Id)
		return bot.FireCannon()
	}
	return false
}

func (self *Scene) onScan(client *conn.Client) [][]int {
	if bot := self.bots[client.Id]; bot != nil {
		self.serv.Broadcast("scan", bot.client.Id)
		return bot.Scan()
	}
	return [][]int{}
}

func (self *Scene) onShield(client *conn.Client) bool {
	if bot := self.bots[client.Id]; bot != nil {
		return bot.Shield()
	}
	return false
}

func (self *Scene) Start() {
	go self.serv.Listen(map[string]string{
		"host":    "localhost:3333",
		"pattern": "/ws",
	})
}
