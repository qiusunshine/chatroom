package user

import (
	"strings"
	"math/rand"
	"strconv"
)

var heads,tails []string
func NewName()string{
	if heads==nil{
		loadNames()
	}
	headsRand:=rand.Intn(len(heads))
	tailsRand:=rand.Intn(len(tails))
	return heads[headsRand]+tails[tailsRand]+strconv.Itoa(rand.Intn(500))
}
func loadNames() {
	s1:="陈,文,梁,谭,周,陆,吕,冯,甘,何,黄,胡,蒋,雷,黎,李,廖,林,赵,刘,莫,牟,宁,欧,庞,苏,王,谢,徐,杨,叶,张,郑,钟,朱"
	s2:="梅,阳,林,妮,博,宝,冰,波,贝,才,超,初,成,程,晨,德,富,福,枫,梵,刚,国,桂,罡,华,红,宏,辉,恒,慧,河,鸿,惠,桦,骅,剑,俊,杰,健,嘉,静,洁,娇,纪,宽,苛,珂,灵,兰,良,玲,磊,明,玛,媚,娜,朋,秋,青,琪,勤,晴,容,睿,蓉,胜,烁,堂,唯,伟,威,韦,雯,苇,香,兴,霞,萱,裕,颖,严,勇,阅,彦,宇,韵,燕,艳,乐,雨,洋,志,忠,宗,震"
	heads=strings.Split(s1,",")
	tails=strings.Split(s2,",")
}