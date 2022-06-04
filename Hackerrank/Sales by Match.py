#!/bin/python3

import math
import os
import random
import re
import sys

def sockMerchant(n, ar):
    socks = {}
    for color in ar:
        if color in socks:
            socks[color] += 1
        else:
            socks[color] = 1
        
    pair = 0
    for k in socks:
        pair += socks[k] // 2
        
    return pair
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    ar = list(map(int, input().rstrip().split()))

    result = sockMerchant(n, ar)

    fptr.write(str(result) + '\n')

    fptr.close()
