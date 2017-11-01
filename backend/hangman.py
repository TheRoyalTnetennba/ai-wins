

class Hangman(object):
	"""Hangman Game State. AI_role can either be 'g' for guesser or 's' for setter'"""
	def __init__(self, board, ai_role, misses, guess = None):
		self.board = board
		self.ai_role = ai_role
		self.misses = misses
		self.guess = None

	def ai_move(self):
		if self.ai_role = 'g'
