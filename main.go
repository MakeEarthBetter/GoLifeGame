//生命游戏
//1.形参不能放不确定长度的数组,其指针也不行
//2.go的内建函数不如python那么全,有时需要自己写,多多复用代码
package main
import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width int = 80
	height int = 30
	seedNum int = 300
)

//定义一个新类型
type Line []bool
type Universe []Line

//新建世界
func NewUniverse() Universe{
	universe := make(Universe,height)
	for i,_ := range universe{
		universe[i] = make(Line,width)
	}
	return universe
}

//展示世界
func (u Universe) Show(){
	for _,v:= range u{
		lineStr :=""
		for _,z := range v{
			switch z {
			case true:
				lineStr+="o "
			default:
				lineStr+="- "
			}
		}
		fmt.Println(lineStr)
	}
}
//判断元素是否在切片里
func numInArr(n int,arr []int) bool{
	for _,v := range arr {
		if v==n{
			return true
		}
	}
	return false
}

//生成不重复的随机切片
func RandomIntSlice(arrlen,start,end int) []int{
	if end < start || (end-start) < arrlen {
		return nil
	}
	retarr := make([]int, arrlen)
	for i:=0;i<len(retarr);i+=0{
		//为随机数更新种子
		rand.Seed(time.Now().UnixNano())
		randNumber:=rand.Intn(end-start)+start
		//判断是否在数组中,不在就加进去
		if numInArr(randNumber,retarr)==false{
			retarr[i] = randNumber
			i++
		}
	}		
	return retarr
}

//随机激活25percent生命,因为是切片所以可以直接在其上改动
func (u Universe) Seed()  {
	randLifeSlice := RandomIntSlice(seedNum,0,height*width)
	for _,v := range randLifeSlice{
		//计算所在位置
		ruotient,remainder:=v/int(width),v%int(width)
		//设为存活
		u[ruotient][remainder]=true
	}
}

//判断周围有多少存活的棋子
func(u Universe) LiveCount(i,j int) int{
	count:=0
	for x:=i-1;x<=i+1;x++{
		for y:=j-1;y<=j+1;y++{
			//越界处理并且不加自身
			if !(x<0 || x>=len(u) || y<0 || y>=len(u[x])) && !(x==i && y==j){
				if u[x][y]==true{
					count+=1
				}
			}
		}
	}
	return count
}

//统计周围细胞存活个数
func (u Universe) Next() Universe{
	//边遍历边改,先拷贝一份原Universe
	newUniverse := make(Universe,height) //开辟空间
	copy(newUniverse,u)
	//遍历修改元素下一世代的情况
	for i,line := range u{
		for j,_ :=range line{
			//判断周围棋子,通过copy出来的宇宙判断
			liveCount:=newUniverse.LiveCount(i,j)
			//低于两个死亡
			if liveCount<2{
				u[i][j] = false
			} else if liveCount>3 { //细胞存活时,两个或三个时保持原样
				u[i][j] = false //大于三个时死亡
			}else if (liveCount==3 && newUniverse[i][j]==false){//细胞死亡时,三个时保持变活
				u[i][j] = true
			}
			}
		}
	return u
	}

func main(){
	//新宇宙
	var universe Universe = NewUniverse()
	universe.Seed()
	universe.Show()
	for i:=0;i<=1;i+=0 { //死循环
		newu := universe.Next()
		fmt.Println("===================================================")
		newu.Show()
		time.Sleep(time.Duration(2)*time.Second)
	}
}


