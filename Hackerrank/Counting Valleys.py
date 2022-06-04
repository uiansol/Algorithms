#!/bin/python3

import math
import os
import random
import re
import sys

def countingValleys(steps, path):
    level = 0
    valleys = 0
    
    for step in path:
        if step == 'U':
            if level == -1:
                valleys += 1
            level += 1
        else:
            level -= 1
            
    return valleys
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    steps = int(input().strip())

    path = input()

    result = countingValleys(steps, path)

    fptr.write(str(result) + '\n')

    fptr.close()
