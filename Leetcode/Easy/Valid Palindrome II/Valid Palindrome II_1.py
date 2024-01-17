# Accepted  113 ms  17.81 MB

class Solution:
    def validPalindrome(self, s: str) -> bool:
        return self.front(s) or self.back(s)

    def front(self, s: str) -> bool:
        i, j = 0, len(s) - 1
        canJump = True

        while i < j:
            if s[i] != s[j]:
                if canJump:
                    canJump = False
                    i += 1
                    continue
                else:
                    return False
            
            i, j = i + 1, j - 1

        i, j = 0, len(s) - 1
        canJump = True

        return True

    def back(self, s: str) -> bool:
        i, j = 0, len(s) - 1
        canJump = True

        while i < j:
            if s[i] != s[j]:
                if canJump:
                    canJump = False
                    j -= 1
                    continue
                else:
                    return False
            
            i, j = i + 1, j - 1

        i, j = 0, len(s) - 1
        canJump = True

        return True