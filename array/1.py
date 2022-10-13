
from typing import List

def twoSum(nums: List[int], target: int) -> List[int]:
    d = dict()

    for index, num in enumerate(nums):
        if target - num in d:
            return [index, d[target - num]]
        else:
            d[num] = index
    
    print(d)



print(twoSum([3,3], 6))