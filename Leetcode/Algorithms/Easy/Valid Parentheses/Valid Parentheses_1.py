# Accepted  31 ms  17.42 MB

class Solution:
    def isValid(self, s: str) -> bool:
        openers = ['(', '{', '[']
        stack = []

        for c in s:
            if c in openers:
                stack.append(c)
            else:
                if len(stack) == 0:
                    return False

                lastOpen = stack.pop()

                if (lastOpen == '(' and c != ')') or (lastOpen == '{' and c != '}') or (lastOpen == '[' and c != ']'):
                    return False
                
        if len(stack) > 0:
            return False
          
        return True