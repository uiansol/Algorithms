#!/bin/python3

import math
import os
import random
import re
import sys

def appendAndDelete(s, t, k):
    l = len(s) + len(t)
    
    i = 0
    while i < len(s) and i < len(t) and s[i] == t[i]:
        i += 1

    return 'Yes' if l <= k + i*2 and l % 2 == k % 2 or l < k else 'No'
    
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    t = input()

    k = int(input().strip())

    result = appendAndDelete(s, t, k)

    fptr.write(result + '\n')

    fptr.close()
