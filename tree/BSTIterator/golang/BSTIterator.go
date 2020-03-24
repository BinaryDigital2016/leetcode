package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

调用 next() 将返回二叉搜索树中的下一个最小的数。

提示：

    next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
    你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
*/

type BSTIterator struct {
	s MyStack
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
		MyStack{make([]*TreeNode, 0)},
	}
	left_inorder(root, &iter.s)
	return iter
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	cur := this.s.Pop()
	if cur.Right != nil {
		left_inorder(cur.Right, &this.s)
	}
	return cur.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return !this.s.Empty()
}

func left_inorder(root *TreeNode, s *MyStack) {
	for root != nil {
		s.Push(root)
		root = root.Left
	}
}

type MyStack struct {
	e []*TreeNode
}

func (this *MyStack) Push(x *TreeNode) {
	this.e = append(this.e, x)
}

func (this *MyStack) Pop() *TreeNode {
	n := len(this.e)
	if n == 0 {
		return nil
	}
	r := this.e[n-1]
	this.e = this.e[:n-1]
	return r
}

func (this *MyStack) Empty() bool {
	return len(this.e) == 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
