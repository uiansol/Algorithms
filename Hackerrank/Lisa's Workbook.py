#!/bin/python3

import math
import os
import random
import re
import sys


def workbook(n, k, arr):
    result = 0
    chapter = 0
    page = 1
    problem = 1
    while chapter < n:
        if problem <= page and page <= min(problem + k - 1, arr[chapter]):
            result += 1
        page += 1
        problem += k
        if problem > arr[chapter]: # next chapter
            chapter += 1
            problem = 1
            
    return result
    
  
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    k = int(first_multiple_input[1])

    arr = list(map(int, input().rstrip().split()))

    result = workbook(n, k, arr)

    fptr.write(str(result) + '\n')

    fptr.close()
