from typing import Tuple
from aocd import submit, get_data
from functools import partial

def level_ab(day: int, test: Tuple = None, methods=(0,1), quiet=False, print_res=True, sep="\n"):
    def solver(solve):
        for method in methods:
            # solve with testdata
            if test and (actual := test[1+method]) != None:
                result = solve(test[0].split(sep), method)
                assert result == actual, f"Test failed {result} != {actual} with data\n{repr(test[0])}"
            
            # solve real
            sol = solve(get_data(day=day).split(sep), method)
            submit(sol, 'ab'[method], day=day, quiet=quiet)
        return solve # return the original function, so you can put another annotation on it
    return solver

level_a = partial(level_ab, methods=(0,))
level_b = partial(level_ab, methods=(1,))