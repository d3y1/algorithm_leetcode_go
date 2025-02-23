// LC141 环形链表
// @author d3y1

package easy

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    // return hasCycle1(head)
    return hasCycle2(head)
}

// 哈希
func hasCycle1(head *ListNode) bool {
    existed := map[*ListNode]struct{}{}
    for head != nil {
        if _, ok := existed[head]; ok {
            return true
        }else{
            existed[head] = struct{}{}
            head = head.Next
        }
    }
    return false
}

// 快慢指针
func hasCycle2(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head
    fast := head
    for fast!=nil && fast.Next!=nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            return true
        }
    }

    return false
}