# Question 1: Sorting Algorithm (partial flips)

Given an array of integers `arr`, sort the array by performing a series of `partial flips`.

In one `partial flip` we do the following steps:
- Choose an integer k where 1 <= k <= length(arr)
- Reverse the sub-array arr[0...k-1] (0-indexed).

For example, if arr = [3,2,1,4] and we performed a partial flip choosing k = 3, we reverse the sub-array [3,2,1], so arr = [1,2,3,4] after the partial flip at k = 3.

The input array must contain all positive integers from 1 to length(arr)-1 once and only one. For example, for an array of length 7, a valid input is [3,6,1,2,4,7,5]. [1,2,4,5,2,3,7] is invalid because it contains the value 2 twice and it doesn’t contain the value 6.

`First`, implement the partial flip function.

`Second`, implement a sorting algorithm that uses the partial flip function to sort an array from the smaller number to the larger number.

This sorting function must output,
- The sorted array
- The list of k values (the sequence of partial flips) used to sort the array.

Any valid answer that sorts the array within 10 * length(arr) flips will be judged as correct.

Lastly, do not forget to write some tests.