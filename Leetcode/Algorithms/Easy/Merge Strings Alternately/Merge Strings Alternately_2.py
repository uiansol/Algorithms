# Accepted  33 ms  17.37 MB

class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        l1, l2 = len(word1), len(word2)
        l = max(l1, l2)
        merged = []
        
        for i in range(l):
            if i < l1:
                merged += word1[i]
            if i < l2:
                merged += word2[i]
            
        return ''.join(merged)