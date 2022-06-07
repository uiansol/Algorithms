#!/bin/python3

import math
import os
import random
import re
import sys


def viralAdvertising(n):
    shared = 5
    liked = 2
    cumulative = 2
    for i in range(n - 1):
        shared = liked * 3
        liked = shared // 2
        cumulative += liked
        
    return cumulative
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    result = viralAdvertising(n)

    fptr.write(str(result) + '\n')

    fptr.close()
