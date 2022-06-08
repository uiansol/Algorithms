#!/bin/python3

import math
import os
import random
import re
import sys


def jumpingOnClouds(c):
    jumps = 0
    position = 0
    last = len(c) - 1
    while position < last:
        if position + 2 <= last and c[position + 2] == 0:
            position += 2
        else:
            position += 1
        
        jumps += 1
        
    return jumps
    
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    c = list(map(int, input().rstrip().split()))

    result = jumpingOnClouds(c)

    fptr.write(str(result) + '\n')

    fptr.close()
