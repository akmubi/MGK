package main

import "./math"

func main() {
	mat := math.InitMatrix()
	mat.Read("check.txt")
	mat.Write()
	new_mat, mat_T := mat.JacobiProcedure(0.01)
	new_mat.Write()
	mat_T.Write()
}