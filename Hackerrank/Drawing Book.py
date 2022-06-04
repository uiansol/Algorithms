#!/bin/python3

import math
import os
import random
import re
import sys

def pageCount(n, p):
    if n % 2 == 0:
        n += 1
        
    min1 = p // 2
    min2 = (n - p) // 2
    
    return min(min1, min2)
    
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    p = int(input().strip())

    result = pageCount(n, p)

    fptr.write(str(result) + '\n')

    fptr.close()
