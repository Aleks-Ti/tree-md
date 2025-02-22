import points
import rectangle
# Создание прямоугольника:
# p - левая верхняя точка
# 5 - ширина
# 4 - высота
#
# p    5
# -----------
# |         |
# |         | 4
# |         |
# -----------
p = points.make(0, 1)
rectangle_1 = rectangle.make(p, 5, 4)
print(rectangle.get_square(rectangle_1))  # 20
print(rectangle.get_perimeter(rectangle_1))  # 18
print(rectangle.contains_the_origin(rectangle_1))  # False
rectangle_2 = rectangle.make(points.make(-4, 3), 5, 4)
print(rectangle.contains_the_origin(rectangle_2))  # True
print(rectangle.contains_the_origin(rectangle.make(points.make(-4, 4), 5, 2)))  # False
print(rectangle.contains_the_origin(rectangle.make(points.make(-4, 3), 2, 8)))  # False