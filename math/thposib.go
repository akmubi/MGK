package math

import "math"

// Средние значения столбцов матрицы
func (mat Matrix) GetAverages() Vector {
	// Создаём вектор средних значений столбцов
	averages := InitVector()
	averages.New(mat.Column_count)
	// Вычисляем средние по столбцам
	for j := 0; j < mat.Column_count; j++ {
		aver := 0.0
		for i := 0; i < mat.Row_count; i++ {
			aver += mat.Array[i][j]
		}
		averages.Array[j] = aver / float64(mat.Row_count)
	}
	return averages
}

// Дисперсии столбцов матрицы
func (mat Matrix) GetDispersions() Vector {
	// Создаём вектор дисперсий
	dispersions := InitVector()
	dispersions.New(mat.Column_count)

	// Вычисляем средние значения столбцов
	averages := mat.GetAverages()

	// Вычисляем дисперсии
	for j := 0; j < mat.Column_count; j++ {
		sum := 0.0
		for i := 0; i < mat.Row_count; i++ {
			sum += (mat.Array[i][j] - averages.Array[j]) * (mat.Array[i][j] - averages.Array[j]) 
		}
		dispersions.Array[j] = math.Sqrt(sum / float64(mat.Row_count)) 
	}
	return dispersions
}

// Выполнить стандартизацию матрицы
func (mat *Matrix) Standartize() {
	// Вычисляем средние и дисперсии каждого столбца матрицы 
	averages := mat.GetAverages()
	dispersions := mat.GetDispersions()

	for i := range mat.Array {
		for j := range mat.Array[i] {
			mat.Array[i][j] = (mat.Array[i][j] - averages.Array[j]) / dispersions.Array[j]
		}
	}
}

// Вычислить ковариационную матрицу
func (mat Matrix) GetCovariation() Matrix {
	result := InitMatrix()
	result.New(mat.Column_count, mat.Column_count)

	for i := 0; i < mat.Column_count; i++ {
		for j := 0; j < mat.Column_count; j++ {
			for k := 0; k < mat.Row_count; k++ {
				result.Array[i][j] += mat.Array[k][i] * mat.Array[k][j]
			}
		}
	}
	return result
}
