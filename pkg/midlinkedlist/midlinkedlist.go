package midlinkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	for l != nil {
		s += fmt.Sprintf("%d ", l.Val)
		l = l.Next
	}
	return s
}

func buildNodes(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	result := ListNode{
		Val: nums[0],
	}
	head := &result
	for _, num := range nums[1:] {
		head.Next = &ListNode{
			Val: num,
		}
		head = head.Next
	}

	return &result
}

func Run() {
	fmt.Println("[main] midlinkedlist")
	nodeInts := []int{1, 2, 3, 4, 5}
	fmt.Println(MiddleNode(buildNodes(nodeInts)))
}

func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func BuildAndSearchMiddleNode(nums []int) *ListNode {
	return MiddleNode(buildNodes(nums))
}
