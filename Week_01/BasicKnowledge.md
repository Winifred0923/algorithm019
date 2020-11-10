# 第一周  五大算法思想
前中后定义



前序：对于当前节点，先输出本身，然后输出他的左孩子，最后输出右孩子。val: 1—>2—>4—>6—>7—>3—>5
中序：对于当前节点，先输出左孩子，然后输出本身，最后输出右孩子.   val: 2—>4—>6—>7—>1—>3—>5
后序：对于当前节点，现输出左孩子，然后输出右孩子， 最后输出本身   val: 2—>4—>6—>7—>3—>5—>1
所以，各叶子结点的次序都是相同的


递归：
递归的优化：
尾递归：在函数返回时调用自身本身，并且return语句不能有表达式，这样，编辑器就可以吧尾递归做优化，使递归本身无论调用多少次，都只占用一个栈

遍历的思想：深度先序，广度层次

堆排序：https://www.runoob.com/w3cnote/heap-sort.html
对于取出TOPK 问题，比较常见的算法为堆排序，堆排序是指利用堆这种数据结构所设计的一种排序算法，堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：子结点的键值或索引总是小于（或者大于）他的父节点，分为大顶堆和小顶堆(降序or升序)

二分查找： 时间复杂度logn
前提为有序数列,当为偶数时，取下整数

常见的算法中，冒泡和插入时间复杂度都为o(n*2), 快排最差为o(n*2)，只有堆排序为o(nlogn)

动态规划 
这个算法主要解决的事多阶段决策问题  初始状态→│决策１│→│决策２│→…→│决策ｎ│→结束状态
将待求解的问题分解为若干个子问题（阶段），按顺序求解子阶段，前一子问题的解，为后一子问题的求解提供了有用的信息。在求解任一子问题时，列出各种可能的局部解，通过决策保留那些有可能达到最优的局部解，丢弃其他局部解。依次解决各子问题，最后一个子问题就是初始问题的解。
由于动态规划解决的问题多数有重叠子问题这个特点，为减少重复计算，对每一个子问题只解一次，将其不同阶段的不同状态保存在一个二维数组中。
与分治法最大的差别是：适合于用动态规划法求解的问题，经分解后得到的子问题往往不是互相独立的（即下一个子阶段的求解是建立在上一个子阶段的解的基础上，进行进一步的求解）
Day02
深度优先DFS
从顶点V，沿着一条路一直走到底，再从这条路的尽头的节点退回到上一个节点，再从另一条路走到底，不断递归重复，直到所有定点遍历完成，不撞南墙不回头

遍历：
1.从根节点1开始，像临界点有2，3，4，顺序为1—>2—>5—>9
2.9后面没有别的节点，往上退，退回5—>2因为都没有其余节点，则从3开始，3—>6—>10
3.到10后回退到3，3有右子节点，则--->7
4.4---->8
这个是完整的深度优先遍历
实现方式(实现方式：前中后序遍历)
type Node struct {
    Val int
    Left *Node
    Right *Node
}

//栈
type Stack struct {
    list *list.List
} 

func NewStack() *Stack {
    list := list.New()
    return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
    stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
    if e := stack.list.Back(); e!= nil {
        stack.list.Remove(e)
        return e.Value
    }
    
    return nil
}

func (stack *Stack) Len() int {
    return stack.list.Len()
}

func (stack *Stack) Empty() bool {
    return stack.Len() == 0
}
前序遍历.  利用栈的特性，先将右子树进行压栈，再将左子树进行压栈
func (root *Node) PreTravesal() {
    if root == nil {
        return
    }
    
    s := stack.NewStack()
    s.push(root)
    
    for !s.Empty() {
        cur := s.Pop().(*Node)
        fmt.Println(cur.Val)
        
        if cur.Right != nil {
            s.Push(cur.Right)
        }
        if cur.Left != nil {
            s.Push(cur.Left)
        }
    }
}
中序遍历
func (root *Node) InTravesal() {
    if root == nil {
        return
    }
    
    s := stack.NewStack()
    cur := root
    for {
        for cur != nil {
            s.Push(cur)
            cur = cur.Left
        }
        
        if s.Empty() {
            break
        }
        
        cur = s.Pop().(*Node)
        fmt.Println(cur.Val)
        cur = cur.right
    }
}
后序遍历：则按照先进后出，将根节点先进行压栈，然后右节点，左节点
func (root *Node) PostTravesal() {
    if root == nil {
        return
    }
    
    s := stack.NewStack()
    out := stack.NewStack()
    s.Push(root)
    
    for !s.Empty() {
        cur := s.Pop().(*Node)
        out.Push(cur)
        
        if cur.Right != nil {
            s.Push(cur.Right)
        }
        if cur.Left != nil {
            s.Push(cur.Left)
        }
    }
    
    for !out.Empty() {
        cur := out.Pop().(*Node)
        fmt.Println(cur.Val)
    }
}

