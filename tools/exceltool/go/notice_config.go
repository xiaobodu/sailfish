package config
type Notice struct{
	Items []*Notice_Item
}

type Notice_Item struct{
	Id int
	Type int
}

var InstNotice *Notice

func init(){
	item := &Notice_Item{}
	item.Id = 1001
	item.Type = 1
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1002
	item.Type = 2
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1003
	item.Type = 3
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1004
	item.Type = 4
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1005
	item.Type = 5
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1006
	item.Type = 6
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1007
	item.Type = 7
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1008
	item.Type = 8
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1009
	item.Type = 9
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1010
	item.Type = 10
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1011
	item.Type = 10
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1012
	item.Type = 11
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1013
	item.Type = 12
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1014
	item.Type = 13
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1015
	item.Type = 14
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1016
	item.Type = 14
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1050
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1051
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1052
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1053
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1054
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1055
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1056
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1057
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1058
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1059
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1060
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1061
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1062
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1063
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1064
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1065
	item.Type = 15
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1098
	item.Type = 16
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1099
	item.Type = 16
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1150
	item.Type = 17
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1151
	item.Type = 17
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1152
	item.Type = 17
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1153
	item.Type = 18
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1154
	item.Type = 19
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1155
	item.Type = 20
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1156
	item.Type = 21
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1157
	item.Type = 22
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1158
	item.Type = 23
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1159
	item.Type = 24
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1160
	item.Type = 25
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1161
	item.Type = 26
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1162
	item.Type = 27
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1163
	item.Type = 28
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1164
	item.Type = 29
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1165
	item.Type = 30
	InstNotice.Items = append(InstNotice.Items,item)
	item = &Notice_Item{}
	item.Id = 1166
	item.Type = 31
	InstNotice.Items = append(InstNotice.Items,item)
}