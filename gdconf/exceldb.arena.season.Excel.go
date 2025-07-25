package gdconf

import (
	"time"

	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadArenaSeasonExcel() {
	g.GetExcel().ArenaSeasonExcel = make([]*sro.ArenaSeasonExcel, 0)
	name := "ArenaSeasonExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().ArenaSeasonExcel)
}

type ArenaSeasonExcel struct {
	ArenaSeasonExcelMap map[int64]*sro.ArenaSeasonExcel
	CurArenaSeason      *sro.ArenaSeasonExcel
}

func (g *GameConfig) gppArenaSeasonExcel() {
	g.GetGPP().ArenaSeasonExcel = &ArenaSeasonExcel{
		ArenaSeasonExcelMap: make(map[int64]*sro.ArenaSeasonExcel),
	}

	for _, v := range g.GetExcel().GetArenaSeasonExcel() {
		g.GetGPP().ArenaSeasonExcel.ArenaSeasonExcelMap[v.UniqueId] = v
	}

	logger.Info("处理竞技场赛季排期配置表完成,大决战总分奖励配置:%v个",
		len(g.GetGPP().ArenaSeasonExcel.ArenaSeasonExcelMap))
}

func genCurArenaSeason() *sro.ArenaSeasonExcel {
	for _, v := range GC.GetExcel().GetArenaSeasonExcel() { // 读取原始文件,保证顺序
		startTime, err := time.Parse("2006-01-02 15:04:05", v.SeasonStartDate)
		endTime, err := time.Parse("2006-01-02 15:04:05", v.SeasonEndDate)
		if err != nil {
			logger.Error("竞技场排期表时间格式错误")
			return nil
		}
		if time.Now().After(startTime) && time.Now().Before(endTime) { // 上期结束且下期未开启
			return v
		}
	}
	logger.Warn("找不到当前竞技场排期")
	return nil
}

func GetCurArenaSeason() *sro.ArenaSeasonExcel {
	info := GC.GetGPP().ArenaSeasonExcel
	if info.CurArenaSeason == nil {
		info.CurArenaSeason = genCurArenaSeason()
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", info.CurArenaSeason.SeasonEndDate)
	if err != nil && time.Now().After(endTime) {
		info.CurArenaSeason = genCurArenaSeason()
	}

	return info.CurArenaSeason
}

func GetArenaSeasonExcel(id int64) *sro.ArenaSeasonExcel {
	return GC.GetGPP().ArenaSeasonExcel.ArenaSeasonExcelMap[id]
}
