{-# OPTIONS_GHC -Wall #-}

module DNA (hammingDistance) where

hammingDistance :: String -> String -> Int
hammingDistance [] [] = 0
hammingDistance (x:xs) (y:ys) = if x /= y then
                                  1 + hammingDistance xs ys
                                else
                                  hammingDistance xs ys
hammingDistance _ _ = error "sequences must have equal length."