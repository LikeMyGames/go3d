package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Mesh struct {
		File         string
		Name         string
		faces        []Face
		Location     Vector3D
		CenterOffset Vector3D
		verticies    []Vertex
		normals      []Vector3D
	}

	Vertex struct {
		XYZ   Vector3D
		W     *float64
		Color *Vector3D
	}

	Face struct {
		Verticies []uint
		Normal    *Vector3D
	}

	Vector3D struct {
		X float64
		Y float64
		Z float64
	}

	Command struct {
		name string
		data []uint8
	}
)

func (mesh *Mesh) GetFaces() []Face {
	return mesh.faces
}

func (cmd Command) GetName() string {
	return cmd.name
}

func (cmd Command) GetData() []uint8 {
	return cmd.data
}

func main() {
	fmt.Println(Parse("test.obj"))
}

func Parse(filepath string) Mesh {
	mesh := Mesh{
		File: filepath,
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	lines := [][][]uint8{}
	for _, v := range data {
		// fmt.Println("reading byte ", i+1)
		if len(lines) == 0 {
			lines = append(lines, [][]uint8{})
			fmt.Println("added starting line")
		}
		if len(lines[len(lines)-1]) == 0 {
			lines[len(lines)-1] = append(lines[len(lines)-1], []uint8{})
			fmt.Println("added first token in line")
		}
		switch v {
		case 10:
			lines = append(lines, [][]uint8{})
			fmt.Println("added new line after ", len(lines))
		case 32:
			lines[len(lines)-1] = append(lines[len(lines)-1], []uint8{})
			fmt.Println("added new token in line ", len(lines))
		default:
			lines[len(lines)-1][len(lines[len(lines)-1])-1] = append(lines[len(lines)-1][len(lines[len(lines)-1])-1], v)
		}
	}

	for i, v := range lines {
		fmt.Println(v)
		if len(v) == 0 {
			continue
		}
		if string(v[0]) == "#" {
			continue
		}
		switch string(v[0]) {
		case "v":
			// parsing of "v" or "vertex" commands from the obj file
			fmt.Println("verticies loaded into mesh: ", len(mesh.verticies))
			x, err := strconv.ParseFloat(string(v[1]), 64)
			if err != nil {
				log.Fatal(err)
			}

			y, err := strconv.ParseFloat(string(v[2]), 64)
			if err != nil {
				log.Fatal(err)
			}

			z, err := strconv.ParseFloat(string(v[3]), 64)
			if err != nil {
				log.Fatal(err)
			}
			switch len(v) - 1 {
			case 3:
				mesh.verticies = append(mesh.verticies, Vertex{
					XYZ: Vector3D{
						X: x,
						Y: y,
						Z: z,
					},
				})
			case 4:
				w, err := strconv.ParseFloat(string(v[4]), 64)
				if err != nil {
					log.Fatal(err)
				}
				mesh.verticies = append(mesh.verticies, Vertex{
					XYZ: Vector3D{
						X: x,
						Y: y,
						Z: z,
					},
					W: &w,
				})
			case 6:
				r, err := strconv.ParseFloat(string(v[5]), 64)
				if err != nil {
					log.Fatal(err)
				}

				g, err := strconv.ParseFloat(string(v[6]), 64)
				if err != nil {
					log.Fatal(err)
				}

				b, err := strconv.ParseFloat(string(v[7]), 64)
				if err != nil {
					log.Fatal(err)
				}
				mesh.verticies = append(mesh.verticies, Vertex{
					XYZ: Vector3D{
						X: x,
						Y: y,
						Z: z,
					},
					Color: &Vector3D{
						X: r,
						Y: g,
						Z: b,
					},
				})
			case 7:
				r, err := strconv.ParseFloat(string(v[0]), 64)
				if err != nil {
					log.Fatal(err)
				}

				g, err := strconv.ParseFloat(string(v[1]), 64)
				if err != nil {
					log.Fatal(err)
				}

				b, err := strconv.ParseFloat(string(v[2]), 64)
				if err != nil {
					log.Fatal(err)
				}
				w, err := strconv.ParseFloat(string(v[3]), 64)
				if err != nil {
					log.Fatal(err)
				}
				mesh.verticies = append(mesh.verticies, Vertex{
					XYZ: Vector3D{
						X: x,
						Y: y,
						Z: z,
					},
					W: &w,
					Color: &Vector3D{
						X: r,
						Y: g,
						Z: b,
					},
				})
			}
		case "f":
			// parsing of "f" or "face" commands from the obj file
			fmt.Println("faces loaded into mesh: ", len(mesh.faces))
			verts := []uint{}
			for _, v := range v[1:] {
				tokens := strings.Split(string(v), "/")

				i, err := strconv.ParseInt(string(tokens[0]), 10, 0)
				if err != nil {
					log.Fatal(err)
				}
				verts = append(verts, uint(i))
			}
			mesh.faces = append(mesh.faces, Face{
				Verticies: verts,
			})
			// switch len(v[1]) {
			// case 1:
			// 	// verticies only
			// 	verts := []uint{}
			// 	for _, v := range v[1:] {
			// 		i, err := strconv.ParseInt(string(v[1]), 10, 0)
			// 		if err != nil {
			// 			log.Fatal(err)
			// 		}
			// 		verts = append(verts, uint(i))
			// 	}
			// 	mesh.faces = append(mesh.faces, Face{
			// 		Verticies: verts,
			// 	})
			// case 3:
			// 	// verticies and vertex texture coordinates
			// 	// if len(mesh.normals) == 0 && len(mesh.{

			// 	// } else {
			// 	verts := []uint{}
			// 	norms := []Vector3D{}
			// 	for _, v := range v[1:] {
			// 		tokens := strings.Split(string(v), "/")

			// 		i, err := strconv.ParseInt(string(tokens[0]), 10, 0)
			// 		if err != nil {
			// 			log.Fatal(err)
			// 		}
			// 		verts = append(verts, uint(i))

			// 		i, err = strconv.ParseInt(string(tokens[1]), 10, 0)
			// 		if err != nil {
			// 			log.Fatal(err)
			// 		}
			// 		norms = append(norms, mesh.normals[uint(i)])
			// 	}
			// 	nv := AddVector(norms...)
			// 	mesh.faces = append(mesh.faces, Face{
			// 		Verticies: verts,
			// 		Normal:    &nv,
			// 	})
			// 	// }
			// case 5:
			// 	// verticies, vertex texture coordinates, and vector normals
			// }
		}
		fmt.Println("lines loaded into mesh: ", i)
	}

	return mesh
}

func AddVector(vect ...Vector3D) Vector3D {
	sum := Vector3D{}
	for _, v := range vect {
		sum.X += v.X
		sum.Y += v.Y
		sum.Z += v.Z
	}
	return sum
}
