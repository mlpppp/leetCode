import random

# ------------------Solution-----------------

class Solution:
    kSequence = []
    def isSorted(self, arr): 
        # return True if the input array is sorted
        if len(arr) ==0 :
            return True
        p_OrderSetEnd = 0
        while (p_OrderSetEnd < len(arr)-1):
            if arr[p_OrderSetEnd+1] == arr[p_OrderSetEnd]+1:
                p_OrderSetEnd += 1
            else:
                break
        
        if (p_OrderSetEnd == len(arr)-1):        #! the array has already in sorted order, return the result
            return True
        else:
            return False 

    # partialFlip: flip partial array from index 0 to "k-1"
    def partialFlip(self, arr, k):
        self.kSequence.append(k)  # ! add the filp operation to solution
        p_head = 0
        p_tail = k-1
        while (p_head < p_tail):
            temp = arr[p_head]
            arr[p_head] = arr[p_tail]
            arr[p_tail] = temp
            p_head += 1
            p_tail -= 1


    def partialFlipSort(self, arr):
        # input: unsorted integer array
        # output: [sortedArray, kValueList]
        # p_OrderSetEnd: is a pointer used to search a index:
            #  "index of the end of partial order group started with 0", if the ordered partial array exist 
        # p_NextInt: is a pointer used to search a index when p_OrderSetEnd==0:
            #  index of value "arr[0]+1" (exploit the input's "non-repetitive positive integer from 1 to len-1" property)
        self.kSequence = []
        if len(arr) == 0:
            return [arr, self.kSequence]

        partialFlipCount = 0
        while(1):
            # ! iteration control
            if (partialFlipCount > int(10*len(arr))):
                raise Exception("too many partialFlipCount")
            partialFlipCount+=1

            #! find the end of partial order group from the start of array
            p_OrderSetEnd = 0
            while (p_OrderSetEnd < len(arr)-1):
                if arr[p_OrderSetEnd+1] == arr[p_OrderSetEnd]+1:
                    p_OrderSetEnd += 1
                else:
                    break
            
            if (p_OrderSetEnd == len(arr)-1):        #! the array has already in sorted order, return the result
                return [arr, self.kSequence]

            if (p_OrderSetEnd > 0):
                self.partialFlip(arr, p_OrderSetEnd+1)
            else:
                if arr[0] == len(arr):          #! when a[0] is argmax(arr), flip the whole array
                    self.partialFlip(arr, len(arr))

                else:                           #! otherwise find the index of a[0]+1, flip in index p_NextInt
                    p_NextInt = 0
                    while (p_NextInt < len(arr)):
                        if (arr[p_NextInt] == (arr[0] + 1)):
                            break
                        p_NextInt += 1
                    self.partialFlip(arr, p_NextInt)    



# ------------Tests----------------------



def PartialFlip_printTest(Solution):
    size = 50
    tests = [random.sample(range(1, size+1), size) for num in range(10)]
    for test_case in tests: 
        print("unsorted: " + str(test_case))
        _ = Solution.partialFlipSort(test_case)
        print("sorted:   " + str(test_case))
        print("seq len:   " + str(len(Solution.kSequence)))
        print('-------------------------------')



def PartialFlip_performanceTest(Solution):   
    print( "100 random arrays in each array length group: [0, 10, 20, 50, 100, 1000]")
    print('-------------------------------')

    def generateTestData():
        # generate 100 tests for array size of [0, 10, 20, 50, 100]
        sizes = [0, 10, 20, 50, 100, 1000]
        tests = {size: [] for size in sizes}
        tests[0] = [[]]
        for size in sizes:
            tests[size] = [random.sample(range(1, size+1), size) for num in range(100)]

        return tests


    tests = generateTestData()
    for size, testSuit in tests.items():
        print("problem size: {}".format(size))
        success_list = []
        flipNum_list = []

        for test_case in testSuit:
            bad_iteration = False
            try:
                _ = Solution.partialFlipSort(test_case)
            except:     # ! the number of iteration exceed 10*len(arr), failed
                bad_iteration = True   

            bad_sorting = not Solution.isSorted(test_case) # ! returned array is not sorted, failed
            if (bad_iteration or bad_sorting):
                success_list.append(0)
            else:
                success_list.append(1)
                flipNum_list.append(len(Solution.kSequence))
        if(len(success_list)>0) :
            print("success rate: {}".format(sum(success_list)/len(success_list)))
        if (len(flipNum_list)>0):
            print("average number of flip: {}".format(sum(flipNum_list)/len(flipNum_list)))
        print('-------------------------------')



# 100 random arrays in each array length group: [0, 10, 20, 50, 100, 1000]            
# -------------------------------
# problem size: 0
# success rate: 1.0
# average number of flip: 0.0
# -------------------------------
# problem size: 10
# success rate: 1.0
# average number of flip: 10.66
# -------------------------------
# problem size: 20
# success rate: 1.0
# average number of flip: 26.57
# -------------------------------
# problem size: 50
# success rate: 1.0
# average number of flip: 76.72
# -------------------------------
# problem size: 100
# success rate: 1.0
# average number of flip: 166.55
# -------------------------------
# problem size: 1000
# success rate: 1.0
# average number of flip: 1943.83
# -------------------------------


            



    
 
def main():
    random.seed( 10 )

    sol = Solution()
    PartialFlip_printTest(sol)
    PartialFlip_performanceTest(sol)

            
if __name__ == "__main__":
    main()


