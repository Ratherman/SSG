import time
import sys


def Test_func(a):
    print(sys.modules.keys())
    print('hello world')
    print(a)

    return a


def main():
    Test_func(1.0)


if __name__ == "__main__":
    main()
