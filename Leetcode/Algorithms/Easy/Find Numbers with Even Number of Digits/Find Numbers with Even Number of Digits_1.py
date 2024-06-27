class Solution:
    def findNumbers(self, nums: List[int]) -> int:
        count = 0

        for v in nums:
            if len(str(v)) % 2 == 0:
                count += 1

        return count
        