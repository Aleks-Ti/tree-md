def make(numer, denom):
    def f(message):
        if message == "numer":
            return numer
        if message == "denom":
            return denom
    return f

def numer(rationals):
    return rationals("numer")


def denom(rationals):
    return rationals("denom")


def to_string(rationals):
    return f"{numer(rationals)} / {denom(rationals)}"


def is_equal(rationals_1, rationals_2):
    return numer(rationals_1) * denom(rationals_2) == numer(rationals_2) * denom(rationals_1)



def add(rationals_1, rationals_2):
    return make(numer(rationals_1) * denom(rationals_2) + denom(rationals_1) * numer(rationals_2), denom(rationals_1) * denom(rationals_2))


def sub(rationals_1, rationals_2):
    left = (numer(rationals_1) * denom(rationals_2)) - (denom(rationals_1) * numer(rationals_2))
    right =  denom(rationals_1) * denom(rationals_2)
    return make(left, right)


def mul(rationals_1, rationals_2):
    left = numer(rationals_1) * numer(rationals_2)
    right = denom(rationals_1) * denom(rationals_2)
    return make(left, right)


def div(rationals_1, rationals_2):
    left = numer(rationals_1) * denom(rationals_2)
    right = denom(rationals_1) * numer(rationals_2)
    return make(left, right)
