package main
import "fmt"
import "strings"
import "strconv"
import "bytes"

type ListNode struct {
     Val int
     Next *ListNode
}

func str2ListNode(input string) *ListNode {
    intVals := strings.Split(input, ", ")
    head := &ListNode{0, nil}
    cur := head
    for i:=0; i < len(intVals); i++ {
        intVal,_ := strconv.Atoi(intVals[i])
        cur.Next = &ListNode{intVal, nil}
        cur = cur.Next
    }
    return head.Next
}

func listNode2Str(t *ListNode) string {
    if t == nil {
        return "[]"
    }

    var buf bytes.Buffer
    for {
        if t == nil {
            ret := buf.String()
            return "[" + ret[:len(ret)-2] + "]"
        }
        buf.WriteString(strconv.Itoa(t.Val))
        buf.WriteString(", ")
        t = t.Next
    }
    ret := buf.String()
    return "[" + ret[:len(ret)-2] + "]"
}

func add2Num(l1 *ListNode, l2 *ListNode) *ListNode{
    head := &ListNode{0, nil}
    cur := head
    carry := 0
    for {
        if l1 != nil || l2 != nil {
            var v1, v2 int
            if l1 != nil {
                v1 = l1.Val
            } else {
                v1 = 0
            }
            if l2 != nil {
                v2 = l2.Val
            } else {
                v2 = 0
            }

            sum := carry + v1 + v2
            carry = sum / 10
            ln := &ListNode{sum % 10, nil}
            cur.Next = ln
            cur = ln

            if l1 != nil {
                l1 = l1.Next
            } else {
                l1 = nil
            }
            if l2 != nil {
                l2 = l2.Next
            } else {
                l2 = nil
            }
        } else {
            break
        }
    }
    if carry > 0 {
        ln := &ListNode{carry, nil}
        cur.Next = ln
    }
    return head.Next
}

func main() {
    l1 := "2, 4, 3"
    l2 := "5, 6, 4"
    fmt.Print(l1, "\n")
    fmt.Print(l2, "\n")

    ret := listNode2Str(add2Num(str2ListNode(l1), str2ListNode(l2)))
    fmt.Print(ret, "\n")
}
