class Hangman(object):
	"""Hangman Game State. AI_role can either be 'g' for guesser or 's' for setter"""
	def __init__(self, sowpods, word, board, ai_role, misses, guess = None):
		self.word = word
		self.board = board
		self.ai_role = ai_role
		self.misses = misses
		# find a better way to share sowpods
		self.sowpods = sowpods
		self.guess = None

	def ai_move(self):
		if self.ai_role = 's' and len(self.word) == 0: self.word = sowpods.get_random()
		if self.ai_role = 'g':
			matches = self.sowpods.find_matches(self.board, self.misses)
			letters = {}
			for word in matches:
				for a in word.letters:
					if a in letters: letters[a] += word.letters[a]
					else: letters[a] = word.letters[a]
			self.guess = max(letters, key=lambda k: letters[k])
		if self.guess in self.word:
			for i in range(len(self.word)):
				if self.word[i] == self.guess: self.board[i] = self.guess
		else: self.misses.append(self.guess)