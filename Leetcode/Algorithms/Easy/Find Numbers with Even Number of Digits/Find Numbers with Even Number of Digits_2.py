class Solution:
    def findNumbers(self, nums: List[int]) -> int:
        count = 0

        for v in nums:
            if (v >= 10 and v <= 99) or (v >= 1000 and v <= 9999) or v == 100000:
                count += 1

        return count
        