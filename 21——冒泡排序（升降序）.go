// package main		//冒泡排序算法

// import "fmt"

// func main() {
//     values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27,12}
//     fmt.Println(values)
//     BubbleAsort(values)
//     BubbleZsort(values)
// }

// func BubbleAsort(values []int) {
//     for i := 0; i < len(values)-1; i++ {
//         for j := i+1; j < len(values); j++ {
//             if  values[i]>values[j]{
//                 values[i],values[j] = values[j],values[i]
//             }
//         }
//     }
//     fmt.Println(values)
// }

// func BubbleZsort(values []int) {
//     for i := 0; i < len(values)-1; i++ {
//         for j := i+1; j < len(values); j++ {
//             if  values[i]<values[j]{
//                 values[i],values[j] = values[j],values[i]
//             }
//         }
//     }
//     fmt.Println(values)
// }

package main

import "fmt"

func main(){
    var arr = [] int{1, 8, 9, 6, 3, 4, 7, 2, 0, 5} 
    fmt.Println("原始数组为",arr)
    maoPao1(arr)
    maoPao2(arr)
}

func maoPao1(arr []int){
    for i :=  0; i < len(arr); i ++{
        for j := i; j < len(arr); j ++{
            if arr[i] > arr[j]{
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
    fmt.Println("升序排序后的数组为：", arr)
}

func maoPao2(arr  [] int){
    n := len(arr)
    for i := 0; i < n; i ++{
        for j := i; j < n; j ++{
            if arr[i] < arr[j]{
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
    fmt.Println("降序排列后的数组为：", arr)
}