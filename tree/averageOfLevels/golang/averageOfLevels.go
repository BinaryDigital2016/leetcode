package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	ret := make([]float64, 0)
	if root == nil {
		return ret
	}
	q := Queue{make([]*TreeNode, 0)}
	q.Append(root)
	for !q.Empty() {
		n := q.Size()
		sum := 0
		for i := 0; i < n; i++ {
			cur := q.Pop()
			sum += cur.Val
			if cur.Left != nil {
				q.Append(cur.Left)
			}
			if cur.Right != nil {
				q.Append(cur.Right)
			}
		}
		ret = append(ret, float64(sum)/float64(n))
	}
	return ret
}

type Queue struct {
	elem []*TreeNode
}

func (q *Queue) Append(e *TreeNode) {
	if q.elem == nil {
		q.elem = make([]*TreeNode, 0)
	}
	q.elem = append(q.elem, e)
}

func (q *Queue) Pop() *TreeNode {
	if q.elem == nil || len(q.elem) == 0 {
		return nil
	}

	t := q.elem[0]
	q.elem = q.elem[1:]
	return t
}

func (q *Queue) Empty() bool {
	return len(q.elem) == 0
}

func (q *Queue) Size() int {
	return len(q.elem)
}
