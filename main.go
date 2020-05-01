package main

import "./math"

const eps = 10e-5

func main() {
	mat := math.InitMatrix()
	mat.Read("check.txt")
	mat.Write()
	mat_A, mat_T := mat.JacobiProcedure(eps)
	mat_A.Write()
	mat_T.Write()
	_, eigenvectors := math.SortEigenMatrices(mat_A, mat_T)
	for _, vector := range eigenvectors {
		vector.Write()
	}
}