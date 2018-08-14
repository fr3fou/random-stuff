function binarySearch(arr, num) {
    let startIndex = 0; // first element
    let endIndex = arr.length - 1; // last element

    while (startIndex < endIndex) {
        let guessedIndex = Math.floor((startIndex + endIndex) / 2); // the middle num between our start and end index is our guess
        let guessedNumber = arr[guessedIndex];

        if (guessedNumber == num) { // if we have found the num
            return guessedIndex; // return the index
        } else {
            if (guessedNumber > num) { // if the guessed number is too big
                endIndex = guessedIndex - 1; // start again until the index BEFORE the guessed num
            } else { // else the guessed number was too small
                startIndex = guessedIndex + 1; // start again at the index AFTER the guessed num
            }
        }
    }
    return undefined;
}

console.log(binarySearch([2, 5, 7, 11, 12, 15, 20, 30, 34, 36], 34));