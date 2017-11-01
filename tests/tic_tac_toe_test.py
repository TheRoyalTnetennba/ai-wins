import sys 
import os
sys.path.append(os.path.abspath("/home/graham/ai-wins/backend"))
from tic_tac_toe import *


def test1():
	board = [['x','o','x'], ['','',''],['','','']]
	state = TTT(board, 'x')
	print('TTT takes middle if available:', state.board[1][1] == 'o')

def test2():
	board = [['x', 'o', 'x'],['x', 'o', ''], ['', '', '']]
	state = TTT(board, 'x')
	print('TTT prioritizes win over preventing loss:', state.board[2][1] == 'o')

def test3():
	board = [['x', '', ''],['x', 'o', ''], ['', '', '']]
	state = TTT(board, 'x')
	print('TTT prevents 3-verticals:', state.board[2][0] == 'o')

def test4():
	board = [['x', 'x', ''], ['', 'o', ''], ['', '', '']]
	state = TTT(board, 'x')
	print('TTT prevents 3-horizontals:', state.board[0][2] == 'o')

def test5():
	board = [['x', 'o', 'x'],['', 'x', ''], ['', '', '']]
	state = TTT(board, 'x')
	print('TTT prevents 3-down-right', state.board[2][2] == 'o')

def test6():
	board = [['', 'o', ''], ['', 'x', ''], ['x', '', '']]
	state = TTT(board, 'x')
	print('TTT prevents 3-up-right', state.board[0][2] == 'o')

def test7():
	board = [['', 'o', ''], ['o', 'x', ''], ['x', '', '']]
	state = TTT(board, 'x')
	print('TTT waits its goddamned turn when its o:', state.board[0][2] == '')

def test8():
	board = [['', 'o', ''], ['', 'x', ''], ['x', '', '']]
	state = TTT(board, 'o')
	print('TTT waits its goddamned turn when its x:', state.board[0][2] == '')

def test9():
	board = [['x', '', ''], ['', 'o', ''], ['', '', 'x']]
	state = TTT(board, 'x')
	print('TTT is all over that sneaky corner strat', state.board[0][1] == 'o')

def test10():
	board = [['x', '', ''], ['', 'o', ''], ['', '', '']]
	state = TTT(board, 'o')
	print('TTT sets itself up for that sneaky corner strat', state.board[2][2] == 'x')

def main():
	test_list = [test1(), test2(), test3(), test4(), test5(), test6(), test7(), test8(), test9(), test10()]


main()