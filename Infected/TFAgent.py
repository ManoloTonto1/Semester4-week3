import numpy as np
from tf_agents.environments import py_environment
from tf_agents.specs import array_spec
from tf_agents.trajectories import time_step as ts

BOARD_SIZE = 5

# Clicks per turn 
CLICKS_PER_TURN = 2

# move directions
UP = 0
DOWN = 1
LEFT = 2
RIGHT = 3
UP_LEFT = 4
UP_RIGHT = 5
DOWN_LEFT = 6
DOWN_RIGHT = 7

JUMP_UP = 8
JUMP_DOWN = 9
JUMP_LEFT = 10
JUMP_RIGHT = 11
JUMP_UP_LEFT = 12
JUMP_UP_RIGHT = 13
JUMP_DOWN_LEFT = 14
JUMP_DOWN_RIGHT = 15

# rewards
GAIN_INFECTED = 1
LOSE_INFECTED = -1
WIN_GAME_REWARD = 10
LOSE_GAME_REWARD = -10

# CELL TYPES
UNOCCUPIED = 0
OCCUPIED_BY_PLAYER = 1
OCCUPIED_BY_ENEMY = 2

# players
PLAYERS = 2
PLAYER = 0
ENEMY = 1


class infectedGameEnvironment(py_environment.PyEnvironment):
    """Environment for the Infected game."""

    def __init__(self, boardSize=BOARD_SIZE, discount=0.9,clicksPerTurn= CLICKS_PER_TURN) -> None:

        super(infectedGameEnvironment, self).__init__()
        assert boardSize > 0
        self.discount = discount
        self._boardSize = BOARD_SIZE
        self._totalMoves = 0
        self._turnEnded = False
        self._clicksPerTurn = clicksPerTurn
        self._actionSpec = array_spec.BoundedArraySpec(
            shape=(), dtype=np.int32, minimum=0, maximum=4**2, name='action')

        self._observationSpec = array_spec.BoundedArraySpec(
            shape=(self._boardSize, self._boardSize),
            dtype=np.int32, minimum=UNOCCUPIED, maximum=[OCCUPIED_BY_PLAYER, OCCUPIED_BY_ENEMY])

        self.setBoard()

    def setBoard(self):
        self._board = np.zeros(
            (self._boardSize, self._boardSize), dtype=np.int32)
        # set the first and last cell to be occupied
        # one for the player and one for the enemy
        self._board[0][0] = OCCUPIED_BY_PLAYER
        self._board[self._boardSize -
                    1][self._boardSize - 1] = OCCUPIED_BY_ENEMY
        self.instances = 1
        self._enemyInstances = 1
        self.instanceSelected = np.array([-1, -1])

    def action_spec(self):
        return self._actionSpec

    def observation_spec(self):
        return self._observationSpec

    def _reset(self):
        self._totalMoves = 0
        self._turnEnded = False
        self.setBoard()
        return ts.restart(self._board)
    
    def raiseMoves(self):
        self._totalMoves += 1

    def checkUnder(self, action_x, action_y):
        try:
            if self._board[action_x][action_y-1] == OCCUPIED_BY_ENEMY:
                self.self._board[action_x][action_y-1] == OCCUPIED_BY_PLAYER
                self.instances += 1
                self._enemyInstances -= 1
                return ts.transition(self._board, reward=GAIN_INFECTED, discount=self.discount)
        except: 
            return
        
    def checkAbove(self, action_x, action_y):
        try:
            if self._board[action_x][action_y+1] == OCCUPIED_BY_ENEMY:
                self.self._board[action_x][action_y+1] == OCCUPIED_BY_PLAYER
                self.instances += 1
                self._enemyInstances -= 1
                return ts.transition(self._board, reward=GAIN_INFECTED, discount=self.discount)

        except: 
            return
        
    def checkLeft(self, action_x, action_y):
        try:
            if self._board[action_x-1][action_y] == OCCUPIED_BY_ENEMY:
                self.self._board[action_x-1][action_y] == OCCUPIED_BY_PLAYER
                self.instances += 1
                self._enemyInstances -= 1
                return ts.transition(self._board, reward=GAIN_INFECTED, discount=self.discount)

        except: 
            return
    def checkRight(self, action_x, action_y):
        try:
            if self._board[action_x+1][action_y] == OCCUPIED_BY_ENEMY:
                self.self._board[action_x+1][action_y] == OCCUPIED_BY_PLAYER
                self.instances += 1
                self._enemyInstances -= 1
                return ts.transition(self._board, reward=GAIN_INFECTED, discount=self.discount)
        except: 
            return
    def checkDiagonal(self, action_x, action_y):
        try:
            if self._board[action_x+1][action_y+1] == OCCUPIED_BY_ENEMY:
                self.self._board[action_x+1][action_y+1] == OCCUPIED_BY_PLAYER
                self.instances += 1
                self._enemyInstances -= 1
                return ts.transition(self._board, reward=GAIN_INFECTED, discount=self.discount)
                
        except:
            """continue"""
        try:
            if self._board[action_x+1][action_y-1] == OCCUPIED_BY_ENEMY:
                self.self._board[action_x+1][action_y-1] == OCCUPIED_BY_PLAYER
                self.instances += 1
                self._enemyInstances -= 1
                return ts.transition(self._board, reward=GAIN_INFECTED, discount=self.discount)
        except:
            """continue"""

        try:
            if self._board[action_x-1][action_y+1] == OCCUPIED_BY_ENEMY:
                self.self._board[action_x-1][action_y+1] == OCCUPIED_BY_PLAYER
                self.instances += 1
                self._enemyInstances -= 1
                return ts.transition(self._board, reward=GAIN_INFECTED, discount=self.discount)
        except:
            """continue"""
        
        try:
            if self._board[action_x-1][action_y-1] == OCCUPIED_BY_ENEMY:
                self.self._board[action_x-1][action_y-1] == OCCUPIED_BY_PLAYER
                self.instances += 1
                self._enemyInstances -= 1
                return ts.transition(self._board, reward=GAIN_INFECTED, discount=self.discount)
        except:
            """continue"""

    def infect(self, action_x, action_y) -> int:
        points = 0
        points += self.checkUnder(action_x, action_y)
        points += self.checkAbove(action_x, action_y)
        points += self.checkLeft(action_x, action_y)
        points += self.checkRight(action_x, action_y)
        points += self.checkDiagonal(action_x, action_y)
        return points

