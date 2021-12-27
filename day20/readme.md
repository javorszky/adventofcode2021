# Day 20

This was fairly easy, essentially flip some stuff.

## Task 1

Very similar to conway's game of life, which we also had numerous ones last year.

Parse the enhance algorithm into a struct with methods, parse the image into a 2 dimensional grid, and on each tick, construct an entirely new grid based on the old one, and a 1px all around growth, and reassign.

## Task 2

This was still fast, but the result was wrong. The issue was that when all pixels are dark, the middle one becomes light, and when all pixels are light, the middle one becomes dark, which means the outside infinity pixels would keep flashing between dark and light between each tick.

To fix that I had to keep track of what the outside pixel was, and then assign those when the calculation was out of bounds, and at the end, reassign the outside pixel to the other one.
