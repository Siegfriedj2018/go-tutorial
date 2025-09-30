import math
def merge_sort(nums):
    if len(nums) < 2:
        return nums

    midpoint = math.trunc(len(nums)/2)
    first_half = nums[:midpoint]
    second_half = nums[midpoint:]
    split_first = merge_sort(first_half)
    split_second = merge_sort(second_half)
    final_list = merge(split_first, split_second)
    return final_list


def merge(first, second):
    final = []
    i = 0
    j = 0
    while i < len(first) and j < len(second):
        print(first)
        print(second)
        if first[i] <= second[j]:
            final.append(first[i])
            i += 1
        else:
            final.append(second[j])
            j += 1
        

    final.append(first[i:])
    final.append(second[j:])
    
    return final
