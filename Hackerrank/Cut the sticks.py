#!/bin/python3

import math
import os
import random
import re
import sys


def cutTheSticks(arr):
    result = []
    
    while countArr(arr) > 0:
        result.append(countArr(arr))
        short = minArr(arr)
        for i in range(len(arr)):
            arr[i] -= short
    
    return result        
    
def minArr(arr):
    res = 1000
    for n in arr:
        if n > 0 and n < res:
            res = n
            
    return res
    
def countArr(arr):
    res = 0
    for n in arr:
        if n > 0:
            res += 1
            
    return res
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    result = cutTheSticks(arr)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
