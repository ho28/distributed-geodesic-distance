package master

import (
    "errors"
)

type stack interface {

    Push(node int) error

    Pop() int

    Top() int

    Length() int

    Create(maxLength int)
}

type Stack struct {

    nodes *[]int
    maxLength int
    nextAvailableIndex *int
}

// stack := CreateStack(10)
func CreateStack(maxLength int) *Stack {
    nodes := make([] int, maxLength)
    nI := 0
    stack := Stack{&nodes, maxLength, &nI}
    return &stack
}

func (s Stack) Push (node int) error {

    if (*s.nextAvailableIndex) < s.maxLength {

        (*s.nodes)[(*s.nextAvailableIndex)] = node
        (*s.nextAvailableIndex)++
        return nil

    } else {
        return errors.New("could not push to stack")
    }
}

func (s Stack) Pop () int {

    if (*s.nextAvailableIndex) > 0 {

        (*s.nextAvailableIndex)--
        temp := (*s.nodes)[(*s.nextAvailableIndex)]
        (*s.nodes)[(*s.nextAvailableIndex)] = 0
        return  temp

    } else {
        return -1

    }
}

func (s Stack) Top () int {

    if (*s.nextAvailableIndex) > 0 {
        return (*s.nodes)[(*s.nextAvailableIndex) - 1]
    } else {
        return -1
    }
}

func (s Stack) Length () int {
    return (*s.nextAvailableIndex)
}


