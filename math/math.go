package math

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"errors"
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