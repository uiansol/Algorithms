#!/bin/python3

import math
import os
import random
import re
import sys


def fairRations(B):
    odds = countOdds(B)
    if odds % 2 != 0:
        return 'NO'
    
    loafs = 0
    for i in range(len(B) - 1):
        if B[i] % 2 != 0:
            B[i] += 1
            B[i + 1] += 1
            loafs += 2
        
    return str(loafs)

def countOdds(B):
    count = 0
    for n in B:
        if n % 2 != 0:
            count += 1
            
    return count
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    N = int(input().strip())

    B = list(map(int, input().rstrip().split()))

    result = fairRations(B)

    fptr.write(result + '\n')

    fptr.close()
