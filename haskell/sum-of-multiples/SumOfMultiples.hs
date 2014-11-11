{-# OPTIONS_GHC -Wall #-}

module SumOfMultiples (sumOfMultiples, sumOfMultiplesDefault) where

sumOfMultiplesDefault :: Int -> Int
sumOfMultiplesDefault = sumOfMultiples [3, 5]

sumOfMultiples :: [Int] -> Int -> Int
sumOfMultiples ns n = sum multiples
  where
    multiples = [x | x <- [1..n-1], isMultiple x]
    isMultiple x = any (\f -> f x) multiples'
    multiples' = map isMultipleOf ns
    isMultipleOf y x = (x `mod` y == 0)
