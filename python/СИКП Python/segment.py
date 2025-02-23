from pairs import car, cdr, cons
from points import get_x, get_y
from points import make as make_point


class segment:
    def __init__(self) -> None:
        self.array: list = []


def get_begin(segment):
    return car(segment)


def get_end(segment):
    return cdr(segment)


def make(point_1, point_2):
    return cons(point_1, point_2)


def to_string(segment):
    return str([(get_x(car(segment)), get_y(car(segment))), (get_x(cdr(segment)), get_y(cdr(segment)))])


def middle_point(segment):
    x = (get_x(car(segment)) + get_x(cdr(segment))) / 2
    y = (get_y(car(segment)) + get_y(cdr(segment))) / 2
    return make_point(x, y)
