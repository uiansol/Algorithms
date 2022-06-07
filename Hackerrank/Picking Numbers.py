#!/bin/python3

import math
import os
import random
import re
import sys

def pickingNumbers(a):
    result = 0
    for i in range(len(a)):
        lenght_up = 1
        lenght_down = 1
        for j in range(i + 1, len(a)):
            if a[j] == a[i] or a[j] == a[i] + 1:
                lenght_up += 1
            if a[j] == a[i] or a[j] == a[i] - 1:
                lenght_down += 1
         
        if lenght_up > result:
            result = lenght_up
        if lenght_down > result:
            result = lenght_down
            
    return result
            
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    a = list(map(int, input().rstrip().split()))

    result = pickingNumbers(a)

    fptr.write(str(result) + '\n')

    fptr.close()
