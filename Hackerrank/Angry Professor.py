#!/bin/python3

import math
import os
import random
import re
import sys


def angryProfessor(k, a):
    amount = 0
    for student in a:
        if student <= 0:
            amount += 1
            
    return 'YES' if amount < k else 'NO'
    
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input().strip())

    for t_itr in range(t):
        first_multiple_input = input().rstrip().split()

        n = int(first_multiple_input[0])

        k = int(first_multiple_input[1])

        a = list(map(int, input().rstrip().split()))

        result = angryProfessor(k, a)

        fptr.write(result + '\n')

    fptr.close()
