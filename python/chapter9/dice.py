import random


class Dice:
    def __init__(self, sides: int) -> None:
        self.sides = sides

    def roll_die(self):
        print(random.randint(1, self.sides) + 1)


dice = Dice(6)
for i in range(11):
    dice.roll_die()

dice10 = Dice(10)
for i in range(10):
    dice10.roll_die()

dice20 = Dice(20)
for i in range(10):
    dice20.roll_die()
