from hashlib import md5, sha1, sha256
from sys import argv
from os import _exit
if len(argv) <= 1:
    print("no string provided")
    _exit(1)

hashString = argv[1]

md5_ = md5(hashString.encode()).digest().hex()
sha1_ = sha1(hashString.encode()).digest().hex()
sha256_ = sha256(hashString.encode()).digest().hex()


print(f"md5 hash: {md5_}")
print(f"sha1 hash: {sha1_}")
print(f"sha256 hash: {sha256_}")
