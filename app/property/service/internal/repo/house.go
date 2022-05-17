package repo

type ListUnitsReq struct {
	VillageId int
}

type ListUnitsResp struct {
	Data []Building
}


type Building struct {
	Units []Unit
}

type Unit struct {
	houses []House
}

type House struct {
	Floor      int //楼层
	Kind       int //1住宅2 办公室3 宿舍
	FloorAge   int //建筑面积
	CarpetArea int //室内面积
	State      int //房屋状态 1未交房 2已交房 3装修中 4 已装修 5已交房 6 已入住
	FineDecoration int  //是否精装修
	Remark string //备注信息
	HouseType int //房屋户型
	Number int //房屋编号
	UnitId int //单元
	RateNum float32 //计费比率
}