BFS层次遍历

核心思想：每次出队一个元素，就将该元素的孩子节点加入队列中，直至队列中元素个数为0时，出队的顺序就是该二叉树的层次遍历结果
1.根节点入队列，出队列时，将孩子节点即B，C入队列
2.B节点出队列时，将孩子节点D入队列
3.C节点出队列时，孩子节点F，G入队列
4，以此类推
广度优先：数组实现
var result [][]int
func levelOrder1(root *TreeNode) [][]int {
result = make([][]int, 0)
if root == nil {
return result
}
dfsHelper(root, 0)
return result
}
func dfsHelper(node *TreeNode, level int) {
if node == nil {
return
}
if len(result) < level + 1 {
result = append(result, make([]int, 0))
}
result[level] = append(result[level], node.Val)
dfsHelper(node.Left, level + 1)
dfsHelper(node.Right, level + 1)
}

Day04   五大常用算法思想：分治算法/动态规划/贪心算法/回溯法(DFS)/分支线算法(BFS)
DFS/BFS已分享
分治算法：分而治之
常用到的算法: 二分，大整数乘法，Strassen矩阵乘法，棋盘覆盖，合并排序，快排。线性时间选择，最接近点对问题。循环赛日程表，汉诺塔
任何一个可以用计算机求解的问题的计算时间都与规模有关
分治算法的策略：对于一个规模为n的问题，如果容易解决则直接解决，否则将其分解为k个规模较小的子问题，字问题间相互独立。递归的解这些子问题，将各个子问题的解合并得到原问题的解，由分治法产生的子问题往往是原问题的较小模式，分治法与递归像一对孪生兄弟，经常同时应用在算法中！！！！
分治法所能解决的问题都有四个特征：
1.该问题的规模缩小到一定的程度就可以容易地解决
2.该问题可以分解为若干个规模较小的相同问题，即该问题具有最有子结构性质
3.利用该问题分解出的子问题的解可以合并为该问题的解
4.该问题所分解出的各个子问题是相互独立的，即子问题之间不包含公共的子子问题
第二个特征是应用分治法的前提，反映了递归思想的应用，第三条特征为关键，能否用分治法取决于是否有第三条,不具备第三条特征则可以考虑用贪心或者动态算法。第四条涉及到了分治法的效率，如果不具备，一般考虑使用动态规划

步骤，三个步骤（分解，解决，合并）
1. 将原问题分解为若干个规模较小，相互独立，与原问题形式相同的子问题
2. 若子问题规模较小，而容易被解决则直接解，否则递归的解各个子问题
3. 将各个子问题的解合并为原问题的解
典型例题
二分查找：前提为有序，且告知升序降序
思想:二分查找的核心思想是将 n 个元素分成大致相等的两部分，取中间值 a[n/2] 与 x 做比较，如果 x=a[n/2]，则找到 x，算法中止，如果 x<a[n/2]，则只要在数组 a 的左半部分继续搜索 x，如果 x>a[n/2]，则只要在数组 a 的右半部搜索 x
func BinaryFind(nums []int, left, right, val int)int{
   middle := left + right / 2

   if nums[middle] > val {
      BinaryFind(nums[0:middle], left, middle-1, val)
   }else if nums[middle] < val{
      BinaryFind(nums[0:middle], middle+1, right, val)
   }else {
      return middle
   }
   return 0
}
快速排序
func partition(nums []int, left, right int) int {
   val := nums[left]

   for left < right {
      for nums[right] >= val && left < right { // 依次查找大于等于基准值的位置
         right --
      }
      nums[left] = nums[right]

      for nums[left] < val && left < right { // 依次查找小于基准值的位置
         left ++
      }
      nums[right] = nums[left]
   }
   nums[left] = val
   //最终 left == right 就是基准值的位置
   return left
}

