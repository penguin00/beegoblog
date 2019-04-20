package help

import "time"

func TimeNow() time.Time {
	//var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	var cstSh = time.FixedZone("CST", 8*3600)
	return time.Now().In(cstSh)
}
