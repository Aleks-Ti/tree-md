class Pairs:
    def __init__(self) -> None:
        self.left = None
        self.right = None


def to_string(pair):
    return pair.left, pair.right


def cons(a, b):
    def f(message):
        if message == "car":
            return a
        if message == "cdr":
            return b
    return f

def car(pair):
    return pair("car")

def cdr(pair):
    return pair("cdr")

def is_pair(pair):
    return isinstance(pair, Pairs)
