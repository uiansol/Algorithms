#!/bin/python3

import math
import os
import random
import re
import sys


def minimumDistances(a):
    dist = len(a)
    for i in range(len(a) - 1):
        for j in range(i + 1, len(a)):
            if a[i] == a[j]:
                if j - i < dist:
                    dist = j - i
                    
    if dist == len(a):
        return -1
    return dist
    
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    a = list(map(int, input().rstrip().split()))

    result = minimumDistances(a)

    fptr.write(str(result) + '\n')

    fptr.close()
