package main

import "fmt"

func minDistance2( a string ,  b string ) int {
	dp := make([][]int, len(a)+1)
	for i:=range dp{
		dp[i] = make([]int, len(b)+1)
	}
	for i:=0; i<len(dp); i++{dp[i][0] = i}
	for i:=0; i<len(dp[0]); i++{dp[0][i] = i}
	for i:=0; i<len(a); i++{
		for j:=0; j<len(b); j++{
			if a[i] == b[j] {
				dp[i+1][j+1] = dp[i][j]

			}else{
				if dp[i+1][j] == dp[i][j+1] && dp[i][j]>dp[i+1][j]{
					dp[i+1][j+1] = dp[i+1][j]
				}else{
					dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1])+1
				}
			}
			fmt.Print(dp[i+1][j+1], "\t")

		}
		fmt.Println()
	}
	return dp[len(a)][len(b)]
}

func min(a,b int) int{
	if a<b{return a}
	return b
}

func minDistance1( a string ,  b string ) int {
	// write code here
	count:=0
	a1 := []byte(a)
	b1 := []byte(b)
	m := make(map[byte]int, len(b1))
	for i:=0; i<len(b1); i++{
		m[b1[i]] = i
	}
	newA := []byte{}//删除了多余的a
	for i:=0; i<len(a1); i++{
		if _, ok := m[a1[i]]; !ok{
			count++
		}else{
			newA = append(newA, a1[i])
		}
	}
	//新数组hash
	m1 := make(map[byte]int)
	for i:=0;i<len(newA); i++{
		m1[newA[i]] = i
	}
	flag := false
	needCha := []byte{}
	for i:=0; i<len(b1); i++{
		if _, ok := m1[b1[i]]; !ok{
			flag = true
			needCha = append(needCha, b1[i])
			count++
		}
	}
	//一个都没有，都得重新插入
	if len(needCha) >= len(b1)-1 {return count}
	for i:=0; i<len(newA); i++{
		fmt.Print(string(newA[i]))
	}
	for i:=0; i<len(b1); i++{
		fmt.Print(string(b1[i]))
	}
	fmt.Println()
	if flag{//插
		fmt.Println("不会")
	}else{
		//l := len(b1)
		m2 := make(map[byte]int)
		for i:=0; i<len(newA); i++{
			if newA[i] == b1[i]{
				m2[newA[i]] = i
			}
		}
		fmt.Println(m2)
		newA1 := []byte{}
		for i:= 0; i<len(newA); i++{
			if i!= m2[newA[i]]{
				newA1 = append(newA1, newA[i])
			}
		}
		b2 := []byte{}
		for i:= 0; i<len(b1); i++{
			if i!= m2[b1[i]]{
				b2 = append(b2, b1[i])
			}
		}

		fmt.Println(string(b2[0]))
		//剩余数组计算交叉元素个数
		for i:=0; i<len(newA1); i++{
			for j:=0; j<len(b2); j++{
				if newA1[i] == b2[j] && newA1[j] == b2[i]{//一次交换
					if i<j{
						newA1 = append(newA1[:j], newA1[j+1:]...)
						b2 = append(b2[:j], b2[j+1:]...)
						newA1 = append(newA1[:i], newA1[i+1:]...)
						b2 = append(b2[:i], b2[i+1:]...)
						fmt.Println(b2)
					}else{
						newA1 = append(newA1[:i], newA1[i+1:]...)
						b2 = append(b2[:i], b2[i+1:]...)
						newA1 = append(newA1[:j], newA1[j+1:]...)
						b2 = append(b2[:j], b2[j+1:]...)
						fmt.Println(b2)
					}
					i,j = 0,0
					count++
				}
			}
		}
		count += len(newA1)*2
	}
	return count
}

func main() {
	//fmt.Println(minDistance1("horse", "rose"))
	fmt.Println(minDistance2("horse", "rose"))
}