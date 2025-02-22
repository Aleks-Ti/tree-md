from points import get_x, get_y
from pairs import cons, car, cdr

class Rectangles:
    def __init__(self):
        self.l_top_point = None
        self.width = None
        self.height = None


def make(l_top_point, width, height):
    return cons(l_top_point, cons(width, height))


def get_start_point(pair):
    return car(pair)


def get_width(rectangle):
    w_h = cdr(rectangle)
    return car(w_h)


def get_height(rectangle):
    w_h = cdr(rectangle)
    return cdr(w_h)


def get_square(rectangle):
    w = get_width(rectangle)
    h = get_height(rectangle)
    return w * h


def get_perimeter(pair):
    w = get_width(pair)
    h = get_height(pair)
    return 2 * (w + h)


def contains_the_origin(rectangle):
    x_l_top_point = get_x(car(rectangle))
    y_l_top_point = get_y(car(rectangle))

    w = get_width(rectangle)
    h = get_height(rectangle)
    if y_l_top_point < 0:
        return False
    if x_l_top_point < 0:
        if y_l_top_point - h < 0 and  x_l_top_point + w > 0:
            return True
        return False
    return False
