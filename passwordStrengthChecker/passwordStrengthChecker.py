from sys import argv

def contiansUpper(password: str):
    for digit in password:
        if digit > "A" and digit < "Z":
            return True
    return False

def contiansLower(password: str):
    for digit in password:
        if digit > "a" and digit < "z":
            return True
    return False

def containSpecial(passowrd: str):
    for digit in passowrd:
        if digit in "!@#$%&*":
            return True
    return False

if len(argv) <= 1:
    print("Append a password when calling a script")

password = argv[1]


if password.__len__() < 8:
    print("password is to short, must be 8 digits long")
elif contiansLower(password):
    print("password must contain at least one lower case")
elif contiansUpper(password):
    print("password must contain at least one upper case")
elif not containSpecial(password):
    print("password must contain at least one of these '!@#$%&*'")
else:
    print("password is perfecto")

