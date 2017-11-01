class TTT(object):
	"""TTT Game State"""
	
	def __init__(self, board, user_marker):
		self.board = board
		self.user = user_marker
		self.ai = 'o' if user_marker == 'x' else 'x'
		self.result = False
		self.loss = []
		self.win = []
		self.empty = []
		self.board_status()
		self.ai_move()

	def three_in_row(self, row):
		return len(row[0]) > 0 and all(c == row[0] for c in row)

	def two_in_row(self, row, marker):
		played = [i for i in row if len(i) > 0]
		return len(played) == 2 and all(c != marker for c in played)

	def opposite_corners(self):
		mid_sides = [[0, 1], [1, 0], [1, 2], [2, 1]]
		empty_mid_sides = [i for i in self.empty if i in mid_sides]
		if len(empty_mid_sides) == 0: return False
		elif self.board[0][0] == self.board[2][2] and self.board[0][0] == self.user:
			return empty_mid_sides[0]
		elif self.board[2][0] == self.board[0][2] and self.board[0][2] == self.user:
			return empty_mid_sides[0]
		else:
			return False

	def available_corners(self):
		corners = [[0, 0], [0, 2], [2, 0], [2, 2]]
		empty_corners = [i for i in self.empty if i in corners]
		ai_corners = [i for i in corners if self.board[i[0]][i[1]] == self.ai]
		if len(ai_corners) > 0 and len(empty_corners) > 0:
			for i in ai_corners:
				opposite = [2 if i[0] == 0 else 0, 2 if i[1] == 0 else 0]
				if opposite in empty_corners: return opposite
		if len(empty_corners) > 0: return empty_corners[0]
		else: return False

	def board_status(self):
		x, o = 0, 0
		for r in range(3):
			col = []
			for c in range(3):
				col.append(self.board[c][r])
				if self.board[r][c] == 'x': x += 1
				elif self.board[r][c] == 'o': o += 1
				else: self.empty.append([r, c])
			if self.three_in_row(self.board[r]): self.result = self.board[r][0]
			if self.three_in_row(col): self.result = self.board[0][r]
			if self.two_in_row(self.board[r], self.ai): self.loss.append([r, self.board[r].index("")])
			if self.two_in_row(self.board[r], self.user): self.win.append([r, self.board[r].index("")])
			if self.two_in_row(col, self.ai): self.loss.append([col.index(""), r])
			if self.two_in_row(col, self.user): self.win.append([col.index(""), r])
		down_r = [self.board[i][i] for i in range(3)]
		up_r = [self.board[2 - i][i] for i in range(3)]
		if self.three_in_row(down_r): self.result = down_r[0]
		if self.three_in_row(up_r): self.result = up_r[0]
		if self.two_in_row(down_r, self.ai): self.loss.append([down_r.index(""), down_r.index("")])
		if self.two_in_row(down_r, self.user): self.win.append([down_r.index(""), down_r.index("")])
		if self.two_in_row(up_r, self.ai): self.loss.append([2 - up_r.index(""), up_r.index("")])
		if self.two_in_row(up_r, self.user): self.win.append([2 - up_r.index(""), up_r.index("")])
		if not self.result and x + o == 9: self.result = 't'
		self.ai_turn = (x > o and self.ai == 'o') or (x == o and self.ai == 'x')

	def ai_move(self):
		opposite = self.opposite_corners()
		corner = self.available_corners()
		if not self.ai_turn:
			return
		if self.result:
			return
		if len(self.board[1][1]) == 0 and self.ai_turn:
			self.board[1][1] = self.ai
		elif len(self.win) > 0:
			self.board[self.win[0][0]][self.win[0][1]] = self.ai
		elif len(self.loss) > 0:
			self.board[self.loss[0][0]][self.loss[0][1]] = self.ai
		elif opposite:
			self.board[opposite[0]][opposite[1]] = self.ai
		elif corner:
			self.board[corner[0]][corner[1]] = self.ai
		else:
			for r in range(3):
				for c in range(3):
					if len(self.board[r][c]) == 0:
						self.board[r][c] = self.ai
						return
