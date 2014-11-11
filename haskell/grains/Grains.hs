module Grains where

import Data.Bits(shiftL)

square :: Int -> Integer
square = (shiftL 1) . pred

total :: Integer
total = pred $ 1 `shiftL` 64