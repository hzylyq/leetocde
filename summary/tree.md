## 二叉树

## 二叉搜索树(bst)

1. 中序遍历有序

## 树的前中后序遍历

```
// 中序
func InOrderTraverse(tree *Node) {

    InOrderTraverse(tree.Left)
    // your code func
    InOrderTraverse(tree.Right)
}

// 前序
func PreOrderTraverse(tree *Node) {
    // your code func
    PreOrderTraverse(tree.Left)
    PreOrderTraverse(tree.Right)
}

// 后序
func PostOrderTraverse(array []int) []int {
    PostOrderTraverse(tree.Left)
    PostOrderTraverse(tree.Right)
    // your code func
}
```
* 1.递归写法只要记住code func的执行位置即可
* 2.非递归写法比较麻烦, 重点记住morris算法
  
Morris 中序遍历的方法把中序遍历的空间复杂度优化到 O(1)O(1)。 我们在中序遍历的时候，一定先遍历左子树，然后遍历当前节点，最后遍历右子树。在常规方法中，我们用递归回溯或者是栈来保证遍历完左子树可以再回到当前节点，但这需要我们付出额外的空间代价。我们需要用一种巧妙地方法可以在 O(1)O(1) 的空间下，遍历完左子树可以再回到当前节点。我们希望当前的节点在遍历完当前点的前驱之后被遍历，我们可以考虑修改它的前驱节点的 \textit{right}right 指针。当前节点的前驱节点的 \textit{right}right 指针可能本来就指向当前节点（前驱是当前节点的父节点），也可能是当前节点左子树最右下的节点。如果是后者，我们希望遍历完这个前驱节点之后再回到当前节点，可以将它的 \textit{right}right 指针指向当前节点。
Morris 中序遍历的一个重要步骤就是寻找当前节点的前驱节点，并且 Morris 中序遍历寻找下一个点始终是通过转移到 \textit{right}right 指针指向的位置来完成的。

如果当前节点没有左子树，则遍历这个点，然后跳转到当前节点的右子树。
如果当前节点有左子树，那么它的前驱节点一定在左子树上，我们可以在左子树上一直向右行走，找到当前点的前驱节点。
如果前驱节点没有右子树，就将前驱节点的 \textit{right}right 指针指向当前节点。这一步是为了在遍历完前驱节点后能找到前驱节点的后继，也就是当前节点。
如果前驱节点的右子树为当前节点，说明前驱节点已经被遍历过并被修改了 \textit{right}right 指针，这个时候我们重新将前驱的右孩子设置为空，遍历当前的点，然后跳转到当前节点的右子树。
因此我们可以得到这样的代码框架


二叉数的深度优先搜索(dfs), 广度优先搜索(bfs)
