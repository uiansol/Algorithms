# Accepted  42 ms  17.36 MB

class Solution:
    def validWordAbbreviation(self, word: str, abbr: str) -> bool:
        numbers = '0123456789'
        l1, l2 = len(word), len(abbr)
        i, j = 0, 0

        while i < l1 and j < l2:
            if word[i] == abbr[j]:
                i, j = i + 1, j + 1
                continue
            elif abbr[j] not in numbers:
                return False

            sub = ''
            while j < l2 and abbr[j] in numbers:
                sub += abbr[j]
                j += 1

            if sub[0] == '0':
                return False

            if i + int(sub) > l1:
                return False

            i += int(sub)

        if (i >= l1 and j < l2) or (j >= l2 and i < l1):
            return False

        return True