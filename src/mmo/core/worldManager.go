package core

import (
	"fmt"
	"mmm/mmo/conf"
	"strconv"
	"sync"
)

var Worldmanager = NewWorldManager()

type WorldManager struct {
	AOI         *AOIManager
	Players     map[int]*Player
	Wmlock      sync.RWMutex
	PlayerCount int
}

func (w *WorldManager) AddPlayer(p *Player) {
	w.AOI.AddPlayerByXY(p.Pid, int(p.X), int(p.Z))
	w.Wmlock.Lock()
	defer w.Wmlock.Unlock()
	w.Players[p.Pid] = p
	w.PlayerCount++
}

func (w *WorldManager) RemovePlayer(p *Player) {
	w.AOI.RemovePlaysByXY(p.Pid, int(p.X), int(p.Z))
	w.Wmlock.Lock()
	fmt.Println("remove前的wm players", Worldmanager.Players, "长度", len(Worldmanager.Players))
	delete(w.Players, p.Pid)
	fmt.Println("remove后的wm players", Worldmanager.Players, "长度", len(Worldmanager.Players))
	w.PlayerCount--
	w.Wmlock.Unlock()

}

func (w *WorldManager) RemovePlayerByPid(p int) {
	player := w.Players[p]
	w.RemovePlayer(player)
}

func (w *WorldManager) GetPlayerByPid(pid int) *Player {
	w.Wmlock.RLock()
	defer w.Wmlock.RUnlock()
	return w.Players[pid]
}

func (w *WorldManager) GetAllPlayer() []*Player {
	ans := make([]*Player, 0)
	w.Wmlock.RLock()
	defer w.Wmlock.RUnlock()
	for _, v := range w.Players {
		fmt.Println("添加玩家号; 总共有"+strconv.Itoa(w.PlayerCount)+"wmplayers:", w.Players)
		ans = append(ans, v)
	}
	return ans
}

func NewWorldManager() *WorldManager {
	aoi := NewAOIManager(conf.MMOConfig.MinX, conf.MMOConfig.MaxX, conf.MMOConfig.CountX, conf.MMOConfig.MinY, conf.MMOConfig.MaxY, conf.MMOConfig.CountY)
	return &WorldManager{
		AOI:         aoi,
		Players:     make(map[int]*Player, 0),
		PlayerCount: 0,
	}
}
