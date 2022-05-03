from typing import List


def findMin(nums : List[int]) -> int :
    # Define the min value
    minValue:int = nums[0]

    # Define left and right sides of the array
    left, right =  0 , len(nums) -1

    # Loop until left is no longer greater than right
    while left <= right :
        # Check if list is sorted : nums[left] should be less than nums[right]
        if nums[left] < nums[right]:
            # array is aleady sorted, get the min value
            minValue = min(minValue, nums[left])
            break

        # Get the mid point of the array
        mid = (left + right) //2

        # Get the min value between minValue and value at mid point
        minValue = min(minValue, nums[mid])

        if nums[left] > nums[mid]:
            right = mid -1
        else:
            left = mid + 1

    return minValue

if __name__ == "__main__":
    nums:List[int] = [2,1]
    minValue = findMin(nums)
    print(minValue)
