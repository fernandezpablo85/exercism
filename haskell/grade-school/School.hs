module School (School, empty, grade, sorted, add) where

import qualified Data.Map as M
import Data.List(sort)

data School = School { impl :: M.Map Int [String] }

add :: Int -> String -> School -> School
add grade name school = School $ M.insert grade newValue imp
  where
    imp = impl school
    newValue = name : M.findWithDefault [] grade imp


empty :: School
empty = School $ M.fromList []

sorted :: School -> [(Int, [String])]
sorted = M.toAscList . M.map sort . impl

grade :: Int -> School -> [String]
grade n school = M.findWithDefault [] n (impl school)