{-# OPTIONS_GHC -Wall #-}

module Strain (keep, discard) where

keep :: (a -> Bool) -> [a] -> [a]
keep p xs = [x | x <- xs, p x]

discard :: (a -> Bool) -> [a] -> [a]
discard p xs = keep (not . p) xs