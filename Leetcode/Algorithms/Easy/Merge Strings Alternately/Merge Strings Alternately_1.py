# Accepted  34 ms  17.28 MB

class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        l1, l2 = len(word1), len(word2)
        i, j = 0, 0
        coin = True
        merged = []

        while i + j < l1 + l2:
            if coin:
                merged.append(word1[i])
                i += 1
                if j < l2:
                    coin = False
            else:
                merged.append(word2[j])
                j += 1
                if i < l1:
                    coin = True

        return ''.join(merged)