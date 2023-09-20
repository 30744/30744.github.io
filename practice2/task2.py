try:
    print(undefined_variable)
except NameError as e:
    print("Exception:", e)


var_int = 42
print("var_int =", var_int)


var_flt = 42.0
print("var_flt =", var_flt)


try:
    var_int_cast_to_float = float(var_int)
    print("Successfully cast var_int to float:", var_int_cast_to_float)
except ValueError as e:
    print("Exception:", e)


try:
    var_int_cast_to_str = str(var_int)
    print("Successfully cast var_int to str:", var_int_cast_to_str)
except ValueError as e:
    print("Exception:", e)
