package main

//描述
//金闪闪死后，红A拿到了王之财宝，里面有n个武器，长度各不相同。红A发现，拿其中三件武器首尾相接，组成一个三角形，进行召唤仪式，就可以召唤出一个山寨金闪闪。（例如，三件武器长度为10、15、20，可以召唤成功。若长度为10、11、30，首尾相接无法组成三角形，召唤失败。）红A于是开了一个金闪闪专卖店。他把王之财宝排成一排，每个客人会随机抽取到一个区间[l,r],客人可以选取区间里的三件武器进行召唤（客人都很聪慧，如果能找出来合适的武器，一定不会放过）。召唤结束后，客人要把武器原样放回去。m个客人光顾以后，红A害怕过多的金闪闪愉悦太多男人，于是找到了你，希望你帮他统计出有多少山寨金闪闪被召唤出来。
//
//输入描述：
//第一行武器数量:n <= 1*10^7
//第二行空格分隔的n个int，表示每件武器的长度。
//第三行顾客数量：m <= 1*10^6
//后面m行，每行两个int l，r，表示每个客人被分配到的区间。（l<r）
//输出描述：
//山寨金闪闪数量。
//示例1
//输入：
//5
//1 10 100 95 101
//4
//1 3
//2 4
//2 5
//3 5
//复制
//输出：
//3

//解题关键：（区间l,r 是武器的下标）
//1.算出l，r在什么区间，三角形是一定可以构成的，用三角形a+b == c的临界情况可用Fibonacci数列构造这题的区间，1，1，2，3，5，8，13（区间里能包含的武器长度数量；如区间长度7，只有13种武器长度，如果区间长度47，超过int32种武器长度了）...
//2.武器范围是int32，就是我们构造的数组要包含int32的范围即可，满足所有能构造三角形的情况
//3.Fibonacci数列构造到47个下标就包含了int32所有数了，即区间长度>47就能构造出这题的三角形
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
const maxn = 1e7+5
var reader *bufio.Reader
var a [maxn]int

func HasJinShanShan(l,r int) bool{

	if r-l+1<3{
		return false
	}else if r-l+1>47{
		return true
	}else{
		weapon := []int{}
		for i:=l; i<=r;i++{
			weapon = append(weapon, a[i])
		}
		sort.Ints(weapon)
		for i:=0; i<len(weapon)-2;i++{
			//             前两条边大于第三条，能构成
			if weapon[i]+weapon[i+1]>weapon[i+2]{
				return true
			}
		}
	}
	return false
}

// func 02mapInterface(){
//     reader = bufio.NewScanner(os.Stdin)
//     var n int
//     fmt.Scanf("%v", &n)

//     for i:=1; i<=n; i++{
//         fmt.Scanf("%v", &a[i])
//     }

// //     fmt.Println(a[1])
//     var m int
//     var l,r ,cnt int
// //     fmt.Scanf("%v", &m)
//     reader.Scan()
//     m,_ = strconv.Atoi(reader.Text())
// //     fmt.Println(m)
//     for i:=0; i<m; i++{
//         reader.Scan()
//         lr := strings.Split(reader.Text(), " ")
// //         fmt.Println(lr)
//         l,_ = strconv.Atoi(lr[0])
//         r,_ = strconv.Atoi(lr[1])
// //         fmt.Println(l,r,HasJinShanShan(l,r))
//         if HasJinShanShan(l,r){cnt++}
//     }
//     fmt.Println(cnt)
// }
func main(){
	reader = bufio.NewReader(os.Stdin)
	var n int

	nStr := splitStdString('\n')
	n,_ = strconv.Atoi(nStr)
	for i:=1; i<n; i++{
		//         fmt.Print(buf, " ")
		buf := splitStdString(' ')
		a[i],_ = strconv.Atoi(buf)
	}
	buf := splitStdString('\n')

	a[n],_ = strconv.Atoi(buf)
	//     fmt.Println(a[:n+1])

	var m int
	var l,r ,cnt int
	mStr := splitStdString('\n')
	m,_ = strconv.Atoi(mStr)
	//     fmt.Println(m)
	for i:=0; i<m; i++{
		lrStr := splitStdString('\n')
		lr := strings.Split(lrStr, " ")

		//         fmt.Println("lr",lr)
		l,_ = strconv.Atoi(lr[0])
		r,_ = strconv.Atoi(lr[1])
		//         fmt.Println(l,r,HasJinShanShan(l,r))
		if HasJinShanShan(l,r){cnt++}
	}
	fmt.Println(cnt)
}

// 以split字符分割标准输入的字串
func splitStdString(split byte) string{
	text,_ := reader.ReadString(split)
	text = strings.Trim(text, string(split))
	return text
}