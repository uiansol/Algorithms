#!/bin/python3

import math
import os
import random
import re
import sys


def superReducedString(s):
    temp = list(s)
    while True:
        exit = True
        for i in range(len(temp) - 1):
            if temp[i] != '0' and temp[i] == temp[i + 1]:
                temp[i] = '0'
                temp[i + 1] = '0'
                temp = list(filter(('0').__ne__, temp))
                exit = False
                break
        
        if exit:
            break
        
    if len(temp) == 0:
        return 'Empty String'
    return ''.join(temp)

        
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = superReducedString(s)

    fptr.write(result + '\n')

    fptr.close()
