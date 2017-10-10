export const sameArr = (rc1, rc2) => {
  if (rc1 === rc2) return true;
  if (rc1 == null || rc2 == null) return false;
  if (rc1.length !== rc2.length) return false;
  for (let i = 0; i < rc1.length; i += 1) {
    if (rc1[i] !== rc2[i]) return false;
  }
  return true;
};

export const emptyMatrix = (numRows, numCols, char = () => '') => {
  const grid = [];
  for (let r = 0; r < numRows; r += 1) {
    let row = [];
    for (let c = 0; c < numCols; c += 1) {
      row.push(char());
    }
    grid.push(row);
  }
  return grid;
};

export const isEmptyMatrix = (matrix) => {
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix.length; j++) {
      if (matrix[i][j].length) {
        return false;
      }
    }
  }
  return true;
}

export const copyMatrix = (matrix) => {
  const newMatrix = [];
  for (let i = 0; i < matrix.length; i++) {
    let row = []
    for (let j = 0; j < matrix[i].length; j++) {
      row.push(matrix[i][j]);
    }
    newMatrix.push(row);
  }
  return newMatrix;
}

export const isEmptyBoard = (board, r, c) => {
  for (let i = 0; i < r; i++) {
    for (let j = 0; j < c; j++) {
      if (board[r][c].length) {
        return false;
      }
    }
  }
  return true;
}

export const emptyBoard = (r, c) => {
  const board = {}
  for (let i = 0; i < r; i++) {
    let row = []
    for (let j = 0; j < c; j++) {
      row.push('');
    }
    board[`${i}`] = row;
  }
  return board;
}

export const matrixCount = (matrix, n) => {
  let count = 0;
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (matrix[i][j] === n) {
        count += 1;
      }
    }
  }
  return count;
}

export const mapCount = (map, n) => {
  let count = 0;
  Object.keys(map).forEach(key => {
    for (let i = 0; i < map[key].length; i++) {
      if (map[key][i] === n) {
        count += 1;
      }
    }
  });
  return count
}

export const mapFromMatrix = (matrix) => {
  let map = {};
  for (let i = 0; i < matrix.length; i++) {
    let row = [];
    for (let j = 0; j < matrix[i].length; j++) {
      row.push(matrix[i][j])
    }
    map[i] = row
  }
  return map
} 