package main

import "fmt"

type face struct{
	box_no int
//	l,b,h int
	dimensions []int
}

var face_arr []face

func main(){
	box := [][]int{ {4, 6, 7}, {1, 2, 3}, {4, 5, 6}, {10, 12, 32} }
	generatefaces(box)
	face_arr = sort(face_arr)
	height := make([]int, len(face_arr))
	global_max_index := len(face_arr)-1
	for i, _ := range face_arr{
		height[i] = face_arr[i].dimensions[2]
	}

	for i := len(height)-2; i >= 0;i--{
		max := height[i]
		for j := i; j<len(height); j++{
			if face_arr[i].dimensions[0] < face_arr[j].dimensions[0] && face_arr[i].dimensions[1] < face_arr[j].dimensions[1] && height[i] + height[j] > max{
				max = height[i] + height[j]
			}
		}
		height[i] = max
		if height[i] > height[global_max_index]{
			global_max_index = i
		}
	}
	fmt.Println("max height-", height[global_max_index])
	Printfaces(height, global_max_index)
}

func generatefaces(box [][]int){
	for i, v := range box{
		createface(i, []int{v[0], v[1], v[2]})
		createface(i, []int{v[0], v[2], v[1]})
                createface(i, []int{v[1], v[2], v[0]})
	}
}

func createface(i int, dim []int){
	var f face
	f.box_no = i
	f.dimensions = dim
	face_arr = append(face_arr, f)
}

func sort(f []face) []face {
	if len(f) == 1{
		return f
	}
	left := sort(f[:len(f)/2])
	right := sort(f[len(f)/2:])
	return merge(left, right)
}

func merge(left, right []face) []face{
	m := make([]face, (len(left) + len(right)))
	a := func(b []face, ind int)int{
		return b[ind].dimensions[0]*b[ind].dimensions[1]
		}
	i, j, k := 0,0,0
	for (i < len(left) && j < len(right)){
		if a(left, i) > a(right, j){
			m[k] = right[j]
			j += 1; k += 1

		}else {
			m[k] = left[i]
			i += 1; k += 1
		}
	}

	for y:=i; y<len(left);y++{
		m[k] = left[y]
		k += 1
	}

	for y:=j; y<len(right);y++{
                m[k] = right[y]
		k += 1
        }
	return m
}



func Printfaces(height []int, max_index int){
	fmt.Println(" L B H", " box_no")
	v := height[max_index]

	for i:= max_index; i<len(face_arr) && v >= 0;i++{
		if v == height[i]{
			v = height[i]-face_arr[i].dimensions[2]
			fmt.Println(face_arr[i].dimensions, (i/3)+1)
		}
	}
}

