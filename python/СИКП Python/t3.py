from points import make as make_point, to_string as point_to_string
import segment

# не важно, чем является current_segment с точки зрения реализации, главное, что с ним можно
# работать используя функции для работы с отрезками
current_segment = segment.make(make_point(1, 2), make_point(-4, -2))
print(segment.to_string(current_segment))  # [(1, 2), (-4, -2)]
point1 = segment.get_begin(current_segment)
print(point_to_string(point1))  # (1, 2)
point2 = segment.get_end(current_segment)
print(point_to_string(point2))  # (-4, -2)
print(point_to_string(segment.get_begin(current_segment)) == point_to_string(make_point(1, 2)))  # True
print(point_to_string(segment.middle_point(current_segment)))  # (-1.5, 0)