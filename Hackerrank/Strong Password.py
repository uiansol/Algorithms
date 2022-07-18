#!/bin/python3

import math
import os
import random
import re
import sys

numbers = "0123456789"
lower_case = "abcdefghijklmnopqrstuvwxyz"
upper_case = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
special_characters = "!@#$%^&*()-+"

def minimumNumber(n, password):
    count = 0
    
    if needDigit(password):
        count += 1
    
    if needLowerCase(password):
        count += 1
    
    if needUpperCase(password):
        count += 1
        
    if needSpecial(password):
        count += 1
        
    if n + count < 6:
        count += 6 - (n + count)   
    
    return count
        
def needDigit(password):
    for c in password:
        if c in numbers:
            return False
    
    return True
 
def needLowerCase(password):
    for c in password:
        if c in lower_case:
            return False
    
    return True    
 
def needUpperCase(password):
    for c in password:
        if c in upper_case:
            return False
    
    return True    

def needSpecial(password):
    for c in password:
        if c in special_characters:
            return False
    
    return True
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    password = input()

    answer = minimumNumber(n, password)

    fptr.write(str(answer) + '\n')

    fptr.close()
