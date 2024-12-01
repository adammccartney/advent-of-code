#!/usr/bin/python3

import random


def main():
    for _ in range(0, 10000):
        l = random.randint(0, 100)
        r = random.randint(0, 100)
        print(l, r)

if __name__ == '__main__':
    main()
