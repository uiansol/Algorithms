class Solution:
    def findSpecialInteger(self, arr: List[int]) -> int:
        counter = {}
        for n in arr:
            if n not in counter:
                counter[n] = 1
            else:
                counter[n] += 1

        result, most_freq = 0, 0
        for key in counter:
            if counter[key] > most_freq:
                result = key
                most_freq = counter[key]

        return result