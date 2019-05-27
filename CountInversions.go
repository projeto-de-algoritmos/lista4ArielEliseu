package main

import "fmt"

/**
 * @brief      { This function sorts the input array and returns the
 * number of inversions in the array }
 *
 * @param      arr         The array to be sorted
 * @param      array_size  The array size
 *
 * @return     { the number of inversions }
 */
func mergeSort(arr [100]int, array_size int) int {
	//temp := make([]int, 100) //array_size)
	var temp [100]int
	return _mergeSort(arr, temp, 0, array_size-1)
}

/**
 * @brief      { An auxiliary recursive function that sorts the input array and
 * returns the number of inversions in the array.}
 *
 * @param      arr    The array of data
 * @param      temp   The temporary array
 * @param      left   Initial index
 * @param      right  Initial index
 *
 * @return     {  the number of inversions }
 */
func _mergeSort(arr [100]int, temp [100]int, left int, right int) int {
	var mid int
	var inv_count = 0
	if right > left {
		/* Divide the array into two parts and
		   call _mergeSortAndCountInv()
		   for each of the parts */
		mid = (right + left) / 2

		/* Inversion count will be sum of
		   inversions in left-part, right-part
		   and number of inversions in merging */
		inv_count = _mergeSort(arr, temp, left, mid)
		inv_count = inv_count + _mergeSort(arr, temp, (mid+1), right)

		/*Merge the two parts*/
		inv_count = inv_count + merge(arr, temp, left, (mid+1), right)
	}
	return inv_count
}

/**
 * @brief      {  This funt merges two sorted arrays
 * and returns inversion count in the arrays. }
 *
 * @param      arr    The array to be merged
 * @param      temp   The temporary array to be merged
 * @param      left   The left index
 * @param      mid    The middle index
 * @param      right  The right index
 *
 * @return     { the number of inversions }
 */
func merge(arr [100]int, temp [100]int, left int, mid int, right int) int {
	var i int
	var j int
	var k int
	var inv_count = 0

	i = left /* i is index for left subarray*/
	j = mid  /* j is index for right subarray*/
	k = left /* k is index for resultant merged subarray*/
	for (i <= mid-1) && (j <= right) {
		if arr[i] <= arr[j] {
			k++
			i++
			temp[k] = arr[i]
		} else {
			k++
			j++
			temp[k] = arr[j]

			/* this is tricky -- see above
			   explanation/diagram for merge()*/
			inv_count = (inv_count + (mid - i))
		}
	}

	/* Copy the remaining elements of left subarray
	(if there are any) to temp*/
	for i <= (mid - 1) {
		k++
		i++
		temp[k] = arr[i]
	}
	/* Copy the remaining elements of right subarray
	(if there are any) to temp*/
	for j <= right {
		k++
		j++
		temp[k] = arr[j]
	}
	/*Copy back the merged elements to original array*/
	for i := left; i <= right; i++ {
		arr[i] = temp[i]
	}

	return inv_count
}

/* Driver code */
func main() {
	//:= [5]int{1, 2, 3, 4, 5}
	arr := [100]int{1, 2, 4, 3, 5}
	fmt.Println(" Number of inversions are ", mergeSort(arr, 5))
	fmt.Println(arr)
}

// This is code is contributed by rathbhupendra and Ariel Serafim;;
