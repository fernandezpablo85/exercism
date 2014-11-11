module Grains where

import Data.Bits(shiftL)

square :: Int -> Integer
square x = 2 ^ (x - 1)

total :: Integer
total = pred $ 1 `shiftL` 64