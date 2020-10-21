import timeit


class Matrix(list):
    def __init__(self, matrix):
        """
        Yes, numpy does the same job and is faster, I know
        The Matrix class is zero-based, so self.matrix[0][0] is the first value
        """
        super().__init__(matrix)
        self.matrix = list(matrix)

    def __add__(self, other):
        if len(self.matrix) != len(other.matrix) or len(self.matrix[0]) != len(other.matrix[0]):
            raise ValueError("Matrices aren't the same size")

        output = [[0] * len(self.matrix[0]) for _ in range(len(self.matrix))]
        for i in range(len(self.matrix)):
            for j in range(len(self.matrix[0])):
                output[i][j] = self.matrix[i][j] + other.matrix[i][j]
        return Matrix(output)

    def __neg__(self):
        return -1 * self

    def __sub__(self, other):
        return self + (-other)

    def __mul__(self, other):
        if type(other) == type(self):
            if len(self.matrix[0]) != len(other.matrix):
                raise ValueError(
                    "Matrices aren't compatible for matrix multiplication")

            n = len(self.matrix[0])
            output = [[0] * len(other.matrix[0])
                      for _ in range(len(self.matrix))]

            for i in range(len(self.matrix)):
                for j in range(len(other.matrix[0])):
                    output[i][j] = sum(self.matrix[i][k] *
                                       other.matrix[k][j] for k in range(n))

            return Matrix(output)

        return self.__rmul__(other)

    def __rmul__(self, other):
        """__rmul__ is called by __mul__, so the code here will run on both sides of the * operator"""
        if type(other) in (int, float):
            output = [[0] * len(self.matrix[0])
                      for _ in range(len(self.matrix))]
            for i in range(len(self.matrix)):
                for j in range(len(self.matrix[0])):
                    output[i][j] = self.matrix[i][j] * other
            return Matrix(output)

        raise TypeError("Not a compatible type")

    def __repr__(self):
        wrapper = "Matrix [\n  {}\n]"
        return wrapper.format(",\n  ".join(f"[{' '.join(str(cell) for cell in col)}]" for col in self.matrix))

    def minor(self, i, j):
        if i >= len(self.matrix) or j >= len(self.matrix[0]):
            raise IndexError("i or j can't be bigger than the actual matrix")

        output = [row[:] for row in self.matrix]
        output.pop(i)
        for row in output:
            row.pop(j)

        return det(Matrix(output))

    def cofactor(self, i, j):
        return (-1) ** (i+j) * self.minor(i, j)

    def inv(self):
        """inv returns the inverted matrix"""
        determinant = det(self)
        if determinant == 0:
            raise ZeroDivisionError(
                "The determinant is null, the matrix is singular")

        return 1 / determinant * adj(self)

    @property
    def T(self):
        """T returns the tranposed matrix"""
        output = zip(*self.matrix)
        return Matrix([list(row) for row in output])


def adj(matrix: Matrix):
    """adj returns the adjacent of the given Matrix"""
    output = [[0] * len(matrix[0]) for _ in range(len(matrix))]

    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            output[i][j] = matrix.cofactor(i, j)

    return Matrix(output).T

def det(A: Matrix):
    """det returns the determinant of the given Matrix"""
    if len(A) != len(A[0]):
        raise ValueError("You need a square matrix to find the determinant")

    if len(A) == 1:
        return A[0][0]
    elif len(A) == 2:
        return A[0][0] * A[1][1] - A[1][0] * A[0][1]

    # https://integratedmlai.com/find-the-determinant-of-a-matrix-with-pure-python-without-numpy-or-scipy/
    # Section 1: Establish n parameter and copy A
    n = len(A)
    AM = [[0] * len(A[0]) for _ in range(len(A))]
    for i in range(len(AM)):
        AM[i] = A[i][:]

 
    # Section 2: Row ops on A to get in upper triangle form
    for fd in range(n): # A) fd stands for focus diagonal
        for i in range(fd+1,n): # B) only use rows below fd row
            if AM[fd][fd] == 0: # C) if diagonal is zero ...
                AM[fd][fd] == 1.0e-18 # change to ~zero
            # D) cr stands for "current row"
            crScaler = AM[i][fd] / AM[fd][fd] 
            # E) cr - crScaler * fdRow, one element at a time
            for j in range(n): 
                AM[i][j] = AM[i][j] - crScaler * AM[fd][j]
     
    # Section 3: Once AM is in upper triangle form ...
    product = 1.0
    for i in range(n):
        # ... product of diagonals is determinant
        product *= AM[i][i] 
 
    return product

