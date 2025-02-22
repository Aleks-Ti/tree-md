def cons(head, tail):
    def f(message):
        if message == "head":
            return head
        if message == "tail":
            return tail
    return f

def car(pair):
    return pair("head")


def cdr (pair):
    return pair("tail")

pair = cons(5, 3)
print(car(pair))
print(cdr(pair))
