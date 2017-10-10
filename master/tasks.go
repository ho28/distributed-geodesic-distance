package master

import (
    "container/list"
)

type Tasks struct {
    notAssigned *Stack
    inProgress *list.List
    elements *[]*list.Element
}

func StoreTasks (tasks []int) Tasks {

    tasksLen := len(tasks)
    stack := CreateStack(tasksLen)
    for _, task := range tasks {
        stack.Push(task)
    }

    l := list.New()
    elements := make([] *list.Element, tasksLen + 1)

    return Tasks{stack, l, &elements}
}

func (self *Tasks) RemoveTaskFromInProgress (task int) {
    e := (*self.elements)[task]
    self.inProgress.Remove(e)
}

func (self *Tasks) FinishedTask (task int) {
    self.RemoveTaskFromInProgress(task)
}

func (self *Tasks) NextTask() int {

    if self.notAssigned.Length() > 0 {

        // Move from notAssigned to inProgress
        task := self.notAssigned.Pop()
        self.inProgress.PushBack(task)

        // Set element location for quick removal
        (*self.elements)[task] = self.inProgress.Back()

        return task

    } else if self.inProgress.Len() > 0 {

        lastInserted := self.inProgress.Back()
        NextTaskE := self.inProgress.Front()

        self.inProgress.MoveAfter(NextTaskE, lastInserted)

        return NextTaskE.Value.(int)

    } else {
        return -1
    }
}

// If a node did not complete a task, move it back to available tasks
func (self *Tasks) AddBack(task int) {

    if self.notAssigned.Length() > 0 {
        self.notAssigned.Push(task)
        self.RemoveTaskFromInProgress(task)
    }

}



