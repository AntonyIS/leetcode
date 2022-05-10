from typing import List


def threeSum(nums:List[int]) -> List:
    # Define lists to store final results
    results:List[int] = []
    # Sort list in ascending order
    nums.sort()

    # Loop through values in list
    for index, value in enumerate(nums):
        # Check if we are processing the same value twice
        if index > 0 and value == nums[index -1] :
            # Skip the current iterator and move to the next
            continue

        # Define left and right pointers of the list
        left, right = index + 1, len(nums) -1
        # left pointer starts after the first items (nums[0])
        # right pointer starts at the last item in the list

        # Loop until left !> right
        while left < right :
            sum = value + nums[left] + nums[right]
            
            if sum > 0 :
                # Check if sum > 0 : move right pointer back by one position
                right -= 1
            elif sum < 0 :
                # Check if sum < 0 : move left pointer forward by one position
                left += 1
            else:
                # At this point sum == 0
                results.append([value, nums[left], nums[right]])
                # Move the left pointer forward
                left += 1
                # Check if we are repeating the same value 
                while nums[left] == nums[left -1] and left < right :
                    left += 1

    return results

if __name__ == "__main__":
    nums:List[int] =[-1,0,1,2,-1,-4]
    response = threeSum(nums)
    print(response)





    