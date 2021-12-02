from aocd import submit, get_data

def imag2tup(i: complex) -> tuple:  # 1+2j -> (1, 2)
    return (int(i.real), int(i.imag))

def tup2imag(t: tuple) -> complex:
    return t[0] + t[1] * 1j

def __level(method):
    def annotation(day, quiet=True, print_res=True, do_example=True):
        def solver(solve):
            part = 'ab'[method]
            td = testdata[day-1]
            if(td and do_example):
                if type(td) == list:
                    esol = solve(td[method], method=method)
                else:
                    esol = solve(td, method=method)
                if print_res:
                    print(f"Ex. {day}{part}", esol, sep="\t")
            sol = solve(get_data(day=day), method=method)
            if print_res:
                print(f"Lvl {day}{part}", sol, sep="\t")
            submit(sol, part, day=day, quiet=quiet)
            return solve # return the original function, so you can put another annotation on it
        return solver
    return annotation

level_a, level_b = __level(0), __level(1)
def level_ab(lvl):
    def solver(solve):
        level_a(lvl)(solve)
        level_b(lvl)(solve)
    return solver

def setdata(d):
    testdata.clear()
    testdata.extend(d)

testdata = []