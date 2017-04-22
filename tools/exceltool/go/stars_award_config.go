package config
type Stars_award struct{
	Items []*Stars_award_Item
}

type Stars_award_Item struct{
	Id int
	Areas int
	Type int
	Stars int
	Award [][]int
}

var InstStars_award *Stars_award

func init(){
	item := &Stars_award_Item{}
	item.Id = 1001
	item.Areas = 1
	item.Type = 1
	item.Stars = 6
	item.Award = [][]int{{1,1,50},{2,2,2000},{201,310020,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1002
	item.Areas = 1
	item.Type = 1
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,310020,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1003
	item.Areas = 1
	item.Type = 1
	item.Stars = 18
	item.Award = [][]int{{1,1,100},{2,2,5000},{201,310020,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1004
	item.Areas = 1
	item.Type = 2
	item.Stars = 6
	item.Award = [][]int{{1,1,50},{2,2,2000},{201,310020,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1005
	item.Areas = 1
	item.Type = 2
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,310020,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1006
	item.Areas = 1
	item.Type = 2
	item.Stars = 18
	item.Award = [][]int{{1,1,100},{2,2,5000},{201,310020,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1007
	item.Areas = 2
	item.Type = 1
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,2000},{201,310042,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1008
	item.Areas = 2
	item.Type = 1
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,310042,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1009
	item.Areas = 2
	item.Type = 1
	item.Stars = 24
	item.Award = [][]int{{1,1,100},{2,2,5000},{201,310042,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1010
	item.Areas = 2
	item.Type = 2
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,2000},{201,310042,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1011
	item.Areas = 2
	item.Type = 2
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,310042,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1012
	item.Areas = 2
	item.Type = 2
	item.Stars = 24
	item.Award = [][]int{{1,1,100},{2,2,5000},{201,310042,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1013
	item.Areas = 3
	item.Type = 1
	item.Stars = 9
	item.Award = [][]int{{1,1,50},{2,2,2000},{201,310049,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1014
	item.Areas = 3
	item.Type = 1
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,310049,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1015
	item.Areas = 3
	item.Type = 1
	item.Stars = 21
	item.Award = [][]int{{1,1,100},{201,432003,2},{201,310049,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1016
	item.Areas = 3
	item.Type = 2
	item.Stars = 9
	item.Award = [][]int{{1,1,50},{2,2,2000},{201,310049,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1017
	item.Areas = 3
	item.Type = 2
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,310049,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1018
	item.Areas = 3
	item.Type = 2
	item.Stars = 21
	item.Award = [][]int{{1,1,100},{201,432003,2},{201,310049,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1019
	item.Areas = 4
	item.Type = 1
	item.Stars = 9
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,1}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1020
	item.Areas = 4
	item.Type = 1
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1021
	item.Areas = 4
	item.Type = 1
	item.Stars = 21
	item.Award = [][]int{{1,1,100},{5,5,500},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1022
	item.Areas = 4
	item.Type = 2
	item.Stars = 9
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,1}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1023
	item.Areas = 4
	item.Type = 2
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1024
	item.Areas = 4
	item.Type = 2
	item.Stars = 21
	item.Award = [][]int{{1,1,100},{5,5,500},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1025
	item.Areas = 5
	item.Type = 1
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,1}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1026
	item.Areas = 5
	item.Type = 1
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1027
	item.Areas = 5
	item.Type = 1
	item.Stars = 24
	item.Award = [][]int{{1,1,100},{5,5,500},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1028
	item.Areas = 5
	item.Type = 2
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,1}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1029
	item.Areas = 5
	item.Type = 2
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1030
	item.Areas = 5
	item.Type = 2
	item.Stars = 24
	item.Award = [][]int{{1,1,100},{5,5,500},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1031
	item.Areas = 6
	item.Type = 1
	item.Stars = 9
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1032
	item.Areas = 6
	item.Type = 1
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1033
	item.Areas = 6
	item.Type = 1
	item.Stars = 21
	item.Award = [][]int{{1,1,100},{5,5,1000},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1034
	item.Areas = 6
	item.Type = 2
	item.Stars = 9
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1035
	item.Areas = 6
	item.Type = 2
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1036
	item.Areas = 6
	item.Type = 2
	item.Stars = 21
	item.Award = [][]int{{1,1,100},{5,5,1000},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1037
	item.Areas = 7
	item.Type = 1
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1038
	item.Areas = 7
	item.Type = 1
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1039
	item.Areas = 7
	item.Type = 1
	item.Stars = 24
	item.Award = [][]int{{1,1,100},{5,5,1500},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1040
	item.Areas = 7
	item.Type = 2
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1041
	item.Areas = 7
	item.Type = 2
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1042
	item.Areas = 7
	item.Type = 2
	item.Stars = 24
	item.Award = [][]int{{1,1,100},{5,5,1500},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1043
	item.Areas = 8
	item.Type = 1
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1044
	item.Areas = 8
	item.Type = 1
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1045
	item.Areas = 8
	item.Type = 1
	item.Stars = 24
	item.Award = [][]int{{1,1,100},{5,5,2000},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1046
	item.Areas = 8
	item.Type = 2
	item.Stars = 12
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1047
	item.Areas = 8
	item.Type = 2
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1048
	item.Areas = 8
	item.Type = 2
	item.Stars = 24
	item.Award = [][]int{{1,1,100},{5,5,2000},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1049
	item.Areas = 9
	item.Type = 1
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1050
	item.Areas = 9
	item.Type = 1
	item.Stars = 24
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1051
	item.Areas = 9
	item.Type = 1
	item.Stars = 33
	item.Award = [][]int{{1,1,100},{5,5,2500},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1052
	item.Areas = 9
	item.Type = 2
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1053
	item.Areas = 9
	item.Type = 2
	item.Stars = 24
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1054
	item.Areas = 9
	item.Type = 2
	item.Stars = 33
	item.Award = [][]int{{1,1,100},{5,5,2500},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1055
	item.Areas = 10
	item.Type = 1
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1056
	item.Areas = 10
	item.Type = 1
	item.Stars = 24
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1057
	item.Areas = 10
	item.Type = 1
	item.Stars = 33
	item.Award = [][]int{{1,1,100},{5,5,3000},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1058
	item.Areas = 10
	item.Type = 2
	item.Stars = 15
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1059
	item.Areas = 10
	item.Type = 2
	item.Stars = 24
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1060
	item.Areas = 10
	item.Type = 2
	item.Stars = 33
	item.Award = [][]int{{1,1,100},{5,5,3000},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1061
	item.Areas = 11
	item.Type = 1
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1062
	item.Areas = 11
	item.Type = 1
	item.Stars = 24
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1063
	item.Areas = 11
	item.Type = 1
	item.Stars = 36
	item.Award = [][]int{{1,1,100},{5,5,3000},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1064
	item.Areas = 11
	item.Type = 2
	item.Stars = 18
	item.Award = [][]int{{1,1,50},{2,2,3000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1065
	item.Areas = 11
	item.Type = 2
	item.Stars = 24
	item.Award = [][]int{{1,1,50},{2,2,5000},{201,432003,2}}
	InstStars_award.Items = append(InstStars_award.Items,item)
	item = &Stars_award_Item{}
	item.Id = 1066
	item.Areas = 11
	item.Type = 2
	item.Stars = 36
	item.Award = [][]int{{1,1,100},{5,5,3000},{201,432003,3}}
	InstStars_award.Items = append(InstStars_award.Items,item)
}