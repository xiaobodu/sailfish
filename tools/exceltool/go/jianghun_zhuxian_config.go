package config
type Jianghun_zhuxian struct{
	Items []*Jianghun_zhuxian_Item
}

type Jianghun_zhuxian_Item struct{
	Id int
	Get_way map[string]int
}

var InstJianghun_zhuxian *Jianghun_zhuxian

func init(){
	item := &Jianghun_zhuxian_Item{}
	item.Id = 310001
	item.Get_way = map[string]int{"1":6015,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310002
	item.Get_way = map[string]int{"1":3007,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310003
	item.Get_way = map[string]int{"1":11018,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310004
	item.Get_way = map[string]int{"1":3018,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310005
	item.Get_way = map[string]int{"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310006
	item.Get_way = map[string]int{"1":9013,"2":1,"4":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310007
	item.Get_way = map[string]int{"2":1,"3":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310008
	item.Get_way = map[string]int{"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310009
	item.Get_way = map[string]int{"1":8005,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310012
	item.Get_way = map[string]int{"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310013
	item.Get_way = map[string]int{"1":11007,"2":1,"4":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310016
	item.Get_way = map[string]int{"1":5019,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310017
	item.Get_way = map[string]int{"1":5014,"2":1,"3":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310018
	item.Get_way = map[string]int{"1":10025,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310019
	item.Get_way = map[string]int{"1":9007,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310020
	item.Get_way = map[string]int{"1":1011,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310021
	item.Get_way = map[string]int{"1":6007,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310022
	item.Get_way = map[string]int{"1":9024,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310023
	item.Get_way = map[string]int{"1":4012,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310024
	item.Get_way = map[string]int{"1":7006,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310025
	item.Get_way = map[string]int{"1":4017,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310026
	item.Get_way = map[string]int{"1":11026,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310027
	item.Get_way = map[string]int{"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310036
	item.Get_way = map[string]int{"1":2006,"2":1,"3":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310037
	item.Get_way = map[string]int{"1":8018,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310038
	item.Get_way = map[string]int{"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310039
	item.Get_way = map[string]int{"1":7018,"2":1,"4":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310040
	item.Get_way = map[string]int{"1":6017,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310042
	item.Get_way = map[string]int{"1":4007,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310043
	item.Get_way = map[string]int{"1":10016,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310044
	item.Get_way = map[string]int{"1":5007,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
	item = &Jianghun_zhuxian_Item{}
	item.Id = 310049
	item.Get_way = map[string]int{"1":2016,"2":1}
	InstJianghun_zhuxian.Items = append(InstJianghun_zhuxian.Items,item)
}