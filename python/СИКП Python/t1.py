import math

class Points:
    def __init__(self) -> None:
        self.x = None
        self.y = None

    def get_x(self, dot):
        return dot.x
    
    def get_y(self, dot):
        return dot.y
    
    def make(self, x, y):
        self.x = x
        self.y = y
        return self

    def to_string(self, coords):
        return str(coords)

points = Points()

def get_quadrant(point):
    x = points.get_x(point)
    y = points.get_y(point)
    if x == 0 or y == 0:
        return None
    if x < 0:
        if y < 0:
            return 3
        return 2
    else:
        if y < 0:
            return 4
        return 1


def get_symmetrical_point(point):
    x = points.get_x(point)
    y = points.get_y(point)
    return points.make(-x, -y)


def calculate_distance(point_1, point_2):
    x_1 = points.get_x(point_1)
    y_1 = points.get_y(point_1)
    x_2 = points.get_x(point_2)
    y_2 = points.get_y(point_2)
    delta_x = x_2 - x_1
    delta_y = y_2 - y_1
    return math.sqrt(delta_x ** 2 + delta_y ** 2)



point1 = points.make(1, 5)
print(get_quadrant(point1))  # 1
point2 = points.make(1, -5)
print(get_quadrant(point2))  # 4

point1 = points.make(0, 7)
print(get_quadrant(point1) is None)  # True
point2 = points.make(2, 0)
print(get_quadrant(point2) is None)  # True

print(points.to_string(get_symmetrical_point(points.make(1, 5))))  # '(-1, -5)'


print(calculate_distance(
    points.make(3, 2),
    points.make(-1, -1),
))


pp: Points = points.make(10, -10)


object_point = get_symmetrical_point(pp)  # (-10, 10)
print(object_point.get_x(object_point))
print(object_point.get_y(object_point))
