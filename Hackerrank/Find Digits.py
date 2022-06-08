#!/bin/python3

import math
import os
import random
import re
import sys


def findDigits(n):
    nString = str(n)
    result = 0
    
    for digit in nString:
        if digit != '0' and n % int(digit) == 0:
            result +=1

            
    return result
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input().strip())

    for t_itr in range(t):
        n = int(input().strip())

        result = findDigits(n)

        fptr.write(str(result) + '\n')

    fptr.close()
