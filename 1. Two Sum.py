class Solution(object):
    def twoSum(self, nums, target):
        seen={}
        for i in range(len(nums)):
            wanted = target - nums[i]
            if wanted in seen:
                return [seen[wanted], i]
            seen[nums[i]] = i

#with hashmap 0ms runtime beats 100%