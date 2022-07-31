def linearSearch(list, target):
    """
    Returns the index of the target value if found, else return None
    """

    for i in range(0, len(list)):
        if list[i] == target :
            return i
    return None


def verify(result):
    if result is not None:
        print("Target found in index: ", result)
    else:
        print("Target not found in list")



nums =[1,2,3,4,5,6,7,8,9.10]

result = linearSearch(nums,5)
verify(result)