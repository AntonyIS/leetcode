from typing import List


# Solution 1
def maxArea_1(nums : List[int])-> int :
    # Define maxArea as the 0
    maxArea = 0

    # define left and right pointers
    left, right = 0, len(nums) - 1

    # Loop through height values 
    while left < right :
        width = min(nums[left], nums[right])
        length = right -left
        area = length * width
        maxArea = max(maxArea, area)

        if nums[left] < nums[right] :
            left += 1
        else:
            right -= 1

    return maxArea
# Solution 2
def maxArea_2(nums: List[int]) -> int :
    maxArea , left, right = 0, 0, len(nums) -1

    while left < right :
        width = right - left 
        height = 0

        if nums[left] < nums[right] :
            height = nums[left]
            left += 1
        else:
            height = nums[right]
            right -= 1

        tempArea = width * height
        maxArea = max(maxArea, tempArea)

    return maxArea

if __name__ == "__main__":
    nums:List[int] = [1,8,6,2,5,4,8,3,7]
    max_area = maxArea_2(nums)
    print(max_area)
