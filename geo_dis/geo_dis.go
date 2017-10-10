package geo_dis

import (
  "math"
  "fmt"
  "proj/graph"
)

const (
  //maxn = 1000002
  //maxn = 500001
  maxn = 250000
  //maxm = 80000002
  //maxm = 40000001
  maxm = 20000000
  //limit = 1000000000
  //limit = 500000000
  limit = 250000000
)

type dijkstraHeap struct {
  g []int
  num []int
  next []int
  w []int
  where []int
  tt int
  v []float64
  d []float64
}

func (self *dijkstraHeap) clear() {
  for i := range self.g {
    self.g[i] = 0
  }
}

func (self *dijkstraHeap) add(a int, b int, c float64) {
  self.tt = self.tt+1
  self.num[self.tt] = b
  self.v[self.tt] = c
  self.next[self.tt] = self.g[a]
  self.g[a] = self.tt
}

func (self *dijkstraHeap) add2(a int, b int, c float64) {
  self.add(a,b,c)
  self.add(b,a,c)
}

func (self *dijkstraHeap) chg(a int, b int) {
  tmp := self.w[a]
  self.w[a] = self.w[b]
  self.w[b] = tmp
  self.where[self.w[a]] = a
  self.where[self.w[b]] = b
}

func (self *dijkstraHeap) up(r int) {
  for (r>1 && self.d[self.w[r]] < self.d[self.w[r>>1]]) {
    self.chg(r, r>>1)
    r = r>>1
  }
}

func (self *dijkstraHeap) down(r int, l int){
  j := r+r
  for j<=l {
    if (j<l&&self.d[self.w[j+1]]<self.d[self.w[j]]) {
      j++
    }
    if (self.d[self.w[j]]>=self.d[self.w[r]]) {
      break
    }
    self.chg(r,j)
    r = j
    j = r+r
  }
}

func (self *dijkstraHeap) get(s int, t int, n int) float64 {
  i := n+1
  k := n+1
  l := n+1
  for i:=0; i<=n; i++ {
    self.d[i] = limit
  }
  self.d[s] = 0
  for i:=0; i<=n; i++ {
    self.w[i+1]=i
    self.where[self.w[i+1]]=i+1
  }
  self.up(self.where[s])
  for (self.w[1] != t) && (l!=0) {
    k = self.w[1]
    self.chg(1,l)
    l = l-1
    self.down(1,l)
    for i=self.g[k];i!=0;i=self.next[i] {
      if (self.d[k]+self.v[i])<(self.d[self.num[i]]) {
        self.d[self.num[i]]=self.d[k]+self.v[i]
        self.up(self.where[self.num[i]])
      }
    }
  }
  if t < 0 {
    return self.d[0]//ALERT!!ALERT!! MODIFIED FROM ORIGINAL!! return self.d[t]
  }
  return self.d[t]
}

func newDijkstraHeap() *dijkstraHeap{
  g := new(dijkstraHeap)
  g.g = make([]int, maxn)
  g.num = make([]int, maxm*2)
  g.next = make([]int,maxm*2)
  g.w = make([]int,maxn)
  g.where = make([]int,maxn)
  g.tt = 0
  g.v = make([]float64,maxm*2)//!!!DONO IF THIS IS OK VALUE TO USE AS SIZE
  g.d = make([]float64,maxm*2)//SAME HERE
  return g
}

func round(f float64) int {
  return int(math.Floor(f + 0.5))
}

func dis(p graph.Point, q graph.Point) float64{
  return math.Sqrt((p.X-q.X)*(p.X-q.X)+(p.Y-q.Y)*(p.Y-q.Y)+(p.Z-q.Z)*(p.Z-q.Z))
}

func Geo_Dis(vertex []graph.Point, edge []float64, index int, n int, m int, killChan <-chan bool) ([]float64,error) {
  //vertexp is the point matrix in single column form matrix(:)
  //edge is the edge4geo matrix in single column form edge4geo(:)
  //if the csv is loaded as a single dimensional array, you can find n and m by dividing len(point) by 3 and len(edge) by 2
  //n is the # of columns in point, m is the # of columns in edge4geo
  //k1 := 1
  select {
    case <-killChan:
    return nil, fmt.Errorf("killed")

    default:
  k2 := n
  ind1 := float64(index)
  //ind2 := float64(n)//in original this is a range 1:n
  /*
  vertex := make([]point, n)
  for i:=0; i<n; i++ {
    vertex[i] = point{x:vertexp[i*3],y:vertexp[i*3+1],z:vertexp[i*3+2]}
  }
  *///vertex[] creation moved outside function to avoid unnecessary recreation
  g := newDijkstraHeap()
  for i:=0; i<m; i++ {
    //fmt.Printf("edge %d\n",i)
    //a := round(edge[i+i])
    //b := round(edge[i+i+1])
    //fmt.Printf("indexing into vertex\n")
    //fmt.Printf("indx 1: %d\n", int(edge[i+i]-0.5))
    //fmt.Printf("indx 1: %d\n", int(edge[i+i+1]-0.5))
    //c := dis(vertex[int(edge[i+i]-0.5)],vertex[int(edge[i+i+1]-0.5)])
    //g.add2(a,b,c)
        g.add2(round(edge[i+i]),round(edge[i+i+1]),dis(vertex[int(edge[i+i]-0.5)],vertex[int(edge[i+i+1]-0.5)]))
  }
  out := make([]float64,n)
  g.get(round(ind1),-1,n)
  for j:=0; j<k2; j++ {
        out[j] = g.d[j+1]
  }
  g = nil//make eligible for garbage collection
  return out,nil

  }

}
