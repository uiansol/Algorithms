#!/bin/python3

import math
import os
import random
import re
import sys

def kaprekarNumbers(p, q):
    exists = False
    for i in range(p, q + 1):
        square = str(i**2)
        n1 = square[len(square) // 2:]
        n2 = square[:len(square) // 2]
        
        if n1 == '': n1 = '0'
        if n2 == '': n2 = '0'
        
        if int(n1) + int(n2) == i:
            print(str(i) + ' ', end='')
            exists = True
        
    if not exists:
        print('INVALID RANGE')
    
    
if __name__ == '__main__':
    p = int(input().strip())

    q = int(input().strip())

    kaprekarNumbers(p, q)
