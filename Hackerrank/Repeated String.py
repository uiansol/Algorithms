#!/bin/python3

import math
import os
import random
import re
import sys

def repeatedString(s, n):
    result = s.count('a') * (n // len(s))
    result += s.count('a', 0, n % len(s))

    return result
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    n = int(input().strip())

    result = repeatedString(s, n)

    fptr.write(str(result) + '\n')

    fptr.close()
