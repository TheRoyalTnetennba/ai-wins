class Word(object):
    """Object for Word. Used by the Sowpod class"""
    def __init__(self, entry):
        self.entry = entry
        self.sorted = ''.join(sorted(entry))
        self.anagrams = []

    def get_contents(self):


    def size(self):
        return len(self.entry)

class Sowpod(object):
    """Holds a collection of Scrabble dict words along with some useful methods for sorting and retreiving them"""
    def __init__(self):

        self.word = word

    def load(self):
        with open('resources/sowpods.txt') as sp:
            for line in sp:
                self.words = [Word(i) for i in sp]

