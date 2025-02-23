from pairs import car, cdr, cons, is_pair, to_string


def reverse_pair(pair):
    lef = cdr(pair)
    rig = car(pair)
    return cons(lef, rig)


def sum_of_pairs(pair_1, pair_2):
    return cons(car(pair_1) + car(pair_2), cdr(pair_1) + cdr(pair_2))


def find_primitive_box(pair):
    if pair is None:
        return
    if not is_pair(car(pair)) and not is_pair(cdr(pair)):
        return pair
    left_result = find_primitive_box(car(pair)) if is_pair(car(pair)) else None
    if left_result:
        return left_result
    right_result = find_primitive_box(cdr(pair)) if is_pair(cdr(pair)) else None
    if right_result:
        return right_result


pair = cons("one", "two")
print(to_string(reverse_pair(pair)))  # ('two', 'one')


pair1 = cons(4, 10)
pair2 = cons(100, 0)
print(to_string(sum_of_pairs(pair1, pair2)))


pair = cons(None, cons("one", "two"))
print(to_string(find_primitive_box(pair)))  # ('one', 'two')

pair2 = cons(cons(None, cons(1, 5)), None)
print(to_string(find_primitive_box(pair2)))  # (1, 5)
