import random

"""use pickle to avoid runnning everytime"""
class Word(object):
    """Object for Word. Used by the Sowpod class"""
    def __init__(self, entry):
        self.entry = entry
        self.sorted = ''.join(sorted(entry))
        self.anagrams = []
        self.letters = {}
        self.count_letters()

    def count_letters(self):
        for i in self.entry:
            if i in self.letters: self.letters[i] += 1
            else: self.letters[i] = 1


    def size(self):
        return len(self.entry)

class Sowpod(object):
    """Holds a collection of Scrabble dictionary words along with some useful methods for sorting and retreiving them"""
    def __init__(self):
        self.words = []
        self.anagrams = {}
        self.letter_frequency = {}
        self.read_txt()

    def read_txt(self):
        with open('resources/sowpods.txt') as sp:
            sorts = {}
            for line in sp:
                word = Word(line)
                if word.sorted in self.anagrams: self.anagrams[word.sorted].append(word)
                else: self.anagrams[word.sorted] = [word]
                for letter in word:
                    if letter in self.letter_frequency: self.letter_frequency[letter] = 1
                    else: self.letter_frequency[letter] += 1

    """optimize for speedy lookup perhaps with words_containing and words_excluding hashmaps or poly-tree"""
    def find_matches(self, board, misses):
        letters = {}
        for letter in board:
            if letter in letters: letters[letter] += 1
            else: letters[letter] = 1
        matches = []
        for word in self.words:
            if (word.size == len(board) and all(i not in misses for i in word.letters) and
                all(board[i] == word.entry[i] or board[i] == '' for i in range(len(board)))):
                matches.append(word)
        return matches

    def get_random(self, size = False, containing = False, excluding = False):
        """containing and excluding should be lists of letters"""
        # switch optional args to lamdas
        sub_list = []
        for word in self.words:
            include = True
            if size: include = include and word.size = size
            if containing: include = include and all(i in word.letters for i in containing)
            if excluding: include = include and none(i in word.letters for i in excluding)
            if include: sub_list.append(word)
        return random.choice(sub_list)