from typing import List


def search(nums:List[int], target: int) -> int :
    # Get the left and right side of the list
    left, right = 0, len(nums) - 1

    # Loop through values in list
    while left <= right :
        # Get the midpoint of the list
        midpoint = (left+right)//2

        # Check if the value is at the middle
        if target == nums[midpoint]:
            return midpoint

        # Process values in the left side
        if nums[left] <= nums[midpoint]:
            if target > nums[midpoint] or target < nums[left] :
                left = midpoint + 1
            else:
                right = midpoint -1
        else:
            if target < nums[midpoint] or target > nums[right]:
                right = midpoint -1 
            else :
                left = midpoint + 1

    return -1


if __name__ == "__main__":
    nums:List[int] =[4,5,6,7,0,1,2]
    findValue = search(nums, 3)
    print(findValue)
