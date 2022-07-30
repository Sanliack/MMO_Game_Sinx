package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	Gid     int
	MinX    int
	MaxX    int
	MinY    int
	MaxY    int
	players map[int]int
	PidLock sync.RWMutex
}

func (g *Grid) AddPlayer(pid int) {
	g.PidLock.Lock()
	defer g.PidLock.Unlock()
	if status := g.players[pid]; status != 0 {
		fmt.Printf("player ID:%d 已在GID:=%d的格子内\n", pid, g.Gid)
		return
	}
	g.players[pid] = 1
}

func (g *Grid) RemovePlayer(pid int) {
	g.PidLock.Lock()
	defer g.PidLock.Unlock()
	if _, ok := g.players[pid]; !ok {
		fmt.Printf("player ID:%d 不在GID:=%d的格子内，无需remove\n", pid, g.Gid)
		return
	}
	delete(g.players, pid)
}

func (g *Grid) GetAllPlayers() []int {
	ans := make([]int, 0)
	//fmt.Println("GetAllPLay:g.playermap:", g.players)
	for k, _ := range g.players {
		ans = append(ans, k)
	}
	return ans
}

func (g *Grid) String() string {
	ans := fmt.Sprintf("GIRD信息: Gid:=%d,minx:=%d,maxx:=%d,miny:=%d,maxy:=%d\n", g.Gid, g.MinX, g.MaxX, g.MinY, g.MaxY)
	for _, i := range g.GetAllPlayers() {
		ans += fmt.Sprintf("player: Id:%d\n", i)
	}
	return ans
}

func NewGrid(gid, minx, maxx, miny, maxy int) *Grid {
	return &Grid{
		Gid:     gid,
		MinX:    minx,
		MaxX:    maxx,
		MinY:    miny,
		MaxY:    maxy,
		players: make(map[int]int, 10),
	}
}
