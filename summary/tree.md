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
