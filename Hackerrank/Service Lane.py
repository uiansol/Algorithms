#!/bin/python3

import math
import os
import random
import re
import sys


def serviceLane(n, cases):
    result = []
    for case in cases:
        res = width[case[0]]
        for segment in range(case[0], case[1] + 1):
            if width[segment] < res:
                res = width[segment]
                
        result.append(res)
    
    return result
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    t = int(first_multiple_input[1])

    width = list(map(int, input().rstrip().split()))

    cases = []

    for _ in range(t):
        cases.append(list(map(int, input().rstrip().split())))

    result = serviceLane(n, cases)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
