{-# OPTIONS_GHC -Wall #-}

module Strain (keep, discard) where

keep :: (a -> Bool) -> [a] -> [a]
keep p xs = [x | x <- xs, p x]

keep' :: (a -> Bool) -> [a] -> [a]
keep' _ [] = []
keep' p (x : xs) = if p x then x : keep' p xs else keep' p xs

discard :: (a -> Bool) -> [a] -> [a]
discard p = keep (not . p)