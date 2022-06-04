#!/bin/python3

import math
import os
import random
import re
import sys

def formingMagicSquare(s):
    minCost = 10000
    
    ms1 = [[8, 1, 6],[3, 5, 7],[4, 9, 2]]
    ms2 = [[8, 3, 4],[1, 5, 9], [6, 7, 2]]
    ms3 = [[6, 1, 8], [7, 5, 3], [2, 9, 4]]
    ms4 = [[4, 3, 8], [9, 5, 1], [2, 7, 6]]
    ms5 = [[4, 9, 2], [3, 5, 7], [8, 1, 6]]
    ms6 = [[6, 7, 2], [1, 5, 9], [8, 3, 4]]
    ms7 = [[2, 9, 4], [7, 5, 3], [6, 1, 8]]
    ms8 = [[2, 7, 6], [9, 5, 1], [4, 3, 8]]
    
    mss = [ms1, ms2, ms3, ms4, ms5, ms6, ms7, ms8]
    for ms in mss:
        msCost = cost(s, ms)
        if msCost < minCost:
            minCost = msCost
    
    return minCost
    
def cost(s1, s2):
    return abs(s1[0][0] - s2[0][0]) +\
           abs(s1[0][1] - s2[0][1]) +\
           abs(s1[0][2] - s2[0][2]) +\
           abs(s1[1][0] - s2[1][0]) +\
           abs(s1[1][1] - s2[1][1]) +\
           abs(s1[1][2] - s2[1][2]) +\
           abs(s1[2][0] - s2[2][0]) +\
           abs(s1[2][1] - s2[2][1]) +\
           abs(s1[2][2] - s2[2][2])
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = []

    for _ in range(3):
        s.append(list(map(int, input().rstrip().split())))

    result = formingMagicSquare(s)

    fptr.write(str(result) + '\n')

    fptr.close()
