#!/bin/python3

import math
import os
import random
import re
import sys

def beautifulTriplets(d, arr):
    count = 0
    for i in range(len(arr) - 2):
        for j in range(i + 1, len(arr) - 1):
            if arr[j] - arr[i] != d:
                continue
            
            for k in range(j + 1, len(arr)):
                if arr[k] - arr[j] == d:
                    count += 1
                    break
                    
    return count
    
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    d = int(first_multiple_input[1])

    arr = list(map(int, input().rstrip().split()))

    result = beautifulTriplets(d, arr)

    fptr.write(str(result) + '\n')

    fptr.close()
