package main

type ListNode struct {
     Val int
     Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	var (
		min **ListNode
		head *ListNode
		cur *ListNode
	)

	for {
		min = &lists[0]
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if *min == nil || (*min).Val > lists[i].Val {
				min = &lists[i]
			}
		}
		if *min == nil {
			break
		}
		if head == nil {
			head = *min
			cur = head
		} else {
			cur.Next = *min
			cur = cur.Next
		}
		*min = (*min).Next
	}

	return head
}

func main() {

}
