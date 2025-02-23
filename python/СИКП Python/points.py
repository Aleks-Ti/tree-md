class Points:
    def __init__(self) -> None:
        self.x = None
        self.y = None


def get_x(dot):
    return dot.x


def get_y(dot):
    return dot.y


def make(x, y):
    point = Points()
    point.x = x
    point.y = y
    return point


def to_string(point):
    return point.x, point.y


def get_quadrant(point):
    return get_x(point) * get_y(point)
