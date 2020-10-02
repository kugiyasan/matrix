import numpy as np
import timeit


# print(A - B)
# print(6 * A)
# print(np.dot(A, B))
# print(np.dot(B, A))
# # print(A.minor(2, 2))
# # print(A.cofactor(2, 2))
# print(np.linalg.inv(E))
# # print(np.inv(E))
# print(np.transpose(A))
# print(np.linalg.det(A))
# print(np.trace(A))

def setup():
    import numpy as np


def loop():
    A = np.array([
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ])
    B = np.array([
        [1, 0, 0],
        [0, 0, 1],
        [0, 1, 0]
    ])
    E = np.array([
        [1, 2],
        [2, -3]
    ])

    A + B
    A - B
    6 * A
    np.dot(A, B)
    np.dot(B, A)
    # A.minor(2, 2)
    # np.linalg.inv(A).T * np.linalg.det(A) # cofactor
    np.linalg.inv(E)
    # np.inv(E)
    # np.transpose(A)
    A.T
    np.linalg.det(A)
    np.trace(A)


times = timeit.repeat(stmt=loop, setup=setup, repeat=5, number=10000)
print("\n".join(str(t*1000) for t in times))
