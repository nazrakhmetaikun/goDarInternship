package main
import (
    "fmt"
)
type node struct {
    data int
    next  *node
}
var head *node = nil
func (n *node) add(data int) *node {
    if head == nil {
        n.data = data
        n.next = nil
        head = n
        return n
    } else {
        for n.next != nil {
            if n.data == data {
                return n
            }
            n = n.next
        }
        n.next = new(node)
        n.next.data = data
        n.next.next = nil
        return n
    }
}
func (n *node) delete() *node {
    if head == nil {
        return n
    }
    tmp := new(node)
    tmp = n
    for tmp.next.next != nil {
        tmp = tmp.next
    }
    tmp.next = nil
    return n
}
func (n *node) display() {
    for head.next != nil {
        fmt.Println(head.data)
        head = head.next
    }
    fmt.Println(head.data)
}
func main() {
    val := new(node)
	val.add(22)
	val.add(25)
	val.delete()
	val.display()
}