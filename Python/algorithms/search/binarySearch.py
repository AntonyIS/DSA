def binarySearch(list, target):
    """
    Returns the index of the target if found , else returns None
    """
    first = 0
    last = len(list) - 1

    while first <= last :
        midpoint = (first + last)//2

        if list[midpoint] == target :
            return midpoint
        elif list[midpoint] < target :
            first = midpoint + 1
        else:
            last = midpoint - 1

    return None

def verify(result):
    if result is not None:
        print("Target found in index: ", result)
    else:
        print("Target not found in list")



nums =[1,2,3,4,5,6,7,8,9.10]

result = binarySearch(nums,5)
verify(result)