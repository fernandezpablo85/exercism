{-# OPTIONS_GHC -Wall #-}

module SumOfMultiples (sumOfMultiples, sumOfMultiplesDefault) where

sumOfMultiplesDefault :: Int -> Int
sumOfMultiplesDefault = sumOfMultiples [3, 5]

sumOfMultiples :: [Int] -> Int -> Int
sumOfMultiples divisors n = sum multiples
  where
    multiples = [x | x <- [1..n-1], anyMultiple x]
    anyMultiple x = isMultipleOf `any` divisors
      where
        isMultipleOf = (==0) . (mod x)
