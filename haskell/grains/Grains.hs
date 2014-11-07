module Grains where

square :: Integral a => a -> a
square 1 = 1
square n = 2 * square (n - 1)

total :: Integral a => a
total = sum (map square board)
  where board = [1..64]