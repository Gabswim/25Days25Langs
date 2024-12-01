from typing import List, Tuple

def get_total_distance_between_list_sol_1(list_a: List[str], list_b: List[str]):
    total_distance = 0

    list_a.sort()
    list_b.sort()

    assert(len(list_a) == len(list_b))

    for a, b in zip(list_a, list_b):
        total_distance += abs(a - b)

    return total_distance


def sol1():
    list_a, list_b = build_list("input.txt")

    return get_total_distance_between_list_sol_1(list_a, list_b)

def sol2():
    list_a, list_b = build_list("input.txt")
    return get_total_distance_between_list_sol_2(list_a, list_b)


def get_total_distance_between_list_sol_2(list_a: List[str], list_b: List[str]):
    total_distance = 0

    b_occurrences = {}
    for b in list_b:
        count = b_occurrences.get(b, 0)
        b_occurrences[b] = count + 1

    for a in list_a:
        total_distance += b_occurrences.get(a, 0) * a

    return total_distance

def build_list(file_name: str) -> Tuple[List[str], List[str]]:
    file = open("input.txt", "r")
    list_a = []
    list_b = []
    for line in file:
        itemA = int(line[0:5])
        list_a.append(itemA)
        itemB = int(line[8:13])
        list_b.append(itemB)
    return list_a, list_b

def test(actual, expected):
    if actual == expected:
        print("✅ Test pass")
    else:
        print(f"❌ Test fail, Expected: {expected}, Actual: {actual}")

test(get_total_distance_between_list_sol_1([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3]), 11)

test(sol1(), 2430334)

test(get_total_distance_between_list_sol_2([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3]), 31)

test(sol2(), 28786472)