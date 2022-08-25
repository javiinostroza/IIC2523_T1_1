package main

import (
	"fmt"
    "io/ioutil"
	"os"
	str "strings"
	conv "strconv"
)

var rows int
var cols int
var file_id string


func check(e error) { // Based on https://gobyexample.com/writing-files
    if e != nil {
        panic(e)
    }
}

func numIslands(grid [][]int) string {
	/* 
	Based on 
	https://leetcode.com/problems/number-of-islands/discuss/2446300/Number-of-Islands-oror-Python3-oror-DFS 
	*/
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ { // Armamos una matriz para guardar los lugares visitados
		visited[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ { // Llenamos con false todos los lugares
		for j := 0; j < cols; j++ {
			visited[i][j] = false
		}
	}
	var islands = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ { // 
			if grid[i][j] == 1 && visited[i][j] == false {
				islands = islands + 1
				dfs(i, j, grid, visited)
			}
		}
	}
	return conv.Itoa(islands)
}

func dfs(row int, col int, grid [][]int, visited [][]bool) {
	/* 
	Based on 
	https://leetcode.com/problems/number-of-islands/discuss/2446300/Number-of-Islands-oror-Python3-oror-DFS 
	*/
	visited[row][col] = true
	var positions = [4][2]int{ {-1, 0}, {0, 1}, {1, 0}, {0, -1} }
	for i := 0; i < len(positions); i++ {
		var nx = positions[i][0] + row
		var ny = positions[i][1] + col
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || visited[nx][ny] == true || grid[nx][ny] != 1 {
			continue
		}
		dfs(nx, ny, grid, visited)
	}
}

func read_file(name string) string{
	/*
	Lee el mapa en el archivo *.txt y retorna un string de sólo números 1s y 0s.
	*/
	map_, err := ioutil.ReadFile(name) // Lectura del mapa en el archivo .txt
	check(err)
	var clean_map = str.Replace(string(map_), ",", "", -1) // Elimina ',' del string
	var map_name = str.Replace(name, ".txt", "", -1) // Elimina .txt del nombre del archivo
	var splitted_map_name = str.Split(map_name, "_") // Separa el nombre del archivo por '_'
	intCols, err := conv.Atoi(splitted_map_name[len(splitted_map_name) - 1]) // Toma la cantidad de columnas del mapa
	file_id = splitted_map_name[1] // Toma el id del mapa
	cols = intCols // Seteamos la cantidad de columnas en variable global
	rows = len(clean_map) / intCols // Seteamos la cantidad de filas en variable global
	return clean_map
}

func build_matrix(map_info string) [][]int { 
	/*
	Recibe un string de 1s y 0s y retorna un array 2D de enteros de tamaño rowXcol,
	que representa el mapa como matriz
	*/
	var i_row = 0
	var i_col = 0
	arr_map := make([][]int, rows)
	for i := 0; i < rows; i++ { // Armamos la matriz
		arr_map[i] = make([]int, cols)
	}
	for i := 0; i < len(map_info); i++ { // Llenamos la matriz
		i_, err := conv.Atoi(string(map_info[i]))
		check(err)
		arr_map[i_row][i_col] = i_
		i_col = i_col + 1
		if i_col == cols {
			i_col = 0
			i_row = i_row + 1
		}
	}
	return arr_map
}

func write_file(id_, ans string) {
	/*
	Esta función escribe la respuesta en un archivo dentro de la carpeta 'respuestas'
	Se asume que la carpeta ya existe en el directorio.
	*/
	var base_name = "respuestas/numberOfISlands_" + id_ + ".txt"
	f, err := os.Create(base_name)
	check(err)
	defer f.Close()
	f.WriteString(ans)
	check(err)
}

func main() {
	matrices_files, err := ioutil.ReadDir("matrices")
	check(err)
	if len(matrices_files) == 0 {
		fmt.Println("La carpeta 'matrices' se encuentra vacía!")
	}
	for _, file := range matrices_files {
		fileName := "matrices/" + file.Name()
		var file_data = read_file(fileName)
		var map_matrix = build_matrix(file_data)
		write_file(file_id, numIslands(map_matrix))
    }
	fmt.Println("Ejecución terminada")
}

