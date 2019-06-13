# The Conway Game

## Introduction
The conway game was designed by Horton Conway in 1970.

## Rules
 - Death: if the count is less than 2 or greater than 3, the current cell is switched off.

- Survival: if (a) the count is exactly 2, or (b) the count is exactly 3 and the current cell is on, the current cell is left unchanged.

- Birth: if the current cell is off and the count is exactly 3, the current cell is switched on. 

## Use

go run main.go numerGenerations ?animated(true)

Enjoy!!
