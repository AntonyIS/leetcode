from typing import List



def maxArea(nums : List[int])-> int :
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

if __name__ == "__main__":
    nums:List[int] = [1,8,6,2,5,4,8,3,7]
    max_area = maxArea(nums)
    print(max_area)
