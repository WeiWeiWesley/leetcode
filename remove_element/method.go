package removeelement

// int[] nums = [...]; // Input array
// int val = ...; // Value to remove
// int[] expectedNums = [...]; // The expected answer with correct length.
//                             // It is sorted with no values equaling val.

// int k = removeElement(nums, val); // Calls your implementation

// assert k == expectedNums.length;
// sort(nums, 0, k); // Sort the first k elements of nums
// for (int i = 0; i < actualLength; i++) {
//     assert nums[i] == expectedNums[i];
// }

//這題除了回傳值外，也會檢查 nums 最終內容，所以需要改變 nums(same address)
func removeElement(nums []int, val int) int {

    index := 0


    for _, v := range nums {
        if v != val {
            //題目測試實會 sort，所以可以透過將留下來了 val 往前移得到答案
            nums[index] = v
            index++
        }
    }

    return index
}