const checkRow = (row) => {
  const exes = row.filter(el => el === 'x');
  const oes = row.filter(el => el === 'o');
  if (exes.length === 3) {
    return 'x';
  } else if (oes.length === 3) {
    return 'o';
  }
  return false;
};


const winner = (grid) => {
  let win = false;
  for (let row = 0; row < grid.length; row += 1) {
    win = checkRow(row);
    if (win) {
      return win;
    }
  }
  for (let col = 0; col < grid[0].length; col += 1) {
    const column = [];
    for (let row = 0; row < grid.length; row += 1) {
      column.push(grid[row][col]);
    }
    win = checkRow(column);
    if (win) {
      return win;
    }
  }
  const downRight = [grid[0][0], grid[1][1], grid[2][2]];
  const upRight = [grid[2][0], grid[1][1], grid[0][2]];
  win = checkRow(downRight);
  if (win) {
    return win;
  } else if (checkRow(upRight)) {
    return checkRow(upRight);
  }
  return false;
};

export default winner;