def oldDet(matrix: Matrix):
    """det returns the determinant of the given Matrix"""
    if len(matrix) != len(matrix[0]):
        raise ValueError("You need a square matrix to find the determinant")

    if len(matrix) == 1:
        return matrix[0][0]
    elif len(matrix) == 2:
        return matrix[0][0] * matrix[1][1] - matrix[1][0] * matrix[0][1]

    i = 0
    n = len(matrix)
    return sum(matrix[i][k] * matrix.cofactor(i, k) for k in range(n))


def tr(matrix):
    """returns the trace of the matrix"""
    if len(matrix) != len(matrix[0]):
        raise ValueError("You need a square matrix to find the determinant")

    n = len(matrix)
    return sum(matrix[k][k] for k in range(n))


def I(n):
    """returns a identity matrix of size n"""
    return Matrix([[int(i == j) for j in range(n)] for i in range(n)])


# A = Matrix([
#     [1, 2, 3],
#     [4, 5, 6],
#     [7, 8, 9]
# ])
# B = Matrix([
#     [1, 0, 0],
#     [0, 0, 1],
#     [0, 1, 0]
# ])

# C = Matrix([
#     [1, 2, 3],
#     [4, 5, 6]
# ])
# D = Matrix([
#     [7, 8],
#     [9, 10],
#     [11, 12]
# ])
# E = Matrix([
#     [1, 2],
#     [2, -3]
# ])
# F = Matrix([
#     [0, 8, 0],
#     [0, 3, 4],
#     [0, 6, 7]
# ])

# G = Matrix((1, 4, 54))

# H = Matrix([
#     [4, 2, 3, 9, 9],
#     [-2, 4, 7, -7, -7],
#     [2, 3, 11, 1, 1],
#     [1, 1, 2, -3, -1],
#     [1, 1, 2, 0, 1]
# ])
# H = Matrix([
#     [2, 3, 9, 9],
#     [4, 7, -7, -7],
#     [3, 11, 1, 1],
#     [1, 2, -3, -1],
# ])

# print(A - B)
# print(6 * A)
# print(A * B)
# print(B * A)
# print(A.minor(2, 2))
# print(A.cofactor(2, 2))
# # print(E.inv())
# print(A.T)
# print(det(A))
# # print(det(F))
# print(tr(A))
# print(I(4))
# print(det(H))


# A[2][2] = 69
# print(A)
# print(A.pop())


def loop():
    A = Matrix([
        [1, 2, 1, 0, 23, 45, -12, 6, -1, 3],
        [4, 5, 2, -5, -1, -5, 0, 6, 21, 4],
        [4, 7, -7, 8, -8, 9, 4, 1, 6, 5],
        [7, 2, 5, 11, 2, 43, 41, 13, 15, 67],
        [-14, 15, 1, -7, 6, -81, 1, 456, -4, -464],
        [-2, 5, 144, 3, 10, 17, 78, 29, 65, 9],
        [4, 46, -561, 145, -5, 54, 613, 41, 854, 6],
        [4, 21, -56, 41, 84, 879, 12, 84, 6, 45],
        [4, 71, 54, 12051, 65, 98, 7, 86951, 1, 51],
        [-100, -1, -8, 4, 2, 7, 4, -42, 7, 42],
    ])

    # A + B
    # A - B
    # A * B
    # B * A
    # 6 * A
    # A.minor(2, 2)
    # A.cofactor(2, 2)
    # A.T
    # print(det(A))
    det(A)
    # print(oldDet(A))
    # tr(A)
    # E.inv()


times = timeit.repeat(stmt=loop, repeat=5, number=10000)
print("[" + " ".join("{:.4f}ms".format(t*1000) for t in times) + "]")
