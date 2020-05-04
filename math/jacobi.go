package math

import (
	"math"
	"errors"
)

//
// МЕТОД ЯКОБИ С ПРЕГРАДАМИ ПО НАХОЖДЕНИЮ СОБСТВЕННЫХ ЗНАЧЕНИЙ И СОБСТВЕННЫХ ВЕКТОРОВ 
//

// Шаг 1
func (mat_A Matrix) JacobiProcedure(eps float64) (Matrix, Matrix) {
	mat_A.checkSimmetry()
	// Матрица, столбцы которой являются собственными векторами
	mat_T := Identity(mat_A.Row_count)

	// Вычисляем начальную преграду
	a0 := mat_A.calculateBarrier()
	for ak := a0; !mat_A.checkAllElementsLessThanEpsBarrier(eps, a0); ak /= float64(mat_A.Row_count * mat_A.Row_count) {
		// Находим элемент по модулю больший преграды
		p, q, err := mat_A.findGreaterThanBarrier(ak)
		check(err)
		// Находим синус и косинус
		sin, cos := mat_A.calculateSinAndCos(p, q)
		// Находим следующую матрицу A[k]
		mat_A.calculateNextMatrix(p, q, sin, cos)
		for i := range mat_T.Array {
			z3 := mat_T.Array[i][p]
			z4 := mat_T.Array[i][q]
			mat_T.Array[i][q] = z3 * sin + z4 * cos
			mat_T.Array[i][p] = z3 * cos - z4 * sin
		}
	}
	return mat_A, mat_T
}

// Шаг 2
func (mat Matrix) calculateBarrier() float64 {
	barrier := 0.0
	// Cуммируем внедиагональные элементы матрицы
	for i := 0; i < mat.Row_count - 1; i++ {
		for j := i + 1; j < mat.Column_count; j++ {
			barrier += mat.Array[i][j] * mat.Array[i][j]
		}
	}
	return math.Sqrt(barrier * 2.0) / float64(mat.Row_count)
}

// Шаг 3
// Функция возвращает номер строки и номер столбца найденного элемента, если 
// поиск дал результаты, и выдаёт ошибку в противном случае
func (mat Matrix) findGreaterThanBarrier(barrier float64) (int, int, error) {
	max := barrier
	p, q := -1, -1
	var err error = nil
	for i := range mat.Array {
		for j, value := range mat.Array[i] {
			if i == j {
				continue
			}
			absolute_val := math.Abs(value)
			if absolute_val > max {
				max = absolute_val
				p, q = i, j
			}
		}
	}
	if p == -1 && q == -1 {
		err = errors.New("Не найден элемент больший текущей преграды")
	}
	return p, q, err
}

// Шаг 4
func (mat Matrix) calculateSinAndCos(p, q int) (sin, cos float64) {
	y := (mat.Array[p][p] - mat.Array[q][q]) / 2.0
	var x, sign float64
	// Знак y
	if y < 0.0 {
		sign = -1.0
	} else {
		sign = 1.0
	}

	if y == 0.0 {
		x = -1.0
	} else {
		x = -sign * mat.Array[p][q] / math.Sqrt(mat.Array[p][q] * mat.Array[p][q] + y * y)
	}
	sin = x / math.Sqrt(2.0 * (1.0 + math.Sqrt(1.0 - x * x)))
	cos = math.Sqrt(1.0 - sin * sin)
	return
}

// Продолжние шага 4
func (mat *Matrix) calculateNextMatrix(p, q int, sin, cos float64) {
	// Преобразуем матрицу
	for i := range mat.Array {
		if i != p && i != q {
			z1 := mat.Array[i][p]
			z2 := mat.Array[i][q]
			mat.Array[q][i] = z1 * sin + z2 * cos
			mat.Array[i][q] = mat.Array[q][i]
			mat.Array[i][p] = z1 * cos - z2 * sin
			mat.Array[p][i] = mat.Array[i][p]
		}
	}
	z5 := sin * sin
	z6 := cos * cos
	z7 := sin * cos
	v1 := mat.Array[p][p]
	v2 := mat.Array[p][q]
	v3 := mat.Array[q][q]

	mat.Array[p][p] = v1 * z6 + v3 * z5 - 2.0 * v2 * z7
	mat.Array[q][q] = v1 * z5 + v3 * z6 + 2.0 * v2 * z7
	mat.Array[p][q] = (v1 - v3) * z7 + v2 * (z6 - z5)
	mat.Array[q][p] = mat.Array[p][q]
}

// Шаг 5
func (mat Matrix) checkAllElementsLessThanEpsBarrier(eps, barrier float64) bool {
	for i := 0; i < mat.Row_count - 1; i++ {
		for j := i + 1; j < mat.Column_count; j++ {
			if mat.Array[i][j] > eps * barrier {
				return false
			}
		}
	}
	return true
}


// Сортировка собственных значений и собственных векторов по убыванию
func SortEigenMatrices(A, T Matrix) (Vector, []Vector) {
	A.checkSimmetry()
	T.checkSquareness()

	eigenvalues := A.GetDiagonal()
	eigenvectors := T.ConvertToVec()

	for step := eigenvalues.Size / 2; step > 0; step /= 2 {
		for i := step; i < eigenvalues.Size; i++ {
			for j := i - step; j >= 0 && eigenvalues.Array[j] < eigenvalues.Array[j + step]; j-= step {
				temp_val, temp_vec := eigenvalues.Array[j], eigenvectors[j]
				eigenvalues.Array[j], eigenvectors[j] = eigenvalues.Array[j + step], eigenvectors[j + step]
				eigenvalues.Array[j + step] , eigenvectors[j + step] = temp_val, temp_vec
			}
		}
	}
	for i := 1; i < len(eigenvectors); i += 2 {
		eigenvectors[i].MulScalar(-1.0)
	}
	return eigenvalues, eigenvectors
}