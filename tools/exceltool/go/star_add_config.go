package config
type Star_add struct{
	Items []*Star_add_Item
}

type Star_add_Item struct{
	Star int
	Attribute map[string]int
	Attribute_add map[string]int
	Consume map[string]int
}

var InstStar_add *Star_add

func init(){
	item := &Star_add_Item{}
	item.Star = 7
	item.Attribute = map[string]int{"4":12,"5":5,"6":40}
	item.Attribute_add = map[string]int{"4":12,"5":5,"6":40}
	item.Consume = map[string]int{"2":2000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 11
	item.Attribute = map[string]int{"4":12,"5":5,"6":40}
	item.Attribute_add = map[string]int{"4":24,"5":10,"6":80}
	item.Consume = map[string]int{"2":5000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 18
	item.Attribute = map[string]int{"4":12,"5":5,"6":40}
	item.Attribute_add = map[string]int{"4":36,"5":15,"6":120}
	item.Consume = map[string]int{"2":10000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 24
	item.Attribute = map[string]int{"4":45}
	item.Attribute_add = map[string]int{"4":81,"5":15,"6":120}
	item.Consume = map[string]int{"2":20000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 30
	item.Attribute = map[string]int{"5":30}
	item.Attribute_add = map[string]int{"4":81,"5":45,"6":120}
	item.Consume = map[string]int{"2":30000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 36
	item.Attribute = map[string]int{"6":180}
	item.Attribute_add = map[string]int{"4":81,"5":45,"6":300}
	item.Consume = map[string]int{"2":40000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 42
	item.Attribute = map[string]int{"4":70}
	item.Attribute_add = map[string]int{"4":151,"5":45,"6":300}
	item.Consume = map[string]int{"2":50000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 51
	item.Attribute = map[string]int{"5":45}
	item.Attribute_add = map[string]int{"4":151,"5":90,"6":300}
	item.Consume = map[string]int{"2":50000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 60
	item.Attribute = map[string]int{"6":240}
	item.Attribute_add = map[string]int{"4":151,"5":90,"6":540}
	item.Consume = map[string]int{"2":80000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 69
	item.Attribute = map[string]int{"1":1}
	item.Attribute_add = map[string]int{"4":151,"5":90,"6":540,"1":1}
	item.Consume = map[string]int{"2":80000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 78
	item.Attribute = map[string]int{"3":1}
	item.Attribute_add = map[string]int{"4":151,"5":90,"6":540,"1":1,"3":1}
	item.Consume = map[string]int{"2":100000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 87
	item.Attribute = map[string]int{"2":1}
	item.Attribute_add = map[string]int{"4":151,"5":90,"6":540,"1":1,"3":1,"2":1}
	item.Consume = map[string]int{"2":100000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 96
	item.Attribute = map[string]int{"4":100}
	item.Attribute_add = map[string]int{"4":251,"5":90,"6":540,"1":1,"3":1,"2":1}
	item.Consume = map[string]int{"2":150000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 105
	item.Attribute = map[string]int{"5":60}
	item.Attribute_add = map[string]int{"4":251,"5":150,"6":540,"1":1,"3":1,"2":1}
	item.Consume = map[string]int{"2":150000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 114
	item.Attribute = map[string]int{"6":360}
	item.Attribute_add = map[string]int{"4":251,"5":150,"6":900,"1":1,"3":1,"2":1}
	item.Consume = map[string]int{"2":200000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 123
	item.Attribute = map[string]int{"1":1,"3":1}
	item.Attribute_add = map[string]int{"4":251,"5":150,"6":900,"1":2,"3":2,"2":1}
	item.Consume = map[string]int{"2":200000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 132
	item.Attribute = map[string]int{"6":480,"2":1}
	item.Attribute_add = map[string]int{"4":251,"5":150,"6":1380,"1":2,"3":2,"2":2}
	item.Consume = map[string]int{"2":250000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 141
	item.Attribute = map[string]int{"10004":200}
	item.Attribute_add = map[string]int{"4":251,"5":150,"6":1380,"1":2,"3":2,"2":2,"10004":200}
	item.Consume = map[string]int{"2":300000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 150
	item.Attribute = map[string]int{"10005":200}
	item.Attribute_add = map[string]int{"4":251,"5":150,"6":1380,"1":2,"3":2,"2":2,"10004":200,"10005":200}
	item.Consume = map[string]int{"2":400000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 155
	item.Attribute = map[string]int{"10006":300}
	item.Attribute_add = map[string]int{"4":251,"5":150,"6":1380,"1":2,"3":2,"2":2,"10004":200,"10005":200,"10006":300}
	item.Consume = map[string]int{"2":400000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 160
	item.Attribute = map[string]int{"6":540,"10004":200}
	item.Attribute_add = map[string]int{"4":251,"5":150,"6":1920,"1":2,"3":2,"2":2,"10004":400,"10005":200,"10006":300}
	item.Consume = map[string]int{"2":500000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 165
	item.Attribute = map[string]int{"4":150,"10005":200}
	item.Attribute_add = map[string]int{"4":401,"5":150,"6":1920,"1":2,"3":2,"2":2,"10004":400,"10005":400,"10006":300}
	item.Consume = map[string]int{"2":500000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 170
	item.Attribute = map[string]int{"5":100,"10006":300}
	item.Attribute_add = map[string]int{"4":401,"5":250,"6":1920,"1":2,"3":2,"2":2,"10004":400,"10005":400,"10006":600}
	item.Consume = map[string]int{"2":500000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 175
	item.Attribute = map[string]int{"1":1,"3":1}
	item.Attribute_add = map[string]int{"4":401,"5":250,"6":1920,"1":3,"3":3,"2":2,"10004":400,"10005":400,"10006":600}
	item.Consume = map[string]int{"2":500000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 180
	item.Attribute = map[string]int{"2":1,"10004":200}
	item.Attribute_add = map[string]int{"4":401,"5":250,"6":1920,"1":3,"3":3,"2":3,"10004":600,"10005":400,"10006":600}
	item.Consume = map[string]int{"2":500000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 185
	item.Attribute = map[string]int{"1":2}
	item.Attribute_add = map[string]int{"4":401,"5":250,"6":1920,"1":5,"3":3,"2":3,"10004":600,"10005":400,"10006":600}
	item.Consume = map[string]int{"2":750000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 190
	item.Attribute = map[string]int{"3":2}
	item.Attribute_add = map[string]int{"4":401,"5":250,"6":1920,"1":5,"3":5,"2":3,"10004":600,"10005":400,"10006":600}
	item.Consume = map[string]int{"2":750000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 195
	item.Attribute = map[string]int{"2":1,"10005":200}
	item.Attribute_add = map[string]int{"4":401,"5":250,"6":1920,"1":5,"3":5,"2":4,"10004":600,"10005":600,"10006":600}
	item.Consume = map[string]int{"2":750000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 200
	item.Attribute = map[string]int{"4":125,"5":90,"6":500}
	item.Attribute_add = map[string]int{"4":526,"5":340,"6":2420,"1":5,"3":5,"2":4,"10004":600,"10005":600,"10006":600}
	item.Consume = map[string]int{"2":1000000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 205
	item.Attribute = map[string]int{"1":3}
	item.Attribute_add = map[string]int{"4":526,"5":340,"6":2420,"1":8,"3":5,"2":4,"10004":600,"10005":600,"10006":600}
	item.Consume = map[string]int{"2":1000000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 210
	item.Attribute = map[string]int{"3":3}
	item.Attribute_add = map[string]int{"4":526,"5":340,"6":2420,"1":8,"3":8,"2":4,"10004":600,"10005":600,"10006":600}
	item.Consume = map[string]int{"2":1250000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 215
	item.Attribute = map[string]int{"2":2,"10006":400}
	item.Attribute_add = map[string]int{"4":526,"5":340,"6":2420,"1":8,"3":8,"2":6,"10004":600,"10005":600,"10006":1000}
	item.Consume = map[string]int{"2":1250000}
	InstStar_add.Items = append(InstStar_add.Items,item)
	item = &Star_add_Item{}
	item.Star = 220
	item.Attribute = map[string]int{"10004":100,"10005":100,"10006":300}
	item.Attribute_add = map[string]int{"4":526,"5":340,"6":2420,"1":8,"3":8,"2":6,"10004":700,"10005":700,"10006":1300}
	item.Consume = map[string]int{"2":1500000}
	InstStar_add.Items = append(InstStar_add.Items,item)
}