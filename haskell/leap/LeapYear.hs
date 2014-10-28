{-# OPTIONS_GHC -Wall #-}

module LeapYear (isLeapYear) where

isLeapYear :: Integral a => a -> Bool
isLeapYear year = (divisibleBy 4) && ( not (divisibleBy 100) || (divisibleBy 400) )
  where
    divisibleBy n = year `mod` n == 0