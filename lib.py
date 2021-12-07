from numpy import array
from typing import Tuple
from aocd import submit, get_data
from functools import partial


def level_ab(day: int, test: Tuple = None, levels=(0,1), quiet=False, print_res=True, sep="\n", apply=lambda a: a):
    def solver(solve):
        for level in levels:
            parse = lambda dd: array([apply(d) for d in dd.split(sep)])
            
            # solve with testdata
            if test and (actual := test[1+level]) != None:
                result = solve(parse(test[0]), level)
                assert result == actual, f"Test failed {result} != {actual} with data\n{repr(test[0])}"
            
            # solve real
            sol = solve(parse(get_data(day=day)), level)
            submit(sol, 'ab'[level], day=day, quiet=quiet)
        return solve # return the original function, so you can put another annotation on it
    return solver

level_a = partial(level_ab, levels=(0,))
level_b = partial(level_ab, levels=(1,))