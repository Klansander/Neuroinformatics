package main 
import (
	"fmt"
)
func main(){
	selection:=[300][15] int{
		{1,1,1,1,0,1,1,0,1,1,0,1,1,1,1},  
		{0,0,1,0,0,1,0,0,1,0,0,1,0,0,1}, 
		{1,1,1,0,0,1,1,1,1,1,0,0,1,1,1}, 
		{1,1,1,0,0,1,1,1,1,0,0,1,1,1,1},  
		{1,0,1,1,0,1,1,1,1,0,0,1,0,0,1},  
		{1,1,1,1,0,0,1,1,1,0,0,1,1,1,1},  
		{1,1,1,1,0,0,1,1,1,1,0,1,1,1,1},  
		{1,1,1,0,0,1,0,0,1,0,0,1,0,0,1},  
		{1,1,1,1,0,1,1,1,1,1,0,1,1,1,1}, 
		{1,1,1,1,0,1,1,1,1,0,0,1,1,1,1},  
		
	}
	vectorW:=[10][15]int {
		{7,6,8,1,4,4,7,2,6,2,3,4,7,1,1},
		{3,2,1,8,5,4,9,6,1,5,6,2,0,1,2},
		{7,2,8,1,10,4,6,2,6,2,3,4,6,1,3},
		{2,4,7,3,4,5,5,2,-1,2,4,4,7,1,3},
		{4,3,8,1,8,7,3,2,-2,4,3,4,2,1,5},
		{6,2,9,-1,4,5,2,2,3,2,3,4,3,1,4},
		{1,4,0,1,4,4,3,1,2,7,1,2,5,1,6},
		{2,5,1,2,1,8,1,5,6,9,3,3,6,1,1},
		{6,1,5,3,4,9,8,0,6,2,10,4,2,1,0},
	}
	count:=0
	ok:=false
	num:=0
	for count <100{
	for i,val:=range selection {
		eduSelection(val,&vectorW ,i)
	}
	for i:=0;i<10;i++ {
		ok=check(selection[i],&vectorW ,i)
		if ok==true {
			num++
		}
	}
	fmt.Println("Прогон №",count,"Количество разспознанных цифр: ", num)
	if num==10 {
		break
	}else{
		count++
	}
	num=0
}
fmt.Println(count)
	
}
func check(selection [15]int,vectorW *[10][15]int, num int ) bool {
	sumArray:=[]int{}
	for _,val := range vectorW {
		sum:=0
		for i,pr :=range val {
			sum+=(selection[i]*pr)
		}
		sumArray=append(sumArray,sum)
	}
	count:=false
	a:=sumArray[num%10]
	
	max:=sumArray[0]
	for _,val := range sumArray {
		if max<val {
			max=val
		}
	}

	if max==a{
		count=true
	}
	
	return count

}
func eduSelection(selection [15]int,vectorW *[10][15]int, num int ){
	sumArray:=[]int{}
	for _,val := range vectorW {
		sum:=0
		for i,pr :=range val {
			sum+=(selection[i]*pr)
		}
		sumArray=append(sumArray,sum)
	}
	a:=sumArray[num%10]
	dec:=0
	max:=sumArray[0]
	for i,val := range sumArray {
		if max<val {
			max=val
			dec=i
		}
	}

	if max>a{
		for j:=0;j<15;j++ {
			if(selection[j]==1){
				vectorW[dec][j]-=1
				vectorW[num%10][j]+=1
			}
		
		}
	}
	

}