func QuickSort(nums []int, left, right int) {
   if left < right {
      middle := partition(nums, left, right)
      QuickSort(nums, left, middle-1)
      QuickSort(nums, middle+1, right)
   }
}
归并排序：将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。
func MergeSort(nums []int) []int {
   if len(nums) < 2 {
      return nums
   }
   mid := len(nums)/ 2

   leftNums := MergeSort(nums[:mid])
   rightNums := MergeSort(nums[mid:])

   res := merge(leftNums, rightNums)
   return  res
}

func merge(left, right []int)[]int {
   l, r := len(left), len(right)
   lIndex,rIndex := 0, 0
   res := make([]int, 0)

   for lIndex < l && rIndex < r {
      if left[lIndex] < right[rIndex] {
         res = append(res, left[lIndex])
         lIndex++
         continue
      }
      res = append(res, rIndex)
      rIndex++
   }

   if lIndex < l {
      res = append(res, left[lIndex:]...)
   }
   if rIndex < r {
      res = append(res, right[rIndex:]...)
   }
   return res
}


Day04 贪心算法
定义：在对问题求解时，总是做出在当前看老师最好的选择，也就是说，从不从整体最有上加以考虑，他所作出的仅仅是在某种意义上的局部最优解
贪心算法没有固定的算法框架，算法设计的关键是贪心策略的选择，必须注意的是贪心算法不是对所有问题都能得到整体最优解，选择的贪心策略必须具备无后效性(即某个状态以后的过程不会影响以前的状态，只与当前状态有关)
基本思路：
1.建立数学模型描述问题
2.把求解的问题分成若干个子问题
3.对每个子问题求解，得到子问题的局部最优解
4.把子问题的解局部最优解合成原来问题的一个解(与分治不同的是，分治的子问题不存在最优解，只有一个解)

存在的问题：
1.不能保证求得的最后解是最佳的
2.一般用来求最大值或最小值问题
3.只能求满足某些约束条件的可行解的范围

适用问题
适用前提：局部最优策略导致产生全局最优解
选择性质
当考虑做何种选择的时候，我们只考虑对当前问题最佳的选择而不考虑子问题的结果。这是贪心算法可行的第一个基本要素，对于一个具体问题，要确定它是否具有贪心选择性质，必须证明每一步所作的贪心选择最终导致问题的整体最优解

实现框架：
从某一问题的初始解出发
while(朝给定总目标前进一步){
    利用可行的决策，求出可行解的一个解元素
}
由所有解元素组合成问题的一个可行解

例题：
找零问题/背包问题/活动选择问题/多机制调度问题
https://blog.csdn.net/mayifan_blog/article/details/85063336?utm_medium=distribute.pc_relevant.none-task-blog-title-10&spm=1001.2101.3001.4242
找零问题：指定币值和相应的数量，用最少的数量凑齐某金额， 利用贪心算法，我们优先选择面值大的钱币，以此类推，直到凑齐总金额
func greedy() {
   values := []int{1, 2, 5, 10, 20, 50, 100}  //面额
   counts := []int{3, 3, 2, 1, 1, 3, 3}  //数量
   result := getNumber(446, values, counts)   //获取需要各种面值多少张
   fmt.Println(result)
}

func getNumber(sum int, values, counts []int)[]int{
   result := make([]int, 0)
   add := 0 //当前凑的a

   for i := len(values)-1; i>=0; i-- {
      num := (sum-add)/values[i]
      if num > counts[i] {
         num=counts[i]
      }
      add = add + num * values[i]
      result[i] = num
   }
   return result
}

总结：贪心算法追求局部最优，拿到问题之后先分析我们需要达到什么目标，是否适合采用贪心算法，并且使得什么最优以及实现的方法
最后附上go学习算法的路程https://www.ctolib.com/ningskyer-Algorithms-Learning-With-Go.html