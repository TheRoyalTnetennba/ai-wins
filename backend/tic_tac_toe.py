class ClassName(Object):
  """TTT Game State"""
    def __init__(self, board, user_marker):
        self.board = board
        self.user_marker = user_marker
        self.ai_marker = 'o' if user_marker == 'x' else 'x'
        self.result = False
        self.threats = []
        self.board_status()

    def three_in_row(self, row):
        return len(row[0]) > 0 and all(c == row[0] for c in row)

    def two_in_row(self, row):
        played = [i for i in row if len(i) > 0]
        return len(played) == 2 and all(c != ai_marker for c in played)

    def board_status(self):
        x, o = 0, 0
        for r in range(3):
            col = []
            for c in range(3):
                col.append(self.board[c][r])
                if self.board[r][c] == 'x': x += 1
                elif self.board[r][c] == 'o': o += 1
            self.result = self.board[r][0] if three_in_row(self, self.board[r]) else self.result
            self.result = self.board[0][r] if three_in_row(self, col) else self.result
        down_right = [self.board[i][i] for i in range(3)]
        up_right = [self.board[2 - i][i] for i in range(3)]
        self.result = down_right[0] if three_in_row(self, down_right) else self.result
        self.result = up_right[0] if three_in_row(self, up_right) else self.result
        self.result = 'tie' if !self.result and x + o == 9
        self.ai_turn = (x > o and self.ai_marker == 'o') or (x == o and self.ai_marker == 'x')

    def ai_move(self):
        if !self.ai_turn:
            return
        if self.result:
            return
        if len(self.board[1][1]) == 0 and self.ai_turn:
            self.board[1][1] = self.ai_marker
        elif 
