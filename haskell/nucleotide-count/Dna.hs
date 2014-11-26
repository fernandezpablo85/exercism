module DNA(count, nucleotideCounts) where

import Data.Map

count :: Char -> String -> Int
count n chain = case Data.Map.lookup n (nucleotideCounts chain) of
                Nothing -> error "invalid nucleotide 'X'"
                Just x -> x

nucleotideCounts :: String -> Map Char Int
nucleotideCounts xs = nucleotideCounts' xs defaults

nucleotideCounts' :: String -> Map Char Int -> Map Char Int
nucleotideCounts' [] acc = acc
nucleotideCounts' (x:xs) acc = nucleotideCounts' xs (insert x next acc)
  where next = (succ (acc ! x))

defaults :: Map Char Int
defaults = fromList [('A',0),('C',0),('G',0),('T',0)]