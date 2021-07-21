package dao

/**
Member

Type 1: 群员 2: 管理 3: 群主
State 状态位 0000 : 0-0-通知开关-被禁言
*/
type Member struct {
	Uid      int64
	Nickname string
	Avatar   string
	Type     uint8
	State    uint8
}

var GroupDao = new(groupDao)

type groupDao struct{}

func (d *groupDao) GetGroup(gid uint64) (string, []int64) {

	return "", []int64{}
}

func (d *groupDao) RemoveMember(gid uint64, uint642 int64) error {

	return nil
}

func (d *groupDao) AddMember(gid uint64, uint642 int64) error {

	return nil
}

func (d *groupDao) GetUserGroup(uid int64) []uint64 {

	return []uint64{}
}