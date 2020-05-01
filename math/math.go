package math

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"errors"
	"math"
)

// Структура матрицы
type Matrix struct {
	Array [][]float64
	Row_count, Column_count int
}

// Структура вектора
type Vector struct {
	Array []float64
	Size int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Инициализация матрицы
func InitMatrix() Matrix {
	return Matrix{ Array: nil, Row_count: 0, Column_count: 0 }
}

// Инициализация вектора
func InitVector() Vector {
	return Vector{ Array: nil, Size: 0 }
}

func (vec *Vector) New(size int) {
	vec.Array = make([]float64, size)
	vec.Size = size
}

func (mat *Matrix) New(row_count, column_count int) {
	// Выделяем память под матрицу
	mat.Array = make([][]float64, row_count)
	for i := range mat.Array {
		mat.Array[i] = make([]float64, column_count)
	}
	mat.Row_count = row_count
	mat.Column_count = column_count
}

// Единичная матрица
func Identity(mat_size int) Matrix {
	mat := InitMatrix()
	mat.New(mat_size, mat_size)
	for i := range mat.Array {
		for j := range mat.Array[i] {
			if i == j {
				mat.Array[i][j] = 1.0
			}
		}
	}
	return mat
}

// Проверить матрицу на "квадратность"
// Если матрица не квадратная, то выдать ошибку
func (mat Matrix) checkSquareness() {
	if mat.Column_count != mat.Row_count {
		check(errors.New("Матрица должна быть квадратной!"))
	} 
}

// Проверка симметричности матрицы
func (mat Matrix) IsSimmetric() bool {
	mat.checkSquareness()
	for i := 0; i < mat.Row_count - 1; i++ {
		for j := i + 1; j < mat.Column_count; j++ {
			if mat.Array[i][j] != mat.Array[j][i] {
				return false
			} 
		}
	}
	return true
}

// Проверить симметричность матрицы
// Если матрица не симметрична, то выдать ошибку
func (mat Matrix) checkSimmetry() {
	if !mat.IsSimmetric() {
		check(errors.New("Матрица не симметрична"))
	}
}

// Чтение из файла
func (mat *Matrix) Read(filepath string) {
	// Открываем файл
	b, err := ioutil.ReadFile(filepath)
	check(err)

	// Построчно делим текст
	lines := strings.Split(string(b), "\r\n")
	
	var Rows, Columns int
	Rows = len(lines)

	// Создаём срез срезов и задаём ему размер (количество строк)
	array := make([][]float64, Rows)

	// Проходимся по каждой строке
	for i, line := range lines {
		// Делим каждую строку на отдельные числовые строки
		string_nums := strings.Split(string(line), " ")
		Columns = len(string_nums)

		// Задаём размер срезу (количество столбцов)
		array[i] = make([]float64, Columns)

		// Проходимся по каждой строке с числом
		for j, number := range string_nums {

			// Конвертируем строку в число
			array[i][j], err = strconv.ParseFloat(number, 64)
			check(err)
		}
	}
	mat.Array = array
	mat.Row_count = Rows
	mat.Column_count = Columns
}

// Вывод содержимого матрицы
func (mat Matrix) Write() {
	fmt.Println("[")
	for i := range mat.Array {
		fmt.Print("\t[ ")
		for _, value := range mat.Array[i] {
			fmt.Print(value, " ")
		}
		fmt.Println("]")
	}
	fmt.Println("]")
}

// Вывод содержимого вектора
func (vec Vector) Write() {
	fmt.Println("[")
	for _, v := range vec.Array {
		fmt.Println("\t", v)
	}
	fmt.Println("]")
}

// Транспонирование матрицы
func (mat *Matrix) Transpose() {
	
	new_row_count := mat.Column_count
	new_column_count := mat.Row_count

	// Создание временной матрицы для хранения промежуточных данных
	temp_array := make([][]float64, new_row_count)
	for i := range temp_array {
		temp_array[i] = make([]float64, new_column_count)  
	}

	for i := 0; i < new_row_count; i++ {
		for j := 0; j < new_column_count; j++ {
			temp_array[i][j] = mat.Array[j][i]
		}
	}

	mat.Array = temp_array
	mat.Row_count = new_row_count
	mat.Column_count = new_column_count
}

// Умножение матриц
func (first *Matrix) Mul(second Matrix) {
	
	if first.Row_count != second.Column_count {
		check(errors.New("Количество строк первой матрицы и количество столбцов второй матрицы не совпадают!"))
	}

	// Временная матрица
	result := InitMatrix()
	result.Array = make([][]float64, first.Row_count)
	for i := range result.Array {
		result.Array[i] = make([]float64, second.Column_count)
	}

	for i := 0; i < first.Row_count; i++ {
		for j := 0; j < second.Column_count; j++ {
			var accum float64
			for k := 0; k < second.Row_count; k++ {
				accum += first.Array[i][k] * second.Array[k][j]
			}
			result.Array[i][j] = accum
		}
	}
	first = &result
}

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
func (mat Matrix) MakeCovariation() Matrix {
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
	fmt.Println("Начальная преграда -", a0)
	for ak := a0; !mat_A.checkAllElementsLessThanEpsBarrier(eps, a0); ak /= float64(mat_A.Row_count * mat_A.Row_count) {
		// Находим элемент по модулю больший преграды
		p, q, err := mat_A.findGreaterThanBarrier(ak)
		if err != nil {
			continue
		}
		// Находим синус и косинус
		sin, cos := mat_A.calculateSinAndCos(p, q)
		// Находим матрицу следующую A[k]
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
	for j := 1; j < mat.Column_count; j++ {
		for i := 0; i < j; i++ {
			barrier += mat.Array[i][j] * mat.Array[i][j]
		}
	}
	return math.Sqrt(barrier * 2.0) / float64(mat.Row_count)
}

// Шаг 3
// Функция возвращает номер строки и номер столбца найденного элемента, если 
// поиск дал результаты, и выдаёт ошибку в противном случае
func (mat Matrix) findGreaterThanBarrier(barrier float64) (int, int, error) {
	for j := 1; j < mat.Column_count; j++ {
		for i := 0; i < j; j++ {
			if math.Abs(mat.Array[i][j]) > barrier {
				return i, j, nil
			}
		}
	}
	return -1, -1, errors.New("Не удалось найти наибольший по модулю вне диагональный элемент, превосходящий преграду")
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
