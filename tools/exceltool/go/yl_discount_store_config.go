package config
type Yl_discount_store struct{
	Items []*Yl_discount_store_Item
}

type Yl_discount_store_Item struct{
	Id int
	Item_name string
	Day int
	Prize_type int
	Item_id int
	Amount int
	Currency_type int
	Price_sell int
	Buy_limit int
	Vip_lv int
	Level int
	Lattice_position int
	Onsale int
}

var InstYl_discount_store *Yl_discount_store

func init(){
	item := &Yl_discount_store_Item{}
	item.Id = 1
	item.Item_name = "粮食"
	item.Day = 1
	item.Prize_type = 3
	item.Item_id = 3
	item.Amount = 16000
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 0
	item.Level = 30
	item.Lattice_position = 1
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 2
	item.Item_name = "烈酒"
	item.Day = 1
	item.Prize_type = 201
	item.Item_id = 431001
	item.Amount = 1
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 0
	item.Level = 30
	item.Lattice_position = 2
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 3
	item.Item_name = "止血绷带"
	item.Day = 1
	item.Prize_type = 201
	item.Item_id = 110201
	item.Amount = 100
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 0
	item.Level = 30
	item.Lattice_position = 3
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 4
	item.Item_name = "便携扳手"
	item.Day = 1
	item.Prize_type = 201
	item.Item_id = 110202
	item.Amount = 100
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 0
	item.Level = 30
	item.Lattice_position = 4
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 5
	item.Item_name = "喀秋莎自行火炮营新兵"
	item.Day = 1
	item.Prize_type = 201
	item.Item_id = 110301
	item.Amount = 100
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 2
	item.Level = 30
	item.Lattice_position = 5
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 6
	item.Item_name = "重型喀秋莎火炮营新兵"
	item.Day = 1
	item.Prize_type = 201
	item.Item_id = 110302
	item.Amount = 100
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 3
	item.Level = 30
	item.Lattice_position = 6
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 7
	item.Item_name = "虎式E型装甲营新兵"
	item.Day = 1
	item.Prize_type = 201
	item.Item_id = 310007
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 4
	item.Level = 30
	item.Lattice_position = 7
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 8
	item.Item_name = "虎式H型装甲营新兵"
	item.Day = 1
	item.Prize_type = 201
	item.Item_id = 310008
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 5
	item.Level = 30
	item.Lattice_position = 8
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 9
	item.Item_name = "勃朗宁A6机枪营新兵"
	item.Day = 2
	item.Prize_type = 201
	item.Item_id = 310009
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 0
	item.Level = 35
	item.Lattice_position = 1
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 10
	item.Item_name = "88MM高炮营新兵"
	item.Day = 2
	item.Prize_type = 201
	item.Item_id = 310012
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 0
	item.Level = 35
	item.Lattice_position = 2
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 11
	item.Item_name = "卡尔重型火炮营新兵"
	item.Day = 2
	item.Prize_type = 201
	item.Item_id = 310013
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 0
	item.Level = 35
	item.Lattice_position = 3
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 12
	item.Item_name = "203MM榴弹炮兵营新兵"
	item.Day = 2
	item.Prize_type = 201
	item.Item_id = 310016
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 0
	item.Level = 35
	item.Lattice_position = 4
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 13
	item.Item_name = "RPG-1战斗营新兵"
	item.Day = 2
	item.Prize_type = 201
	item.Item_id = 310017
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 2
	item.Level = 35
	item.Lattice_position = 5
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 14
	item.Item_name = "铁拳100型战斗组新兵"
	item.Day = 2
	item.Prize_type = 201
	item.Item_id = 310018
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 3
	item.Level = 35
	item.Lattice_position = 6
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 15
	item.Item_name = "41型波波沙冲锋队新兵"
	item.Day = 2
	item.Prize_type = 201
	item.Item_id = 310019
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 4
	item.Level = 35
	item.Lattice_position = 7
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
	item = &Yl_discount_store_Item{}
	item.Id = 16
	item.Item_name = "M1型汤姆逊冲锋营新兵"
	item.Day = 2
	item.Prize_type = 201
	item.Item_id = 310020
	item.Amount = 5
	item.Currency_type = 1
	item.Price_sell = 20
	item.Buy_limit = 3
	item.Vip_lv = 5
	item.Level = 35
	item.Lattice_position = 8
	item.Onsale = 1
	InstYl_discount_store.Items = append(InstYl_discount_store.Items,item)
}