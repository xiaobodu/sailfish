package config
type City_spot struct{
	Items []*City_spot_Item
}

type City_spot_Item struct{
	Spot_id int
	City_id int
}

var InstCity_spot *City_spot

func init(){
	item := &City_spot_Item{}
	item.Spot_id = 10011
	item.City_id = 1001
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10012
	item.City_id = 1001
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10021
	item.City_id = 1002
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10022
	item.City_id = 1002
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10031
	item.City_id = 1003
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10032
	item.City_id = 1003
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 0
	item.City_id = 1004
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10051
	item.City_id = 1005
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10052
	item.City_id = 1005
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10061
	item.City_id = 1006
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10062
	item.City_id = 1006
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10071
	item.City_id = 1007
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10072
	item.City_id = 1007
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10081
	item.City_id = 1008
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10082
	item.City_id = 1008
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10091
	item.City_id = 1009
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10092
	item.City_id = 1009
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10101
	item.City_id = 1010
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10102
	item.City_id = 1010
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10111
	item.City_id = 1011
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10112
	item.City_id = 1011
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10113
	item.City_id = 1011
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10114
	item.City_id = 1011
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10121
	item.City_id = 1012
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10122
	item.City_id = 1012
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10131
	item.City_id = 1013
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10132
	item.City_id = 1013
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10141
	item.City_id = 1014
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10142
	item.City_id = 1014
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10151
	item.City_id = 1015
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10152
	item.City_id = 1015
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10161
	item.City_id = 1016
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10162
	item.City_id = 1016
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10171
	item.City_id = 1017
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10172
	item.City_id = 1017
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10181
	item.City_id = 1018
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10182
	item.City_id = 1018
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10191
	item.City_id = 1019
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10192
	item.City_id = 1019
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10201
	item.City_id = 1020
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10202
	item.City_id = 1020
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10211
	item.City_id = 1021
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10212
	item.City_id = 1021
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10221
	item.City_id = 1022
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10222
	item.City_id = 1022
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10231
	item.City_id = 1023
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10232
	item.City_id = 1023
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10241
	item.City_id = 1024
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10242
	item.City_id = 1024
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10251
	item.City_id = 1025
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10252
	item.City_id = 1025
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10261
	item.City_id = 1026
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10262
	item.City_id = 1026
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10263
	item.City_id = 1026
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10264
	item.City_id = 1026
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10271
	item.City_id = 1027
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10271
	item.City_id = 1027
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10281
	item.City_id = 1028
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10282
	item.City_id = 1028
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10291
	item.City_id = 1029
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10292
	item.City_id = 1029
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10301
	item.City_id = 1030
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10302
	item.City_id = 1030
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10311
	item.City_id = 1031
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10312
	item.City_id = 1031
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10313
	item.City_id = 1031
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10314
	item.City_id = 1031
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10321
	item.City_id = 1032
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10322
	item.City_id = 1032
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10331
	item.City_id = 1033
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10332
	item.City_id = 1033
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10341
	item.City_id = 1034
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10342
	item.City_id = 1034
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10351
	item.City_id = 1035
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10352
	item.City_id = 1035
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10361
	item.City_id = 1036
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10362
	item.City_id = 1036
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10363
	item.City_id = 1036
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10364
	item.City_id = 1036
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10371
	item.City_id = 1037
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10372
	item.City_id = 1037
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10381
	item.City_id = 1038
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10382
	item.City_id = 1038
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10383
	item.City_id = 1038
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10384
	item.City_id = 1038
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10391
	item.City_id = 1039
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10392
	item.City_id = 1039
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10401
	item.City_id = 1040
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10402
	item.City_id = 1040
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10411
	item.City_id = 1041
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10412
	item.City_id = 1041
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10421
	item.City_id = 1042
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10422
	item.City_id = 1042
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10431
	item.City_id = 1043
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10432
	item.City_id = 1043
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10441
	item.City_id = 1044
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10442
	item.City_id = 1044
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10451
	item.City_id = 1045
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10452
	item.City_id = 1045
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10461
	item.City_id = 1046
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10462
	item.City_id = 1046
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10471
	item.City_id = 1047
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10472
	item.City_id = 1047
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10481
	item.City_id = 1048
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10482
	item.City_id = 1048
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10491
	item.City_id = 1049
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10492
	item.City_id = 1049
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10501
	item.City_id = 1050
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10502
	item.City_id = 1050
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10503
	item.City_id = 1050
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10504
	item.City_id = 1050
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10511
	item.City_id = 1051
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10512
	item.City_id = 1051
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10521
	item.City_id = 1052
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10522
	item.City_id = 1052
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10531
	item.City_id = 1053
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10532
	item.City_id = 1053
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10541
	item.City_id = 1054
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10542
	item.City_id = 1054
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10551
	item.City_id = 1055
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10552
	item.City_id = 1055
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10561
	item.City_id = 1056
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10562
	item.City_id = 1056
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10571
	item.City_id = 1057
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10572
	item.City_id = 1057
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10581
	item.City_id = 1058
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10582
	item.City_id = 1058
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10591
	item.City_id = 1059
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10592
	item.City_id = 1059
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10601
	item.City_id = 1060
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10602
	item.City_id = 1060
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10611
	item.City_id = 1061
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10612
	item.City_id = 1061
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10621
	item.City_id = 1062
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10622
	item.City_id = 1062
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10623
	item.City_id = 1062
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10624
	item.City_id = 1062
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10631
	item.City_id = 1063
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10632
	item.City_id = 1063
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10641
	item.City_id = 1064
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10642
	item.City_id = 1064
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10651
	item.City_id = 1065
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10652
	item.City_id = 1065
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10653
	item.City_id = 1065
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10654
	item.City_id = 1065
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10661
	item.City_id = 1066
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10662
	item.City_id = 1066
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10671
	item.City_id = 1067
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10672
	item.City_id = 1067
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10681
	item.City_id = 1068
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10682
	item.City_id = 1068
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10691
	item.City_id = 1069
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10692
	item.City_id = 1069
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10701
	item.City_id = 1070
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10702
	item.City_id = 1070
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10711
	item.City_id = 1071
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10712
	item.City_id = 1071
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10721
	item.City_id = 1072
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10722
	item.City_id = 1072
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10731
	item.City_id = 1073
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10732
	item.City_id = 1073
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10741
	item.City_id = 1074
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10742
	item.City_id = 1074
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10743
	item.City_id = 1074
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10744
	item.City_id = 1074
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10751
	item.City_id = 1075
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10752
	item.City_id = 1075
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10761
	item.City_id = 1076
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10762
	item.City_id = 1076
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10771
	item.City_id = 1077
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10772
	item.City_id = 1077
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10781
	item.City_id = 1078
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10782
	item.City_id = 1078
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10791
	item.City_id = 1079
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10792
	item.City_id = 1079
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10793
	item.City_id = 1079
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10794
	item.City_id = 1079
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10801
	item.City_id = 1080
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10802
	item.City_id = 1080
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10811
	item.City_id = 1081
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10812
	item.City_id = 1081
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10821
	item.City_id = 1082
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10822
	item.City_id = 1082
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10831
	item.City_id = 1083
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10832
	item.City_id = 1083
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10841
	item.City_id = 1084
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10842
	item.City_id = 1084
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10851
	item.City_id = 1085
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10852
	item.City_id = 1085
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10853
	item.City_id = 1085
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10854
	item.City_id = 1085
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10861
	item.City_id = 1086
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10862
	item.City_id = 1086
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10871
	item.City_id = 1087
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10872
	item.City_id = 1087
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10881
	item.City_id = 1088
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10882
	item.City_id = 1088
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10891
	item.City_id = 1089
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10892
	item.City_id = 1089
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10901
	item.City_id = 1090
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10902
	item.City_id = 1090
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10911
	item.City_id = 1091
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10912
	item.City_id = 1091
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10921
	item.City_id = 1092
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10922
	item.City_id = 1092
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10931
	item.City_id = 1093
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10932
	item.City_id = 1093
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10941
	item.City_id = 1094
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10942
	item.City_id = 1094
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10951
	item.City_id = 1095
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10952
	item.City_id = 1095
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10953
	item.City_id = 1095
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10954
	item.City_id = 1095
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10961
	item.City_id = 1096
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10962
	item.City_id = 1096
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10971
	item.City_id = 1097
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10972
	item.City_id = 1097
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10973
	item.City_id = 1097
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10974
	item.City_id = 1097
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10981
	item.City_id = 1098
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10982
	item.City_id = 1098
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10991
	item.City_id = 1099
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 10992
	item.City_id = 1099
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11001
	item.City_id = 1100
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11002
	item.City_id = 1100
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11011
	item.City_id = 1101
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11012
	item.City_id = 1101
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11021
	item.City_id = 1102
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11022
	item.City_id = 1102
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11031
	item.City_id = 1103
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11032
	item.City_id = 1103
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11041
	item.City_id = 1104
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11042
	item.City_id = 1104
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11051
	item.City_id = 1105
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11052
	item.City_id = 1105
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11061
	item.City_id = 1106
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11062
	item.City_id = 1106
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11071
	item.City_id = 1107
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11072
	item.City_id = 1107
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11081
	item.City_id = 1108
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11082
	item.City_id = 1108
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11091
	item.City_id = 1109
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11092
	item.City_id = 1109
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11101
	item.City_id = 1110
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11102
	item.City_id = 1110
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11111
	item.City_id = 1111
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11112
	item.City_id = 1111
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11113
	item.City_id = 1111
	InstCity_spot.Items = append(InstCity_spot.Items,item)
	item = &City_spot_Item{}
	item.Spot_id = 11114
	item.City_id = 1111
	InstCity_spot.Items = append(InstCity_spot.Items,item)
}