// Code generated by gucooing DO NOT EDIT.

package proto

type HexaUnit struct {
	EntityId int64
	Id       int64
	Rotate   *Vector3
	Location *Vector3

	BuffGroupIds  []string
	SkillCardHand *SkillCardHand
	PlayAnimation bool
	RewardItems   map[TacticEntityType][]int64
}
