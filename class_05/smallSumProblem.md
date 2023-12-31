# 小和问题

---
## 1.小和问题介绍
给定一个数组，每个数中，所有左边比它小的数加起来的结果，叫做这个数的“**小和**”，求一个数组中所有小和相加的结果，就是小和问题。

举例：
> 给定一个数组`arr[2 6 1 3 5 0 4 8]`，这个数组中有8个数，每个数的小和分别是：<br>
> - 2 的小和是0（左边没数） 
> - 6 的小和是2
> - 1 的小和是0（左边没有比0小的数）
> - 3 的小和是3（2+1）
> - 5 的小和是6（2+1+3）
> - 0 的小和是0
> - 4 的小和是6（2+1+3）
> - 8 的小和是21（2+6+1+3+5+4）
>
> 最后，整体的结果是 2+3+6+6+21=38

---
## 2.小和问题求解思路
> 这个问题如果用`O(N^2)`的算法，很好实现，遍历每个数，循环找左边小的数累加，最后累加到总和上。这里不过多赘述。我们这里要讨论的是，如何用`O(N*LogN)`的算法解决这个问题。

### 2.1 用mergeSort的流程求解小和问题
首先将小和问题的视角转换一下，我们求每个数**左边所有比它小的数累加**，和求一个数**右边有N个数比它大，然后将这个数累加N次**是完全等价的。

我们举个例子：
> 假设有个数组，我们以任意中间某个数x的视角来看，我们求最终结果的过程中，x被累加了多少次？答案是在遍历到x右边的数时，只要比x大，x就会被累加到最终结果中，也就是**x的右边有多少个数比它大，x就会被累加多少次**。
> 
> 还是上面的例子`arr[2 6 1 3 5 0 4 8]`，我们以2这个数的视角来看，2被累加了5次，因为2的右边有`6 3 5 4 8`5个数的求小和过程中都累加了2

明白了这个道理，我们就可以利用`mergeSort`的流程解决，在执行`merge`流程的过程中，顺便统计右边有多少个数比x大。<br>
思路是：在`merge`的过程中，左组数和右组数比较大小时（merge的流程不熟悉的可以去看前面的介绍和代码）：
1. 如果左组数大，则将右组数拷贝到`help`数组中，并且**不累加**
2. 如果左组数和右组数相等，则将右组数拷贝到`help`数组中，并且**不累加**（注意之前拷贝左右组都行，但这里一定要拷贝右组，因为右组的右边其他数可能还比左组当前数大，拷贝左组数就会统计漏了）
3. 如果左组数小，则将左组数拷贝到`help`数组中，并将**左组数累加右组数剩下数的个数次**。

我们以一个具体的示例来看：

仍然是数组`arr[2 6 1 3 5 0 4 8]` 我们设 `ans = 0`
1. step为1时，分为了4组 `[2]与[6]` `[1]与[3]` `[5]与[0]` `[4]与[8]`, help数组长度是2 
   - 2 pk 6, 2小，拷贝2，2累加1次，此时左侧越界，拷贝6。help=[2 6]，ans=2
   - 1 pk 3, 1小，拷贝1，1累加1次，此时左侧越界，拷贝3。help=[1 3], ans = 1 + 2 = 3
   - 5 pk 0, 5大，拷贝0，不累加。 `help=[0 5]` ans = 3
   - 4 pk 8，4小，拷贝4，4累加1次。`help=[4 8]` ans = 3+4 = 7<br>此时第一轮执行完，ans=7，arr变成了 `arr[2 6 1 3 0 5 4 8]`
2. step为2时，分为了两组`[2 6]与[1 3]` `[0 5]与[4 8]`，help数组长度是4
   - 2 pk 1, 2大，拷贝1，不累加。`help=[1]` ans = 7
   - 2 pk 3, 2小，拷贝2，2累加1次。`help=[1 2]` ans = 7+2 = 9
   - 6 pk 3, 6大，拷贝3，不累加，此时右侧越界，拷贝左边6。`help=[1 2 3 6]` ans = 9
   - 0 pk 4, 0小，拷贝0，0累加2次。`help=[0]` ans = 9
   - 5 pk 4, 5大，拷贝4，不累加。`help=[0 4]` ans = 9
   - 5 pk 8, 5小，拷贝5，5累加1次，左侧越界，拷贝右边8`help=[0 4 5 8]` ans = 9+5 = 14<br>此时第二轮执行完，ans=14,arr变成了`arr[1 2 3 6 0 4 5 8]`
3. step为4时，只有1组`[1 2 3 6]与[0 4 5 8]` help数组长度是8
   - 1 pk 0,1大，拷贝0，不累加。`help=[0]` ans = 14
   - 1 pk 4,1小，拷贝1，1累加3次。`help=[0 1]` ans = 14+3 = 17
   - 2 pk 4,2小，拷贝2，2累加3次。`help=[0 1 2]` ans = 17 + 6 = 23
   - 3 pk 4,3小，拷贝3, 3累加3次。`help=[0 1 2 3]` ans = 23 + 9 = 32
   - 6 pk 4,6大，拷贝4, 不累加。`help=[0 1 2 3 4]` ans = 32
   - 6 pk 5,6大，拷贝5, 不累加。`help=[0 1 2 3 4 5]` ans = 32
   - 6 pk 8,6小，拷贝6, 6累加1次，左侧越界，8拷贝，结束。`help=[0 1 2 3 4 5 6 8]` ans = 32+6=38

可以看到，最后小和计算正确，值是38，整个流程走下来，数组排好序的同时，小和也计算出来了。

这是什么原理呢？当我们执行mergeSort时，实际上是在**不断扩大范围求一个数**被累加了多少次。<br>
以上面数组`arr[2 6 1 3 5 0 4 8]`例子中的数字2为例:
1. 当`step==1`时，2与`[6]`pk，我们实际上统计了 `1-1`范围2要累加的次数 1次
2. 当`step==2`时，2与`[1 3]`pk，我们实际上统计了 `2-3`范围2要累加的次数 1次
3. 当`step==4`时，2与`[0 4 5 8]`pk，我们实际上统计了 `4-7`范围2要累加的次数 3次，整体2被累加5次。

由于归并排序每次merge，都将一个小范围内的顺序固定下来了，所以在pk 2与4 时，2比4小，就已经知道2要累加3次了，因为右边数组是有序的。<br>
这样相当于节省了大量比较的过程，所以时间复杂度是`O(N*logN)`,这也是归并排序能做到`O(N*logN)`的原因。

具体代码详见[code03_small_sum_problem.go](code03_small_sum_problem.go)
