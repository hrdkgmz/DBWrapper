package cache

import (
	"errors"
	"strconv"
)

const (
	archiveDb int = 0
	taskDb    int = 1
)

func GetFabUserBySysUser(sysUser string) (string, error) {
	fabUser, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_user_info:hash:sys_user:"+sysUser, "fab_user")
	return fabUser, err
}

func GetOrgByFabUser(fabUser string) (string, error) {
	org, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_user_info:hash:fab_user:"+fabUser, "org_name")
	return org, err
}

func GetOrgBySysUser(sysUser string) (string, error) {
	org, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_user_info:hash:sys_user:"+sysUser, "org_name")
	return org, err
}

func GetPeerByOrg(org string) ([]string, error) {
	return getSetStringVals("org_peers:" + org)
}

func GetPeerByChannel(channel string) ([]string, error) {
	return getSetStringVals("chan_peers:" + channel)
}

func GetCommonPeerByOrgChannel(org string, channel string) ([]string, error) {
	peer, err := GetInstance().GetSetInterWitchDb(archiveDb, "chan_peers:"+channel, "org_peers:"+org)
	strVals := make([]string, 0)
	for _, p := range peer {
		if str, ok := p.([]uint8); ok {
			strVals = append(strVals, string(str))
		}
	}
	if len(peer) != len(strVals) {
		return nil, errors.New("insufficient number of values")
	}
	return strVals, err
}

func GetPeersByCCID(ccid string) ([]string, error) {
	return getSetStringVals("cc_peers:" + ccid)
}

func GetChansByCCID(ccid string) ([]string, error) {
	return getSetStringVals("cc_chans:" + ccid)
}

func GetCCInfoByCCID(ccid string) (string, string, error) {
	name, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_cc_info:hash:cc_id:"+ccid, "cc_name")
	if err != nil {
		return name, "", err
	}
	ver, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_cc_info:hash:cc_id:"+ccid, "cc_ver")
	return name, ver, nil
}
func GetOrderNames() ([]string, error) {
	return getSetStringVals("orderers")
}

func GetChannelTxByChannel(channel string) (string, error) {
	channelTxPath, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_channel_info:hash:chan_name:"+channel, "channeltx_path")
	return channelTxPath, err
}

func GetCCPathByCCID(ccID string) (string, error) {
	ccPath, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_cc_info:hash:cc_id:"+ccID, "cc_path")
	return ccPath, err
}

func GetChanDetailByName(name string) (map[string]string, error) {
	channel, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_channel_info:hash:chan_name:"+name)
	return channel, err
}

func GetAllChanName() ([]string, error) {
	return getSetStringVals("channels")
}

func GetAllChanDetail() ([]map[string]string, error) {
	chanNames, err := GetAllChanName()
	if err != nil {
		return nil, err
	}
	chanDetails := make([]map[string]string, 0)
	for _, v := range chanNames {
		chanDetail, err := GetChanDetailByName(v)
		if err != nil {
			return nil, err
		}
		chanDetails = append(chanDetails, chanDetail)
	}
	return chanDetails, nil
}

func GetCCDetailById(id string) (map[string]string, error) {
	cc, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_cc_info:hash:cc_id:"+id)
	return cc, err
}

func GetAllCCId() ([]string, error) {
	return getSetStringVals("ccs")
}

func GetAllCCDetail() ([]map[string]string, error) {
	ccIds, err := GetAllCCId()
	if err != nil {
		return nil, err
	}
	ccDetails := make([]map[string]string, 0)
	for _, v := range ccIds {
		ccDetail, err := GetCCDetailById(v)
		if err != nil {
			return nil, err
		}
		ccDetails = append(ccDetails, ccDetail)
	}
	return ccDetails, nil
}

func GetPeerDetailByName(name string) (map[string]string, error) {
	peer, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_peer_info:hash:peer_fullname:"+name)
	return peer, err
}

func GetChanPeerDetail(channel string, peer string) (map[string]string, error) {
	chanPeer, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "s_channel_peer:hash:chan_name&peer_fullname:"+channel+"&"+peer)
	return chanPeer, err
}

func GetPeerJoinTime(channel string, peer string) (string, error) {
	time, err := GetInstance().GetHashStringWitchDb(archiveDb, "s_channel_peer:hash:chan_name&peer_fullname:"+channel+"&"+peer, "join_time")
	return time, err
}

func GetChanCCDetail(channel string, ccid string) (map[string]string, error) {
	chanCc, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "s_channel_cc:hash:chan_name&cc_id:"+channel+"&"+ccid)
	return chanCc, err
}

func GetPeerCCDetail(peer string, ccid string) (map[string]string, error) {
	chanCc, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "s_peer_cc:hash:peer_fullname&cc_id:"+peer+"&"+ccid)
	return chanCc, err
}

func SetTaskDataHash(taskID int64, val map[string]interface{}, timeoutSecond int) error {
	_, err := GetInstance().SetHashAndExpireWitchDb(taskDb, strconv.FormatInt(taskID, 10), val, timeoutSecond)
	if err != nil {
		return err
	}
	return nil
}

func GetTaskDataHashAll(taskID int64) (vals map[string]string, err error) {
	return GetInstance().GetHashMapStringWitchDb(taskDb, strconv.FormatInt(taskID, 10))
}

func GetTaskDataHashSingle(taskID int64, field string) (val string, err error) {
	return GetInstance().GetHashStringWitchDb(taskDb, strconv.FormatInt(taskID, 10), field)
}

func SetTaskDataString(taskID int64, val interface{}, timeoutSecond int) error {
	if _, err := GetInstance().SetStringAndExpireWitchDb(taskDb, strconv.FormatInt(taskID, 10), val, timeoutSecond); err != nil {
		return err
	}
	return nil
}

func GetTaskDataString(taskID int64) (val string, err error) {
	return GetInstance().GetStringWitchDb(taskDb, strconv.FormatInt(taskID, 10))
}

func SetTaskDataList(taskID int64, val []string, timeoutSecond int) error {
	_, err := GetInstance().SetListToTailAndExpireWitchDb(taskDb, strconv.FormatInt(taskID, 10), val, timeoutSecond)
	if err != nil {
		return err
	}
	return nil
}

func GetTaskDataListByRange(taskID int64, start int, stop int) (val []string, err error) {
	return GetInstance().GetListByRangeWitchDb(taskDb, strconv.FormatInt(taskID, 10), start, stop)
}

func GetTaskDataListLen(taskID int64) (len int, err error) {
	return GetInstance().GetListLenWitchDb(taskDb, strconv.FormatInt(taskID, 10))
}

func IsTaskDataTimeout(taskID int64) (bool, error) {
	ttl, err := GetInstance().GetTTLWitchDb(taskDb, strconv.FormatInt(taskID, 10))
	if err != nil {
		return false, err
	}
	if ttl >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func RenewTaskDataTimout(taskID int64, timeoutSecond int) error {
	if _, err := GetInstance().SetExpireWitchDb(taskDb, strconv.FormatInt(taskID, 10), timeoutSecond); err != nil {
		return err
	}
	return nil
}

func getSetStringVals(key string) ([]string, error) {
	vals, err := GetInstance().GetSetMembersWitchDb(archiveDb, key)
	strVals := make([]string, 0)
	for _, o := range vals {
		if str, ok := o.([]uint8); ok {
			strVals = append(strVals, string(str))
		}
	}
	if len(vals) != len(strVals) {
		return nil, errors.New("insufficient number of values")
	}
	return strVals, err
}
