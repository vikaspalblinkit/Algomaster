"""
    Generator in Python : Don't hold the integrate results in memory 
    it basically yields on the way, one results at one times. 
    Memory and CPUs performs boosting while working with the large datasets
"""

# function to add perform the square 
def squareNums(nums):
    results = [] 
    for i in nums:
        results.append(i*i) 
    return results


def squareNumWithGenerator(nums):
    for i in nums:
        yield (i*i)
    


nums = squareNums([1,2,3,4]) 

new_results = [x*x for x in [1,2,3,4]] # Shortcut in python to do square in one inline functions 
generator_obj = squareNumWithGenerator([1,2]) 

print(next(generator_obj)) # Don't hold the results in memory generate the no on the run basically. 
print(new_results)
print(nums) # StopIteration exception it will be once all the elements are iterated over the data 

for num in generator_obj:
    print(num) # We don't in-countered the stop iteration exception while working with generators 