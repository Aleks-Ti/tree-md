from rationals import make, to_string, is_equal, add,sub,mul,div

rat23 = make(2, 3)
rat14 = make(1, 4)
rat32 = make(3, 2)
rat1 = make(2, 3)
rat12 = make(4, 6)
rat2 = make(7, 2)
print(to_string(rat12))  # '4 / 6'
print(is_equal(rat1, rat12))  # True
print(to_string(add(rat1, rat2)))  # 25/6
print(to_string(sub(rat2, rat1)))  # 17/6
print(to_string(mul(rat1, rat2)))  # 14/6
print(to_string(div(rat1, rat2)))  # 4/21
print(to_string(mul(rat32, rat14)))
assert to_string(mul(rat32, rat14)) == to_string(make(3, 8))