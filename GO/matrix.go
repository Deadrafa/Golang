package main

import (
	"os"
	"os/exec"
)

// func main() {
// 	Clear_skrean()
// 	// Считывание матрицы Scanf и печать
// 	// var num_matrix [3][3]int
// 	// for i := 0; i < 3; i++ {
// 	// 	for j := 0; j < 3; j++ {
// 	// 		fmt.Scanf("%d", &num_matrix[i][j])
// 	// 	}
// 	// 	fmt.Printf("\n")
// 	// }
// 	num_matrix1 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// 	num_matrix2 := [3][3]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
// 	mat := Sum_Matrix(num_matrix1, num_matrix2)
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("%d ", mat[i][j])
// 		}
// 		fmt.Printf("\n")
// 	}

// }
func Sum_Matrix(matrix1 [3][3]int, matrix2 [3][3]int) [3][3]int {
	var num_matrix_sum [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num_matrix_sum[i][j] += matrix1[i][j] + matrix2[i][j]
		}
	}
	return num_matrix_sum
}

func Clear_skrean() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
