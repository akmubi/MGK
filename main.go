package main

import "./math"
import stdio "fmt"

const eps = 10e-5

func main() {
	// Исходная матрица
	mat := math.InitMatrix()
	mat.Read("data.txt")
	stdio.Println("Исходная матрица данных:")
	mat.Write()

	// Средние значения столбцов
	avers := mat.GetAverages()
	stdio.Println("Средние:")
	avers.Write()

	// Дисперсии столбцов
	dispers := mat.GetDispersions()
	stdio.Println("Дисперсии:")
	dispers.Write()

	// Стандартизирование матрицы
	mat.Standartize()
	stdio.Println("Стандартизованная матрица:")
	mat.Write();
	stdio.Println("Дисперсии:")
	disps := mat.GetDispersions()
	disps.Write()
	stdio.Println("Средние:")
	avers = mat.GetAverages()
	avers.Write()

	// Корреляционная матрица
	mat_corel := mat.GetCorrelation()
	stdio.Println("Корреляционная матрица:")
	mat_corel.Write()
	
	// Величина D, показыващая отклонение корреляционной матрицы от единичной
	d := mat_corel.ExistDifference(mat.Row_count)
	stdio.Println("D -", d)
	// stdio.Println("rows -", mat_corel.Row_count)
	// stdio.Println("columns -", mat_corel.Column_count)

	// Нахождение собственных значений и собственных векторов
	mat_A, mat_T := mat_corel.JacobiProcedure(eps)
	stdio.Println("Матрица с собственными значениями:")
	mat_A.Write()
	stdio.Println("Матрица с собственными векторами:")
	mat_T.Write()

	// Сортировка собственных значений и собственных векторов по убыванию
	eigenvalues, eigenvectors := math.SortEigenMatrices(mat_A, mat_T)
	stdio.Println("Собственные значения:")
	eigenvalues.Write()
	stdio.Println("Собственные векторы:")
	printVectors(eigenvectors)

	// Проекции объектов на главные компоненты
	stdio.Println("Проекции на главные компоненты:")
	main_components := math.CalculateMainComponents(mat, eigenvectors)
	printVectors(main_components)
	stdio.Println("Дисперсии:")
	stdio.Println("[")
	for _, v := range main_components {
		stdio.Println("\t", v.GetDispersion())
	}
	stdio.Println("]")

	// Проверка равенства дисперсий
	sum1, sum2 := math.CheckDispersionEquality(mat.ConvertToVec(), main_components)
	stdio.Println("sum1, sum2 =", sum1, sum2)

	// Относительная доля разброса
	part_size, I := math.CalculateIValue(eigenvalues)
	stdio.Println("part size -", part_size)
	stdio.Println("I value -", I)

	// Ковариационна матрица вычисленных проекций
	stdio.Println("Ковариационная матрица проекций на главные компоненты:")
	main_comp_mat := math.ConvertToMat(main_components)
	covar_main_components := main_comp_mat.GetCovariation()
	covar_main_components.Write()

}

func printVectors(vecs []math.Vector) {
	for _, vector := range vecs {
		vector.Write()
	}
}