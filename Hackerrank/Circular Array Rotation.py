#!/bin/python3

import math
import os
import random
import re
import sys

def circularArrayRotation(a, k, queries):
    result = []
    start = len(a) - (k % len(a))
    if start == len(a):
        start = 0
    rotated = a[start:] + a[:start]
    
    for q in queries:
        result.append(rotated[q])

    return result

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    k = int(first_multiple_input[1])

    q = int(first_multiple_input[2])

    a = list(map(int, input().rstrip().split()))

    queries = []

    for _ in range(q):
        queries_item = int(input().strip())
        queries.append(queries_item)

    result = circularArrayRotation(a, k, queries)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
