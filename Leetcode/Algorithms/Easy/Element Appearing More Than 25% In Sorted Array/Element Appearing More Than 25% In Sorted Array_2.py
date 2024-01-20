class Solution:
    def findSpecialInteger(self, arr: List[int]) -> int:
        counter = defaultdict(int)
        target = len(arr) / 4

        for n in arr:
            counter[n] +=1
            if counter[n] > target:
                return n

        return -1