def maxIncreaseKeepingSkyline(grid):
    row, col = map(max, grid), map(max, zip(*grid))
    print row, col
    return sum(min(i, j) for i in row for j in col) - sum(map(sum, grid))

if __name__ == '__main__':
    grid = [[3, 0, 8, 4], [2, 4, 5, 7], [9, 2, 6, 3], [0, 3, 1, 0]]
    max = maxIncreaseKeepingSkyline(grid)
    print max
