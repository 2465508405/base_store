/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-10 18:46:00
 * @LastEditors: ykk ykk@qq.com
*
*对选择排序的一种优化， 非稳定性排序
*
*/
package main

import "fmt"

//满二叉树
//一棵深度为k，且有2^k-1个节点的树是满二叉树。
//完全二叉树是由满二叉树而引出来的。对于深度为K的，有n个结点的二叉树，
//当且仅当其每一个结点都与深度为K的满二叉树中编号从1至n的结点
//一一对应时称之为完全二叉树。

// 堆排序的步骤分为两步：1、构建大（小）根堆 2、调整根堆
// 1、构建堆，把最值元素放到父节点，从最后一个非叶子节点开始调整，直到i=0(非叶子节点=0...(n/2-1))
// 2、把堆顶和未调整堆的最后一个元素交换，然后i--继续执行1和2步骤
// 由于和选择排序一样是交换排序，所以堆排序也是不稳定排序
func main() {

	defer fmt.Println("heap sort complete")
	cha1n := make(chan []int)
	var array = []int{10, 2, 7, 9, 4, 11}
	go HeapSort(cha1n, array)
	fmt.Println(<-cha1n)
}

func HeapSort(cha1n chan<- []int, nums []int) {
	// 1、构建堆(这里用大顶堆构建升序)
	// 2、调整堆，把堆顶元素和第i-1个元素交换，这样0....i-2就又成为一个堆，继续对这个堆进行构建，调整
	Hepify(nums, len(nums)) // 先构建n个元素的大顶堆
	// fmt.Println(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i] // 调整堆顶元素，把堆顶元素和最后一个元素交换
		Hepify(nums, i)
	}

	cha1n <- nums
}

// 构建堆，一般从最后一个非叶子节点开始构建，即从下往上调整，从下往上能让最大（小）值元素转移到堆顶
func Hepify(nums []int, unsortCapacity int) {
	for i := (unsortCapacity / 2) - 1; i >= 0; i-- { // 非叶子节点的i范围从0...(n/2-1)个
		// 调整左子树
		leftIndex := 2*i + 1
		if leftIndex < unsortCapacity && nums[i] < nums[leftIndex] {
			nums[i], nums[leftIndex] = nums[leftIndex], nums[i] // 左孩子值大于父节点，交换
		}
		// 调整右子树
		rightIndex := 2*i + 2
		if rightIndex < unsortCapacity && nums[i] < nums[rightIndex] {
			nums[i], nums[rightIndex] = nums[rightIndex], nums[i] // 右孩子值大于父节点，交换
		}
	}
}
