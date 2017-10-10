package graph

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
  X float64
  Y float64
  Z float64
}

type Graph struct {
  Edges []float64
  Vertices []Point
  NV int//number of vertices in the graph
  NE int//number of edges in the graph
  Name string
}


func (self *Graph) LoadFromFile(name string, pointsFile string, edgeFile string) error {
  f, _ := os.Open(pointsFile)
  defer f.Close()
  r := csv.NewReader(f)
  points,err := r.ReadAll()
  if err != nil {
    return err
  }
  vertices := make([]float64,0,len(points)*len(points[0]))
  for i := range points[0] {
    x,_ := strconv.ParseFloat(points[0][i],64)
    vertices = append(vertices,x)
    y,_ := strconv.ParseFloat(points[1][i],64)
    vertices = append(vertices,y)
    z,_ := strconv.ParseFloat(points[2][i],64)
    vertices = append(vertices,z)
  }
  self.NV = len(points[0])
  self.Vertices = make([]Point, self.NV)
  for i:=0; i<self.NV; i++ {
    self.Vertices[i] = Point{X:vertices[i*3],Y:vertices[i*3+1],Z:vertices[i*3+2]}
  }
  fmt.Println("length of point array: ",len(self.Vertices))

  f1,_ := os.Open(edgeFile)
  defer f.Close()
  r1 := csv.NewReader(f1)
  edges, err1 := r1.ReadAll()
  if err1 != nil {
    return err1
  }
  self.Edges = make([]float64,0,len(edges)*len(edges[0]))
  for i := range edges[0] {
    u,_ := strconv.ParseFloat(edges[0][i],64)
    self.Edges = append(self.Edges,u)
    v,_ := strconv.ParseFloat(edges[1][i],64)
    self.Edges = append(self.Edges,v)
  }
  self.NE = len(edges[0])
  fmt.Println("length of edges array: ", len(self.Edges))

  self.Name = name
  return nil
}

