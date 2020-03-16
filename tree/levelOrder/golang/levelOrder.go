package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	q := Queue{make([]*TreeNode, 0)}
	q.Push(root)
	for !q.Empty() {
		n := q.Size()
		t := make([]int, 0)
		for i := 0; i < n; i++ { //确定同层节点
			cur := q.Pop()
			t = append(t, cur.Val)
			if cur.Left != nil {
				q.Push(cur.Left)
			}
			if cur.Right != nil {
				q.Push(cur.Right)
			}
		}
		ret = append(ret, t)
	}
	return ret
}

type Queue struct {
	queue []*TreeNode
}

func (q *Queue) Push(x *TreeNode) {
	if q.queue == nil {
		q.queue = make([]*TreeNode, 0)
	}
	q.queue = append(q.queue, x)
}

func (q *Queue) Pop() *TreeNode {
	if len(q.queue) <= 0 {
		return nil
	}

	t := q.queue[0]
	if len(q.queue) == 1 {
		q.queue = nil
	} else {
		q.queue = q.queue[1:]
	}

	return t
}

func (q *Queue) Empty() bool {
	return q.queue == nil || len(q.queue) == 0
}

func (q *Queue) Size() int {
	if q.queue == nil {
		return 0
	}
	return len(q.queue)
}
