{-# OPTIONS_GHC -Wall #-}

module Gigasecond (fromDay) where

import Data.Time.Calendar (Day, addDays)

fromDay :: Day -> Day
fromDay day = addDays daysToGigasecond day
  where daysToGigasecond = gigasecond `div` secondsInDay
        gigasecond = 10 ^ (9 :: Int)
        secondsInDay = 24 * 60 * 60