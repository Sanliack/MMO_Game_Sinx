package core

import "fmt"

type AOIManager struct {
	MinX    int
	MaxX    int
	CountX  int
	MinY    int
	MaxY    int
	CountY  int
	GridMap map[int]*Grid
}

func (a *AOIManager) String() string {
	ans := fmt.Sprintf("AOImanager信息: MinX：%d,MaxX:%d,countX:%d---MinY:%d,MaxY:%d,countY:%d\n", a.MinX, a.MaxX, a.CountX, a.MinY, a.MaxY, a.CountY)
	ans += fmt.Sprintf("aoimanager存在以下grid：\n")

	for k, v := range a.GridMap {
		ans += fmt.Sprintf("Grid:%d == minx:%d maxx:%d miny:%d maxy:%d\n", k, v.MinX, v.MaxX, v.MinY, v.MaxY)
	}
	return ans
}

func (a *AOIManager) AddPlayerByGid(pid, gid int) {
	a.GetGridById(gid).AddPlayer(pid)
}

func (a *AOIManager) RemovePlayerFromGridByID(pid, gid int) {
	a.GetGridById(gid).RemovePlayer(pid)
}

func (a *AOIManager) GetAllPlayersIDInOneGrid(gid int) []int {
	return a.GetGridById(gid).GetAllPlayers()
}

func (a *AOIManager) AddPlayerByXY(pid, x, z int) {
	gid := a.GetGidByXY(x, z)
	a.GetGridById(gid).AddPlayer(pid)
}

func (a *AOIManager) RemovePlaysByXY(pid, x, z int) {
	gid := a.GetGidByXY(x, z)
	a.GetGridById(gid).RemovePlayer(pid)
}

func (a *AOIManager) GetGridById(gid int) *Grid {
	return a.GridMap[gid]
}

func (a *AOIManager) GetXLength() int {
	return (a.MaxX - a.MinX) / a.CountX
}

func (a *AOIManager) GetYLength() int {
	return (a.MaxY - a.MinY) / a.CountY
}

func (a *AOIManager) GetSurroundGrid(gid int) []*Grid {
	gridrowids := a.getrowIds(gid)
	//fmt.Println(gridrowids)
	allSurroundid := a.getcolumnIds(gridrowids)
	//fmt.Println(allSurroundid)
	surroundGrid := make([]*Grid, 0)
	for _, v := range allSurroundid {
		onegrid := a.GetGridById(v)
		surroundGrid = append(surroundGrid, onegrid)
	}
	return surroundGrid
}

func (a *AOIManager) GetGidByXY(x, y int) int {
	xno := x / a.GetXLength()
	yno := y / a.GetYLength()
	return xno + yno*a.CountX
}

func (a *AOIManager) getrowIds(gid int) []int {
	ans := []int{gid}
	if a.CountX == 1 {
		return ans
	}
	if gid%a.CountX == 0 {
		ans = append(ans, gid+1)
	} else if (gid+1)%a.CountX == 0 {
		ans = append(ans, gid-1)
	} else {
		ans = append(ans, gid+1)
		ans = append(ans, gid-1)
	}
	return ans
}

func (a *AOIManager) getcolumnIds(gids []int) []int {
	newans := []int{}
	for _, v := range gids {
		upnum := v - a.CountX
		downnum := v + a.CountX
		if upnum >= 0 {
			newans = append(newans, upnum)
		}
		if downnum <= (a.CountY*a.CountX)-1 {
			newans = append(newans, downnum)
		}
	}
	gids = append(gids, newans...)
	return gids
}

func (a *AOIManager) GetSurroundGridPlayersIDsByXY(x int, y int) []int {
	gid := a.GetGidByXY(x, y)
	surroundgridlist := a.GetSurroundGrid(gid)
	playersID := make([]int, 0)
	for _, v := range surroundgridlist {
		playersID = append(playersID, v.GetAllPlayers()...)
	}
	return playersID
}

func NewAOIManager(minx, maxx, countx, miny, maxy, county int) *AOIManager {
	gm := make(map[int]*Grid, county*countx)
	var x, y int

	mingridx := (maxx - minx) / countx
	mingridy := (maxy - miny) / county
	for x = 0; x < countx; x++ {
		for y = 0; y < county; y++ {
			newgrid := NewGrid(x+countx*y, x*mingridx, (x+1)*mingridx, y*mingridy, (y+1)*mingridy)
			gm[x+countx*y] = newgrid
		}
	}
	return &AOIManager{
		MinX:    minx,
		MaxX:    maxx,
		CountX:  countx,
		MinY:    miny,
		MaxY:    maxy,
		CountY:  county,
		GridMap: gm,
	}
}