# y is left and right is x
    def getCoordinatesFromAction(self, action):
        if action == UP:
            return -1, 0
        elif action == DOWN:
            return 1, 0
        elif action == LEFT:
            return 0, -1
        elif action == RIGHT:
            return 0, 1
        elif action == UP_LEFT:
            return -1, -1
        elif action == UP_RIGHT:
            return -1, 1
        elif action == DOWN_LEFT:
            return 1, -1
        elif action == DOWN_RIGHT:
            return 1, 1
        elif action == JUMP_UP:
            return -2, 0
        elif action == JUMP_DOWN:
            return 2, 0
        elif action == JUMP_LEFT:
            return 0, -2
        elif action == JUMP_RIGHT:
            return 0, 2
        elif action == JUMP_UP_LEFT:
            return -2, -1
        elif action == JUMP_UP_RIGHT:
            return -2, 1
        elif action == JUMP_DOWN_LEFT:
            return 2, -1
        elif action == JUMP_DOWN_RIGHT:
            return 2, 1
    
    def _step(self, action):
        """Move the player based on the action"""
        if self.instances == 0:
            self._turnEnded = True
            self.reset()
            return ts.termination(np.array(self._board, dtype=np.int32), LOSE_GAME_REWARD)

        if self._enemyInstances == 0:
            self._turnEnded = True
            self.reset()
            return ts.termination(np.array(self._board, dtype=np.int32), WIN_GAME_REWARD)

        action_y, action_x = self.getCoordinatesFromAction(action)
        action_x += self.instanceSelected[0]
        action_y += self.instanceSelected[1]

        if self._board[action_x][action_y] == OCCUPIED_BY_PLAYER & self.instanceSelected == np.array([-1, -1]):
            self.instanceSelected = np.array([action_x, action_y])
            self._turnEnded = False
            return ts.transition(np.array(self._board, dtype=np.int32), 1)
            
        if self.board[action_x][action_y] == UNOCCUPIED & self._movesExecuted == 1:

            # move to position
            # don't touch enemies bc it does not matter
            if self._board[action_x][action_y] == OCCUPIED_BY_ENEMY:
                self._turnEnded = False
                return ts.transition(np.array(self._board, dtype=np.int32), 0)

            if self._board[action_x][action_y] == UNOCCUPIED:
                if action > 7:
                    self._board[self.instanceSelected[0]][self.instanceSelected[1]] = UNOCCUPIED
                    self.instanceSelected = np.array([-1, -1])
                    
                self._board[action_x][action_y] = OCCUPIED_BY_PLAYER
                self.instances += 1
                self._turnEnded = True
                return self.infect(action_x, action_y)
        def render(self, mode: "human") -> np.ndarray:
            if mode != "human":
                 raise ValueError(
                    "Only rendering mode supported is 'human', got {} instead.".format(
                        mode))
            return self._visible_board
