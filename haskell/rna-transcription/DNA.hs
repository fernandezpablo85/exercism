{-# OPTIONS_GHC -Wall #-}

module DNA (toRNA) where

toRNA :: String -> String
toRNA = map transcribe
  where
    transcribe 'G' = 'C'
    transcribe 'C' = 'G'
    transcribe 'T' = 'A'
    transcribe 'A' = 'U'
    transcribe other = other
