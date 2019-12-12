quickSort :: (Ord a) => [a] -> [a]
quickSort [] = []
quickSort (x:xs) = low ++ [x] ++ high
          where low = quickSort [a | a <- xs, a <= x]
                high = quickSort [a | a <- xs, a > x]

main :: IO ()
main = quickSort [4, 5, 12, 6, 3, 8, 11, 15, 1, 2]

