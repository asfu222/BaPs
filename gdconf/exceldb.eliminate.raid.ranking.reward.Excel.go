package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEliminateRaidRankingRewardExcel() {
	g.GetExcel().EliminateRaidRankingRewardExcel = make([]*sro.EliminateRaidRankingRewardExcel, 0)
	name := "EliminateRaidRankingRewardExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().EliminateRaidRankingRewardExcel)
}

type EliminateRaidRankingRewardExcel struct {
	EliminateRaidRankingRewardExcelMap map[int64][]*sro.EliminateRaidRankingRewardExcel
}

func (g *GameConfig) gppEliminateRaidRankingRewardExcel() {
	g.GetGPP().EliminateRaidRankingRewardExcel = &EliminateRaidRankingRewardExcel{
		EliminateRaidRankingRewardExcelMap: make(map[int64][]*sro.EliminateRaidRankingRewardExcel),
	}

	for _, v := range g.GetExcel().GetEliminateRaidRankingRewardExcel() {
		if g.GetGPP().EliminateRaidRankingRewardExcel.EliminateRaidRankingRewardExcelMap[v.RankingRewardGroupId] == nil {
			g.GetGPP().EliminateRaidRankingRewardExcel.EliminateRaidRankingRewardExcelMap[v.RankingRewardGroupId] =
				make([]*sro.EliminateRaidRankingRewardExcel, 0)
		}
		g.GetGPP().EliminateRaidRankingRewardExcel.EliminateRaidRankingRewardExcelMap[v.RankingRewardGroupId] =
			append(g.GetGPP().EliminateRaidRankingRewardExcel.EliminateRaidRankingRewardExcelMap[v.RankingRewardGroupId], v)
	}

	logger.Info("处理大决战结算奖励配置表完成,奖励配置:%v个",
		len(g.GetGPP().EliminateRaidRankingRewardExcel.EliminateRaidRankingRewardExcelMap))
}

func GetEliminateRaidRankingRewardExcel(gid, ranking int64) *sro.EliminateRaidRankingRewardExcel {
	for _, conf := range GC.GetGPP().EliminateRaidRankingRewardExcel.EliminateRaidRankingRewardExcelMap[gid] {
		if conf.RankStart <= ranking && (conf.RankEnd >= ranking || conf.RankEnd == 0) {
			return conf
		}
		if ranking <= 0 && conf.RankEnd == 0 {
			return conf
		}
	}
	return nil
}
