amean :: Fractional a => a -> a -> a
amean x y = (x + y) / 2

hmean :: Fractional a => a -> a -> a
hmean x y = (2 * x * y) / (x + y)

max' :: Ord a => a -> a -> a
max' x y | x > y     = x
         | otherwise = y
        
min' :: Ord a => a -> a -> a
min' x y | x < y     = x
         | otherwise = y

sgn :: (Ord a, Num a) => a -> Int
sgn x | x > 0   = 1
      | x == 0  = 0
      | x < 0   = -1

type Point = (Int, Int)

origin :: Point
origin = (0, 0)

strafe ::  Int -> Point -> Point
strafe dist (x, y) = (x + dist, y)

walk ::  Int -> Point -> Point
walk dist (x, y) = (x, y + dist)

head' :: [a] -> a
head' []    = error "empty list"
head' (x:_) = x

tail' :: [a] -> [a]
tail' []     = error "empty list"
tail' (_:xs) = xs

length' :: [a] -> Int
length' []     = 0
length' (_:xs) = 1 + length' xs

sum' :: Num a => [a] -> a
sum' []     = 0
sum' (x:xs) = head' xs + (sum' xs)

elem' :: Eq a => a -> [a] -> Bool
elem' _ []     = False
elem' y (x:xs) 
      | x == y    = True 
      | otherwise =  elem' y xs

maximum' :: Ord a => [a] -> a
maximum' []     = error "empty list"
maximum' [x]    = x
maximum' (x:xs) = max' x (maximum' xs) 

replicate' :: (Num i, Ord i) => i -> a -> [a]
replicate' n x
    | n < 0 = []
    | otherwise = x:replicate' (n-1) x

take' :: (Num i, Ord i) => i -> [a] -> [a]
take' _ [] = []
take' n _
    | n <= 0 = []
take' n (x:xs) = x:take' (n-1) xs

reverse' :: [a] -> [a]
reverse' [] = []
reverse' (x:xs) = (reverse' xs) ++ [x]

main :: IO ()
main = print "ok"
