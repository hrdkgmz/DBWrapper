package cache

import (
	"errors"
	"strconv"
)

func GetFabUserBySysUser(sysUser string) (string, error) {
	fabUser, err := GetInstance().GetHashString("b_user_info:sys_user:"+sysUser, "fab_user")
	return fabUser, err
}

func GetOrgByFabUser(fabUser string) (string, error) {
	org, err := GetInstance().GetHashString("b_user_info:fab_user:"+fabUser, "org_name")
	return org, err
}

func GetOrgBySysUser(sysUser string) (string, error) {
	org, err := GetInstance().GetHashString("b_user_info:sys_user:"+sysUser, "org_name")
	return org, err
}

func GetPeerByOrg(org string) ([]string, error) {
	return getSetStringVals("org_peers:" + org)
}

func GetPeerByChannel(channel string) ([]string, error) {
	return getSetStringVals("chan_peers:" + channel)
}

func GetCommonPeerByOrgChannel(org string, channel string) ([]string, error) {
	peer, err := GetInstance().GetSetInters("chan_peers:"+channel, "org_peers:"+org)
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

func GetPeersByCC(ccid string) ([]string, error) {
	return getSetStringVals("cc_peers:" + ccid)
}
func GetCCInfoByCCID(ccid string) (string, string, error) {
	name, err := GetInstance().GetHashString("b_cc_info:cc_id:"+ccid, "cc_name")
	if err != nil {
		return name, "", err
	}
	ver, err := GetInstance().GetHashString("b_cc_info:cc_id:"+ccid, "cc_ver")
	return name, ver, nil
}
func GetOrderNames() ([]string, error) {
	return getSetStringVals("orderers")
}

func GetChannelTxByChannel(channel string) (string, error) {
	channelTxPath, err := GetInstance().GetHashString("b_channel_info:chan_name:"+channel, "channeltx_path")
	return channelTxPath, err
}

func GetCCPathByCC(ccID string) (string, error) {
	ccPath, err := GetInstance().GetHashString("b_cc_info:cc_id:"+ccID, "cc_path")
	return ccPath, err
}

func SetTaskDataHash(taskID int64, val map[string]interface{}, timeoutSecond int) error {
	_, err := GetInstance().SetHashAndExpire(strconv.FormatInt(taskID, 10), val, timeoutSecond)
	if err != nil {
		return err
	}
	return nil
}

func GetTaskDataHashAll(taskID int64) (vals map[string]string, err error) {
	return GetInstance().GetHashMapString(strconv.FormatInt(taskID, 10))
}

func GetTaskDataHashSingle(taskID int64, field string) (val string, err error) {
	return GetInstance().GetHashString(strconv.FormatInt(taskID, 10), field)
}

func SetTaskDataString(taskID int64, val interface{}, timeoutSecond int) error {
	if _, err := GetInstance().SetStringAndExpire(strconv.FormatInt(taskID, 10), val, timeoutSecond); err != nil {
		return err
	}
	return nil
}

func GetTaskDataString(taskID int64) (val string, err error) {
	return GetInstance().GetString(strconv.FormatInt(taskID, 10))
}

func IsTaskDataTimeout(taskID int64) (bool, error) {
	ttl, err := GetInstance().GetTTL(strconv.FormatInt(taskID, 10))
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
	if _, err := GetInstance().SetExpire(strconv.FormatInt(taskID, 10), timeoutSecond); err != nil {
		return err
	}
	return nil
}

func getSetStringVals(key string) ([]string, error) {
	vals, err := GetInstance().GetSetMembers(key)
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
