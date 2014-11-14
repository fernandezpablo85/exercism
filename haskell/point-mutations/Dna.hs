{-# OPTIONS_GHC -Wall #-}

module DNA (hammingDistance) where

hammingDistance :: String -> String -> Int
hammingDistance first second = sum differences
  where
    differences = zipWith areDifferent first second
    areDifferent a b = if a /= b then 1 else 0